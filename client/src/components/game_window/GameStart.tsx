import {
  GameStartContainer,
  HudWindow,
  LabelContainer,
} from "../styled_components/gameControllerStyles";
import { useGameContext } from "../../GameController";
import { FriendlyWatersGrid } from "./FriendlyWatersGrid";
import { EnemyWatersGrid } from "./EnemyWatersGrid";
//import { useState } from "react";
import Announcement from "./Announcement";

export const GameStart = () => {
  const {
    state: { game },
  } = useGameContext();
  //const [showAnn, setShowAnn] = useState(false);

  const handleAction = () => {};

  let message = "Welcome to battleship";
  if (game) {
    message = game.message;
  }
  const show = true;

  return (
    <GameStartContainer>
      <HudWindow>{message}</HudWindow>
      <LabelContainer row="4">
        <h1 style={{ margin: "auto auto 0" }}>Friendly waters</h1>
      </LabelContainer>
      <LabelContainer row="2">
        <h1 style={{ margin: "auto auto 0" }}>Enemy waters</h1>
      </LabelContainer>
      <FriendlyWatersGrid />
      <EnemyWatersGrid />
      {show && (
        <Announcement
          text="Nice hit!"
          _fn={handleAction}
          duration={5000}
          stayOnScreen={false}
        />
      )}
    </GameStartContainer>
  );
};
