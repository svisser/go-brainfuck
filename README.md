# Brainfuck interpreter in Go

## Overview

Brainfuck interpreter written in Go. Written and tested with Go 1.4.

For more information about the Brainfuck programming language, see:

- https://en.wikipedia.org/wiki/Brainfuck
- https://esolangs.org/wiki/Brainfuck

This implementation follows these conventions:

- The data array contains signed 64-bit integers.
- Newlines are expected as a single character ('\n').

## Running a Brainfuck program

You can run a program as follows:

    go run brainfuck.go -path=samples/hello_world.bf

The program's parameters are:

    -path="": path to Brainfuck source
    -size=30000: size of the data array
