import { useState } from "react";
import {
  AxisButton,
  Cell,
  GameBoardGrid,
  GridOverlayContainer,
  LoaderCancelButton,
  LoaderContainer,
  LoaderWrapper,
  SetupGridContainer,
  SetupTitle,
  SetupWindow,
} from "../styled_components/gameControllerStyles";
import { useGameContext } from "../../GameController";
import { GameStatus, Timeline, WSEvents } from "../../types";
import shipTypes from "../../helpers/shipTypes";

export const GameSetup = () => {
  const {
    dispatch,
    sendMessage,
    state: { timeline, game, name },
  } = useGameContext();
  const [axis, setAxis] = useState<"X" | "Y">("X");
  const [loading, setIsLoading] = useState(false);

  const handleSetAxis = () => {
    setAxis((prev) => (prev === "X" ? "Y" : "X"));
  };

  const handleCancel = () => {
    setIsLoading(false);
    sendMessage({ type: WSEvents.EventQuitGame, payload: null });
  };

  const handleLeaveRoom = () => {};

  const handleGoToMainMenu = () => {
    dispatch({ type: "CHANGE_TIMELINE", payload: Timeline.Menu });
  };

  const handleGenShips = () => {
    if (!game) return;
    sendMessage({
      type: WSEvents.EventPlaceShip,
      payload: {
        instruction: "randomize",
        roomID: game.roomID,
        playerIndex: game.index,
      },
    });
  };

  const handleReady = () => {
    if (!game) return;
    sendMessage({
      type: WSEvents.EventShipReady,
      payload: {
        roomID: game.roomID,
        playerIndex: game.index,
      },
    });
  };

  if (!game) {
    return (
      <LoaderWrapper show={true}>
        <LoaderContainer>
          <LoaderCancelButton onClick={handleGoToMainMenu}>
            Main menu
          </LoaderCancelButton>
        </LoaderContainer>
      </LoaderWrapper>
    );
  }

  if (game.status === GameStatus.WaitingForOpponent) {
    return (
      <LoaderWrapper show={true}>
        <LoaderContainer>
          <div>Waiting for player to join room...</div>
          <LoaderCancelButton onClick={handleLeaveRoom}>
            Leave room
          </LoaderCancelButton>
        </LoaderContainer>
      </LoaderWrapper>
    );
  }

  const { players } = game;
  const { board, fleet } = players[game.index];

  return (
    <SetupWindow>
      <SetupTitle>{name}, Place Your Ships</SetupTitle>
      <AxisButton onClick={handleSetAxis}>AXIS: {axis}</AxisButton>
      <AxisButton onClick={handleGenShips}>Randomize</AxisButton>
      <AxisButton onClick={handleReady}>Ready</AxisButton>
      <GridOverlayContainer>
        <SetupGridContainer>
          <GameBoardGrid>
            {fleet.map((ship) => shipTypes[ship.type].getShipWithProps(ship))}
          </GameBoardGrid>
        </SetupGridContainer>
        <SetupGridContainer>
          <GameBoardGrid>
            {board.map((row) => {
              return row.map((cell, i) => {
                return (
                  <Cell
                    key={i}
                    position=""
                    board="friendly"
                    highlight={cell === "Sunk"}
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
      {
        <LoaderWrapper show={loading}>
          <LoaderContainer>
            <div>Loading...</div>
            <LoaderCancelButton onClick={handleCancel}>
              Cancel
            </LoaderCancelButton>
          </LoaderContainer>
        </LoaderWrapper>
      }
    </SetupWindow>
  );
};
