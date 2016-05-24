import React from 'react';
import SummonerNameInput from 'summoner-name-input';
import RecentGamesList from 'recent-games-list';
import $ from 'jquery';

export default class GameBox extends React.Component {
  constructor() {
    super();
    this.state = { recentGames: [] };
    this.onRecentGamesLoad = this.onRecentGamesLoad.bind(this);
  }
  onRecentGamesLoad(data) {
    this.setState({ recentGames: data.results });
    $('.summoner-name-input').hide();
  }
  render() {
    return (
      <div className="game-box centered bordered shadowed">
        <RecentGamesList recentGames={this.state.recentGames} />
        <SummonerNameInput onSuccess={this.onRecentGamesLoad} />
      </div>
    );
  }
}
