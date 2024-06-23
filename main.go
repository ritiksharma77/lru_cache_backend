package main

import (
	"fmt"
	"net/http"
)

var cache *LRUCache

func main() {
	cache = NewLRUCache(10)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/delete", deleteHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
