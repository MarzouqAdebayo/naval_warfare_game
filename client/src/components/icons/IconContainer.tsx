import styled from "styled-components";

export default styled.div<{
  $x: number;
  $y: number;
  $axis: "X" | "Y";
  $length: number;
}>`
  display: flex;
  border: 1px solid #ddd;
  height: 100%;
  grid-row: ${({ $x, $axis, $length }) => {
    return $axis === "Y" ? `${$x + 1} / span ${$length}` : `${$x + 1} / span 1`;
  }};
  grid-column: ${({ $y, $axis, $length }) => {
    return $axis === "X" ? `${$y + 1} / span ${$length}` : `${$y + 1} / span 1`;
  }};
`;
