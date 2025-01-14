import {
  GameStartContainer,
  HudWindow,
  LabelContainer,
} from "../styled_components/gameControllerStyles";
import { useGameContext } from "../../GameController";
import { FriendlyWatersGrid } from "./FriendlyWatersGrid";
import { EnemyWatersGrid } from "./EnemyWatersGrid";

const enum shipTypes {
  Carrier = "Carrier",
  Battleship = "Battleship",
  Cruiser = "Cruiser",
  Submarine = "Submarine",
  Destroyer = "Destroyer",
}

const ships = [
  { Type: shipTypes.Carrier, Size: 5, Hits: 0, Sunk: false },
  { Type: shipTypes.Battleship, Size: 4, Hits: 0, Sunk: false },
  { Type: shipTypes.Cruiser, Size: 3, Hits: 0, Sunk: false },
  { Type: shipTypes.Submarine, Size: 3, Hits: 0, Sunk: false },
  { Type: shipTypes.Destroyer, Size: 2, Hits: 0, Sunk: false },
];

export const GameStart = () => {
  const { sendMessage } = useGameContext();
  return (
    <GameStartContainer>
      <button
        onClick={() => sendMessage({ type: "find_game", payload: null })}
        style={{ position: "absolute", zIndex: 50, top: 0, left: 0 }}
      >
        sendM
      </button>
      <HudWindow>Hi</HudWindow>
      <LabelContainer row="4">
        <h1 style={{ margin: "auto auto 0" }}>Friendly waters</h1>
      </LabelContainer>
      <LabelContainer row="2">
        <h1 style={{ margin: "auto auto 0" }}>Enemy waters</h1>
      </LabelContainer>
      <FriendlyWatersGrid />
      <EnemyWatersGrid />
    </GameStartContainer>
  );
};
