package main

import (
        "net/http"
        "fmt"
)

func main() {
        dir := http.Dir("./files")
        fmt.Println("Server spawning on localhost:8080")
        http.ListenAndServe(":8080", http.FileServer(dir))

}
