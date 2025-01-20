import styled from "styled-components";

const FooterComponent = styled.div`
  display: flex;
  position: relative;
  font-size: 1rem;
  line-height: 1.5rem;
  bottom: 0;
  left: 0;
  right: 0;
  padding: auto;
  font-family: "Big Shoulders Text", cursive;
  background: rgb(14, 57, 78);
  background: linear-gradient(
    90deg,
    rgba(14, 57, 78, 1) 0%,
    rgba(22, 78, 106, 1) 29%,
    rgba(18, 64, 87, 1) 76%,
    rgba(14, 57, 78, 1) 100%
  );
  height: 4rem;
  justify-content: center;
  text-align: center;
  & > :first-child {
    margin-left: auto;
    padding-right: 1rem;
  }
  & > :last-child {
    margin-right: auto;
    padding-left: 1rem;
  }
`;

const FooterText = styled.p`
  color: #b8d8e8;
  margin: auto;
  & > a {
    transition: 0.5s;
  }
  & > a:link,
  & > a:active,
  & > a:visited {
    color: #b8d8e8;
  }
  & > a:hover {
    color: #ffffff;
    transition: 0.5s;
  }
`;

const FooterLinksDiv = styled.div`
  display: flex;
  & > * {
    margin: auto;
  }
`;

export { FooterComponent, FooterText, FooterLinksDiv };
