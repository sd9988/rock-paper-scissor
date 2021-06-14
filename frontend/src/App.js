import React, { useState } from "react";
import Header from "./components/Header";
import Play from "./components/Play";
import Game from "./components/Game";
import Footer from "./components/Footer";
import { Switch, Route } from "react-router-dom";

function App() {
  const [myChoice, setMyChoice] = useState({});
  const [score, setScore] = useState(0);
  const [game, setGame] = useState(0);

  return (
    <>
      <div className="container">
        <Header score={score} game={game} />
        <Switch>
          <Route exact path="/">
            <Play setMyChoice={setMyChoice} />
          </Route>
          <Route path="/game">
            <Game myChoice={myChoice} score={score} setScore={setScore} game={game} setGame={setGame} />
          </Route>
        </Switch>
      </div>
      <Footer />
    </>
  );
}

export default App;
