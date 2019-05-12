package main

import (
    "net/http"
    "fmt"
	"os"
	"io/ioutil"
	
	"github.com/buger/jsonparser"
)

func main() {
	host := os.Getenv("DELPHI_API_SERVER_URL")
	resp, err := http.Get(fmt.Sprintf("http://%s/nodes", host))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	val, _, _, _ := jsonparser.Get(body)
	fmt.Printf("Value: %s\n", string(val))
    // for each node in resp, initialize make etcd request to initialize
    // $MONGO_ID key to "wait"
}
