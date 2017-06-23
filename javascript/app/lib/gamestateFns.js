'use strict';

let GameState = require('../game/gamestate');

module.exports = {
  gamestatesEqual: function (gs1, gs2) {
    return gs1.A == gs2.A && gs1.B == gs2.B && gs1.C == gs2.C;
  },

  gamestatesEqual2_3: function(gs1, gs2) {
    if (gs1.A == gs2.A && gs1.B == gs2.B) {
      return true;
    };
    if (gs1.B == gs2.B && gs1.C == gs2.C) {
      return true;
    };
    if (gs1.A == gs2.A && gs1.C == gs2.C) {
      return true;
    };
    return false;
  },
};
