package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type JsonFormat struct {
    Ok          bool        `json:"ok, omitempty"`
    Status      string      `json:"status, omitempty"`
    Warning     string      `json:"warning, omitempty"`
    Error       string      `json:"error, omitempty"`
    Data        *Outlets    `json:"outlets, omitempty"`
}

type Outlets struct {
   // Location    int        `json:"location, omitempty"`
    Locations   []*Locations  `json:"locations, omitempty"`
}

type Locations struct {
    Id          string      `json:"id, omitempty"`
    Outlet      string      `json:"outlet, omitempty"`
}

// our main function
func main() {
    url := "https://api.xilnex.com/logic/v2/outlets"

    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }

    req.Header.Set("DKey","")
    req.Header.Set("Token","")

    // For control over HTTP client headers,
    // redirect policy, and other settings,
    // create a Client
    // A Client is an HTTP client
    client := &http.Client{}

    // Send the request via a client
    // Do sends an HTTP request and
    // returns an HTTP response
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }

    // Callers should close resp.Body
    // when done reading from it
    // Defer the closing of the body
    defer resp.Body.Close()

    // Fill the record with the data from the JSON
    var record JsonFormat

    // Use json.Decode for reading streams of JSON data
    if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
        log.Println(err)
    }

    fmt.Println("data = ", record.Data)
}