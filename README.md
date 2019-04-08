### Instalação

```go
go get github.com/lucasmdomingues/gocnpj
```

### Exemplo

```go
package main

import (
	"fmt"
	"log"

	"github.com/lucasmdomingues/gocnpj"
)

func main() {

	// Utilize um CNPJ com ou sem os caracteres especiais(Ex: 06813678000170)
	cnpj, err := gocnpj.BuscaCNPJ("06.813.678/0001-70")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cnpj)
}
```

### ReceitaWS
https://www.receitaws.com.br/
