require('bootstrap-webpack');
require('../sass/main.sass');
require('../img/favicon.png');
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
