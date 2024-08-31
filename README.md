# Monkey Interpreter

An interpreter written in Go for the made-up Monkey programming language, from the book ["Writing An Interpreter In Go"](https://interpreterbook.com/).

## The Monkey programming language

Monkey has the following features:

- Variable bindings.
- Integers and Booleans.
- Arithmetic expressions.
- Built-in functions.
- First-class and higher-order functions.
- Closures
- Strings.
- Arrays.
- Hash data structure.

Here is what Monkey code looks like:

```js
// Binding values

let age = 1;
let name = "Monkey";
let result = 10 * (20/2);

// Arrays and hashes

let myArray = [1, 2, 3, 4, 5];
let myHash = {"name": "Tomas", "year": 2023};

myArray[0] // => 1
myHash["name"] // => "Tomas"


// Functions

let add = fn(a, b) { a + b; };
add(1, 2);

let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

// Higher-order functions

let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2); // => 6
```

## Interpreter parts

The interpreter tokenizes and parses Monkey source code in a REPL(Read-Eval-Print-Loop), builds up an internal representation of the code called abstract syntax tree (AST) and evaluates this tree. It features these major parts:

- The lexer: transforms source code into tokens.
- The parser: takes the tokens and constructs the Abstract Syntax Tree (AST).
- The Abstract Syntax Tree (AST): represents the source code, its order and operation priority.
- The internal object system: the representation of the objects that Monkey supports (integer, boolean, null, return value, error, function, string, built-in function, array, hash, quote function, and macro).
- The evaluator: evaluates the code.
