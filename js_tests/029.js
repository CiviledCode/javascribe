/*
    029: Demonstrates nested statements and various declaration types in an if-elif-else statement.
*/

var w = 30;                 // 0
var x = 40;                 // 1
var y = 50;                 // 2
var z = 60;                 // 3

if(y < 100) {
    x = 60;                 // 4
} else if(y < 200) {
    x = 70;                 // 5
} else {
    x = 80;                 // 6
    let z = 300;            // 7
    if(y > 400) {
        z = 900;            // 8
        
        if(y > 1000) {
            let x = 100;    // 9
            x = 500;        // 10
            log(x);
        } else {
            x = 1000;       // 11
            w = 1000;       // 12
        }
    }

    log(z);
}

log(z);
log(x);
log(w);