import styled from "styled-components";

const StyledApp = styled.div`
  display: flex;
  position: relative;
  width: 100%;
  flex-direction: column;
  background: rgb(6, 34, 51);
  background: linear-gradient(
    90deg,
    rgba(6, 34, 51, 1) 0%,
    rgba(14, 57, 78, 1) 29%,
    rgba(18, 64, 87, 1) 76%,
    rgba(8, 45, 66, 1) 100%
  );
  overflow: auto;
`;

const GameWindowContainer = styled.div`
  display: flex;
  position: relative;
  width: 100%;
  margin: 1rem auto 0;
  flex: 1;
`;

export { StyledApp, GameWindowContainer };
