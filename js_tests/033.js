/*
    033: Demonstrates for loop with no conditional.
*/

var x = 10;       // 0
for(
    var i = 0;    // 1
    ;
    i+=50         // 2
) {
    x = 15;       // 3
    break;
}

log(i);
log(x);