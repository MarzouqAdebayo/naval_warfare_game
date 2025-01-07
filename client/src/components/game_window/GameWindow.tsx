import { Timeline, useGameContext } from "../../GameController";
import { MainWindow } from "../styled_components/gameControllerStyles";
import { GameInit } from "./GameInit";
import { GameSetup } from "./GameSetup";
import { GameStart } from "./GameStart";

const renderChild = (timeline: Timeline) => {
  switch (timeline) {
    case Timeline.Init:
      return GameInit({ hey: "Hello World" });
    case Timeline.Setup:
      return GameSetup();
    case Timeline.GameStart:
      return GameStart();
    default:
      return null;
  }
};

export default function GameWindow() {
  const {
    state: { timeline },
  } = useGameContext();
  return <MainWindow>{renderChild(timeline)}</MainWindow>;
}
