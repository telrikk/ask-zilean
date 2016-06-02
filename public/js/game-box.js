import React from 'react';
import RecentGamesList from 'recent-games-list';

export default class GameBox extends React.Component {
  render() {
    if (this.props.recentGames.length === 0) {
      return null;
    }
    return (
      <div id="game-box">
        <RecentGamesList recentGames={this.props.recentGames} />
      </div>
    );
  }
}

GameBox.propTypes = {
  recentGames: React.PropTypes.array,
};
