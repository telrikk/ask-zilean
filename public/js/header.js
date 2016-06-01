import React from 'react';

export default class Header extends React.Component {
  constructor() {
    super();
    this.state = { games: 'unknown', patch: 'unknown' };
  }

  render() {
    return (
      <div className="header row">
        <span>Games Analyzed: </span>{this.state.games}
        <span>Patch: </span>{this.state.patch}
      </div>
    );
  }
}
