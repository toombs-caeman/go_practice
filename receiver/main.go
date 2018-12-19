package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type Event struct {
    SourceID    int
    data        int
}

var events chan string

func logIncomingData(writer http.ResponseWriter, r *http.Request) {
    var event Event
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Println("err:", err)
    }
    err = json.Unmarshal(body, event)
    if err != nil {
        log.Println("err:", err)
    } else {
        events <- fmt.Sprintf("From container %d got data: %d", event.SourceID, event.data)
    }
}


func main() {
    events = make(chan string)
    // from https://yourbasic.org/golang/log-to-file/
    f, err := os.OpenFile("log/event.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer f.Close()
    // create a log which will be used to save the event stream
    output := log.New(f, "", log.LstdFlags)

    // read from the channel and write to the log
    go func() {
        output.Println("receiver is ready")
        for {
            output.Println(<-events)
        }
    }()

    http.HandleFunc("/", logIncomingData)
    log.Fatal(http.ListenAndServe(":80", nil))
}

