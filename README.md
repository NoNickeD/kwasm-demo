# KWasm Demo

This repository provides a comprehensive demo of integrating WebAssembly (Wasm) with Kubernetes using KWasm. The demo includes setting up a Kind cluster, installing the KWasm operator, and deploying Wasm applications.

## Step 1: Create a Kind Cluster

Create a Kind cluster using the provided configuration. Save the following YAML into a file named `kind-config.yaml`:

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    kubeadmConfigPatches:
      - |
        kind: InitConfiguration
        nodeRegistration:
          kubeletExtraArgs:
            node-labels: "ingress-ready=true"
    extraPortMappings:
      - containerPort: 80
        hostPort: 80
        protocol: TCP
      - containerPort: 443
        hostPort: 443
        protocol: TCP
```

## Step 2: Install KWasm Operator

Add the Helm repository and install the KWasm operator:

```bash
helm repo add kwasm http://kwasm.sh/kwasm-operator/
helm install -n kwasm --create-namespace kwasm-operator kwasm/kwasm-operator
```

Annotate the nodes to indicate that they are ready for KWasm:

```bash
kubectl annotate node --all kwasm.sh/kwasm-node=true
```

## Step 3: Install Wasm Runtime and Create a Test Workload

Choose one of the following runtimes to install and test:

### Option 1: WasmEdge

Apply the following YAML to create a RuntimeClass and a test Job using WasmEdge:

```bash
kubectl apply -f - <<EOF
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: wasmedge
handler: wasmedge
---
apiVersion: batch/v1
kind: Job
metadata:
  creationTimestamp: null
  name: wasm-test
spec:
  template:
    metadata:
      annotations:
        module.wasm.image/variant: compat-smart
      creationTimestamp: null
    spec:
      containers:
      - image: wasmedge/example-wasi:latest
        name: wasm-test
        resources: {}
      restartPolicy: Never
      runtimeClassName: wasmedge
  backoffLimit: 1
EOF
```

### Option 2: Spin

Apply the following YAML to create a RuntimeClass and a Deployment using Spin:

```bash
kubectl apply -f - <<EOF
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: wasmtime-spin
handler: spin
EOF

kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wasm-spin
spec:
  replicas: 1
  selector:
    matchLabels:
      app: wasm-spin
  template:
    metadata:
      labels:
        app: wasm-spin
    spec:
      runtimeClassName: wasmtime-spin
      containers:
      - name: spin-hello
        image: ghcr.io/deislabs/containerd-wasm-shims/examples/spin-rust-hello:latest
        command: ["/"]
EOF
```

## Step 4: Test the Deployment

For the Spin runtime, you can test the deployment by port-forwarding and making a request:

```bash
kubectl port-forward deployment/wasm-spin 8000:80
```

In a separate terminal, test the application:

```bash
curl localhost:8000/hello
```

You should see a response from the Wasm application.

# Conclusion

WebAssembly is a transformative technology with significant advantages in performance, security, and portability. The integration of Wasm into the Kubernetes ecosystem through Kwasm marks a pivotal moment in its evolution, enabling Wasm applications to leverage the full power of Kubernetes' robust infrastructure and ecosystem. This advancement not only facilitates the adoption of Wasm in serverless and edge computing but also positions it as a strong contender for a wide range of cloud-native applications. As the ecosystem continues to grow, Wasm is poised to become a foundational technology in the modern software landscape.
