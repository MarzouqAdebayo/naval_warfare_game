import { ChangeEvent, FormEvent, useState } from "react";
import { useGameContext } from "../../GameController";
import {
  InitWindow,
  PlayerForm,
} from "../styled_components/gameControllerStyles";
import { Timeline, WSEvents } from "../../types";

export const GameInit = () => {
  const {
    state: { name },
    dispatch,
    sendMessage,
  } = useGameContext();
  const [error, setError] = useState("");

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    dispatch({ type: "SET_PLAYER_NAME", payload: e.target.value });
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
    sendMessage({ type: WSEvents.EventSetUserData, payload: { name } });
    dispatch({ type: "CHANGE_TIMELINE", payload: Timeline.Menu });
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
        <p style={{ color: "red" }}>{error}</p>
        <button type="submit">Set Name</button>
      </PlayerForm>
    </InitWindow>
  );
};
