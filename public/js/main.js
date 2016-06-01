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

const Application = function render() {
  return (
    <div>
      <Header />
      <GameBox />
    </div>
  );
};

ReactDOM.render(
  <Application />,
  document.getElementById('content')
);
