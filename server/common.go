package server

import (
    "net/http"
    "io"
    "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    resp := map[string]string{"greed": "Hello, friend!", "from": "Layne"}
    
    io.WriteString(w, returnJson(resp))
}

func returnJson(value interface{}) (string){
    json, _ := json.Marshal(value)
    return string(json)
}