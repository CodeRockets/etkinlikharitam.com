package main

import (
	/*	"database/sql"*/
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"io/ioutil"
	"net/http"
	/*	"os"*/)

type Handler struct {
	DB *sqlx.DB
}

func (h *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	dat, err := ioutil.ReadFile("sql_sel_events.sql")
	throwIf(err)
	events := []Event{}
	h.DB.Select(&events, string(dat), string(params["startdate"]), string(params["enddate"]))

	b, err := json.Marshal(events)
	throwIf(err)
	fmt.Fprintf(w, string(b))
}
func throwIf(err error) {
	if err != nil {
		panic(err)
		fmt.Errorf("Error: %s" + err.Error())
	}
}

const (
	CONFIG_PATH = "config.cfg"
)

var (
	DbConStr = config.String("database.conn_str", "")
)

func main() {

	config.Parse(CONFIG_PATH)

	r := mux.NewRouter()
	db, err := sqlx.Connect("postgres", *DbConStr)
	throwIf(err)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	h := &Handler{db}

	r.HandleFunc("/events/date/start/{startdate}/end/{enddate}", h.GetEvents).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)
	//http.ListenAndServe(":3000", nil)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
