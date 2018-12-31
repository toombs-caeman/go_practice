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

func main() {
    idPtr := flag.Int("id", rand.Int(), "the id of this client")
    urlPtr := flag.String("url", "http://127.0.0.1:80/", "the url to send events to")
    seedPtr := flag.Int64("seed", time.Now().UTC().UnixNano(), "the seed for the random value")
    flag.Parse()
    rand.Seed(*seedPtr)

    log.Println("client", *idPtr, "is up, sending traffic to", *urlPtr)
    // set the maximum waiting time (between events) to 2 seconds
    var maxWait int64 = 2000000000
    for {
        event := map[string]int{"id":*idPtr,"value":rand.Int()}
        out, err := json.Marshal(event)
        if err != nil {
            //TODO handle error
            log.Println("err:", err)
        } else {
            // parse address arg
            _, err = http.Post(*urlPtr, "application/json", bytes.NewBuffer(out))
            if err != nil {
                 log.Println("err:", err)
            }
        }
        time.Sleep(time.Duration(rand.Int63n(maxWait)))
    }
}
