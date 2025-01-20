import { createContext, ReactNode, useContext } from "react";
import reducer, { Action } from "./reducer";
import {
  useWebSocket,
  WebSocketMessage,
  ConnectionStatus,
} from "./hooks/useWebSocket";
import { AppState, Timeline } from "./types";

interface GameContextType<T> {
  state: AppState;
  dispatch: React.Dispatch<Action>;
  status: ConnectionStatus;
  connectionString: string;
  isConnected: boolean;
  sendMessage: (data: Omit<WebSocketMessage<T>, "from">) => void;
  reconnect: () => void;
}

interface GameProviderProps {
  children: ReactNode;
}

const initialState: AppState = {
  timeline: Timeline.Init,
  name: "",
  game: null,
};

const GameContext = createContext<GameContextType<unknown>>(
  {} as GameContextType<unknown>,
);

function GameProvider({ children }: GameProviderProps) {
  const {
    state,
    dispatch,
    status,
    connectionString,
    sendMessage,
    reconnect,
    isConnected,
  } = useWebSocket("ws://192.168.1.108:5000/ws", reducer, initialState);

  return (
    <GameContext.Provider
      value={{
        state,
        dispatch,
        status,
        connectionString,
        sendMessage,
        reconnect,
        isConnected,
      }}
    >
      {children}
    </GameContext.Provider>
  );
}

export const useGameContext = () => {
  const context = useContext(GameContext);
  if (!context) {
    throw new Error("useGameContext must be used within a GameProvider");
  }
  return context;
};

export default GameProvider;
