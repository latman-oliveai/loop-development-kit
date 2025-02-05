"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class TestSuite {
    constructor(tests, logger) {
        this.tests = tests;
        this.logger = logger;
    }
    async start(host) {
        const elements = {};
        let i = 0;
        for await (const test of this.tests) {
            try {
                await test.runTest(host, this.logger);
                elements[`${i}`] = {
                    value: test.getStatus(),
                    label: test.getId(),
                    order: i + 1,
                    type: 'pair',
                };
                i += 1;
            }
            catch (err) {
                elements[`${i}`] = {
                    value: test.getStatus(),
                    label: test.getId(),
                    order: i + 1,
                    type: 'pair',
                };
                i += 1;
            }
        }
        const hotkeys = {
            key: '/',
            modifiers: {
                ctrl: true,
            },
        };
        const listWhisper = host.whisper.listWhisper({
            label: 'Self Test - Results',
            markdown: '',
            elements,
        });
        const markdown = host.whisper.markdownWhisper({
            label: 'Self Test - Results',
            markdown: 'Press "Ctrl + /" to bring back up the original whisper',
        });
        const keyboard = host.keyboard.streamHotKey(hotkeys, (error, response) => {
            if (error) {
                // Do nothing
            }
            else {
                listWhisper.stop();
                markdown.stop();
                keyboard.stop();
            }
        });
        return true;
    }
}
exports.default = TestSuite;
