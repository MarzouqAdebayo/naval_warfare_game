import { ShipTypeKeys } from "./helpers/shipTypes";

export enum Timeline {
  Init = "init",
  Menu = "menu",
  Setup = "setup",
  GameStart = "game_start",
}

export enum CellState {
  "Empty" = "Empty",
  "Ship" = "Ship",
  "Hit" = "Hit",
  "Miss" = "Miss",
  "Sunk" = "Sunk",
}

export type Player = {
  board: CellState[][];
  fleet: ShipIconProps[];
};

export type GameData = {
  roomID: string;
  index: number;
  message: string;
  players: Player[];
  currentTurn: 0 | 1;
  winner: string;
  gameOver: boolean;
  mode: 0 | 1;
};

export type AppState = {
  timeline: Timeline;
  name: string;
  game: GameData | null;
};

export interface ShipIconProps {
  type: ShipTypeKeys;
  x: number;
  y: number;
  axis: "X" | "Y";
  length: number;
  sunk: boolean;
}
