package dns

import (
	"fmt"
	"net"
)

// GetMX (domain)
func GetMX(domain string) []*net.MX {
	var err error

	mxs, err := net.LookupMX(domain)

	if err != nil {
		fmt.Println(err)
	}

	// if len(mxs) == 0 {
	// 	return []{domain}
	// }

	return mxs
}
