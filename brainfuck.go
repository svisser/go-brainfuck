package main

import "flag"
import "fmt"
import "io/ioutil"

func main() {
    pathPtr := flag.String("path", "", "path to Brainfuck source")

    flag.Parse()

    path := *pathPtr
    if path == "" {
        fmt.Println("You must specify a path to Brainfuck source: -path=...")
    }

    b, err := ioutil.ReadFile(path)
    if err != nil {
       panic(err)
    }

    var data [30000]int
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
                var i int
                d, err := fmt.Scanf("%d", &i)
                if err != nil {
                    panic(err)
                }
                data[dp] = d
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
