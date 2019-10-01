package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/***************************
*        Handler
***************************/
func hitMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Opsss!!!")
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	array := sort1bilionTimes()
	fmt.Fprintf(w, "%v", array)
}

func readFileHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	n := params.Get("n")

	data, err := ioutil.ReadFile("../data/" + n + ".txt")

	if err != nil {
		fmt.Fprintf(w, "read failed %v: ", err)
		return
	}

	fmt.Fprintf(w, "read done: %v", len(data))
}

func main() {
	httpPort := 8081

	http.HandleFunc("/hitme", hitMeHandler)
	http.HandleFunc("/readfile", readFileHandler)
	http.HandleFunc("/sort", sortHandler)

	fmt.Printf("listening on %v\n", httpPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

/*********************************************
*               Helper Functions
**********************************************/
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("Go: %v:%v:%v - %v - %v\n", now.Hour(), now.Minute(), now.Second(), r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func sort1bilionTimes() []int16 {

	array := []int16{3, 4, 1, 3, 5, 1, 92, 2, 4124, 424, 52, 12}

	for c := 0; c < 1000000; c++ {

		for i := 0; i < len(array); i++ {
			for y := 0; y < len(array)-1; y++ {
				if array[y+1] < array[y] {
					t := array[y]
					array[y] = array[y+1]
					array[y+1] = t
				}
			}
		}
	}

	return array
}
