import {
  GameStartContainer,
  HudWindow,
  LabelContainer,
} from "../styled_components/gameControllerStyles";
import { useGameContext } from "../../GameController";
import { FriendlyWatersGrid } from "./FriendlyWatersGrid";
import { EnemyWatersGrid } from "./EnemyWatersGrid";
import Announcement from "./Announcement";
import { Timeline } from "../../types";

export const GameStart = () => {
  const {
    dispatch,
    state: { game },
  } = useGameContext();

  const handleAction = () => {
    dispatch({
      type: "CHANGE_TIMELINE_AND_RESET_GAME",
      payload: Timeline.Menu,
    });
  };

  let message = "Welcome to battleship";
  if (game) {
    message = game.message;
  }

  return (
    <GameStartContainer>
      <HudWindow>{message}</HudWindow>
      <LabelContainer $row="4">
        <h1 style={{ margin: "auto auto 0" }}>Friendly waters</h1>
      </LabelContainer>
      <LabelContainer $row="2">
        <h1 style={{ margin: "auto auto 0" }}>Enemy waters</h1>
      </LabelContainer>
      <FriendlyWatersGrid />
      <EnemyWatersGrid />
      {game?.gameOver && (
        <Announcement
          text={`Game Over, you ${game.currentTurn === game.index ? "win" : "lose"}!!!`}
          actionButtonText="Go to main menu"
          _fn={handleAction}
          duration={10000}
          stayOnScreen={true}
        />
      )}
    </GameStartContainer>
  );
};
