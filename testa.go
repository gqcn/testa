package testa

import (
    "fmt"
    "github.com/gogf/gf/g"
)

func TestA() {
    err := g.DB().PingMaster()
    if err != nil {
        panic(err)
    }
    fmt.Println("testa ping master ok")
}
