"use strict"

var expect = require('chai').expect;
let GameState = require('../../app/game/gamestate');
let WinningStates = require('../../app/lib/winningStates');

describe('WinningStates', function() {
  it('has length 163', function() {
    expect(WinningStates.length).to.equal(163);
  });
});
