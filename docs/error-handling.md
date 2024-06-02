# Error Handling Strategy in Golang Starter Project

This document outlines the error handling strategy for the Golang Starter Project, adhering to Clean Architecture principles.

## Custom Errors

## Error Handling Strategy

* **Domain Layer**: Define business-specific errors.
* **Application Layer (Use Cases)**: Handle errors returned from the domain layer and convert them into application-specific errors if necessary.
* **Interface Adapters (Controllers/Handlers)**: Translate application errors into HTTP responses or other appropriate formats.
* **Infrastructure Layer**: Return errors related to external systems or persistence mechanisms.

### Domain Layer

Errors specific to the domain logic are defined in the `domain` package.

```go
package domain

import "errors"

var (
    ErrUnimplemented = errors.New("feature not implemented")
    ErrNotSupported  = errors.New("operation not supported")
)
