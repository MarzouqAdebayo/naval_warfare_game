import styled from "styled-components";
import { ShipIconProps } from "../../types";

export default styled.div<ShipIconProps>`
  display: flex;
  border: 1px solid #ddd;
  height: 100%;
  grid-row: ${({ x, axis, length: ship_length }) => {
    return axis === "Y"
      ? `${x + 1} / span ${ship_length}`
      : `${x + 1} / span 1`;
  }};
  grid-column: ${({ y, axis, length: ship_length }) => {
    return axis === "X"
      ? `${y + 1} / span ${ship_length}`
      : `${y + 1} / span 1`;
  }};
`;
