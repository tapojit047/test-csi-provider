# test-csi-provider
```bash
docker build -t tapojit047/test-csi-provider:latest .
docker push tapojit047/test-csi-provider:latest
kubectl apply -f deploy/test-driver.yaml
```