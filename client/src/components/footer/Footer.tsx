import { FooterComponent, FooterText } from "../styled_components/footerStyles";
import FooterLinks from "./FooterLinks";

export default function Footer() {
  return (
    <FooterComponent>
      <FooterLinks />
      <FooterText>
        Created by Marzouq Adebayo and Daniel Oyekunle as part of the{" "}
        <a href="https://www.alxafrica.com">ALX Africa</a> Software Engineering
        track.
      </FooterText>
    </FooterComponent>
  );
}
