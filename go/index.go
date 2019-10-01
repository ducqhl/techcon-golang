package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	//https://planetmath.org/listoffibonaccinumbers
	params := r.URL.Query()
	n := params.Get("n")
	nInt, _ := strconv.ParseInt(n, 0, 64)
	result := fibonacci(nInt)

	fmt.Fprintf(w, "%v", result)
}

func primethHandler(w http.ResponseWriter, r *http.Request) {
	//https://primes.utm.edu/lists/small/1000.txt
	params := r.URL.Query()
	n := params.Get("n")
	nInt, _ := strconv.ParseInt(n, 0, 64)
	result := getNthPrime(nInt)

	fmt.Fprintf(w, "%d", result)
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

	http.HandleFunc("/fibonacci", fibonacciHandler)
	http.HandleFunc("/hitme", hitMeHandler)
	http.HandleFunc("/nthprime", primethHandler)
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

func fibonacci(n int64) int64 {
	var f1, f2, nth int64 = 0, 1, 0

	for i := int64(0); i <= n; i++ {
		nth = f1 + f2
		f1 = f2
		f2 = nth
	}

	return nth
}

// prime check
func isPrime(number int64) bool {

	if number < int64(2) {
		return false
	}

	for i := int64(2); i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

//get nth prime
func getNthPrime(nth int64) int64 {
	if nth < 1 {
		return -1
	}

	var nthPrime, prime int64 = 0, -1
	for i := int64(2); nthPrime < nth; i++ {
		if isPrime(i) {
			prime = i
			nthPrime++
		}
	}

	return prime
}

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
