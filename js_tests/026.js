/*
    026: Demonstrates assignments expiring correcly in if statements.
*/
var y = 200; // 0

var x = 20;  // 1

if(y < 10) {
    let x = 40; // 2
    if(y < 5) {
       x = 50; // 3
       z = 90; // 4
    }
    log(x);
}

log(x);
log(z);