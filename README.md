# getenvor

Simple `Getenv` function with default value support.

Usage:
```golang
import (
  "http"
  . "github.com/ikenfin/getenvor"
)

func main() {
  host := GetenvOr("HOST", "127.0.0.1") + ":" + GetenvOr("PORT", "3000")

  err := http.ListenAndServe(host, nil)
  // ...
}
```
