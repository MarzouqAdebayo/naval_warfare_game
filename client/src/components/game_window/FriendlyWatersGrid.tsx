import { Timeline, useGameContext } from "../../GameController";
import {
  Cell,
  GameBoardGrid,
  SetupGridContainer,
  WatersContainer,
} from "../styled_components/gameControllerStyles";

const fillCells = (timeline: Timeline) => {
  const arr = [];
  for (let i = 0; i < 100; i++) {
    arr.push([i]);
  }
  return Array.from({ length: 100 }).map((_, index) => {
    return (
      <Cell
        key={index}
        position={""}
        highlight={false}
        timeline={timeline}
        board="friendly"
        shot={true}
        cursor={""}
      />
    );
  });
};

export const FriendlyWatersGrid = () => {
  const {
    state: { timeline },
  } = useGameContext();
  return (
    <WatersContainer row="5">
      <SetupGridContainer>
        <GameBoardGrid>{fillCells(timeline)}</GameBoardGrid>
      </SetupGridContainer>
    </WatersContainer>
  );
};
