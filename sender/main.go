package main
import (
    "math/rand"
    "net/http"
    "encoding/json"
    "time"
    "flag"
    "bytes"
    "log"
)
type Event struct {
    SourceID    int
    data        int
}


func main() {
    idPtr := flag.Int("id", rand.Int(), "the id of this client")
    urlPtr := flag.String("url", "127.0.0.1:8080/", "the url to send events to")
    flag.Parse()

    log.Println("client", *idPtr, "is up, sending traffic to", *urlPtr)
    // set the maximum waiting time (between events) to 2 seconds
    var maxWait int64 = 2000000000
    for {
        event := Event{*idPtr, rand.Int()}
        out, err := json.Marshal(event)
        if err != nil {
            //TODO handle error
            log.Println("err:", err)
        } else {
            // parse address arg
            _,_ = http.Post(*urlPtr, "application/json", bytes.NewBuffer(out))
        }
        time.Sleep(time.Duration(rand.Int63n(maxWait)))
    }
}
