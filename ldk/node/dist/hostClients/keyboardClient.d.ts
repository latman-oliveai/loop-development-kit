import { KeyboardClient as KeyboardGRPCClient } from '../grpc/keyboard_grpc_pb';
import BaseClient, { GRPCClientConstructor } from './baseClient';
import { StoppableStream, StreamListener } from './stoppables';
import { HotKeyEvent, HotKeyRequest, KeyboardService, TextStream } from './keyboardService';
/**
 * @internal
 */
export declare class KeyboardClient extends BaseClient<KeyboardGRPCClient> implements KeyboardService {
    streamHotKey(hotKeys: HotKeyRequest, listener: StreamListener<HotKeyEvent>): StoppableStream<HotKeyEvent>;
    streamText(listener: StreamListener<string>): StoppableStream<string>;
    streamChar(listener: StreamListener<TextStream>): StoppableStream<TextStream>;
    protected generateClient(): GRPCClientConstructor<KeyboardGRPCClient>;
    protected serviceName(): string;
}
