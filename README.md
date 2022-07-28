# linenotify
LINE Notify module

## Sample code
```go
package main

import (
	"os"

	"github.com/qoopooh/linenotify/notify"
)

func main() {
	msg := "line_notify"

	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	param := notify.SendOpts{
		Token:   os.Getenv("LINE_NOTIFY_TOKEN"),
		Prefix:  "sample",
		Message: msg,
		Verbose: true,
	}

	notify.Send(param)
}
```

## Test
Follow the instruction [here](https://go.dev/doc/modules/publishing)
```sh
go test ./...
```
