import {
  AppState,
  GameData,
  RandomShipPlacementPayload,
  Timeline,
} from "./types";

const newGame = (roomID: string): GameData => {
  return {
    roomID: roomID,
    index: 0,
    message: "",
    players: [],
    currentTurn: 0,
    gameOver: false,
    mode: 0,
    winner: "",
    status: 0,
  };
};

export type Action =
  | { type: "INITIALIZE"; payload: object }
  | { type: "CHANGE_TIMELINE"; payload: Timeline }
  | { type: "SET_NAME"; payload: string }
  | { type: "SET_GAME_STATE"; payload: { roomID: string } }
  | { type: "SET_RANDOM_SHIP_PLACEMENT"; payload: RandomShipPlacementPayload }
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
      return { ...state, game: newGame(action.payload.roomID) };
    case "SET_RANDOM_SHIP_PLACEMENT": {
      const { message, playerData } = action.payload;
      if (!state.game) {
        return state;
      }
      const playersCopy = [...state.game.players];
      playersCopy[state.game.index] = playerData;
      return {
        ...state,
        game: { ...state.game, message, players: playersCopy },
      };
    }
    case "SET_SERVER_GAME_STATE": {
      const {
        roomID,
        index,
        message,
        players,
        currentTurn,
        gameOver,
        mode,
        status,
      } = action.payload;
      const game: GameData = {
        roomID,
        index,
        message,
        players,
        currentTurn,
        gameOver,
        mode,
        winner: "",
        status,
      };
      return { ...state, game };
    }
    case "UPDATE_SERVER_GAME_STATE": {
      console.log(action.payload);
      const {
        roomID,
        index,
        message,
        players,
        currentTurn,
        gameOver,
        mode,
        status,
      } = action.payload;
      const game: GameData = {
        roomID,
        index,
        message,
        players,
        currentTurn,
        gameOver,
        mode,
        winner: "",
        status,
      };
      return { ...state, game };
    }
    case "STATUS":
      return state;
    default:
      throw new Error("");
  }
}
