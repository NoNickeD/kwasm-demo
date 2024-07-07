A Simple Go Application Running as WebAssembly (WASM) Using Spin, Deployed on a Kubernetes Cluster

### Prerequisites

Before you begin, ensure the following are installed on your system:

- [Go](https://go.dev/dl/)
- [TinyGo](https://tinygo.org/getting-started/install/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Kind](https://kind.sigs.k8s.io)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh)
- [Spin](https://developer.fermyon.com/spin/v2/install)

### Getting Started

- Clone the Repository

```bash
git clone https://github.com/NoNickeD/kwasm-demo.git
cd kwasm-demo
```

- Build the WASM Application

```bash
cd go-demo-app

spin build
```

- Push the WASM Module to a Registry

```bash
spin registry push ttl.sh/wasm-go-demo-app:v0.1.1
```

### Set Up the Kubernetes Cluster

- Create a Kind Cluster

```bash
kind create cluster --config kind-config.yaml
```

- Install the kwasm Operator

```bash
helm install -n kwasm --create-namespace kwasm-operator kwasm/kwasm-operator
```

- Annotate Nodes

```bash
kubectl annotate node kwasm.sh/kwasm-node=true --all
```

### Deploy the WASM Application

- You can deploy using either `kubectl apply` or `kubectl kustomize`.

Using `kubectl apply`

```bash
kubectl apply -f ./k8s/runtimeclass.yaml
kubectl apply -f ./k8s/demo-app.yaml
```

Using `kubectl kustomize`

```bash
kubectl apply -k ./k8s/
```

- Test the Deployment

```bash
kubectl port-forward svc/wasm-service 8080:8080
curl -vvv http://localhost:8080
```

### Repository Structure

```
├── LICENSE
├── README.md
├── go-demo-app
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── main.wasm
│   └── spin.toml
├── k8s
│   ├── demo-app.yaml
│   ├── kustomization.yaml
│   └── runtimeclass.yaml
└── kind-config.yaml
```

### Files Description

- `go-demo-app/main.go` This file contains the main application code written in Go, which runs as a WASM module.
- `go-demo-app/spin.toml` Spin configuration file to build and deploy the WASM module.
- `k8s/demo-app.yaml` Kubernetes deployment and service configuration for the WASM application.
- `k8s/runtimeclass.yaml` Defines the runtime class for running the WASM application using Spin.
- `k8s/kustomization.yaml` Kustomize configuration to apply the Kubernetes manifests.
- `kind-config.yaml`Kind cluster configuration file to set up the local Kubernetes cluster
