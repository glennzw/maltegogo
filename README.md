# Introduction

Translation of the Python Maltego [TDS library](http://www.paterva.com/web6/documentation/developer-tds.php), allowing you to create amazingly fast Go Transforms.

# Installation

`go get github.com/glennzw/maltegogo`

# A taste of maltegogo

To get a taste of maltegogo, try this:

```
//Test local transform
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/glennzw/maltegogo"
)

func main() {

	lt := maltegogo.ParseLocalArguments(os.Args)
	ip := lt.Value
	revIPs, _ := net.LookupAddr(ip)
	TRX := maltegogo.MaltegoTransform{}
	for _, host := range revIPs {
		NewEnt := TRX.AddEntity("maltego.DNSName", host)
		NewEnt.AddProperty("hostname", "Hostname", "stict", host)
		NewEnt.SetIconURL("http://dadgostari-hr.ir/portals/dadhor/pic/4-1.png")
	}
	fmt.Println(TRX.ReturnOutput())

}
```

For a TDS transform, try the following:

```
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/glennzw/maltegogo"
)

func dnsReverse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
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
}

func main() {
  //Transform endpoints below (add to TDS)
	http.HandleFunc("/reverseDNS", dnsReverse)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

```

For the above transform you'd of course need to add the `:9090/reverseDNS` path to your TDS server (the [public one](https://cetas.paterva.com/TDS/), or [your own](http://www.paterva.com/web6/sales/server.php?)). e.g:

![TDS Config](example.png?raw=true "TDS Config")
