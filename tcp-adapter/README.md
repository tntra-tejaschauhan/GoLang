# TCP Adapter Project

A simple TCP server-client implementation in Go that demonstrates core TCP networking concepts.

## Project Structure

```
tcp-adapter/
├── cmd/
│   ├── server/         # TCP server executable
│   │   └── main.go
│   └── client/         # TCP client executable
│       └── main.go
├── pkg/
│   ├── adapter/        # Core TCP adapter logic
│   │   └── adapter.go
│   ├── protocol/       # Message protocol handling
│   │   └── protocol.go
│   └── handler/        # Connection handler
│       └── handler.go
├── go.mod
└── README.md
```

## Core TCP Concepts Demonstrated

1. **TCP Listener** (`net.Listen`) - Creates a TCP server socket
2. **Connection Acceptance** (`listener.Accept`) - Accepts incoming connections
3. **Concurrent Handling** - Each connection handled in a separate goroutine
4. **Buffered I/O** - Using `bufio.Reader` and `bufio.Writer` for efficient communication
5. **Protocol Design** - Simple message format (COMMAND:PAYLOAD)
6. **Graceful Shutdown** - Handling OS signals for clean server shutdown

## How It Works

### Protocol Format
Messages follow a simple format:
```
COMMAND:PAYLOAD\n
```

Example: `ECHO:Hello World\n`

### Server Flow
1. Creates TCP listener on `localhost:8080`
2. Accepts incoming connections
3. Each connection spawns a goroutine
4. Reads messages, processes commands, sends responses
5. Continues until client disconnects

### Available Commands
- `ECHO <text>` - Echo back the text
- `UPPER <text>` - Convert to uppercase
- `LOWER <text>` - Convert to lowercase
- `REVERSE <text>` - Reverse the text
- `PING` - Health check (responds with PONG)
- `QUIT` - Disconnect from server

## Running the Project

### 1. Start the Server
```bash
cd tcp-adapter
go run cmd/server/main.go
```

Output:
```
TCP Adapter listening on localhost:8080
```

### 2. Start the Client (in another terminal)
```bash
cd tcp-adapter
go run cmd/client/main.go
```

### 3. Interact with the Server
```
> ECHO Hello TCP
Server: [ECHO_RESPONSE] Hello TCP

> UPPER golang is awesome
Server: [UPPER_RESPONSE] GOLANG IS AWESOME

> REVERSE hello
Server: [REVERSE_RESPONSE] olleh

> PING
Server: [PONG] alive

> QUIT
Server: [BYE] Goodbye!
```

## Building Executables

```bash
# Build server
go build -o server cmd/server/main.go

# Build client
go build -o client cmd/client/main.go

# Run
./server
./client
```

## Testing with Multiple Clients

Open multiple terminals and run the client in each. The server handles concurrent connections:

**Terminal 1:**
```bash
go run cmd/client/main.go
```

**Terminal 2:**
```bash
go run cmd/client/main.go
```

Both clients can communicate with the server simultaneously!

## Testing with Telnet

You can also test using telnet:
```bash
telnet localhost 8080
```

Then type commands manually:
```
ECHO:Testing with telnet
PING:
QUIT:
```

## Key Learning Points

### 1. TCP Listener Creation
```go
listener, err := net.Listen("tcp", "localhost:8080")
```

### 2. Accepting Connections
```go
conn, err := listener.Accept()
```

### 3. Concurrent Handling
```go
go handleConnection(conn)  // Each connection in its own goroutine
```

### 4. Reading from Connection
```go
reader := bufio.NewReader(conn)
line, err := reader.ReadString('\n')
```

### 5. Writing to Connection
```go
writer := bufio.NewWriter(conn)
writer.WriteString(message)
writer.Flush()  // Important: flush the buffer!
```

## Extending the Project

Ideas for enhancement:
- Add authentication
- Implement custom binary protocol
- Add TLS/SSL encryption
- Implement connection pooling
- Add metrics and monitoring
- Create a chat room functionality
- Add rate limiting

## License

MIT License - Feel free to use and modify!
