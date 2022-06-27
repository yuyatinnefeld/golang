

```bash
# create a vm for k8s
minikube start --vm-driver=hyperkit

# run image from container registry
kubectl run --restart=Never --image=gcr.io/kuar-demo/kuard-amd64:blue kuard
kubectl get pods

# setup port
kubectl port-forward kuard 8080:8080

# Open your browser to 

http://localhost:8080

# clean up
kubectl delete pods/kuard
kubectl get pods
```

```bash
# create a pod with yaml file
kubectl apply -f kuard-pod.yaml
kubectl get pods
kubectl port-forward kuard 8080:8080

# Open your browser to 
http://localhost:8080

# check logs
kubectl logs kuard

# open interactive session
kubectl exec -it kuard ash

# copy data into the container (anti pattern)
vi config.txt
chmod 777 config.txt
kubectl cp config.txt kuard:/tmp/config.txt

#clean up
kubectl delete -f kuard-pod.yaml
```