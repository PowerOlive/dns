package main

import (
        "fmt"
        "dns"
        "dns/resolver"
)

func main() {
	res := new(resolver.Resolver)
	ch := res.NewXfer()

	res.FromFile("/etc/resolv.conf")
	m := new(dns.Msg)
	m.Question = make([]dns.Question, 1)
	m.Question[0] = dns.Question{"atoom.net", dns.TypeAXFR, dns.ClassINET}

        ch <- resolver.Msg{m, nil}
	for dm := range ch {
                fmt.Printf("%v\n",dm.Dns)
        }
}
