# Notas: Implementação de um Servidor HTTP Básico em Go

## Introdução ao Pacote HTTP

- No vídeo anterior, foi usado o pacote `net` para criar um servidor **TCP** básico.
- Conexões **TCP** são de baixo nível, demandando conhecimento detalhado do protocolo **HTTP**, o que pode ser trabalhoso.
- Em *Go*, o pacote `http` fornece um nível mais alto, facilitando a criação de servidores **HTTP**.

## Iniciando o Projeto com o Pacote HTTP

### 1. Importando o Pacote HTTP

- O pacote `http` está dentro de `net`, mas funciona de forma independente.

```go
import "net/http"
```

### 2. Criando o Servidor HTTP

- A função `http.ListenAndServe` inicia um servidor **HTTP**.
- Parâmetros:
  - **Endereço:** porta ou *IP* onde o servidor vai rodar (ex: `":5000"`).
  - **Handler:** responsável por lidar com as requisições (pode ser `nil` inicialmente).

```go
func main() {
    http.ListenAndServe(":5000", nil)
}
```

## Implementando um Handler

1. **O Que é um Handler:** Em Go, um handler é uma interface com o método `ServeHTTP`.
   - O `ServeHTTP` recebe dois parâmetros:
     - **ResponseWriter:** para enviar a resposta ao cliente.
     - **Request:** para acessar os dados da rSequisição.
  
```go
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}
```

2. **Configurando o Handler no Servidor:**

    ```go
    func main() {
        handler := MyHandler{}
        http.ListenAndServe(":5000", handler)
    }
    ```

3. **ResponseWriter:**

   - O `ResponseWriter` permite:
     - Definir o código de status com `WriteHeader`.
     - Escrever a resposta com `Write`.

```go
w.WriteHeader(200)
w.Write([]byte("Hello, World!"))
```

## Observando o Status e Headers da Resposta

- **Código de Status:** Se `WriteHeader` não for chamado, o padrão será `200 OK`.
- **Headers Padrão:** Go adiciona headers como `Content-Type: text/plain; charset=utf-8.`

## Resumo do Funcionamento

1. O servidor `http.ListenAndServe` está configurado para responder em uma porta específica.
2. O handler implementa `ServeHTTP`, lidando com requisições HTTP.
3. **Resultado:** Acessar `localhost:5000` exibe `Hello, World!`.

## Exemplos de Códigos HTTP e Mensagens Personalizadas

1. **Mudando o Status e Corpo da Resposta:**

    ```go
    w.WriteHeader(201)
    w.Write([]byte("Resource created"))
    ```

2. **Verificando Headers de Resposta:**

   - O navegador mostra headers HTTP padrão que o pacote http envia automaticamente, como `Content-Type`.

3. **Tratamento de Erros:**
    - Caso o handler não esteja definido, o servidor retorna `404 Not Found`.
