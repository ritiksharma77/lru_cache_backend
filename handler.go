package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, exists := cache.Get(key)
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	ttl, err := strconv.Atoi(r.URL.Query().Get("ttl"))
	if err != nil {
		http.Error(w, "Invalid TTL", http.StatusBadRequest)
		return
	}
	cache.Set(key, value, time.Duration(ttl)*time.Second)
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	cache.Delete(key)
	w.WriteHeader(http.StatusOK)
}
