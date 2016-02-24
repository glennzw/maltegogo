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
