import { useState } from "react";
import {
  AxisButton,
  Cell,
  GameBoardGrid,
  GridOverlayContainer,
  SetupGridContainer,
  SetupTitle,
  SetupWindow,
} from "../styled_components/gameControllerStyles";
import { useGameContext } from "../../GameController";

const enum shipTypes {
  Carrier = "Carrier",
  Battleship = "Battleship",
  Cruiser = "Cruiser",
  Submarine = "Submarine",
  Destroyer = "Destroyer",
}

const ships = [
  { Type: shipTypes.Carrier, Size: 5, Hits: 0, Sunk: false },
  { Type: shipTypes.Battleship, Size: 4, Hits: 0, Sunk: false },
  { Type: shipTypes.Cruiser, Size: 3, Hits: 0, Sunk: false },
  { Type: shipTypes.Submarine, Size: 3, Hits: 0, Sunk: false },
  { Type: shipTypes.Destroyer, Size: 2, Hits: 0, Sunk: false },
];

export const GameSetup = () => {
  const {
    //data,
    sendMessage,
    state: { timeline, game },
  } = useGameContext();
  const [axis, setAxis] = useState<"X" | "Y">("X");
  //const [shipPosition, setShipPosition] = useState([]);

  const handleGenShips = () => {
    sendMessage({ type: "find_game", payload: { name: axis } });
  };

  const handleSetAxis = () => {
    setAxis((prev) => (prev === "X" ? "Y" : "X"));
  };

  if (!game) return null;

  const { board } = game;

  return (
    <SetupWindow>
      <SetupTitle>, Place Your Ships</SetupTitle>
      <AxisButton onClick={handleSetAxis}>AXIS: {axis}</AxisButton>
      <AxisButton onClick={handleGenShips}>Randomize</AxisButton>
      <GridOverlayContainer>
        <SetupGridContainer>
          <GameBoardGrid>
            {board.map((row) => {
              return row.map((cell, i) => {
                return (
                  <Cell
                    key={i}
                    position=""
                    board="friendly"
                    highlight={cell === "Ship"}
                    cursor={"pointer"}
                    timeline={timeline}
                    shot={true}
                  />
                );
              });
            })}
          </GameBoardGrid>
        </SetupGridContainer>
      </GridOverlayContainer>
    </SetupWindow>
  );
};
