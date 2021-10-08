package main


import (
"crypto/sha256"
"fmt"
"assignment01IBC"
)


func main() {


var a string = "ez"
b := "bz"
fmt.Printf("%x\n", sha256.Sum256([]byte(a)))
fmt.Printf("%x\n", sha256.Sum256([]byte(b)))


}