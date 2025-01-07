import { useState, useEffect, useCallback, useRef } from "react";

export interface WebSocketMessage<T = unknown> {
  type: string;
  content: T;
  from?: string;
  to?: string;
}

export interface WebSocketOptions<T> {
  reconnectAttempts?: number;
  reconnectInterval?: number;
  onOpen?: () => void;
  onClose?: () => void;
  onMessage?: (data: WebSocketMessage<T>) => void;
  onError?: (error: Event) => void;
}

export enum WebSocketStatus {
  CONNECTING = 0,
  OPEN = 1,
  CLOSING = 2,
  CLOSED = 3,
}

export const useWebSocket = <T = unknown>(
  url: string,
  options: WebSocketOptions<T> = {},
) => {
  const {
    reconnectAttempts = 5,
    reconnectInterval = 3000,
    onOpen,
    onClose,
    onMessage,
    onError,
  } = options;

  const [status, setStatus] = useState<WebSocketStatus>(WebSocketStatus.CLOSED);
  const [message, setMessage] = useState<WebSocketMessage<T> | null>(null);
  const ws = useRef<WebSocket | null>(null);
  const reconnectCount = useRef<number>(0);
  const reconnectTimeoutId = useRef<number | null>(null);

  const handleMessage = useCallback(
    (event: MessageEvent) => {
      try {
        //const data: WebSocketMessage<T> = JSON.parse(event.data);
        const data = {} as WebSocketMessage<T>;
        setMessage(data);
        if (onMessage) onMessage(data);
      } catch (error) {
        console.error("Error parsing WebSocket message:", error);
      }
    },
    [onMessage],
  );

  const connect = useCallback(() => {
    if (ws.current?.readyState === WebSocketStatus.OPEN) return;

    ws.current = new WebSocket(url);
    setStatus(WebSocketStatus.CONNECTING);

    ws.current.onopen = () => {
      setStatus(WebSocketStatus.OPEN);
      reconnectCount.current = 0;
      if (onOpen) onOpen();
    };

    ws.current.onclose = () => {
      setStatus(WebSocketStatus.CLOSED);
      if (onClose) onClose();

      if (reconnectCount.current < reconnectAttempts) {
        reconnectCount.current += 1;
        reconnectTimeoutId.current = setTimeout(
          connect,
          reconnectInterval * 2 ** reconnectCount.current,
        );
      }
    };

    ws.current.onerror = (error: Event) => {
      if (onError) onError(error);
    };

    ws.current.onmessage = handleMessage;
  }, [
    url,
    reconnectAttempts,
    reconnectInterval,
    handleMessage,
    onOpen,
    onClose,
    onError,
  ]);

  useEffect(() => {
    if (window.WebSocket) {
      connect();
    }

    return () => {
      if (reconnectTimeoutId.current) {
        clearTimeout(reconnectTimeoutId.current);
      }
      if (ws.current) {
        ws.current.close();
        ws.current = null;
      }
    };
  }, [connect]);

  const sendMessage = useCallback((data: Omit<WebSocketMessage<T>, "from">) => {
    if (ws.current?.readyState === WebSocketStatus.OPEN) {
      ws.current.send(JSON.stringify(data));
    } else {
      console.warn("WebSocket is not connected");
    }
  }, []);

  const reconnect = useCallback(() => {
    if (ws.current) {
      ws.current.close();
    }
    connect();
  }, [connect]);

  return {
    status,
    message,
    sendMessage,
    reconnect,
    isConnected: status === WebSocketStatus.OPEN,
  };
};
