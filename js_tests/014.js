/*
    015: Demonstrates definitions as usages in arithmetic.
*/

var x = 50;         // 0

let z = x * 60;     // 1

const y = 80 * z;   // 2

x = x * y;          // 3

log(y);
log(x);