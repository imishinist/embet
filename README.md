# embet

Golang's "embed" package has only naive features.  
embet is a tool to easily implement use cases such as listing only files
contained in `embed.FS` or exporting files in `embed.FS`.

# Usage

To use embet in your Go code:

```go
import "github.com/imishinist/embet"
```

For further details see [GoDev documentation](https://pkg.go.dev/github.com/imishinist/embet).

# Example

```bash
$ tree assets/
assets/
├── index.html
└── main.css

0 directories, 2 files
```

This is an example of writing an "embed" files to a different directory.

```go
package main

import (
	"embed"
	"log"
	"os"

	"github.com/imishinist/embet"
)

//go:embed assets/*
var assets embed.FS

func main() {
	dest := "dest"
	if err := os.Mkdir(dest, 0755); err != nil {
		log.Fatal(err)
	}
	if err := embet.WriteEmbedFiles(assets, "assets", dest); err != nil {
		log.Fatal(err)
	}
}
```

```bash
$ ls
assets/  go.mod   go.sum   main.go
$ go run main.go
$ ls
assets/  dest/    go.mod   go.sum   main.go
$ tree dest
dest
└── assets
    ├── index.html
    └── main.css
        
1 directory, 2 files        
```


This is an example of displaying a list of "embed" files.

```go
package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/imishinist/embet"
)

//go:embed assets/*
var asset embed.FS

func main() {
	// print embedded files
	list, err := embet.List(asset, "assets")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range list {
		fmt.Println(f)
	}
}
```

