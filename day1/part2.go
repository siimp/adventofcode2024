package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main()  {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Failed to open file: %v", err)
    }
    defer file.Close()

    r := bufio.NewReader(file)
    var list1 []int
    var list2 []int
    for {
        line, _, err := r.ReadLine()
        if err != nil {
            break
        }

        numbers := strings.Split(string(line), "   ")
        n1, _ := strconv.Atoi(numbers[0])
        list1 = append(list1, n1)

        n2, _ := strconv.Atoi(numbers[1])
        list2 = append(list2, n2)
    }

    m := make(map[int]int)
    for _, v := range list2 {
        m[v]++
    }

    sum := 0
    for _, v := range list1 {
        sum += v * m[v]
    }
    fmt.Println(sum)

}