import { Timeline } from "../../GameController";
import {
  Cell,
  GameBoardGrid,
  SetupGridContainer,
} from "../styled_components/gameControllerStyles";

export const ShipPlacementGrid = () => {
  return (
    <SetupGridContainer>
      <GameBoardGrid>
        {Array.from({ length: 100 }).map((_) => {
          const props = {
            position: "",
            highlight: true,
            timeline: Timeline.Setup,
            board: "",
            shot: true,
            cursor: "",
          };
          return <Cell {...props} />;
        })}
      </GameBoardGrid>
    </SetupGridContainer>
  );
};
