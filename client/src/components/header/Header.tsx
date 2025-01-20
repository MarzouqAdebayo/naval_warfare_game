import { useGameContext } from "../../GameController";
import {
  HeaderComponent,
  InfoSection,
  Logo,
  Status,
} from "../styled_components/headerStyles";
import logo from "../../assets/images/bs_logo.png";
import { Timeline } from "../../types";

export default function Header() {
  const {
    state: { timeline },
    isConnected,
    connectionString,
  } = useGameContext();

  return (
    <>
      <HeaderComponent>
        <Logo $large={timeline === Timeline.Init} src={logo} />
        <Status $connected={isConnected}>
          <div className="indicator" />
          {connectionString}
        </Status>
        <InfoSection />
      </HeaderComponent>
      <InfoSection>
        <div>
          <span>{22}</span>
          Ongoing Games
        </div>
        <div>
          <span>{2000}</span>
          Connected Clients
        </div>
      </InfoSection>
    </>
  );
}
