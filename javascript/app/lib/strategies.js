'use strict';

let readline = require('readline');
let GameState = require('../game/gamestate');
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
    while (true) {
      let base = Math.floor(Math.random() * 100);
      let selection = base % 3;

      if (selection == 0 && gs.A > 0) {
	return new GameState(gs.A - 1, gs.B, gs.C);
      };
      if (selection == 1 && gs.B > 0) {
	return new GameState(gs.A, gs.B - 1, gs.C);
      };
      if (selection == 2 && gs.C > 0) {
	return new GameState(gs.A, gs.B, gs.C - 1);
      };

      continue
    };
  },

  getNextMoveFromUser: function(gs) {
    let rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout
    });

    rl.question('Current game state: ' + gs + '. Indicate your move A, B, C: ', (move) => {
      // parse values from move
      // checks for validity
      // return new gs
    });
  },
};
