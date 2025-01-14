import { useState } from "react";
import { useGameContext } from "../../GameController";
import {
  ConnectionIndicator,
  LoaderContainer,
  LoaderWrapper,
  LoaderCancelButton,
  MenuOption,
  MenuOptionsContainer,
  MenuTitle,
  MenuWindow,
} from "../styled_components/gameControllerStyles";

export const GameMenu = () => {
  const {
    isConnected,
    sendMessage,
    state: { game },
  } = useGameContext();
  const [loading, setIsLoading] = useState(true);

  const handleNewGame = () => {
    if (!game) {
      setIsLoading(true);
      sendMessage({ type: "find_game", payload: null });
    }
  };

  const handleCancel = () => {
    setIsLoading(false);
    //sendMessage({ type: "cancel_find_game", payload: null });
  };

  return (
    <MenuWindow>
      <MenuTitle>Menu</MenuTitle>
      <MenuOptionsContainer>
        <MenuOption onClick={handleNewGame}>New Game</MenuOption>
      </MenuOptionsContainer>
      <LoaderWrapper show={loading}>
        <LoaderContainer>
          <div>Loading...</div>
          <LoaderCancelButton onClick={handleCancel}>Cancel</LoaderCancelButton>
        </LoaderContainer>
      </LoaderWrapper>
      <ConnectionIndicator isConnected={isConnected} />
    </MenuWindow>
  );
};
