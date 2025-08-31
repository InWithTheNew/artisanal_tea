# artisanal_tea

A remote proxy for PHP Artisan commands, with a React frontend and Go backend.

## Premise

**artisanal_tea** is a two-part application:
- **artisanal-kettle** (Go backend): Exposes a REST API to securely run remote PHP Artisan commands (and other shell commands) on remote servers or Kubernetes clusters. It manages service discovery, command execution, and integrates with Redis for state or caching.
- **artisanal-tea** (React frontend): Provides a user-friendly web interface to select services, environments, and commands, then submits them to the backend and displays the results.

This setup allows teams to safely and conveniently run remote Artisan commands (or similar) from a browser, with full auditability and control.

---

## Running with Podman Compose

You can run the entire stack (frontend, backend, and Redis) using Podman Compose.

### 1. Prerequisites

- [Podman](https://podman.io/) and [podman-compose](https://github.com/containers/podman-compose) installed
- (Optional) Docker Compose v2+ also works with this file

### 2. Build and Start All Services

From the root of this repository, run:

```sh
podman-compose -f podman-compose.example.yaml up --build
```

This will:
- Build and start the Go backend (`kettle`), React frontend (`tea`), and Redis (`redis`)
- Expose the following ports:
  - **Frontend:** [http://localhost:3000](http://localhost:3000)
  - **Backend API:** [http://localhost:8080](http://localhost:8080)
  - **Redis:** `localhost:6379` (internal use)

### 3. Stopping and Cleaning Up

To stop and remove all containers:

```sh
podman-compose down
```

If you get errors about container names already in use, run:

```sh
podman rm redis kettle tea
```

---

## Environment Variables

- The compose file sets up all required environment variables for each service, other than ssh keys. Only use them in plaintext for tests / non prod environments.

---

## Development

- **Backend:** See [`artisanal-kettle/README.md`](artisanal-kettle/README.md) for backend-specific details.
- **Frontend:** See [`artisanal-tea/README.md`](artisanal-tea/README.md) for frontend-specific details.

---

## License

MIT (or your chosen license)
