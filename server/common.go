package server

import (
    "net/http"
    "io"
    "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    resp := map[string]string{"greet": "Hello, friend!", "from": "Layne"}
    Resp(ReturnJson(resp), w, r)
}

func Resp(content string, w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    io.WriteString(w, content))
}

func ReturnJson(value interface{}) (string){
    json, _ := json.Marshal(value)
    return string(json)
}