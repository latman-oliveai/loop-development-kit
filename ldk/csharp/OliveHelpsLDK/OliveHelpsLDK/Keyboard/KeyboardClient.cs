using System.Threading;
using Grpc.Core;
using OliveHelpsLDK.Logging;
using Proto;

namespace OliveHelpsLDK.Keyboard
{
    internal class KeyboardClient : BaseClient<Proto.Keyboard.KeyboardClient>, IKeyboardService
    {
        internal KeyboardClient(Proto.Keyboard.KeyboardClient client, Session session, ILogger logger) : base(
            client, session, logger, "keyboard")
        {
        }

        internal KeyboardClient(CallInvoker callInvoker, Session session, ILogger logger) : this(
            new Proto.Keyboard.KeyboardClient(callInvoker), session, logger)
        {
        }

        public IStreamingCall<bool> StreamHotKey(HotKey hotkey, CancellationToken cancellationToken = default)
        {
            var req = new KeyboardHotkeyStreamRequest
            {
                Session = CreateSession(),
                Hotkey = ToProto(hotkey),
            };
            var call = Client.KeyboardHotkeyStream(req, CreateOptions(cancellationToken));
            var loggedParser = LoggedParser<KeyboardHotkeyStreamResponse, bool>(response => response.Scanned);
            return new StreamingCall<KeyboardHotkeyStreamResponse, bool>(call, loggedParser);
        }

        public IStreamingCall<string> StreamText(CancellationToken cancellationToken = default)
        {
            var req = new KeyboardTextStreamRequest
            {
                Session = CreateSession()
            };
            var call = Client.KeyboardTextStream(req, CreateOptions(cancellationToken));
            var loggedParser = LoggedParser<KeyboardTextStreamResponse, string>(resp => resp.Text);
            return new StreamingCall<KeyboardTextStreamResponse, string>(call, loggedParser);
        }

        public IStreamingCall<string> StreamChar(CancellationToken cancellationToken = default)
        {
            var req = new KeyboardCharacterStreamRequest
            {
                Session = CreateSession()
            };
            var call = Client.KeyboardCharacterStream(req, CreateOptions(cancellationToken));
            var loggedParser = LoggedParser<KeyboardCharacterStreamResponse, string>(resp => resp.Text);
            return new StreamingCall<KeyboardCharacterStreamResponse, string>(call, loggedParser);
        }

        private static KeyboardHotkey ToProto(HotKey hotKey)
        {
            return new KeyboardHotkey
            {
                Key = hotKey.Key,
                Modifiers = hotKey.Modifiers()
            };
        }
    }
}