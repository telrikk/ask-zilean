import React from 'react';
import SummonerNameInput from 'summoner-name-input';
import RecentGamesList from 'recent-games-list';

export default class GameBox extends React.Component {
  constructor() {
    super();
    this.state = { recentGames: [], isSearchComplete: false };
    this.onRecentGamesLoad = this.onRecentGamesLoad.bind(this);
  }
  onRecentGamesLoad(data) {
    this.setState({ recentGames: data.results, searchComplete: true });
  }
  render() {
    let className = '';
    if (this.state.isSearchComplete) {
      className = 'hidden';
    }
    return (
      <div id="game-box">
        <RecentGamesList recentGames={this.state.recentGames} />
        // FIXME: this relationship is a bit awkward and leads to styling issues
        <SummonerNameInput className={className} onSuccess={this.onRecentGamesLoad} />
      </div>
    );
  }
}
