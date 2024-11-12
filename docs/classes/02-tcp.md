# TCP em Go

## Conceitos

1. Cliente-Servidor: A web funciona com uma arquitetura de cliente-servidor.
2. Pacote net: Em Go, o pacote net fornece as ferramentas para criar conexões de baixo nível, como sockets TCP, embora não lide diretamente com HTTP.

## Passo a Passo para criar o servicr TCP

### 1.  Importando o package `net`

- O pacote net oferece funções para abrir conexões e escutar portas.

```go
import "net"
```

### 2.  Criando o Listener

- Para criar um servidor, usamos o método `Listen` do package `net`, que recebe
  - - Network: o *protocolo*, no caso `tcp`.
  - - Endereço: porta onde o servidor escutará conexões (ex : `":5000"`).

  ```go
  listener, err := net.Listen("tcp", ":5000")
    if err != nil {
    panic(err)
    }
    defer listener.Close() // Fecha o listener ao final
  ```

### 3.  Loop para aceitar connexões

- Um loop infinito permite que o servidor aceite múltiplas conexões simultâneas.
- O método `Accept` retorna uma `Connection` para a comunicação com o cliente.

```go
for {
    conn, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    go handleConnection(conn) // Usa uma Go Routine para processar cada conexão.
}
```

### 4.  Função `handleConnection` para Processar Conexões

- Esta função lida com a comunicação entre servidor e cliente.

```go
func handleConnection(conn net.Conn) {
    defer conn.Close() // Garante o fechamento da conexão ao final.

    // Recebe dados do cliente.
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        panic(err)
    }
    data := string(buffer[:n])
    fmt.Println("Dado recebido:", data)

    // Envia resposta ao cliente.
    _, err = conn.Write([]byte("Sua mensagem foi recebida com sucesso"))
    if err != nil {
        panic(err)
    }
    fmt.Println("Resposta enviada ao cliente.")
}(conn)
```

### 5. Usando Go Routines

- As Go Routines permitem que o servidor atenda várias requisições ao mesmo tempo.
- Cada conexão é tratada em paralelo, o que evita bloquear o servidor.