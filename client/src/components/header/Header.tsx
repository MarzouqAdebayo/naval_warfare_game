import { useGameContext } from "../../GameController";
import { HeaderComponent, Logo } from "../styled_components/headerStyles";
import logo from "../../assets/images/bs_logo.png";
import { Timeline } from "../../types";

export default function Header() {
  const { timeline } = useGameContext().state;

  return (
    <HeaderComponent>
      <Logo large={timeline === Timeline.Init} src={logo} />
    </HeaderComponent>
  );
}
