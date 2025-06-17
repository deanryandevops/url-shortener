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

## Build and Deploy

### Option 1: Local Development

```bash
# Build the Go binary
go build -o url-shortener

# Run locally
./url-shortener
