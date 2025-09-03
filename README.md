# nginx-proxy-golang

NGINX reverse proxy with Golang backend service that demonstrate what would be inclued in request headers.

## ⚠️  DISCLAIMER ⚠️ 

> This is for demonstration purposes only. Do not use in production environments.

## Quick Start

### Option 1: Docker Compose (Recommended for beginners)

1. Navigate to the docker-compose directory:

```bash
cd docker-compose
```

2. Start the services:

```bash
docker-compose up
```

3. Open the page with your browser: http://localhost:8080

You should see your request headers displayed on the page.

### Option 2: Kubernetes

1. Apply the Kubernetes manifests:

```bash
kubectl apply -f kubernetes/
```

2. Forward the port to access locally:

```bash
kubectl port-forward service/nginx 8080:80
```

3. Open the page with your browser: http://localhost:8080

You should see your request headers displayed on the page.

