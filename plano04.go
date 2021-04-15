package main

import (
        f "fmt"
        m "math"
)

func init() {
        f.Println("Primeiro init")
    }

func main() {
        f.Println("Raiz de 64:")
        f.Println(m.Sqrt(64))
}