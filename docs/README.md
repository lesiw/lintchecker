# lesiw.io/lintchecker [![Go Reference](https://pkg.go.dev/badge/lesiw.io/lintchecker.svg)](https://pkg.go.dev/lesiw.io/lintchecker)

Package lintchecker provides a function to run golangci-lint as a Go test.

```go
package main

import (
    "testing"

    "lesiw.io/lintchecker"
)

func TestLint(t *testing.T) {
    lintchecker.Lint(t, "2.2.1") // Run golangci-lint v2.2.1.
}
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
