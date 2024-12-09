package main

import (
    "fmt"
    "slices"
    "log"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main()  {
/*
    list1 := []int{3, 4, 2, 1, 3, 3}
    slices.Sort(list1)
    list2 := []int{4, 3, 5, 3, 3, 9}
    slices.Sort(list2)
*/

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

    slices.Sort(list1)
    slices.Sort(list2)

    result := 0
    for i, _ := range list1 {
        if (list1[i] > list2[i]) {
            result += list1[i] - list2[i]
        } else {
            result += list2[i] - list1[i]
        }
    }

    fmt.Println(result)
}