# Brainfuck interpreter in Go

## Overview

Brainfuck interpreter written in Go.

For more information about the Brainfuck programming language, see:

- https://en.wikipedia.org/wiki/Brainfuck
- https://esolangs.org/wiki/Brainfuck

This implementation follows these conventions:

- The data array contains signed integers.
- Newlines are expected as a single character ('\n').

# Running a Brainfuck program

You can run a program as follows:

    go run brainfuck.go -path=samples/hello_world.bf
