import { GameState } from "./GameController";

export type Action = { type: "INITIALIZE"; payload: object };

export default function reducer(state: GameState, action: Action): GameState {
  switch (action.type) {
    case "INITIALIZE":
      break;
    default:
      throw new Error("");
  }
  return state;
}
