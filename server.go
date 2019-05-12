package main

import (
    "net/http"
    "fmt"
    "os"
)

func main() {
    serverName := os.Getenv("API_SERVER_URL")
    resp, err := http.Get(fmt.Sprintf("http://%s/nodes", serverName))
    // for each node in resp, initialize make etcd request to initialize
    // $MONGO_ID key to "wait"
}
