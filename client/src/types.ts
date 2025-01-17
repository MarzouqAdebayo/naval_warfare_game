import { ShipTypeKeys } from "./helpers/shipTypes";

export enum WSEvents {
  // Outgoing Events
  EventSetUserData = "set_user_data",
  EventAttack = "attack",
  EventFindGame = "find_game",
  EventQuitGame = "quit_game",
  EventPlaceShips = "place_ships",

  // Incoming Events
  EventFindGameWaiting = "find_game_waiting",
  EventFindGameStart = "find_game_start",
  EventShipRandomized = "randomized_place_ship_response",
  EventBroadcastAttack = "broadcast_attack",
  EventPong = "pong",
  EventOpponentQuit = "opponent_quit",
  EventClientDisconnected = "client_disconnected",
}

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

export enum GameStatus {
  Waiting = 0,
  Ready = 1,
}

export enum GameMode {
  ContinousFire = 0,
  SingleFire = 1,
}

export type GameData = {
  roomID: string;
  index: number;
  message: string;
  players: Player[];
  currentTurn: 0 | 1;
  winner: string;
  gameOver: boolean;
  mode: GameMode;
  status: GameStatus;
};

export type AppState = {
  timeline: Timeline;
  name: string;
  game: GameData | null;
};

export type RandomShipPlacementPayload = {
  message: string;
  playerData: Player;
};

export interface ShipIconProps {
  type: ShipTypeKeys;
  x: number;
  y: number;
  axis: "X" | "Y";
  length: number;
  sunk: boolean;
}
