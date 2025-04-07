# Javascribe
Javascribe is a Javascript Data Flow Analysis (DFA) library written in Go. This describes the usage-defintion relationships in javascript code by implementing Reaching Defintion Analysis (RDA), hence the name.

Javascribe depends on the library [go-fAST](https://github.com/t14raptor/go-fast/)

## Usage
Javascribe requires that you parse the file with gofast and pass in the root node to it:
```go
        yourJsCode := "..."
		
        a, err := parser.ParseFile(yourJsCode)
		if err != nil {
			panic(err)
		}

		rdaCtx := dfa.CreateContextRDA(256)

		rdaCtx.Start(a)


		for idx, ud := range rdaCtx.UseDefs {
			// Use-Def chains here
		}
```

## Testing
Javascribe utilizes the power of Golangs "testing" module to test its modules against a variety of JS code and compare the output to precomputed expected output from the V8 JS engine. These tests are found in the `js_tests` directory

Each test will contain two items:
1. Javascript file containing code that is going to be analysed.
2. JSON file describing the correct analysis of the code.

For the javascript file there are a few practices all tests must adhere to:
1. Tests should adhere to the naming system as defined in the section below.
2. A brief description of the test in a comment at the top.
3. Each declaration/assignment must be labeled with an incrementing counter starting from 0.
4. Code must be concise and testing 1 thing.
5. Identifiers should use concise names when possible.

**Example**:
```js
/*
    This test demonstrates definitions as usages.
*/

var x = 50;         // 0

let z = x * 60;     // 1

const y = 80 * z;   // 2

x = x * y;          // 3

log(y);
log(x);
```

Tests are named numerically where the lowest digit is the test number and the upper digits are the type of thing being tested. The upper numbers are referred to as the "test group". Below is the list of test groups:

01. Variables and Arithmetic
02. If Statements
03. For Loops
04. For Each Loops
05. While Loops
06. Try Catch Statements
07. Arrays and Objects
08. Functions and Function Calls
09. Empty Blocks
10. Switch Statements
11. Function Literals

**Examples**:
- `./js_test/11.js`: Variables and Arithmetic Test #2
- `./js_test/45.js`: While Loops Test #6
- `./js_test/103.js`: Switch Statement Test #3

## Todo
- For Loops (In Progress)
- For Each Loops
- While Loops
- Try Catch Statements
- Arrays and Objects
- Functions, Function Literals, and Function Calls
- Empty Blocks
- Switch Statements