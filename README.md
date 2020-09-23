### Installation

```go
go get github.com/lucasmdomingues/gocnpj
```

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/lucasmdomingues/gocnpj"
)

func main() {
	service := gocnpj.NewService()

	company, err := service.Search("06.813.678/0001-70")
	if err != nil {
		log.Fatal(err)
	}
}
```

### ReceitaWS
https://www.receitaws.com.br/
