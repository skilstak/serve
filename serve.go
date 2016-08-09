package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

const USAGE = `usage:
  serve
  serve www
  serve 8080
  serve www 8080
  serve 8080 www
`

const SKILSTAK = `
                 Hello there from ...
        __   .__.__            __          __          
  _____|  | _|__|  |   _______/  |______  |  | __      
 /  ___/  |/ /  |  |  /  ___/\   __\__  \ |  |/ /      
 \___ \|    <|  |  |__\___ \  |  |  / __ \|    <       
/____  >__|_ \__|____/____  > |__| (____  /__|_ \______
     \/     \/            \/            \/     \/_____/
               Coding Arts
`

var ldir string
var port string = ":8080"
var dir string = "www"

// Prints usage to stderr
func usage() {
	fmt.Fprint(os.Stderr, USAGE)
}

// Writes text directly to the browser
func skilstakHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, SKILSTAK)
}

// Sends files to the browser
func serve(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path
	http.ServeFile(w, r, dir+"/"+file)
	log.Print(r.Method, " ", file)

}

func exists(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	} else {
		return false
	}
}

func parseArgs() {
	argc := len(os.Args)
	switch {
	case argc == 1:
		return
	case argc == 2:
		a := os.Args[1]
		if a == "help" || a == "-h" || a == "/h" {
			usage()
			os.Exit(0)
		}
		n, _ := strconv.Atoi(a)
		if n > 0 {
			port = ":" + a
		} else {
			dir = a
		}
	case argc == 3:
		a := os.Args[1]
		b := os.Args[2]
		n, _ := strconv.Atoi(a)
		if n > 0 {
			port = ":" + a
			dir = b
		} else {
			port = ":" + b
			dir = a
		}
	}
}

func portFromUid() {
	u, err := user.Current()
	if err != nil {
		return
	}
	uid := u.Uid
	port = ":8" + uid[len(uid)-3:]
}

func main() {
	wd, _ := os.Getwd()
	portFromUid()
	parseArgs()
	ldir = filepath.Join(wd, dir)
	if !exists(ldir) {
		log.Printf("%s not found, serving current dir", ldir, port)
	}

	// first match wins
	http.HandleFunc("/skilstak", skilstakHello)
	http.HandleFunc("/", serve)

	log.Printf("SERVE %s on %s", ldir, port)

	// loops infinitely listening and serving until interrupted
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
