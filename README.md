### Instalação

```go
go get github.com/lucasmdomingues/gocnpj
```

### Como usar:

```go
package main

import (
	"fmt"
	"log"

	"github.com/lucasmdomingues/gocnpj"
)

func main() {

	cnpj, err := gocnpj.BuscaCNPJ("123456789123")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cnpj)
}
```
