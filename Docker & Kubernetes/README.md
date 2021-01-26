## Docker

```
// build image name and path
docker build -t go/api-hello src/

// show all images
docker image ls

// delete image
docker image rm 'id'

// show all containers
docker container ls

// run container
docker run --name api-hello -d -p 85:3080 go/api-hello

// stop container 
docker container stop 'id'

// delete container
docker container rm 'id'

// Web need login first in our docker account, push repository of image
docker login

// Docker push image 
docker push 'repo name'

```



## Kubernetes

1. A developer perspective 

   - Everything is an object

   - An object represents a desired state

   - YAML is the way to define the state object

   - Cloud providers offer managed Kubernetes

   - Applications will continue running if Kubernetes is down

   - Install 

     - ```
       // download for linux
       curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
       
       // validate the binary
       Download the kubectl checksum file:
       
       curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
       
       Validate the kubectl binary against the checksum file:
       
       echo "$(<kubectl.sha256) kubectl" | sha256sum --check
       
       If valid, the output is:
       kubectl: OK
       
       // install kubectl
       
       sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
       ```

     - ```
       // minikube 
       // is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.
       
       curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
       
       // make, it executable 
       chmod +x ./minikube
       
       // Finally, move it somewhere in your $PATH
       sudo mv ./minikube /usr/local/bin/
       
       // start cluster
       minikube start
       
       
       // show clusters 
       kubectl get all
       ```

2. Pods

   - Smallest compute unit

   - Group one or more containers 

   - Containers share networking, storage, co-scheduled

     ![Screenshot from 2021-01-26 15-02-05](/media/jirohero/Rest/DeskDir/oh-my-learn/Docker & Kubernetes/img/Screenshot from 2021-01-26 15-02-05.png)

   - pod.yaml

     ```yaml
     apiVersion: v1
     kind: Pod
     metadata:
     	name: go-api-hello-pod
     	labels: 
     		app: go-api-hello
     spec:
     	containers: 
     	- name: go-api-hello
     	  image: go-api-hello:lastest
     	  ports:
     	  - containerPort: 3080
     ```

     ```
     // Command:
     
     
     kubectl get pods 
     
     kubectl apply -f 'pod.yaml path'
     
     kubectl logs 'name'
     
     kubectl delete  pod 'name'
  
     //get for info
  
     kubectl describe pod 'name'
  ```
   
   - Deployments
   
     ![Screenshot from 2021-01-26 15-14-33](/media/jirohero/Rest/DeskDir/oh-my-learn/Docker & Kubernetes/img/Screenshot from 2021-01-26 15-14-33.png)
   
     ```yaml
     apiVersion: apps/v1
     kind: Deployment
     metadata:
     	name: go-api-hello-deployment
     	labels: 
     		app: go-api-hello
     spec:
     	replicas: 10
     	selector:
     		matchLabels:
     			app: go-api-hello
     		template:
     			metadata:
     				labels:
     					app: go-api-hello
     			spec:
     				containers: 
  				- name: go-api-hello
     	  			  image: go-api-hello:lastest
     	  			  ports:
     	  			  - containerPort: 3080
     ```
   
     ```
     kubectl apply -f 'deployment.yaml path'
     
     kubectl get all 
     
  // show all pods and we can see 10 pods was created
     kubectl get pods
  
     // show developments
  kubectl get deployments
     
     // delete deployments
      kubectl delete deployment 'name'
     
     // if we delete one, kubernetes will create a new one
     ```
   
   - Services
   
     ![Screenshot from 2021-01-26 15-27-56](/media/jirohero/Rest/DeskDir/oh-my-learn/Docker & Kubernetes/img/Screenshot from 2021-01-26 15-27-56.png)
   
     ```yaml
     apiVersion: v1
     kind: Service
  metadata:
     	name: go-api-hello-service
     spec:
        	type: LoadBalancer
        	selector:
        		app: go-api-hello
     	ports:
        		- protocol: TCP
        		  port: 80
        		  targetPort: 3080
     ```
   
     ```
     kubectl apply -f 'service.yaml path'
     
     kubectl get svc
     // delete service 
   kubectl delete svc go-api-hello-service   
     
     // if we use in local pc, the external-ip is pending, we can use this command for test localhost
     minikube tunnel
     ```
     
     ```
     // to test our service with request loop 
     
     for ((i=1;i<=100;i++)); do   curl -v --header "Connection: keep-alive" "localhost or service ip"; done
     
     ```
     
     