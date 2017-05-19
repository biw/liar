# Liar

__Still a WIP.__ Simple Go middleware that serves different HTTP routes to clients
if they are a bot. It currently only uses the client's user agent.

## Example

```go
package main

import (
    "fmt"
    "github.com/719ben/liar"
    "net/http"
)

func humanHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You are a human, WUBBA LUBBA DUB DUB.")
}

func botHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You are a bot, I'm happy.")
}

func main() {
    newLiar := liar.NewFunc(botHandler)
    http.Handle("/", newLiar.HumanFunc(humanHandler))
    http.ListenAndServe(":8000", nil)
}
```

## Contributions

I welcome all kinds of contributions.

## License

MIT