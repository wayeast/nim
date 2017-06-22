"use strict"

var expect = require('chai').expect;
let GameState = require('../../app/gamestate');
let gamestatesEqual = require('../../app/lib/gamestateFns').gamestatesEqual;
let gamestatesEqual2_3 = require('../../app/lib/gamestateFns').gamestatesEqual2_3;

describe('GameState functions:', function() {
  describe('gamestatesEqual', function() {
    it('returns true when A, B, and C equal', function() {
      let gs1 = new GameState(1, 2, 3);
      let gs2 = new GameState(1, 2, 3);

      expect(gamestatesEqual(gs1, gs2)).to.equal(true);
    });

    it('returns false when one of A, B, or C not equal', function() {
      let gs1 = new GameState(0, 2, 3);
      let gs2 = new GameState(1, 2, 3);

      expect(gamestatesEqual(gs1, gs2)).to.equal(false);
    });
  });

  describe('gamestatesEqual2_3', function() {
    it('returns true when A, B, and C equal', function() {
      let gs1 = new GameState(1, 2, 3);
      let gs2 = new GameState(1, 2, 3);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(true);
    });

    it('returns true when A and B equal', function() {
      let gs1 = new GameState(1, 2, 0);
      let gs2 = new GameState(1, 2, 3);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(true);
    });

    it('returns true when B and C equal', function() {
      let gs1 = new GameState(0, 2, 3);
      let gs2 = new GameState(1, 2, 3);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(true);
    });

    it('returns true when A and C equal', function() {
      let gs1 = new GameState(1, 0, 3);
      let gs2 = new GameState(1, 2, 3);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(true);
    });

    it('returns false when only A equal', function() {
      let gs1 = new GameState(1, 2, 3);
      let gs2 = new GameState(1, 3, 2);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(false);
    });

    it('returns false when only B equal', function() {
      let gs1 = new GameState(1, 2, 3);
      let gs2 = new GameState(3, 2, 1);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(false);
    });

    it('returns false when only C equal', function() {
      let gs1 = new GameState(1, 2, 3);
      let gs2 = new GameState(2, 1, 3);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(false);
    });

    it('returns false when none are equal', function() {
      let gs1 = new GameState(1, 2, 3);
      let gs2 = new GameState(3, 1, 2);

      expect(gamestatesEqual2_3(gs1, gs2)).to.equal(false);
    });
  });
});
