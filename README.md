# getenvor

Simple `GetEnvOr` function with default value support.

Usage:
```golang
import (
  "http"
  . "github.com/ikenfin/getenvor"
)

func main() {
  host := GetEnvOr("HOST", "127.0.0.1") + ":" + GetEnvOr("PORT", "3000")

  err := http.ListenAndServe(host, nil)
  // ...
}
```