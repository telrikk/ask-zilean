require('bootstrap-webpack');
require('../sass/main.sass');
require('../img/favicon.png');
require('../img/score.png');
require('../img/items.png');
require('../img/minion.png');
require('../img/gold.png');
import React from 'react';
import ReactDOM from 'react-dom';
import Header from 'header';
import GameBox from 'game-box';
import SearchBox from 'search-box';

class Application extends React.Component {
  constructor() {
    super();
    this.state = { recentGames: [] };
    this.onRecentGamesLoad = this.onRecentGamesLoad.bind(this);
  }
  onRecentGamesLoad(data) {
    this.setState({ recentGames: data.results });
  }
  render() {
    return (
      <div>
        <Header />
        <SearchBox onRecentGamesLoad={this.onRecentGamesLoad} />
        <GameBox recentGames={this.state.recentGames} />
      </div>
    );
  }
}


ReactDOM.render(
  <Application />,
  document.getElementById('content')
);
