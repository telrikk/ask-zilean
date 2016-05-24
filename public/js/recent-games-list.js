import React from 'react';

export default class RecentGamesList extends React.Component {
  render() {
    const gameNodes = this.props.recentGames.map(game =>
    (
      <div className="panel panel-default">
        <h4 className="panel-title">
          <a data-toggle="collapse" href={`#${game.id}`}>
            {game.mapName}
          </a>
        </h4>
        <div id={game.id} className="panel-collapse collapse">
          {game.id}
        </div>
      </div>
    )
    );
    return (
      <div className="recent-games-list">
        {gameNodes}
      </div>
    );
  }
}

RecentGamesList.propTypes = {
  recentGames: React.PropTypes.array,
};
