package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"kubernetes-pod-dependency-handler/Dtos"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	host = flag.String("h", "kubernetes.default", "kubernetes api server hostname")
	selfSigned = flag.Bool("s", false, "self-signed api server, if set to true ca cert will pass" +
		" to crawler library")
	namespace = flag.String("n", "default", "namespace which pod(s) located in")
	labelSelectors = flag.String("l", "", "label selectors to find pods")
	minimumRunningPods= flag.Int("m",1, "minimum running pods to wait for")
	timeout = flag.Int("t", 0, "timeout seconds to wait for pods to bring up. set to 0 for infinity")
	cacert = flag.String("c", "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt", "cacert path " +
		"for self signed certificates")
	token= flag.String("o", "/var/run/secrets/kubernetes.io/serviceaccount/token", "token file for " +
		"JWT authorization")
	port = flag.Int("p",443, "https port to connect to kubernetes api server")
	delay = flag.Int("d", 10, "delay in seconds between each check with kubernetes api server")
	help = flag.Bool("-help",false, "show help message")
)

func main() {
	flag.Parse()

	if *help {
		showHelp()
	}

	if err := validate(); err != nil {
		log.Fatal(err)
	}

	var query_url = fmt.Sprintf("https://%s:%d/api/v1/namespaces/%s/pods%s",
		*host, *port, *namespace,getQueryString())

	var ep = Endpoint{ Cacert:*cacert, SelfSigned:*selfSigned, Url:query_url, Token:*token}

	var started = time.Now()
	for {
		buffer, err := ep.Call()
		if err != nil {
			log.Fatal(err)
		}
		var pods Dtos.Pod
		json.Unmarshal(buffer, &pods)
		var c = 0
		for _, pod := range pods.Items {
			if pod.Status.Phase == "Running" {
				c++
			}
		}

		log.Println("Running Pods: "+strconv.Itoa(c))
		if c >= *minimumRunningPods {
			break
		}

		time.Sleep(time.Second*time.Duration(*delay))

		var elapsedTime = int(time.Now().Sub(started).Seconds())
		if *timeout > 0 && elapsedTime >= *timeout {
			fmt.Println("timeout reached")
			os.Exit(2)
		}
	}
	log.Println("minimum number of running pods condition satisfied, exiting...")
	os.Exit(0)
}

func showHelp() {
	flag.Usage()
	os.Exit(1)
}

func getQueryString() string {
	if *labelSelectors != "" {
		return "?"+ *labelSelectors
	}
	return ""
}