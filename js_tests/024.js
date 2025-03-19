/*
    024: Demonstrates a usage of a block scoped variable outside of the if block.
*/
var x = 40;         // 0

if(x == 50) {
    let y = 200;    // 1
    log(y);
}

log(y);