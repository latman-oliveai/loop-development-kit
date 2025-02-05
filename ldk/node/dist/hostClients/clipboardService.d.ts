import { StoppableStream, StreamListener } from './stoppables';
/**
 * The ClipboardService provides access to the OS's clipboard.
 */
export interface ClipboardService {
    /**
     * @returns A Promise resolving with the current contents of the clipboard.
     */
    queryClipboard(): Promise<string>;
    /**
     * Starts listening to changes to the clipboard.
     *
     * @param listener - A function that's called whenever the clipboard's contents change.
     * @param includeOliveHelpsTraffic - A flag to indicate if clipboard aptitude will be triggered when olive helps window is in focus.
     * @returns A stream object that can be stopped.
     */
    streamClipboard(listener: StreamListener<string>, includeOliveHelpsTraffic?: boolean): StoppableStream<string>;
    /**
     * Writes the provided text into the clipboard.
     *
     * @param text
     */
    writeClipboard(text: string): Promise<void>;
}
