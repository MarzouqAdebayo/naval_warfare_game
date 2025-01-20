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
  | { type: "CHANGE_TIMELINE"; payload: Timeline }
  | { type: "SET_PLAYER_NAME"; payload: string }
  | { type: "INITIALIZE_NEW_ROOM"; payload: { roomID: string } }
  | { type: "FIND_GAME_START"; payload: GameData }
  | { type: "SET_RANDOM_SHIP_PLACEMENT"; payload: RandomShipPlacementPayload }
  | { type: "GAME_START"; payload: GameData }
  | { type: "BROADCAST_ATTACK"; payload: GameData };

export default function reducer(state: AppState, action: Action): AppState {
  switch (action.type) {
    case "CHANGE_TIMELINE":
      return { ...state, timeline: action.payload };
    case "SET_PLAYER_NAME":
      return { ...state, name: action.payload };
    case "INITIALIZE_NEW_ROOM":
      return {
        ...state,
        game: newGame(action.payload.roomID),
        timeline: Timeline.Setup,
      };
    case "FIND_GAME_START": {
      if (state.timeline === Timeline.Setup) {
        return { ...state, game: action.payload };
      }
      return { ...state, game: action.payload, timeline: Timeline.Setup };
    }
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
    case "GAME_START": {
      return { ...state, game: action.payload, timeline: Timeline.GameStart };
    }
    case "BROADCAST_ATTACK": {
      return { ...state, game: action.payload };
    }
    default:
      throw new Error(`Event does not exist`);
  }
}
