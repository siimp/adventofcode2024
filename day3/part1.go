package main

import (
"fmt"
"os"
"bufio"
"strings"
"strconv"
)

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    result := 0

    scanner := bufio.NewScanner(bufio.NewReader(file))
    for scanner.Scan() {
        line := scanner.Text()
        mulInstructions := strings.Split(line, "mul(")
        for _, mulInstruction := range mulInstructions {
            fmt.Printf("inst: %s \n", mulInstruction)

            mulArguments := strings.Split(mulInstruction, ",")


            if len(mulArguments) < 2 {
                fmt.Println("Less than two arguments")
                continue
            }

            firstArgument, err := strconv.Atoi(mulArguments[0])
            if err != nil {
                fmt.Println("Bad first argument")
                continue
            }

            secondArgumentString := mulArguments[1]
            secondArgumentEnd := strings.Index(secondArgumentString, ")")
            if secondArgumentEnd == -1 {
                fmt.Println("Bad second argument")
                continue
            }


            secondArgument, err := strconv.Atoi(secondArgumentString[:secondArgumentEnd])
            if err != nil {
                continue
            }

            fmt.Printf("%d %d \n", firstArgument, secondArgument)
            result += firstArgument * secondArgument
        }
    }

    fmt.Printf("Result: %d\n", result)

}