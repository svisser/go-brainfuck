package main

import "bufio"
import "flag"
import "fmt"
import "io/ioutil"
import "os"

func main() {
	pathPtr := flag.String("path", "", "path to Brainfuck source")
	sizePtr := flag.Int("size", 30000, "size of the data array")

	flag.Parse()

	path := *pathPtr
	if path == "" {
		fmt.Println("You must specify a path to Brainfuck source: -path=...")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	var data = make([]int, *sizePtr)
	var b_length = len(b)
	var dp = 0
	for ip := 0; ip < b_length; {
		var c = string(b[ip])
		switch {
		case c == ">":
			dp += 1
			ip += 1
		case c == "<":
			dp -= 1
			ip += 1
		case c == "+":
			data[dp] += 1
			ip += 1
		case c == "-":
			data[dp] -= 1
			ip += 1
		case c == ".":
			fmt.Print(string(data[dp]))
			ip += 1
		case c == ",":
			s, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if len(s) != 2 {
				panic("You must provide a single character")
			}
			data[dp] = int(s[0])
			ip += 1
		case c == "[":
			if data[dp] == 0 {
				var target = 0
				var cnt = 1
				for tip := (ip + 1); tip < b_length; tip++ {
					cc := string(b[tip])
					if cc == "[" {
						cnt += 1
					} else if cc == "]" {
						cnt -= 1
					}
					if cnt == 0 {
						target = tip
						break
					}
				}
				ip = target + 1
			} else {
				ip += 1
			}
		case c == "]":
			if data[dp] != 0 {
				var target = 0
				var cnt = 1
				for tip := (ip - 1); tip > 0; tip-- {
					cc := string(b[tip])
					if cc == "]" {
						cnt += 1
					} else if cc == "[" {
						cnt -= 1
					}
					if cnt == 0 {
						target = tip
						break
					}
				}
				ip = target + 1
			} else {
				ip += 1
			}
		default:
			ip += 1
		}
	}
}
