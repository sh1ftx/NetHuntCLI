package main

import (
    "log"
    "os"

    app "CLI/Project/App"
)

func main() {
    aplicacao := app.Gerar()
    if erro := aplicacao.Run(os.Args); erro != nil {
        log.Fatal(erro)
    }

}

// By Sh1ftx