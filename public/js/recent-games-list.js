import React from 'react';
import RecentGame from 'recent-game';

export default function RecentGamesList(props) {
  const gameNodes = props.recentGames.map(game =>
  (
    <RecentGame key={game.id} game={game} />
  )
  );
  if (props.recentGames.length === 0) {
    return null;
  }
  return (
    <div className="recent-games-list">
      <div className="panel panel-default">
        <div className="panel-heading">
          <span className="champion">
            <span className="centered-image-helper" />
            <img
              alt="Champion"
              src={props.recentGames[0].championImageURL}
            />
          </span>
          <span className="queue-description" />
          <span className="score">
            <span className="centered-image-helper" />
            <img
              alt="Score"
              src={props.recentGames[0].statsImageURL}
            />
          </span>
          <span className="minions">
            <span className="centered-image-helper" />
            <img
              alt="Minions"
              src={props.recentGames[0].creepScoreImageURL}
            />
          </span>
          <span className="gold">
            <span className="centered-image-helper" />
            <img
              alt="Gold"
              src={props.recentGames[0].goldImageURL}
            />
          </span>
          <span className="items">
            <span className="centered-image-helper" />
            <img
              alt="Items"
              src={props.recentGames[0].itemsImageURL}
            />
          </span>
        </div>
        <ul className="list-group">
          {gameNodes}
        </ul>
      </div>
    </div>
  );
}

RecentGamesList.propTypes = {
  recentGames: React.PropTypes.array,
};
