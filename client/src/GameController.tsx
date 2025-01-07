import { createContext, ReactNode, useContext, useReducer } from "react";
import reducer, { Action } from "./reducer";
import { useWebSocket } from "./hooks/useWebSocket";

export interface GameContextType {
  state: GameState;
  dispatch: React.Dispatch<Action>;
}

export enum Timeline {
  Init = "init",
  Setup = "setup",
  GameStart = "game_start",
}

export type GameState = {
  timeline: Timeline;
  players: object;
  turn: 0 | 1;
  message: string;
  winner: string;
};

const initialState: GameState = {
  timeline: Timeline.GameStart,
  players: {},
  turn: 0,
  message: "",
  winner: "",
};

const GameContext = createContext<GameContextType>({} as GameContextType);

interface GameProviderProps {
  children: ReactNode;
}

function GameProvider({ children }: GameProviderProps) {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { status, message, sendMessage, reconnect, isConnected } = useWebSocket(
    "ws://localhost:5000/ws",
  );

  return (
    <GameContext.Provider value={{ state, dispatch }}>
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
