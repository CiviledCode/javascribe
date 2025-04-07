/*
    031: Demonstrates variables expiring out of scope properly in for loops.
*/

var x = 20;     // 0

for(let i = 0;  // 1 
    i < 10; 
    i++         // 2
) {
    x = 50;     // 3
    z = 30;     // 4
    let y = 20; // 5
}

log(x);
log(y);
log(z);

