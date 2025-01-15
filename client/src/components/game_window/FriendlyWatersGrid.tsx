import { useGameContext } from "../../GameController";
import shipTypes from "../../helpers/shipTypes";
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
        {![CellState.Empty, CellState.Ship].includes(cell) && (
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
  const player = players[index];

  return (
    <WatersContainer row="5">
      <SetupGridContainer>
        <GameBoardGrid>
          {player.fleet.map((ship) =>
            shipTypes[ship.type].getShipWithProps(ship),
          )}
        </GameBoardGrid>
      </SetupGridContainer>
      <SetupGridContainer>
        <GameBoardGrid>{fillCells(timeline, player)}</GameBoardGrid>
      </SetupGridContainer>
    </WatersContainer>
  );
};
