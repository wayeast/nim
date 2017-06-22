'use strict';

let GameState = require('../gamestate');
let WinningStates = require('./winningStates');
let gamestatesEqual2_3 = require('./gamestateFns').gamestatesEqual2_3;

module.exports = {
  playToWin: function(gs) {
    for (var ws of WinningStates) {
      if (gamestatesEqual2_3(gs, ws) && ws.magnitude() < gs.magnitude()) {
	return ws;
      };
    };

    return null;
  },

  justPlaySomething: function(gs) {
    // pick a random number 1<=n<=3
    // case 1: if A > 0 => A--
    // case 2: if B > 0 ...
  },
};
