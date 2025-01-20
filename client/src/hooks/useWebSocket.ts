import { useState, useEffect, useCallback, useRef, useReducer } from "react";
import { Action } from "../reducer";
import { GameData, RandomShipPlacementPayload, WSEvents } from "../types";

export interface WebSocketMessage<T = unknown> {
  type: WSEvents;
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

export function connectionStatusString(status: ConnectionStatus) {
  return {
    [ConnectionStatus.CONNECTING]: "Connecting",
    [ConnectionStatus.OPEN]: "Connected",
    [ConnectionStatus.CLOSING]: "Disconnecting",
    [ConnectionStatus.CLOSED]: "Disconnected",
  }[status];
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
          case WSEvents.EventFindGameWaiting:
            dispatch({
              type: "INITIALIZE_NEW_ROOM",
              payload: data.payload as { roomID: string },
            });
            break;
          case WSEvents.EventFindGameStart:
            dispatch({
              type: "FIND_GAME_START",
              payload: data.payload as GameData,
            });
            break;
          case WSEvents.EventShipRandomized:
            dispatch({
              type: "SET_RANDOM_SHIP_PLACEMENT",
              payload: data.payload as RandomShipPlacementPayload,
            });
            break;
          case WSEvents.EventGameStart:
            dispatch({
              type: "GAME_START",
              payload: data.payload as GameData,
            });
            break;
          case WSEvents.EventBroadcastAttack:
            dispatch({
              type: "BROADCAST_ATTACK",
              payload: data.payload as GameData,
            });
            break;
          default:
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
    connectionString: connectionStatusString(status),
    sendMessage,
    reconnect,
    isConnected: status === ConnectionStatus.OPEN,
  };
};
