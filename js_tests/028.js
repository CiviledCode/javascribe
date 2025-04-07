/*
    028: Demonstrates if, else if, else statement.

*/

var y = 50; // 0

var x = 20; // 1

var z = 30; // 2


if(y < 100) {
    x = 10; // 3
    z = 20; // 4    
} else if(y < 200) {
    x = 20; // 5
} else {
    x = 30; // 6
    z = 10; // 7
}

log(x);
log(z);