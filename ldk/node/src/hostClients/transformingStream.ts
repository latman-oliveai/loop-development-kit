import * as grpc from '@grpc/grpc-js';
import { ServiceError } from '@grpc/grpc-js/build/src/call';
import { StoppableStream, StreamListener } from './stoppables';
import { Logger } from '../logging';

/**
 * @internal
 */
export type StreamTransformer<TInput, TOutput> = (input: TInput) => TOutput | undefined;

/**
 * @internal
 */
export type MessageWithError = {
  getError(): string;
};

/**
 * The TransformingStream is a wrapper class that abstracts the grpc.ClientReadableStream interface away from the
 * user and transforms the input from the grpc format to Node objects.
 *
 * This is used when the Library sensor is providing a stream of events, instead of a one-time response.
 *
 * @internal
 */
export class TransformingStream<TInput extends MessageWithError, TOutput>
  implements StoppableStream<TOutput> {
  private stream: grpc.ClientReadableStream<TInput>;

  private transformer: StreamTransformer<TInput, TOutput>;

  private listener: StreamListener<TOutput> | undefined;

  private logger: Logger;

  /**
   * @param stream - the stream object
   * @param transformer - a transformer function that converts the grpc input to the desired output.
   * @param listener - an optional listener function provided by the consumer that is called whenever events are outputted.
   */
  constructor(
    stream: grpc.ClientReadableStream<TInput>,
    transformer: StreamTransformer<TInput, TOutput>,
    listener?: StreamListener<TOutput>,
  ) {
    this.stream = stream;
    this.transformer = transformer;
    this.listener = listener;
    this.stream.addListener('data', this.streamWatcher);
    this.stream.addListener('error', this.errorWatcher);
    this.logger = new Logger('loop-core');
  }

  setListener(callback: StreamListener<TOutput>): void {
    this.listener = callback;
  }

  streamWatcher = (stream: TInput): void => {
    if (this.listener) {
      const error = stream.getError();
      if (error !== '') {
        this.listener(error);
      } else {
        this.listener(null, this.transformer(stream));
      }
    }
  };

  errorWatcher = (error: ServiceError): void => {
    if (this.listener) {
      this.listener(error.details);

      // Stream was stopped, need to remove the stream here
      // to handle the "error" from cancelling the stream
      if (error.code === 1) {
        this.logger.debug('Cancelling stream');
        this.stream.removeAllListeners('error');
      }
    }
  };

  stop(): void {
    // SIDE-1556: Needs to be wrapped this way so that we don't trigger a race condition
    setImmediate(() => {
      this.stream.cancel();
      this.stream.removeAllListeners('data');
    });
  }
}
