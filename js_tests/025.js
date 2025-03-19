/*
    025: Demonstrates nested if statements
*/
var y = 200; // 0

var z = 30;  // 1
var x = 20;  // 2

if(y > 10) {
    x = 40; // 3
    if(y > 20) {
        z = 50; // 4
        if(y > 30) {
            x = 60; // 5
            if(y > 40) {
                z = 70; // 6
                if(y > 50) {
                    x = 80; // 7
                    if(y > 60) {
                        z = 90; // 8
                        if(y > 70) {
                            x = 100; // 9
                        }
                    }
                }
            }
        }
    }
}

log(z);
log(x);