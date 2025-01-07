import {
  AxisButton,
  GridOverlayContainer,
  SetupTitle,
  SetupWindow,
} from "../styled_components/gameControllerStyles";
import { ShipPlacementGrid } from "./ShipPlacementGrid";

export const GameSetup = () => {
  return (
    <SetupWindow>
      <SetupTitle>Your Name</SetupTitle>
      <AxisButton>AXIS: X</AxisButton>
      <GridOverlayContainer>
        <ShipPlacementGrid></ShipPlacementGrid>
      </GridOverlayContainer>
    </SetupWindow>
  );
};
