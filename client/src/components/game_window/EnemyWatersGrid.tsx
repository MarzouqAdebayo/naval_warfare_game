import { useGameContext } from "../../GameController";
import shipTypes from "../../helpers/shipTypes";
import { CellState, Player, Timeline, WSEvents } from "../../types";
import ShotMarker from "../icons/ShotMarker";
import {
  Cell,
  GameBoardGrid,
  SetupGridContainer,
  WatersContainer,
} from "../styled_components/gameControllerStyles";

const fillCells = (
  timeline: Timeline,
  player: Player,
  _fn: (x: number, y: number) => void,
) => {
  return player.board.map((row, i) =>
    row.map((cell, j) => (
      <Cell
        key={`${i}-${j}`}
        $position={""}
        $highlight={cell === CellState.Ship}
        $timeline={timeline}
        $board="enemy"
        $shot={cell !== CellState.Empty}
        $cursor={cell === CellState.Empty ? "crosshair" : "not-allowed"}
        onClick={() => _fn(i, j)}
      >
        {cell !== CellState.Empty && (
          <ShotMarker hit={cell === CellState.Hit || cell === CellState.Sunk} />
        )}
      </Cell>
    )),
  );
};

export const EnemyWatersGrid = () => {
  const {
    sendMessage,
    state: { timeline, game },
  } = useGameContext();

  if (!game) return null;
  const { players, index } = game;
  const player = players[1 - index];

  const handleCellClick = (x: number, y: number) => {
    if (game.index !== game.currentTurn) return;
    const payload = {
      roomID: game.roomID,
      attackerIndex: game.index,
      attackPosition: { X: x, Y: y },
    };
    sendMessage({ type: WSEvents.EventAttack, payload });
  };

  return (
    <WatersContainer $row="3">
      <SetupGridContainer>
        <GameBoardGrid>
          {player.fleet.map((ship) =>
            shipTypes[ship.type].getShipWithProps(ship),
          )}
        </GameBoardGrid>
      </SetupGridContainer>
      <SetupGridContainer>
        <GameBoardGrid>
          {fillCells(timeline, player, handleCellClick)}
        </GameBoardGrid>
      </SetupGridContainer>
    </WatersContainer>
  );
};
