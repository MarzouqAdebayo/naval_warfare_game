import * as React from "react";
import styled from "styled-components";

interface ShotMarkerProps {
  hit: boolean;
  className?: string;
}

const ShotCell = styled.div`
  display: flex;
  height: 100%;
  width: 100%;
  & > * {
    margin: auto;
  }
`;

const ShotMarker: React.FC<ShotMarkerProps> = (props) => {
  const { hit, ...restProps } = props;
  return (
    <ShotCell>
      <svg
        width={16}
        height={16}
        fill={hit ? "red" : "white"}
        xmlns="http://www.w3.org/2000/svg"
        {...restProps}
      >
        <circle cx={8} cy={8} r={8} />
      </svg>
    </ShotCell>
  );
};

export default ShotMarker;
