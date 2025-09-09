/daily-helper <- project root
│
├─ go.mod
├─ go.sum
│
├─ cmd <- entrypoints for your binaries
│ ├─ daily-helper <- main app/server
│ │ └─ main.go <- wires everything: routes, services, db
│ └─ migrate <- optional, for DB migrations
│ └─ main.go
│
├─ internal <- private application code
│ ├─ domain <- entities & repository interfaces
│ │ └─ task.go
│ │
│ ├─ repository <- repository implementations (DB, memory, etc.)
│ │ └─ task_repo.go
│ │
│ ├─ service <- business logic / use-cases
│ │ └─ task_service.go
│ │
│ ├─ handler <- HTTP handlers / controllers
│ │ └─ task_handler.go
│ │
│ ├─ middlewares <- auth, logging, etc.
│ │ └─ auth.go
│ │
│ └─ utils <- helpers, e.g., JWT, validations
│ └─ jwt.go
│
├─ pkg <- optional, reusable packages (public)
│
└─ configs <- configuration files, env, etc.
└─ config.yaml

Key points for Clean Architecture:

1. Domain Layer (internal/domain)

   - Contains your core entities (Task) and repository interfaces.
   - No dependencies on DB, HTTP, or frameworks.

2. Repository Layer (internal/repository)

   - Implements interfaces defined in domain.
   - Can be memory-based or DB-based (e.g., PostgreSQL).
   - Only depends on domain.

3. Service Layer (internal/service)

   - Implements business logic / use-cases.
   - Calls repository interfaces from domain.

4. Handler Layer (internal/handler)

   - HTTP layer, registers routes and parses requests.
   - Calls services, never repository directly.

5. Cmd (/cmd/daily-helper/main.go)
   - Only wires everything together.
   - Creates repository → service → handler → routes → starts server.
