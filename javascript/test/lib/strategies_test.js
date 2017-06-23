"use strict"

var expect = require('chai').expect;
let GameState = require('../../app/game/gamestate');
let WinningStates = require('../../app/lib/winningStates');
let playToWin = require('../../app/lib/strategies').playToWin;
let justPlaySomething = require('../../app/lib/strategies').justPlaySomething;
let gamestatesEqual = require('../../app/lib/gamestateFns').gamestatesEqual;
let gamestatesEqual2_3 = require('../../app/lib/gamestateFns').gamestatesEqual2_3;

describe('Play strategy function:', function() {
  describe('playToWin', function() {
    it('results in (1, 1, 0) from (1, 2, 0)', function() {
      let gamestate = new GameState(1, 2, 0);
      let expected = new GameState(1, 1, 0);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (1, 2, 3) from (1, 2, 4)', function() {
      let gamestate = new GameState(1, 2, 4);
      let expected = new GameState(1, 2, 3);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (1, 2, 3) from (5, 2, 3)', function() {
      let gamestate = new GameState(5, 2, 3);
      let expected = new GameState(1, 2, 3);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (1, 2, 3) from (1, 4, 3)', function() {
      let gamestate = new GameState(1, 4, 3);
      let expected = new GameState(1, 2, 3);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (2, 4, 6) from (2, 4, 8)', function() {
      let gamestate = new GameState(2, 4, 8);
      let expected = new GameState(2, 4, 6);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (2, 4, 6) from (3, 4, 6)', function() {
      let gamestate = new GameState(3, 4, 6);
      let expected = new GameState(2, 4, 6);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (4, 8, 12) from (5, 8, 12)', function() {
      let gamestate = new GameState(5, 8, 12);
      let expected = new GameState(4, 8, 12);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (0, 0, 0) from (0, 1, 0)', function() {
      let gamestate = new GameState(0, 1, 0);
      let expected = new GameState(0, 0, 0);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });

    it('results in (0, 0, 0) from (0, 3, 0)', function() {
      let gamestate = new GameState(0, 3, 0);
      let expected = new GameState(0, 0, 0);

      let winningstate = playToWin(gamestate);
      expect(gamestatesEqual(winningstate, expected)).to.equal(true);
    });
  });

  describe('playToWin', function() {
    for (var ws of WinningStates) {
      it('returns null from ' + ws, function() {
	expect(playToWin(ws)).to.be.null;
      });
    };
  });

  describe('justPlaySomething', function() {
    for (var ws of WinningStates) {
      it('returns something 2/3 equal from ' + ws, function() {
	let play = justPlaySomething(ws);
	expect(gamestatesEqual2_3(ws, play)).to.equal(true);
      });
    };
  });
});
