package main

import (
    "os"
    "log"
    "net/http"
    "fmt"

    "github.com/codegangsta/cli"

    "github.com/LaynePeng/RestServer/server"
)

func main() {
    var config string
    var port int

    app := cli.NewApp()
    app.Name = "My Server"
    app.Usage = "A simple REST server!"

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "config",
            Value: "./server.conf",
            Usage: "Config file for the server",
            Destination: &config,
        },
    }

    app.Commands = []cli.Command{
        {
            Name:      "start",
            Aliases:     []string{"s"},
            Usage:     "start the web server",
            Action: func(c *cli.Context) {
                portStr := fmt.Sprintf(":%d", port)
                println(fmt.Sprintf("Visit the web: http://127.0.0.1%s", portStr))
                http.HandleFunc("/",server.Index)
                log.Fatal(http.ListenAndServe(portStr, nil))
            },

            Flags: []cli.Flag{
                cli.IntFlag{
                    Name: "port",
                    Value: 8080,
                    Usage: "Listening port",
                    Destination: &port,
                },
            },
        },
    }

    app.Action = func(c *cli.Context) {
        println("Hello friend!")
    }

    app.Run(os.Args)
}

