import { useGameContext } from "../../GameController";
import { Timeline } from "../../types";
import { MainWindow } from "../styled_components/gameControllerStyles";
import { GameInit } from "./GameInit";
import { GameMenu } from "./GameMenu";
import { GameSetup } from "./GameSetup";
import { GameStart } from "./GameStart";

const TimelineComponents = {
  [Timeline.Init]: GameInit,
  [Timeline.Menu]: GameMenu,
  [Timeline.Setup]: GameSetup,
  [Timeline.GameStart]: GameStart,
};

export default function GameWindow() {
  const {
    state: { timeline },
  } = useGameContext();
  const Child = TimelineComponents[timeline];
  return (
    <MainWindow>
      <Child />
    </MainWindow>
  );
}
