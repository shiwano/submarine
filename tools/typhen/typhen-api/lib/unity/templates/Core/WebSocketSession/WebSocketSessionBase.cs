﻿// This file was generated by typhen-api

using System;
using System.Collections;
using System.IO;
using WebSocketSharp;

namespace TyphenApi
{
    public interface IWebSocketSession
    {
        ISerializer MessageSerializer { get; }
        IDeserializer MessageDeserializer { get; }
        void Send(int messageType, IType message);
    }

    public abstract class WebSocketSessionBase<ApiT, ErrorT> : IWebSocketSession, IDisposable
        where ApiT : IWebSocketApi
        where ErrorT : class, IType, new()
    {
        const byte MessageTypeBytesLength = 4;
        volatile bool isOpened;

        protected readonly WebSocket connection;
        protected readonly SynchronizedFunctionCaller synchronizedFunctionCaller = new SynchronizedFunctionCaller();

        public ISerializer MessageSerializer { get; protected set; }
        public IDeserializer MessageDeserializer { get; protected set; }

        public Uri RequestUri { get; private set; }
        public ApiT Api { get; protected set; }
        public bool IsOpened { get { return isOpened; } }

        public abstract void OnConnectionCreate(WebSocket connection);
        public abstract void OnConnectionOpen();
        public abstract void OnConnectionClose(ushort code, string reason, bool wasClean);
        public abstract void OnBeforeMessageSend(IType message);
        public abstract void OnMessageReceive(IType message);
        public abstract void OnError(WebSocketSessionError<ErrorT> error);

        protected WebSocketSessionBase(string requestUri)
        {
            RequestUri = new Uri(requestUri);

            connection = new WebSocket(RequestUri.ToString());
            connection.OnMessage += OnMessage;
            connection.OnOpen += OnOpen;;
            connection.OnClose += OnClose;
            connection.OnError += OnError;

            OnConnectionCreate(connection);
        }

        public void Dispose()
        {
            Close();
        }

        public WebSocketSessionError<ErrorT> Error(Exception error)
        {
            return error as WebSocketSessionError<ErrorT>;
        }

        public void Update()
        {
            synchronizedFunctionCaller.Call();
        }

        public IEnumerator ConnectAsync()
        {
            Connect();

            while (!IsOpened)
            {
                yield return null;
            }
        }

        public void Connect()
        {
            connection.Connect();
        }

        public void Close()
        {
            connection.Close();
        }

        public void Send(int messageType, IType message)
        {
            OnBeforeMessageSend(message);
            var messageData = MessageSerializer.Serialize(message);
            var messageTypeBytes = BitConverter.GetBytes(messageType);

            if (!BitConverter.IsLittleEndian)
            {
                Array.Reverse(messageTypeBytes);
            }

            using (var stream = new MemoryStream(messageData.Length + MessageTypeBytesLength))
            {
                stream.Write(messageTypeBytes, 0, MessageTypeBytesLength);
                stream.Write(messageData, 0, messageData.Length);
                var data = stream.GetBuffer();

                if (isOpened)
                {
                    connection.Send(data);
                }
            }
        }

        void OnMessage(object sender, MessageEventArgs e)
        {
            if (e.IsBinary)
            {
                var messageTypeBytes = new byte[MessageTypeBytesLength];
                var messageData = new byte[e.RawData.Length - MessageTypeBytesLength];

                using (var stream = new MemoryStream(e.RawData))
                {
                    stream.Read(messageTypeBytes, 0, MessageTypeBytesLength);
                    stream.Read(messageData, 0, e.RawData.Length - MessageTypeBytesLength);
                }

                if (!BitConverter.IsLittleEndian)
                {
                    Array.Reverse(messageTypeBytes);
                }

                var messageType = BitConverter.ToInt32(messageTypeBytes, 0);
                var func = CreateDispatchMessageFunction(messageType, messageData);
                synchronizedFunctionCaller.ReserveCall(func);
            }
        }

        SynchronizedFunctionCaller.Function CreateDispatchMessageFunction(int messageType, byte[] messageData)
        {
            return () =>
            {
                IType message;

                try
                {
                    message = Api.DispatchMessageEvent(messageType, messageData);
                }
                catch (SerializationException e)
                {
                    OnError(new WebSocketSessionError<ErrorT>(null, e, e.Message));
                    return;
                }

                if (message == null)
                {
                    return;
                }

                var messageAsError = message as ErrorT;
                if (messageAsError != null)
                {
                    OnError(new WebSocketSessionError<ErrorT>(messageAsError, null, null));
                }
                else
                {
                    OnMessageReceive(message);
                }
            };
        }

        void OnOpen(object sender, EventArgs e)
        {
            isOpened = true;
            synchronizedFunctionCaller.ReserveCall(() => OnConnectionOpen());
        }

        void OnClose(object sender, CloseEventArgs e)
        {
            isOpened = false;
            synchronizedFunctionCaller.ReserveCall(() => OnConnectionClose(e.Code, e.Reason, e.WasClean));
        }

        void OnError(object sender, WebSocketSharp.ErrorEventArgs e)
        {
            var error = new WebSocketSessionError<ErrorT>(null, e.Exception, e.Message);
            synchronizedFunctionCaller.ReserveCall(() => OnError(error));
        }
    }
}
