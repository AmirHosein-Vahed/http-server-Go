# HTTP Server in Go

This is a clean and well-structured HTTP server implementation in Go that handles concurrent requests and provides a solid foundation for building web applications.

## Features

- Clean and modular code structure
- Concurrent request handling
- Graceful shutdown
- Custom request and response models
- Extensible handler system
- Thread-safe handler registration

## Project Structure

```
.
├── server/
│   └── server.go     # Core server implementation with concurrent request handling
├── handlers/
│   └── handlers.go   # HTTP request handlers and routing logic
├── models/
│   ├── request.go    # Request data structures and methods
│   └── response.go   # Response data structures and methods
├── main.go          # Application entry point and server configuration
└── README.md        # Project documentation
```

### Component Details

#### server/server.go
- Contains the main `Server` struct and its methods
- Handles server lifecycle (start/stop)
- Manages concurrent request handling
- Implements thread-safe handler registration
- Provides graceful shutdown mechanism

#### handlers/handlers.go
- Defines HTTP request handlers
- Implements routing logic
- Contains endpoint-specific business logic
- Handles request/response formatting
- Includes built-in handlers for:
  - `/hello` endpoint
  - `/health` endpoint

#### models/request.go
- Defines request-related data structures
- Provides request parsing and validation
- Implements request helper methods
- Standardizes request handling across the application

#### models/response.go
- Defines response-related data structures
- Handles response formatting
- Implements response helper methods
- Ensures consistent response structure

#### main.go
- Application entry point
- Server configuration and setup
- Signal handling for graceful shutdown
- Handler registration
- Error handling and logging setup

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/http-server-Go.git
```

2. Navigate to the project directory:
```bash
cd http-server-Go
```

3. Run the server:
```bash
go run main.go
```

The server will start on port 8080 by default.

## Available Endpoints

- `/hello` - Returns a hello world message
- `/health` - Returns the server health status

## Features Explained

### Concurrent Request Handling
The server uses Go's built-in concurrency features to handle multiple requests simultaneously. Each request is processed in its own goroutine.

### Graceful Shutdown
The server implements a graceful shutdown mechanism that:
- Captures interrupt signals (Ctrl+C)
- Allows ongoing requests to complete
- Times out after 5 seconds if requests don't complete

### Thread-Safe Handler Registration
The server uses a mutex to ensure thread-safe registration of handlers, allowing for dynamic handler registration even while the server is running.

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is licensed under the MIT License - see the LICENSE file for details.