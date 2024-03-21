package main


import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var port = ":5500"

type Status struct {
	Status U `json:"users"`
}

type U struct {
	Wind int `json:"wind"`
	Water int `json:"water"`
	WindStatus string `json:"wind_status"`
	WaterStatus string `json:"water_status"`
}

func main (){
	http.HandleFunc("/", random)
	http.ListenAndServe(port, nil)
}