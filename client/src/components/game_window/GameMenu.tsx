import { useGameContext } from "../../GameController";
import {
  ConnectionIndicator,
  MenuOption,
  MenuOptionsContainer,
  MenuTitle,
  MenuWindow,
} from "../styled_components/gameControllerStyles";
import { Timeline, WSEvents } from "../../types";

export const GameMenu = () => {
  const {
    dispatch,
    isConnected,
    sendMessage,
    state: { game },
  } = useGameContext();

  const handleNewGame = () => {
    if (!game) {
      sendMessage({ type: WSEvents.EventFindGame, payload: null });
      dispatch({ type: "CHANGE_TIMELINE", payload: Timeline.Setup });
    }
  };

  return (
    <MenuWindow>
      <MenuTitle>Menu</MenuTitle>
      <MenuOptionsContainer>
        <MenuOption onClick={handleNewGame}>New Game</MenuOption>
      </MenuOptionsContainer>
      <ConnectionIndicator isConnected={isConnected} />
    </MenuWindow>
  );
};
