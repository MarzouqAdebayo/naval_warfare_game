import styled from "styled-components";

const HeaderComponent = styled.div`
  margin: 1rem auto 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  overflow: hidden;
  width: 100%;
  background: rgba(14, 57, 78, 0.3); // Subtle theme background
  padding: 0.5rem 1rem; // Slight padding for spacing
  border-bottom: 1px solid rgba(255, 255, 255, 0.2); // Subtle separator
`;

const Logo = styled.img<{ $large: boolean }>`
  margin: auto;
  height: ${(props) => (props.$large ? "auto" : "3rem")};
  @media (max-width: 1000px) {
    height: ${(props) => (props.$large ? "3rem" : "2rem")};
  }
  animation: fadeinlogo 1.5s ease-in-out forwards;
  filter: drop-shadow(0 0 4px rgba(184, 216, 232, 0.4)); // Subtle glow effect
`;

const Status = styled.div<{ $connected: boolean }>`
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #b8d8e8;
  font-size: 0.9rem;
  font-weight: 600;

  .indicator {
    width: 0.8rem;
    height: 0.8rem;
    border-radius: 50%;
    background-color: ${(props) => (props.$connected ? "#00ff00" : "#ff0000")};
    box-shadow: 0 0 6px ${(props) => (props.$connected ? "#00ff00" : "#ff0000")};
  }
`;

const InfoSection = styled.div`
  display: flex;
  justify-content: space-around;
  align-items: center;
  padding: 1rem;
  background: rgba(14, 57, 78, 0.4); // Slightly darker theme
  color: #b8d8e8;
  font-size: 1rem;

  div {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 1rem;
  }

  span {
    font-size: 1.2rem;
    font-weight: bold;
    color: #fff; // Emphasize the data
  }

  @media (max-width: 768px) {
    flex-direction: row;
    gap: 1rem;
  }
`;

export { HeaderComponent, Logo, Status, InfoSection };
