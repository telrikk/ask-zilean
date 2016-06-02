import React from 'react';
import SummonerNameInput from 'summoner-name-input';

export default class SearchBox extends React.Component {
  constructor() {
    super();
    this.state = { isSearchComplete: false };
    this.onRecentGamesLoad = this.onRecentGamesLoad.bind(this);
  }
  onRecentGamesLoad(data) {
    this.setState({ isSearchComplete: true });
    this.props.onRecentGamesLoad(data);
  }
  render() {
    if (this.state.isSearchComplete) {
      return null;
    }
    return (
      <div id="search-box">
        <SummonerNameInput onSuccess={this.onRecentGamesLoad} />
      </div>
    );
  }
}

SearchBox.propTypes = {
  onRecentGamesLoad: React.PropTypes.func,
};
