package logs

import (
    "fmt"
)

func Info (msg string) {
    fmt.Println(">> " + msg);
}

func Error (err error) {
    if err == nil {
        return
    }

    panic(err)
}
