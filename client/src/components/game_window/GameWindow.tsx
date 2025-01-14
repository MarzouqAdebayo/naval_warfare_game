import { useGameContext } from "../../GameController";
import { Timeline } from "../../types";
import { MainWindow } from "../styled_components/gameControllerStyles";
import { GameInit } from "./GameInit";
import { GameMenu } from "./GameMenu";
import { GameSetup } from "./GameSetup";
import { GameStart } from "./GameStart";

const renderChild = (timeline: Timeline) => {
  switch (timeline) {
    case Timeline.Init:
      return GameInit({ hey: "Hello World" });
    case Timeline.Menu:
      return GameMenu();
    case Timeline.Setup:
      return GameSetup();
    case Timeline.GameStart:
      return GameStart();
    default:
      throw new Error(`timeline ${timeline} does not exist`);
  }
};

export default function GameWindow() {
  const {
    state: { timeline },
  } = useGameContext();
  return <MainWindow>{renderChild(timeline)}</MainWindow>;
}
