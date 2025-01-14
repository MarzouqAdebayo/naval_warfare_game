import { useGameContext } from "../../GameController";
import { CellState, Player, Timeline } from "../../types";
import ShotMarker from "../icons/ShotMarker";
import {
  Cell,
  GameBoardGrid,
  SetupGridContainer,
  WatersContainer,
} from "../styled_components/gameControllerStyles";

const fillCells = (timeline: Timeline, player: Player) => {
  return player.board.map((row, i) =>
    row.map((cell, j) => (
      <Cell
        key={`${i}-${j}`}
        position={""}
        highlight={false}
        timeline={timeline}
        board="friendly"
        shot={false}
        cursor={""}
      >
        {cell !== CellState.Empty && (
          <ShotMarker hit={cell === CellState.Hit || cell === CellState.Sunk} />
        )}
      </Cell>
    )),
  );
};

export const FriendlyWatersGrid = () => {
  const {
    state: { timeline, game },
  } = useGameContext();

  if (!game) return null;
  const { players, index } = game;
  const board = players[index];

  return (
    <WatersContainer row="5">
      <SetupGridContainer>
        <GameBoardGrid>{fillCells(timeline, board)}</GameBoardGrid>
      </SetupGridContainer>
    </WatersContainer>
  );
};
