package main

import (
    "fmt"
    "reflect"
)

func main() {

    s := []string{"a"}
    modifySlice(s)
    fmt.Println(s)

    modifySliceRef(&s)
    fmt.Println(s)

    s = addToSlice(s)
    fmt.Println(s)
}

func modifySlice(s []string) {
    fmt.Println(reflect.TypeOf(s))
    s[0] = "yolo"
    s = append(s, "second")
}

func modifySliceRef(s *[]string) {
    fmt.Println(reflect.TypeOf(s))
    (*s)[0] = "yolo-ref"
    *s = append(*s, "second-ref")
}

func addToSlice(s []string) ([]string) {
    return append(s, "added")
}