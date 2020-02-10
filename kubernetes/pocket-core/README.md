# How to deploy testnet

kubectl create namespace testnet
kubectl apply -f volumes.yml -n testnet
kubectl apply -f configmaps.yml -n testnet
kubectl apply -f pods-services.yml -n testnet 

