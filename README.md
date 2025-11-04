# Monkey Interpreter

An interpreter for the made-up Monkey programming language. Typed by hand, following Thorsten Ball's [Writing an Interpreter in Go](https://interpreterbook.com/) guide.

This project is all about implementing an interpreter using only golang's standard lib, and zero dependencies. Everything to implement this interpreter is built from the ground up right here.

### Why write an interpreter?

I have had a long standing interest in what is happening at a lower level when I write javascript function like:

```
function add(x, y) {
  x + y
}

add(1,2)
```
How is this evaluated? How does my IDE know to jump to this `add` function in another file when I click on it? How do computer languages actually _work_ under the hood?

I am a relentlessly curious person. The answers to these questions are _knowable_. That is one of my favorite things about programming, I can go as shallow/deep as I want to.

As I continue to grow as an engineer, I have decided to spend some effort in understanding these lower level, _systems_ questions. I know that after completing this project/book, I won't be able to implement an interpreter by myself totally unassisted. That's not the goal. The goal is to type all this code by hand and absorb as much as I can, which admittadly probably will only be a small fraction of the complexity contained therein.

I hope that even if I only absorb a small fraction of this book, that small fraction will impact how I see code and the systems that run that code.
