# Education

## commands

```bash
mkdir -p cmd/api internal/{domain,usecase,interfaces,infra} pkg # Setup the folder structure according to clean architecture

```

## Clean architecture

### Domain Layer

#### Entities

* Value Objects
* Repositories (Interfaces)

#### Use Case/Application Layer

* Services
* DTOs (Data Transfer Objects)
* Interfaces for interacting with the domain layer

#### Interface Adapters

* Controllers
* Presenters
* API Handlers

#### Infrastructure Layer

* Database implementations
* External services/providers
