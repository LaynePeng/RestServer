package main

import (
  "os"
  "log"
  "net/http"

  "github.com/codegangsta/cli"

  "github.com/LaynePeng/RestServer/server"
)

func main() {
  app := cli.NewApp()
  app.Name = "My Server"
  app.Usage = "A simple REST server!"

  app.Commands = []cli.Command{
  	{
   		Name:      "start",
    	Aliases:     []string{"s"},
    	Usage:     "start the web server",
    	Action: func(c *cli.Context) {
    		println("Visit the web: http://127.0.0.1:8080")
    		http.HandleFunc("/",server.Index)
    		log.Fatal(http.ListenAndServe(":8080", nil))
    	},
   	},
  }

  app.Action = func(c *cli.Context) {
    println("Hello friend!")
  }

  app.Run(os.Args)
}

