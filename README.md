# blueprint-go-toolkits

A toolkit of Go utilities and libraries, crafted for the Blueprint stack.

## Overview

This repository contains a collection of reusable Go packages that provide common functionality for Blueprint-based services and applications. Each package is designed to be modular, well-tested, and easy to integrate.

## Available Packages

- **database**: Database connectivity and management utilities
- **env**: Environment variable management
- **goodns**: DNS-related utilities
- **grpc**: gRPC client/server utilities and interceptors
- **health**: Health check implementations
- **interceptor**: Middleware and interceptors for various Go frameworks
- **metric**: Metrics collection and reporting
- **otel**: OpenTelemetry integration for tracing and observability
- **rediscluster**: Redis cluster client
- **zerolog**: Zero-allocation logging with structured logs

## Installation

Each package can be imported independently:

```go
import "github.com/ggsrc/blueprint-go-toolkits/zerolog"
import "github.com/ggsrc/blueprint-go-toolkits/otel"
// ... etc
```

## Usage

Each package has its own documentation and examples. Refer to package-specific README files for more details.

## Development

### Prerequisites

- Go 1.21+
- Make

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

### Common Commands

```bash
# Run linter
make lint

# Auto fix linting issues
make lint-fix

# Build all packages
make build

# Run tests
make test

# Generate code coverage
make codecov
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

### Common Commands

```bash
# Run linter
make lint

# Auto fix linting issues
make lint-fix

# Build all packages
make build

# Run tests
make test

# Generate code coverage
make codecov
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

### Common Commands

```bash
# Run linter
make lint

# Auto fix linting issues
make lint-fix

# Build all packages
make build

# Run tests
make test

# Generate code coverage
make codecov
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
