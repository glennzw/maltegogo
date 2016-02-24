package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/glennzw/maltegogo"
)

func defaultResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No")
}

func dnsReverse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if len(body) > 0 {
		data := string(body)
		m := maltegogo.MaltegoMsg(data)

		ip := m.Value
		revIPs, _ := net.LookupAddr(ip)
		TRX := maltegogo.MaltegoTransform{}
		for _, host := range revIPs {
			fmt.Println(host)
			NewEnt := TRX.AddEntity("maltego.DNSName", host)
			NewEnt.AddProperty("hostname", "Hostname", "stict", host)
		}

		fmt.Fprintf(w, TRX.ReturnOutput()) // send data to client side
	} else {
		fmt.Fprintf(w, "No form data")
	}
}

func main() {
	http.HandleFunc("/", defaultResponse)
	//Add Transform endpoints here (add to TDS)
	http.HandleFunc("/reverseDNS", dnsReverse)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
