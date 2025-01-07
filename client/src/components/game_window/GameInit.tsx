import { ChangeEvent, FormEvent, useState } from "react";
import { useGameContext } from "../../GameController";
import {
  InitWindow,
  PlayerForm,
} from "../styled_components/gameControllerStyles";

export const GameInit = ({ hey }: { hey: string }) => {
  const { dispatch } = useGameContext();
  const [name, setName] = useState("");
  const [error, setError] = useState("");

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };

  const handleFocus = () => {};

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!name.trim()) {
      setError("Name required");
      return;
    }
    if (name.length > 20) {
      setError("Name is too long");
      return;
    }
    setError("");
    console.log(name);
  };

  return (
    <InitWindow>
      <PlayerForm onSubmit={handleSubmit}>
        <label htmlFor="name">Enter player name:</label>
        <input
          type="text"
          name="name"
          id="name"
          placeholder="Battleship combatant"
          onChange={handleChange}
          onFocus={handleFocus}
          autoComplete="off"
          value={name}
        />
        <div>{hey}</div>
        <p style={{ color: "red" }}>{error}</p>
        <button type="submit">Start game</button>
      </PlayerForm>
      <div className="connection"></div>
    </InitWindow>
  );
};
