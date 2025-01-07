import {
  GameStartContainer,
  HudWindow,
  LabelContainer,
} from "../styled_components/gameControllerStyles";
import { FriendlyWatersGrid } from "./FriendlyWatersGrid";

export const GameStart = () => {
  return (
    <GameStartContainer>
      <HudWindow>Hi</HudWindow>
      <LabelContainer row="4">
        <h1 style={{ margin: "auto auto 0" }}>Friendly waters</h1>
      </LabelContainer>
      <FriendlyWatersGrid />
    </GameStartContainer>
  );
};
