# dply
k8s service deployment manager

## Make & Build Docker Image
- docker build -t dply-server .
- docker tag dply-server herryg91/dply-server && docker push herryg91/dply-server


## How To Start
```
1. Setup Docker in Docker (or we can also spawn docker engine in a machine)
- openssl genrsa -out "~/.dply/ca/key.pem" 4096
- openssl req -new -key "~/.dply/ca/key.pem" -out "~/.dply/ca/cert.pem" -subj '/CN=docker:dind CA' -x509 -days 
- kubectl create secret -n dply generic dind-certs --from-file=ca.cert.pem=~/.dply/ca/cert.pem --from-file=ca.key.pem=~/.dply/ca/key.pem 
- Get dind client certificate: kubectl -n dply cp <pod>:/certs/client/* ~/.dply/certs
2. kubectl apply -f k8s-dply.yaml
3. kubectl apply -f k8s-dind.yaml 
4. Access from dply client (https://github.com/herryg91/dply/tree/main/dply)
```
