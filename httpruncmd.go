package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

var cmdFlag = flag.String("cmd", "", "Command to exec")
var pathFlag = flag.String("path", "", "The http://localhost:port/path that will spawn the command")
var portFlag = flag.Int("port", 8000, "HTTP server port")

type ExecStates struct {
	running   bool
	startTime time.Time
}

var states = ExecStates{}

func runHandler(w http.ResponseWriter, r *http.Request) {
	// If the cmd is not finihsed yet return KO and the start date of the last run
	if states.running {
		fmt.Fprintf(w, "KO ")
		fmt.Fprintf(w, states.startTime.Format(time.RFC3339))
		return
	}
	fmt.Fprintf(w, "OK")

	states.running = true
	states.startTime = time.Now()
	out, err := exec.Command(*cmdFlag).Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", out)
	states.running = false
}

func main() {
	flag.Parse()

	if *cmdFlag == "" || *pathFlag == "" {
		log.Fatal("Invalid parameters")
	}

	http.HandleFunc("/"+*pathFlag+"/", runHandler)
	http.ListenAndServe(":"+strconv.Itoa(*portFlag), nil)
}
