package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, `<h1>Hello, World! <span style="color:#55D7E5">ʕ◔ϖ◔ʔ</span></h1>`)
    })

    log.Println("Listening on http://localhost")
    log.Println(http.ListenAndServe(":80", nil))
}
