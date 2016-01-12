package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var checkQuery = "show global status where variable_name='wsrep_local_state'"

func checkGaleraStatus(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", mysqlUser+":"+mysqlPassword+"@tcp("+*galeraHost+":"+strconv.Itoa(*galeraPort)+")/")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error connecting to galera host (%s)\n%s", *galeraHost, err.Error()) // send data to client side
		return
	}
	var value int
	var variableName string
	err = db.QueryRow(checkQuery).Scan(&variableName, &value)
	db.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error checking galera host (%s) status\n%s", *galeraHost, err.Error()) // send data to client side
		return
	}

	if value == 4 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Galera host (%s) is running fine", *galeraHost) // send data to client side
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Galera host (%s) is down with status: %d", *galeraHost, value) // send data to client side
	//fmt.Fprintf(w, "%+v", rows.Columns())
	return
}
