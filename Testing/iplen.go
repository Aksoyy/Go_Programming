package main

import "fmt"

func main() {

	ips := []string{"123", "234"}
	ipsLen := len(ips)
	if ipsLen > 0 && ipsLen < 10 {
		a := 0
		for i := 10 - ipsLen; i > 0; i-- {
			ips = append(ips, ips[a])
			a++
		}
	}
	fmt.Printf("%+v", ips)
}
