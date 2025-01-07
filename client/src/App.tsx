import "./App.css";
import Footer from "./components/footer/Footer";
import GameWindow from "./components/game_window/GameWindow";
import Header from "./components/header/Header";
import {
  GameWindowContainer,
  StyledApp,
} from "./components/styled_components/appStyles";
import GameProvider from "./GameController";

function App() {
  return (
    <GameProvider>
      <StyledApp>
        <Header />
        <GameWindowContainer>
          <GameWindow />
        </GameWindowContainer>
        <Footer />
      </StyledApp>
    </GameProvider>
  );
}

export default App;
