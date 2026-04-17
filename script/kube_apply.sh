minikube kubectl -- delete deployment go-app-deployment

minikube image load go-app:v1

minikube kubectl -- apply -f go-app-deployment.yaml

kubectl get pods