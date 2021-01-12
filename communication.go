package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//var wg = sync.WaitGroup{}

func main() {
	otp := "0"
	http.HandleFunc("/otp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		fmt.Println("Request recieved for OTP")
		fmt.Println("OTP before generating new", otp)

		var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
		b := make([]byte, 6)
		n, err := io.ReadAtLeast(rand.Reader, b, 6)
		if n != 6 {
			panic(err)
		}
		for i := 0; i < len(b); i++ {
			b[i] = table[int(b[i])%len(table)]
		}
		otp = string(b)
		responseData := "Your OTP is  " + string(b)

		time.Sleep(2 * time.Second)
		fmt.Fprintf(w, responseData)
		fmt.Println(responseData)
	})

	http.HandleFunc("/vaidateOTP", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		fmt.Println("Validating OTP methon")

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(string(data))
		res1 := strings.Split(string(data), "-")
		//fmt.Printf("%v %v %v", modi, kejriwal, rahul)
		var responseData string
		if res1[0] == otp && res1[1] != "expired" {
			responseData = "OTP transaction successfull"
		} else if res1[0] != otp {
			responseData = "Your otp is incorrect"
		} else {
			responseData = "Your otp is expired"
		}
		fmt.Println(responseData)

		fmt.Fprintf(w, responseData)

	})

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err.Error())
	}
}
