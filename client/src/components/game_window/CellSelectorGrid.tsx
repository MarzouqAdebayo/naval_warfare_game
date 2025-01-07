import { useGameContext } from "../../GameController";
import {
  Cell,
  GameBoardGrid,
  SetupGridContainer,
} from "../styled_components/gameControllerStyles";

export const CellSelectorGrid = () => {
  const {
    state: { timeline },
  } = useGameContext();
  return (
    <SetupGridContainer>
      <GameBoardGrid>
        {Array.from({ length: 100 }).map((_, i) => (
          <Cell
            key={i}
            //highlight={hovered.includes(i)}
            //cursor={hovered.includes(index) ? "pointer" : "not-allowed"}
            timeline={timeline}
            onClick={() => {}}
            onMouseEnter={() => {}}
            onMouseLeave={() => {}}
          />
        ))}
      </GameBoardGrid>
    </SetupGridContainer>
  );
};
