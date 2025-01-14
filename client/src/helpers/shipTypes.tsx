import Carrier from "../components/icons/CarrierIcon";
import Battleship from "../components/icons/BattleshipIcon";
import Destroyer from "../components/icons/DestroyerIcon";
import Submarine from "../components/icons/SubmarineIcon";
import Patrol from "../components/icons/PatrolIcon";
import { ShipIconProps } from "../types";

export const shipTypes = {
  carrier: {
    name: "carrier",
    length: 5,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Carrier
          key={"carrier"}
          start={props.start}
          axis={props.axis}
          ship_length={5}
          sunk={props.sunk}
        />
      );
    },
  },
  battleship: {
    name: "battleship",
    length: 4,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Battleship
          key={"battleship"}
          start={props.start}
          axis={props.axis}
          ship_length={4}
          sunk={props.sunk}
        />
      );
    },
  },
  destroyer: {
    name: "destroyer",
    length: 3,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Destroyer
          key={"destroyer"}
          start={props.start}
          axis={props.axis}
          ship_length={3}
          sunk={props.sunk}
        />
      );
    },
  },
  submarine: {
    name: "submarine",
    length: 3,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Submarine
          key={"submarine"}
          start={props.start}
          axis={props.axis}
          ship_length={3}
          sunk={props.sunk}
        />
      );
    },
  },
  patrol_boat: {
    name: "patrol boat",
    length: 2,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Patrol
          key={"patrol-boat"}
          start={props.start}
          axis={props.axis}
          ship_length={2}
          sunk={props.sunk}
        />
      );
    },
  },
};

export default shipTypes;
