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
podman-compose -f podman-compose.yaml up --build
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

## Running on Kubernetes

You can deploy the entire stack to Kubernetes using the provided manifest.

### 1. Prerequisites
- A running Kubernetes cluster (local or cloud)
- `kubectl` configured to access your cluster

### 2. Deploy the stack

From the root of this repository, run:

```sh
kubectl apply -f artisanal-stack.yaml
```

Or, to pull directly from GitHub (replace with your repo URL):

```sh
kubectl apply -f https://raw.githubusercontent.com/<your-username>/<your-repo>/main/artisanal-stack.yaml
```

### 3. Notes
- Edit `artisanal-stack.yaml` to set your Docker Hub username and base64-encoded SSH key.
- Services are deployed in the `artisanal` namespace.
- Expose the `tea` or `kettle` service with a LoadBalancer or Ingress for external access if needed.

---

## Environment Variables

- The compose file and manifest set up all required environment variables for each service.
- For sensitive values (like SSH keys), you may want to use a `.env` file or secrets manager.

---

## Development

- **Backend:** See [`artisanal-kettle/README.md`](artisanal-kettle/README.md) for backend-specific details.
- **Frontend:** See [`artisanal-tea/README.md`](artisanal-tea/README.md) for frontend-specific details.

---

## TODO

- Implement authentication and authorisation
- Make services catalogue readable from a repository config file, rather than just REDIS. Make toggleable.
- Will PHP artisan ever need interactive input? How would we implement that.
- Write a proper PHP interface framework - this seems overkill but the I'm not a fan of the implementation of how we're interacting with the remote endpoints.
---
## License

MIT (or your chosen license)
