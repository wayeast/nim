'use strict';

let GameState = require('../gamestate');
let permutation = require('array-permutation');

let winningPositions = [
  [1, 2, 3],
  [1, 4, 5],
  [1, 6, 7],
  [1, 8, 9],
  [2, 4, 6],
  [2, 5, 7],
  [3, 4, 7],
  [3, 5, 6],
  [4, 8, 12],
  [4, 9, 13],
  [5, 8, 13],
  [5, 9, 12],
];

let WinningStates = [
  new GameState(0, 0, 0),
];

for (var winning_position of winningPositions) {
  for (var perm of permutation(winning_position)) {
    WinningStates.push(new GameState(perm[0], perm[1], perm[2]));
  };
};

for (let i = 1; i <= 15; i++) {
  for (var perm of permutation([0, i, i])) {
    WinningStates.push(new GameState(perm[0], perm[1], perm[2]));
  };
};

module.exports = WinningStates;
