import { AppState, GameData, Timeline } from "./types";

const newGame = (): GameData => {
  return {
    roomID: "",
    index: 0,
    message: "",
    players: [],
    currentTurn: 0,
    gameOver: false,
    mode: 0,
    winner: "",
  };
};

export type Action =
  | { type: "INITIALIZE"; payload: object }
  | { type: "CHANGE_TIMELINE"; payload: Timeline }
  | { type: "SET_NAME"; payload: string }
  | { type: "SET_GAME_STATE" }
  | { type: "SET_SERVER_GAME_STATE"; payload: GameData }
  | { type: "UPDATE_SERVER_GAME_STATE"; payload: GameData }
  | { type: "STATUS"; payload: object };

export default function reducer(state: AppState, action: Action): AppState {
  switch (action.type) {
    case "INITIALIZE":
      return state;
    case "CHANGE_TIMELINE":
      return { ...state, timeline: action.payload };
    case "SET_NAME":
      return { ...state, name: action.payload };
    case "SET_GAME_STATE":
      return { ...state, game: newGame() };
    case "SET_SERVER_GAME_STATE": {
      const { roomID, index, message, players, currentTurn, gameOver, mode } =
        action.payload;
      const game: GameData = {
        roomID,
        index,
        message,
        players,
        currentTurn,
        gameOver,
        mode,
        winner: "",
      };
      return { ...state, game };
    }
    case "UPDATE_SERVER_GAME_STATE": {
      console.log(action.payload);
      const { roomID, index, message, players, currentTurn, gameOver, mode } =
        action.payload;
      const game: GameData = {
        roomID,
        index,
        message,
        players,
        currentTurn,
        gameOver,
        mode,
        winner: "",
      };
      return { ...state, game };
    }
    case "STATUS":
      return state;
    default:
      throw new Error("");
  }
}
