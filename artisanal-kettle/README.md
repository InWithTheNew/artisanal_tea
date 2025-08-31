# Artisanal Kettle

Artisanal Kettle is a full-stack system for managing and executing remote commands on services and environments, supporting both traditional servers (via SSH) and Kubernetes clusters (via in-cluster exec). It is designed for DevOps, SREs, and platform engineers who need a unified interface to run commands across diverse infrastructure.

## Features

- **Service Registry:** Register and list services (with metadata: name, server, isKubernetes).
- **Remote Command Execution:**
  - For traditional servers: runs commands via SSH using configured credentials.
  - For Kubernetes: finds running pods across all namespaces and execs commands in-cluster.
- **REST API:**
  - Register new services
  - List all registered services
  - Submit commands to services
- **Swagger/OpenAPI Documentation:**
  - Interactive API docs available at `/swagger/index.html` when the server is running.
- **CORS Support:**
  - Allows cross-origin requests for easy frontend integration.

## API Endpoints

- `POST /admin/submit` — Register a new service
- `GET /list/services` — List all registered service names
- `POST /submit` — Submit a command to a service (by name)

## How It Works

- **Service Registration:**
  - Register a service with a name, server address, and a flag for Kubernetes.
  - Services are stored in Redis for persistence.
- **Command Submission:**
  - When a command is submitted, the backend determines if the service is a Kubernetes or SSH target.
  - For SSH, it connects and runs the command remotely.
  - For Kubernetes, it uses in-cluster permissions to find a running pod and exec the command.
- **Frontend:**
  - (If using artisanal-tea) React app for selecting services, environments, and commands, with live feedback.

## Setup & Usage

1. **Clone the repository**
2. **Configure environment variables:**
   - `REDIS_HOST` (for backend)
   - `SSH_USER`, `SSH_KEY` (for SSH command execution)
3. **Install Go dependencies:**
   ```sh
   go mod tidy
   go get k8s.io/client-go@latest
   go get k8s.io/api@latest
   go get k8s.io/apimachinery@latest
   go get github.com/swaggo/swag/cmd/swag@latest
   go get github.com/swaggo/http-swagger
   go get github.com/swaggo/files
   ```
4. **Generate Swagger docs:**
   ```sh
   swag init
   ```
5. **Run the backend:**
   ```sh
   go run main.go
   ```
6. **Access API docs:**
   - Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Example Service JSON

```json
{
  "name": "my-app",
  "server": "10.0.0.5:22",
  "isKubernetes": false
}
```

## Security
- SSH keys are read from environment variables and never stored in code.
- Kubernetes exec uses in-cluster permissions; ensure your service account is scoped appropriately.

## Contributing
Pull requests and issues are welcome! Please open an issue to discuss your feature or bugfix idea.

## License
MIT
