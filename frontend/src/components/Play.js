import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Pentagon from "../images/bg-pentagon.svg";
import { choices } from "../services/game.service";


const Play = ({ setMyChoice }) => {
  const [allChoices, setAllChoices] = useState([])
  const setChoice = (e) => {
    const id = allChoices.filter(i => i.name === e.target.dataset.id)[0].id
    setMyChoice({
      name: e.target.dataset.id,
      id
    })
  }

  useEffect(() => {
    async function getChoices() {
      const options = await choices()
      setAllChoices(options.data)
    }
    getChoices()
  }, [])

  return (
    <div className="play">
      <img src={Pentagon} alt="" className="triangle" />
      <Link to="/game">
        <div className="items">
          {allChoices.map(i => {
            return (
              <div
                key={i.id}
                data-id={`${i.name}`}
                onClick={setChoice}
                className={`icon icon--${i.name}`}
              ></div>
            )

          })}
        </div>
      </Link>
    </div>
  );
}

export default Play;
