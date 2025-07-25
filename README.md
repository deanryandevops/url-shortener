# URL Shortener with Go and Kubernetes

A lightweight URL shortener service built with Go, containerized with Docker, and deployed to Kubernetes (Minikube).

## Features

- Shorten URLs via REST API
- Redirect from short URLs to original URLs
- Basic metrics endpoint
- Kubernetes deployment with Ingress
- Local development with Minikube

## Prerequisites

- Go (1.21+)
- Docker
- Minikube
- kubectl
- PowerShell (Windows)

## Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### 2. Build and Deploy

```bash
# Build Docker image
docker build -t url-shortener .

# Start Minikube cluster
minikube start

# Load image into Minikube
minikube image load url-shortener:latest

# Deploy to Kubernetes
kubectl apply -f k8s/

# Enable ingress addon
minikube addons enable ingress
```

### 3. Configure Host File

```bash
Add-Content -Path "C:\Windows\System32\drivers\etc\hosts" -Value "`n$(minikube ip) url-shortener.local" -Force
```

### 4. Test Application

```bash
# Create short URL
curl.exe -X POST -H "Content-Type: application/json" -d '{\"url\":\"https://google.com\"}' http://url-shortener.local/shorten

# Test redirect (replace {key} with returned key)
Start-Process "http://url-shortener.local/{key}"

#Check metrics
curl.exe http://url-shortener.local/metrics
```

### 5. Verification Commands

```bash
kubectl get all,ingress -l app=url-shortener
kubectl logs -l app=url-shortener
```

### 6. Cleanup

```bash
kubectl delete -f k8s/
minikube stop
```




