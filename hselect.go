package main

import (
	"flag"
	"fmt"
	"github.com/gptankit/harmonic"
	"github.com/gptankit/hselect/store"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	e string // endpoints
	v bool   // verbosity
)

// main parses command line flags, populates past errors and consults harmonic to select next service. It also pings selected service and updates harmonic with service error if ping returned a failure.
func main() {

	// ready diskv
	store.ReadyDiskv()

	flag.BoolVar(&v, "v", false, "verbosity")
	flag.StringVar(&e, "e", "", "service list csv")
	flag.Parse()

	// parse service list
	printOut("Parsing endpoints...")
	endPointList := strings.Split(e, ",")
	printOut("Done.\n")

	// initialize cluster state with zeroed error state
	printOut("Initializing cluster state...")
	cs, _ := harmonic.InitClusterState(endPointList)
	printOut("Done.\n")

	// load errors from disk
	printOut("Populating errors...")
	for _, endpoint := range cs.GetServices() {
		domain := getHostPort(endpoint)
		epErr, _ := strconv.ParseUint(string(store.ReadDKV(domain)), 10, 64)
		cs.UpdateError(endpoint, epErr)
	}
	printOut("Done.\n")

	// consult harmonic for service selection
	retryIndex, retryLimit, svc := 0, len(endPointList)-1, ""

	for retryIndex <= retryLimit {

		printOut("Selecting service...")
		svc, _ = harmonic.SelectService(cs, retryIndex, svc)
		printOut("Done -> " + svc + "\n")

		printOut("Dialing service " + svc + "...")
		domain := getHostPort(svc)
		_, err := net.DialTimeout("tcp", domain, 2*time.Second)

		if err != nil {
			printOut("Fail.\n")
			domErr, _ := cs.GetError(svc)
			store.WriteDKV(domain, []byte(strconv.FormatUint(domErr+1, 10)))
			cs.IncrementError(svc)
			retryIndex++
		} else {
			printOut("Success.\n")
			store.WriteDKV(domain, []byte("0"))
			cs.ResetError(svc)
			printOut("Connection successful to " + svc + "\n")
			fmt.Fprint(os.Stdout, svc) // final service output
			break
		}
	}

	// nothing gets to stdout if no service is selected
	return
}

// getHostPort returns host:port for a given endpoint
func getHostPort(rawEndPoint string) string {

	parsedEndPoint, err := url.Parse(rawEndPoint)
	if err != nil {
		os.Exit(1)
	}
	_, port, _ := net.SplitHostPort(parsedEndPoint.Host)
	if port == "" {
		if parsedEndPoint.Scheme == "http" {
			return parsedEndPoint.Host + ":80"
		} else if parsedEndPoint.Scheme == "https" {
			return parsedEndPoint.Host + ":443"
		}
	}

	return parsedEndPoint.Host
}

// printOut prints message if verbosity is set to true
func printOut(message string) {

	if v {
		fmt.Fprint(os.Stdout, message)
	}
}
