import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { play } from "../services/game.service";

const Game = ({ score, myChoice, setScore, setGame, game }) => {
  const [gameResult, setGameResult] = useState({});

  const { id, name } = myChoice

  const [counter, setCounter] = useState(3);

  useEffect(() => {
    async function startGame() {
      const result = await play(id)
      setGameResult(result.data)
    }
    startGame()
  }, []);

  const Result = () => {
    setGame(game + 1)
    if (gameResult.result === "win") {
      setScore(score + 1)
    }
    if (gameResult.result === "lose") {
      setScore(score - 1)
    }
  };

  useEffect(() => {
    const timer =
      counter > 0
        ? setInterval(() => {
          setCounter(counter - 1);
        }, 1000)
        : Result();

    return () => {
      clearInterval(timer);
    };
  }, [counter]);

  return (
    <div className="game">
      <div className="game__you">
        <span className="text">You Picked</span>
        {counter == 0 ? (
          <div
            className={`icon icon--${name} ${gameResult.result == "win" ? `icon icon--${name}--winner` : ""
              }`}
          ></div>) : (<div className={`icon icon--${name}`}></div>)}
      </div>
      {counter == 0 && gameResult.result == "win" && (
        <div className="game__play">
          <span className="text">You Win</span>
          <Link to="/" className="play-again">
            Play Again
          </Link>
        </div>
      )}
      {counter == 0 && gameResult.result == "lose" && (
        <div className="game__play">
          <span className="text">You Lose</span>
          <Link to="/" className="play-again">
            Play Again
          </Link>
        </div>
      )}
      {counter == 0 && gameResult.result == "tie" && (
        <div className="game__play">
          <span className="text">Tie</span>
          <Link to="/" className="play-again">
            Play Again
          </Link>
        </div>
      )}

      <div className="game__house">
        <span className="text">The House Picked</span>
        {counter == 0 ? (
          <div
            className={`icon icon--${gameResult.computerChoice} ${gameResult.result == "lose" ? `icon icon--${gameResult.computerChoice}--winner` : ""
              }`}
          ></div>
        ) : (
          <div className="counter">{counter}</div>
        )}
      </div>
    </div>
  );
};

export default Game;
