# Notas: Criando Múltiplos Handlers no Servidor HTTP com Go

## Revisão: O Que É um Handler?

- Um **handler** em Go é qualquer tipo que implemente o método `ServeHTTP`.
- Ele é responsável por processar requisições HTTP e construir respostas.

## Criando Múltiplos Handlers

### 1. Problema com o `http.ListenAndServe`

- Quando usamos http.ListenAndServe, podemos associar apenas um único handler.
- Isso limita a flexibilidade de responder a diferentes URLs.

### 2. Solução: Função `http.Handle`

- A função http.Handle permite associar diferentes URLs (patterns) a handlers distintos.
- Sintaxe:

    ```go
    http.Handle(pattern string, handler http.Handler)
    ```

  - **pattern:** Define a URL ou caminho para o qual o handler será chamado.
  - **handler:** O handler associado ao padrão.
  
    ```go
    http.Handle("/hello", MyHandler{})
    http.Handle("/world", MyOtherHandler{})
    ```

- **Exemplo:**

```go
http.Handle("/hello", MyHandler{})
http.Handle("/world", MyOtherHandler{})
```

## Passo a Passo: Implementando Múltiplos Handlers

### 1. Definindo Handlers

Criamos dois handlers separados, cada um com uma resposta específica.

```go
package main

import (
    "fmt"
    "net/http"
)

// Primeiro Handler
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}

// Segundo Handler
type WorldHandler struct{}

func (w WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("World"))
}
```

### 2. Associando Handlers com `http.Handle`

Usamos a função http.Handle para mapear cada handler a uma URL específica.

```go
func main() {
    http.Handle("/hello", HelloHandler{})
    http.Handle("/world", WorldHandler{})

    fmt.Println("Servidor rodando na porta 5000")
    http.ListenAndServe(":5000", nil)
}
```

## Testando o Servidor

1. **Acessando** `/hello`:
   - URL: `http://localhost:5000/hello`
   - Resposta: `Hello`
2. **Acessando** `/world`:
    - URL: `http://localhost:5000/world`
    - Resposta: `World`
3. **Acessando URLs Não Mapeadas:**
   - URL: `http://localhost:5000/`
   - Resposta: `404 Not Found`

## Entendendo o Pattern

- O **pattern** define o caminho da URL associado ao handler.
- Exemplos:
  - `/hello`: Responde a requisições exatas para `/hello.`
  - `/world/`: Responde a qualquer URL que comece com `/world/`.

## Vantagens e Limitações

### Vantagens

- **Flexibilidade:** Podemos associar diferentes caminhos a handlers específicos.
- **Organização:** Cada handler é responsável por uma funcionalidade única.

### Limitações

- Implementar handlers para cada caminho pode ser trabalhoso.
- Frameworks e routers (como `gorilla/mux`) oferecem maneiras mais eficientes de lidar com múltiplos endpoints.

## Alternativa: Combinar Handlers em Um Único

Se necessário, podemos criar um único handler para gerenciar múltiplos padrões.

### Exemplo

```go
type UnifiedHandler struct{}

func (h UnifiedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/hello":
        w.Write([]byte("Hello"))
    case "/world":
        w.Write([]byte("World"))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("404 Not Found"))
    }
}

func main() {
    handler := UnifiedHandler{}
    fmt.Println("Servidor rodando na porta 5000")
    http.ListenAndServe(":5000", handler)
}
```

## Conclusão

- http.Handle: Facilita a associação de múltiplos handlers para diferentes padrões de URL.
- Handlers Unificados: Podem ser úteis, mas tornam o código mais difícil de gerenciar.
- Próximos Passos: Usar routers ou frameworks para simplificar a lógica de rotas.

Essa abordagem ajuda a entender como construir servidores HTTP mais flexíveis usando apenas a biblioteca padrão do Go.