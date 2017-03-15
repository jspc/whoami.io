package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "os/exec"
    "regexp"
    "strings"
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

type Resp struct {
    User string
}

func main() {
    w := WhoRequest{Hostname: "fuckit"}

    out,_ := exec.Command("who").Output()

    outStr := string(out)
    lines := strings.Split(outStr, "\n")

    for _,l := range lines {
        re := regexp.MustCompile(`(\S*)\W*(\w*)\W*(.*)`)
        matches := re.FindAllStringSubmatch(l, -1)

        u := UserStatus{matches[0][1], matches[0][2], matches[0][3]}
        w.Users = append(w.Users, u)
    }

    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(w)
    res, _ := http.Post("http://localhost:8000", "application/json; charset=utf-8", b)

    resp := Resp{}
    err := json.NewDecoder(res.Body).Decode(&resp)
    if err != nil {
        log.Print(err)
    }
    log.Println(resp.User)
}
