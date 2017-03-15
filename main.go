package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
)

type WhoRequest struct {
    Hostname string
    Users []UserStatus
}

type UserStatus struct {
    User string
    Device string
    Timestamp string
}

func main() {
    flag.Parse()
    log.Println("running")

    http.HandleFunc("/", who)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func who(w http.ResponseWriter, r *http.Request) {
    reqData := WhoRequest{}
    _ = json.NewDecoder(r.Body).Decode(&reqData)

    resp := fmt.Sprintf(`{"user": "%s"}`, Infer(reqData.Users))
    log.Print(resp)
    fmt.Fprintf(w, resp)
}

func Infer(users []UserStatus) string {
    for _,u := range users {
        if u.Device == "console" { return u.User }
    }

    return "fuck knows"
}
