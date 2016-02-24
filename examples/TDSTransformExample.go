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

func parseXML(r *http.Request) string {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if len(body) > 0 {
		return string(body)
	}
	return ""
}

func dnsForward(w http.ResponseWriter, r *http.Request) {
	data := parseXML(r)
	if data != "" {
		m := maltegogo.MaltegoMsg(data)
		hostname := m.Value
		revHosts, _ := net.LookupHost(hostname)
		TRX := maltegogo.MaltegoTransform{}
		for _, ip := range revHosts {
			fmt.Println("Forward resolved " + hostname + " to " + ip)
			TRX.AddEntity("maltego.IPv4Address", ip)
		}
		fmt.Fprintf(w, TRX.ReturnOutput()) // send data to client side
	} else {
		fmt.Fprintf(w, "Error: No form data")
	}
}

func dnsReverse(w http.ResponseWriter, r *http.Request) {
	data := parseXML(r)
	if data != "" {
		m := maltegogo.MaltegoMsg(data)
		ip := m.Value
		revIPs, _ := net.LookupAddr(ip)
		TRX := maltegogo.MaltegoTransform{}
		for _, host := range revIPs {
			fmt.Println("Reverse resolved " + ip + " to " + host)
			NewEnt := TRX.AddEntity("maltego.DNSName", host)
			NewEnt.AddProperty("hostname", "Hostname", "stict", host)
		}
		fmt.Fprintf(w, TRX.ReturnOutput()) // send data to client side
	} else {
		fmt.Fprintf(w, "Error: No form data")
	}
}

func main() {
	http.HandleFunc("/", defaultResponse)
	//Add Transform endpoints here (add to TDS)
	http.HandleFunc("/reverseDNS", dnsReverse)
	http.HandleFunc("/forwardDNS", dnsForward)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
