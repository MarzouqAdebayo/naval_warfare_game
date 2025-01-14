import { useState } from "react";

export const Mypage = () => {
  const [theme, setTheme] = useState("dark");
  return (
    <div>
      <VeryLargeList />
      <div>Very large list of 500 items</div>
      <div
        onClick={() => setTheme((prev) => (prev === "dark" ? "light" : "dark"))}
      >
        {theme}
      </div>
    </div>
  );
};

const VeryLargeList = () => {
  return <div>Very large list of 1000 items</div>;
};
