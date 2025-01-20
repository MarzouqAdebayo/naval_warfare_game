import styled from "styled-components";
import { Timeline } from "../../types";

const MainWindow = styled.div`
  display: flex;
  width: 100%;
  justify-content: center;
  text-align: center;
`;

const InitWindow = styled.div`
  display: flex;
  height: 100%;
  width: 100%;
  text-align: center;
  position: relative;
  padding: 1rem 4rem;
`;

const buttonStyles = `
  background-color: rgba(18, 64, 87, 1);
  color: #b8d8e8;
  border: 1px solid #b8d8e8;
  transition: 0.3s;
  &:hover {
    background-color: rgba(22, 78, 106, 1);
  }
`;

const PlayerForm = styled.form`
  color: #b8d8e8;
  display: flex;
  height: 100%;
  width: 100%;
  max-width: 640px;
  font-size: 2rem;
  flex-direction: column;
  margin: 3rem auto;
  padding: 2rem 0;
  & > * {
    margin-top: 1rem;
  }
  & > button {
    margin: 1rem auto;
    padding: 0.8rem;
    cursor: pointer;
    border-radius: 16px;
    background-color: rgba(18, 64, 87, 1);
    color: #b8d8e8;
    border: 1px solid #b8d8e8;
    &:hover {
      background-color: rgba(22, 78, 106, 1);
    }
  }
  & > input {
    padding: 1rem;
    border-radius: 16px;
    background-color: rgba(6, 34, 51, 0.7);
    color: #b8d8e8;
    border: 1px solid #b8d8e8;
    &::placeholder {
      color: rgba(184, 216, 232, 0.7);
    }
  }
`;

const MenuWindow = styled.div`
  display: flex;
  height: 100%;
  width: 100%;
  text-align: center;
  flex-direction: column;
  animation: fadeinslow 2s ease-in;
  margin: auto;
  position: relative;
`;

const MenuTitle = styled.h1`
  font-size: 3rem;
  @media (max-width: 900px) {
    font-size: 2rem;
  }
`;

const MenuOptionsContainer = styled.div``;

const MenuOption = styled.button`
  ${buttonStyles}
  padding: 0.5rem;
  margin: auto;
  font-size: 1.5rem;
  @media (max-width: 900px) {
    font-size: 1rem;
    padding: 0.5rem 0.8rem;
  }
`;

const LoaderWrapper = styled.div<{ show: boolean }>`
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  z-index: ${(props) => (props.show ? "45" : "-45")};
  background-color: rgba(6, 34, 51, 0.95);
  font-size: 2rem;
  color: #b8d8e8;
`;

const LoaderContainer = styled.div`
  position: relative;
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
`;

const LoaderCancelButton = styled.button`
  ${buttonStyles}
  padding: 0.5rem;
  font-size: 1.5rem;
  cursor: pointer;
  @media (max-width: 900px) {
    font-size: 1rem;
    padding: 0.5rem 0.8rem;
  }
`;

const SetupWindow = styled.div`
  display: flex;
  height: 100%;
  width: 100%;
  text-align: center;
  flex-direction: column;
  animation: fadeinslow 2s ease-in;
  margin: auto;
  position: relative;
`;

const SetupTitle = styled.h1`
  font-size: 3rem;
  @media (max-width: 900px) {
    font-size: 2rem;
  }
`;

const AxisButton = styled.button`
  ${buttonStyles}
  padding: 0.5rem;
  margin: auto;
  font-size: 1.5rem;
  @media (max-width: 900px) {
    font-size: 1rem;
    padding: 0.5rem 0.8rem;
  }
`;

const GridOverlayContainer = styled.div`
  width: 100%;
  height: 36rem;
  position: relative;
  margin: 1rem auto 0;
  @media (max-width: 900px) {
    height: 22rem;
  }
`;

const SetupGridContainer = styled.div`
  position: absolute;
  display: flex;
  justify-content: center;
  left: 0;
  right: 0;
`;

const GameBoardGrid = styled.div`
  display: grid;
  position: relative;
  margin: 0 auto;
  grid-template: repeat(10, 3rem) / repeat(10, 3rem);
  text-align: center;
  gap: 2px;
  @media (max-width: 1050px) {
    grid-template: repeat(10, 2rem) / repeat(10, 2rem);
  }
`;

const Cell = styled.div<{
  position: string;
  highlight: boolean;
  timeline: Timeline;
  board: string;
  shot: boolean;
  cursor: string;
}>`
  border: 1px solid #b8d8e8;
  height: 100%;
  width: 100%;
  transition: 0.3s;
  position: ${(props) => props.position};
  background-color: ${(props) =>
    props.highlight ? "rgba(184, 216, 232, 0.3)" : ""};
  &:hover {
    background-color: ${(props) =>
      props.timeline === Timeline.GameStart && props.board === "friendly"
        ? "transparent"
        : props.board === "enemy" && !props.shot
          ? "rgba(22, 78, 106, 0.6)"
          : props.shot
            ? "rgba(255, 60, 60, 0.6)"
            : props.highlight
              ? ""
              : "rgba(255, 60, 60, 0.6)"};
    cursor: ${(props) => props.cursor};
  }
`;

const GameStartContainer = styled.div`
  position: relative;
  display: grid;
  grid-template-rows: 4rem auto 32rem;
  grid-template-columns: 1fr 1fr;
  margin: 2% auto 4rem;
  width: 100%;
  max-width: 1200px;
  animation: fadein 2s;
  @media (max-width: 1050px) {
    grid-template-rows: auto auto 22rem;
  }
  @media (max-width: 750px) {
    grid-template-columns: 1fr;
    grid-template-rows: 6rem auto 22rem auto 22rem;
  }
`;

const WatersContainer = styled.div<{ row: string }>`
  height: 100%;
  width: 100%;
  position: relative;
  display: flex;
  @media (max-width: 750px) {
    grid-row: ${(props) => props.row};
  }
`;

const HudWindow = styled.div`
  display: flex;
  margin: auto;
  height: 100%;
  text-align: center;
  grid-column: 1 / span 2;
  width: 70%;
  border: 1px solid #b8d8e8;
  border-radius: 1rem;
  background: rgb(6, 34, 51);
  background: linear-gradient(
    90deg,
    rgba(6, 34, 51, 1) 0%,
    rgba(14, 57, 78, 1) 29%,
    rgba(18, 64, 87, 1) 76%,
    rgba(8, 45, 66, 1) 100%
  );
  font-family: "Special Elite", monospace;
  font-size: 1.4rem;
  color: #b8d8e8;
  @media (max-width: 1050px) {
    font-size: 1rem;
    padding: 10px;
  }
  @media (max-width: 750px) {
    grid-column: 1 / span 1;
    grid-row: 1 / span 1;
  }
`;

const VolumeContainer = styled.div<{ timeline: Timeline }>`
  display: flex;
  animation: fadeinslow 5s;
  position: absolute;
  top: ${(props) => (props.timeline === Timeline.Init ? "0" : "-3rem")};
  right: ${(props) => (props.timeline === Timeline.Init ? "" : "3rem")};
  @media (max-width: 450px) {
    right: ${(props) => (props.timeline === Timeline.Init ? "" : "1.5rem")};
  }
`;

const LabelContainer = styled.div<{ row: string }>`
  display: flex;
  width: 100%;
  text-align: center;
  @media (max-width: 750px) {
    grid-row: ${(props) => `${props.row} / span 1`};
  }
`;

export {
  MainWindow,
  InitWindow,
  MenuWindow,
  MenuTitle,
  MenuOptionsContainer,
  MenuOption,
  LoaderWrapper,
  LoaderContainer,
  LoaderCancelButton,
  PlayerForm,
  SetupWindow,
  SetupTitle,
  GameBoardGrid,
  SetupGridContainer,
  AxisButton,
  Cell,
  GridOverlayContainer,
  GameStartContainer,
  HudWindow,
  LabelContainer,
  WatersContainer,
  VolumeContainer,
};
