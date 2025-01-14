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
  start: number;
  axis: "X" | "Y";
  ship_length: number;
  sunk: CellState;
}
