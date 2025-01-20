import { useEffect, useState } from "react";
import styled, { keyframes } from "styled-components";

// Animations
const fadeIn = keyframes`
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
`;

const fadeOut = keyframes`
  from {
    opacity: 1;
    transform: scale(1);
  }
  to {
    opacity: 0;
    transform: scale(0.8);
  }
`;

// Styled Components
const Overlay = styled.div<{ $visible: boolean }>`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  opacity: ${(props) => (props.$visible ? 1 : 0)};
  pointer-events: ${(props) => (props.$visible ? "auto" : "none")};
  animation: ${(props) => (props.$visible ? fadeIn : fadeOut)} 0.5s ease;
`;

const AnnouncementBox = styled.div`
  background: rgba(14, 57, 78, 0.9);
  color: #b8d8e8;
  padding: 2rem 3rem;
  border-radius: 8px;
  text-align: center;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  transform: scale(1.05);
  transition: transform 0.2s;

  &:hover {
    transform: scale(1.1);
  }
`;

const AnnouncementText = styled.p`
  font-size: 2rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
`;

const ActionButton = styled.button`
  background: #00ff00;
  color: #003300;
  font-size: 1rem;
  font-weight: bold;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  box-shadow: 0 4px 8px rgba(0, 255, 0, 0.4);
  transition:
    background 0.2s,
    transform 0.2s;

  &:hover {
    background: #00e600;
    transform: scale(1.05);
  }

  &:active {
    background: #00cc00;
  }
`;

const Announcement = ({
  text,
  actionButtonText,
  _fn,
  duration = 5000,
  stayOnScreen = false,
}: {
  text: string;
  actionButtonText: string;
  _fn: () => void;
  duration?: number;
  stayOnScreen?: boolean;
}) => {
  const [visible, setVisible] = useState(true);

  useEffect(() => {
    if (!stayOnScreen) {
      const timer = setTimeout(() => setVisible(false), duration);
      return () => clearTimeout(timer); // Clean up timer on unmount
    }
  }, [duration, stayOnScreen]);

  if (!visible) return null;

  return (
    <Overlay $visible={visible}>
      <AnnouncementBox>
        <AnnouncementText>{text}</AnnouncementText>
        <ActionButton onClick={_fn}>{actionButtonText}</ActionButton>
      </AnnouncementBox>
    </Overlay>
  );
};

export default Announcement;
