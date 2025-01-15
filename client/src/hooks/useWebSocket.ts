import { useState, useEffect, useCallback, useRef, useReducer } from "react";
import { Action } from "../reducer";
import { GameData, Timeline } from "../types";

export interface WebSocketMessage<T = unknown> {
  type: string;
  payload: T;
}

export interface WebSocketOptions<T> {
  reconnectAttempts?: number;
  reconnectInterval?: number;
  onOpen?: () => void;
  onClose?: () => void;
  onMessage?: (data: WebSocketMessage<T>) => void;
  onError?: (error: Event) => void;
}

export enum ConnectionStatus {
  CONNECTING = 0,
  OPEN = 1,
  CLOSING = 2,
  CLOSED = 3,
}

export const useWebSocket = <T = unknown, U = unknown>(
  url: string,
  reducer: (state: U, action: Action) => U,
  initialState: U,
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

  const [state, dispatch] = useReducer(
    reducer,
    initialState instanceof Function ? initialState() : initialState,
  );
  const [data, setData] = useState<WebSocketMessage<T> | null>(null);
  const [status, setStatus] = useState<ConnectionStatus>(
    ConnectionStatus.CLOSED,
  );
  const ws = useRef<WebSocket | null>(null);
  const reconnectCount = useRef<number>(0);
  const reconnectTimeoutId = useRef<number | null>(null);

  const handleMessage = useCallback(
    (event: MessageEvent) => {
      try {
        const data: WebSocketMessage<T> = JSON.parse(event.data);
        switch (data.type) {
          case "game_start":
            console.log(data.payload);
            dispatch({
              type: "SET_SERVER_GAME_STATE",
              payload: data.payload as GameData,
            });
            break;
          case "game_update":
            console.log(data.payload);
            dispatch({
              type: "SET_SERVER_GAME_STATE",
              payload: data.payload as GameData,
            });
            break;
          case "game_found":
            // TODO Returns a room ID with intial room data
            // TODO Recieve room data from server
            dispatch({ type: "SET_GAME_STATE" });
            // Dispatch timeline instantly
            dispatch({ type: "CHANGE_TIMELINE", payload: Timeline.Setup });
            break;
          default:
            setData(data);
            break;
        }
        if (onMessage) onMessage(data);
      } catch (error) {
        console.error("Error parsing WebSocket message:", error);
      }
    },
    [onMessage],
  );

  const connect = useCallback(() => {
    if (ws.current?.readyState === ConnectionStatus.OPEN) return;

    ws.current = new WebSocket(url);
    setStatus(ConnectionStatus.CONNECTING);

    ws.current.onopen = () => {
      setStatus(ConnectionStatus.OPEN);
      reconnectCount.current = 0;
      if (onOpen) onOpen();
    };

    ws.current.onclose = () => {
      setStatus(ConnectionStatus.CLOSED);
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

  const sendMessage = useCallback((data: WebSocketMessage<T>) => {
    if (ws.current?.readyState === ConnectionStatus.OPEN) {
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
    state,
    dispatch,
    status,
    data,
    sendMessage,
    reconnect,
    isConnected: status === ConnectionStatus.OPEN,
  };
};
