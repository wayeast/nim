"use strict"

class GameState {
  constructor(a, b, c) {
    this._a = a;
    this._b = b;
    this._c = c;
  }

  get A() { return this._a; }
  set A(newA) {
    if (newA < 0 || newA > this._a) {
      return;
    }
    this._a = newA;
  }

  get B() { return this._b; }
  set B(newB) {
    if (newB < 0 || newB > this._b) {
      return;
    }
    this._b = newB;
  }

  get C() { return this._c; }
  set C(newC) {
    if (newC < 0 || newC > this._c) {
      return;
    }
    this._c = newC;
  }
}

module.exports = GameState;
