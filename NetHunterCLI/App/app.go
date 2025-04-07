package app

import (
    "fmt"
    "log"
    "net"

    "github.com/urfave/cli"
)

// * Gerar cria e retorna a aplicação CLI configurada
func Gerar() *cli.App {
    app := cli.NewApp()
    app.Name = "Aplicação CLI"
    app.Usage = "Busca IPs e nomes de servidores na internet"

    flags := []cli.Flag{
        cli.StringFlag{
            Name:  "host",
            Value: "devbooh.com.br",
        },
    }

    app.Commands = []cli.Command{
        {
            Name:   "ip",
            Usage:  "Busca IPs de endereços na internet",
            Flags:  flags,
            Action: buscarIps,
        },
        {
            Name:   "servidores",
            Usage:  "Busca o nome dos servidores na internet",
            Flags:  flags,
            Action: buscarServidores,
        },
    }

    return app
}

func buscarIps(c *cli.Context) {
    //  ! Implementar a lógica para buscar o IP do host
    // ! Exemplo: usar a biblioteca net para resolver o nome do host
    host := c.String("host")

    ips, erro := net.LookupIP(host)
    if erro != nil {
        log.Fatal(erro)
    }

    for _, ip := range ips {
        fmt.Println(ip)
    }
}

func buscarServidores(c *cli.Context) {
    host := c.String("host")

    servidores, erro := net.LookupNS(host) // * Busca os servidores de nomes (DNS) do host
    if erro != nil {
        log.Fatal(erro)
    }
    for _, servidor := range servidores {
        fmt.Println(servidor.Host)
    }
}