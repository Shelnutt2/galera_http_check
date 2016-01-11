package main

import (
	"fmt"
	"os"
	"strconv"

	goopt "github.com/droundy/goopt"
	_ "github.com/go-sql-driver/mysql"
)

var port = goopt.String([]string{"-p", "--port"}, "80", "Port to bind server too")

var galeraHost = goopt.String([]string{"-H", "--host"}, "localhost", "Host to check galera status of")
var galeraPort = goopt.Int([]string{"--mysql_port"}, 3306, "Specify a port to connect to")
var mysqlUser = os.Getenv("MYSQL_USERNAME")
var mysqlPassword = os.Getenv("MYSQL_PASSWORD")

func init() {
	//Parse options
	goopt.Parse(nil)

	// Setup goopts
	goopt.Description = func() string {
		return "Galera http Check"
	}
	goopt.Version = "0.9.0"
	goopt.Summary = "main [-H] [-p]"

}

func main() {
	fmt.Println("Starting server to monitor " + *galeraHost + ":" + strconv.Itoa(*galeraPort) + " on port " + *port)
	runServer(*port)
}
