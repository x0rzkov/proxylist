package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/x0rzkov/proxylist/pkg/proxylist"
)

const service = "https://gimmeproxy.com/api/getProxy"

var export bool
var howmany int
var sorting string
var countries string

func main() {
	flag.Usage = func() {
		fmt.Println("ProxyList")
		fmt.Println("http://cixtor.com/")
		fmt.Println("https://github.com/cixtor/proxylist/cmd/proxylist")
		fmt.Println("https://en.wikipedia.org/wiki/Proxy_server")
		fmt.Println(service)
		fmt.Println()
		fmt.Println("In computer networks, a proxy server is a server (a computer system or an\n" +
			"application) that acts as an intermediary for requests from clients seeking\n" +
			"resources from other servers. A client connects to the proxy server, requesting\n" +
			"some service, such as a file, connection, web page, or other resource available\n" +
			"from a different server and the proxy server evaluates the request as a way to\n" +
			"simplify and control its complexity. Proxies were invented to add structure and\n" +
			"encapsulation to distributed systems. Today, most proxies are web proxies,\n" +
			"facilitating access to content on the World Wide Web and providing anonymity.\n")
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

	flag.BoolVar(&export, "e", false, "Export all data as JSON")
	flag.IntVar(&howmany, "n", 10, "How many proxies to list")
	flag.StringVar(&countries, "c", "US", "Country's proxy (format: ISO2, commas spearated, eg: US,UK)")
	flag.StringVar(&sorting, "s", "speed", "Sort in descending mode")

	flag.Parse()

	countriesSlice := strings.Split(countries, ",")
	p := proxylist.NewProxy(service, countriesSlice...)

	if err := p.Execute(howmany); err != nil {
		fmt.Print(err)
		/* return */
	}

	if p.Success == 0 {
		return
	}

	p.Sort(sorting)

	if export {
		p.Export(os.Stdout)
		return
	}

	p.Print(os.Stdout)
}
