import { useGameContext } from "../../GameController";
import {
  MenuOption,
  MenuOptionsContainer,
  MenuTitle,
  MenuWindow,
} from "../styled_components/gameControllerStyles";
import { WSEvents } from "../../types";

export const GameMenu = () => {
  const {
    sendMessage,
    state: { game },
  } = useGameContext();

  const handleNewGame = () => {
    if (!game) {
      sendMessage({ type: WSEvents.EventFindGame, payload: null });
    }
  };

  return (
    <MenuWindow>
      <MenuTitle>Menu</MenuTitle>
      <MenuOptionsContainer>
        <MenuOption onClick={handleNewGame}>New Game</MenuOption>
      </MenuOptionsContainer>
    </MenuWindow>
  );
};
