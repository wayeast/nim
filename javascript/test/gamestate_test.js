"use strict"

var expect = require('chai').expect;
let GameState = require('../app/gamestate');

describe('GameState class', function() {
  describe('restricts values of A, B, and C', function() {
    let gamestate;

    beforeEach(() => {
      gamestate = new GameState(1, 2, 3);
    });

    it('to non-negative numbers', function() {
      gamestate.A = -1;
      gamestate.B = -1;
      gamestate.C = -1;

      expect(gamestate.A).to.equal(1);
      expect(gamestate.B).to.equal(2);
      expect(gamestate.C).to.equal(3);
    });

    it('to numbers <= current values', function() {
      gamestate.A = 5;
      gamestate.B = 5;
      gamestate.C = 5;

      expect(gamestate.A).to.equal(1);
      expect(gamestate.B).to.equal(2);
      expect(gamestate.C).to.equal(3);
    });
  });
});
