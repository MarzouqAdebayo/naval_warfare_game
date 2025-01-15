import Carrier from "../components/icons/CarrierIcon";
import Battleship from "../components/icons/BattleshipIcon";
import Destroyer from "../components/icons/DestroyerIcon";
import Submarine from "../components/icons/SubmarineIcon";
import Cruiser from "../components/icons/CruiserIcon";
import { ShipIconProps } from "../types";

export type ShipTypeKeys = keyof typeof shipTypes;

export const shipTypes = {
  carrier: {
    name: "carrier",
    length: 5,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Carrier
          key={"carrier"}
          type={props.type}
          x={props.x}
          y={props.y}
          axis={props.axis}
          length={5}
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
          type={props.type}
          key={"battleship"}
          x={props.x}
          y={props.y}
          axis={props.axis}
          length={4}
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
          type={props.type}
          x={props.x}
          y={props.y}
          axis={props.axis}
          length={3}
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
          type={props.type}
          x={props.x}
          y={props.y}
          axis={props.axis}
          length={3}
          sunk={props.sunk}
        />
      );
    },
  },
  cruiser: {
    name: "cruiser",
    length: 2,
    getShipWithProps: (props: ShipIconProps) => {
      return (
        <Cruiser
          key={"cruiser"}
          type={props.type}
          x={props.x}
          y={props.y}
          axis={props.axis}
          length={2}
          sunk={props.sunk}
        />
      );
    },
  },
};

export default shipTypes;
