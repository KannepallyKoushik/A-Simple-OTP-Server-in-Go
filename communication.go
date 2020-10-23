package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//var wg = sync.WaitGroup{}

func main() {
	modi := 0
	kejriwal := 0
	rahul := 0
	http.HandleFunc("/vote", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		// fmt.Println("Request recieved through ajax")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		option := string(data)
		responseData := "Thanks for voting, Your candidate is- " + option

		switch string(data) {
		case "modi":
			modi++
		case "kejriwal":
			kejriwal++
		case "rahul":
			rahul++
		}

		fmt.Printf("%v %v %v", modi, kejriwal, rahul)
		fmt.Fprintf(w, responseData)

	})

	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		// fmt.Println("Request recieved through ajax")

		responseData := fmt.Sprint("modi -", modi, " keriwal -", kejriwal, " rahul- ", rahul)

		//fmt.Printf("%v %v %v", modi, kejriwal, rahul)
		fmt.Fprintf(w, responseData)

	})

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err.Error())
	}
}
