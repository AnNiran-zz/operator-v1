Kubernetes Operator

Operator monitors Kubeneters CRD types and their objects in terms of:
create
delete
update

Types names are set to be literal (no pretty names): 
typ0crd is type 0 of Custom Resource Definition
typ0crd-obj0 is object 0 of typ0crd

Cobra package: the operator contains implemented functions for add/remove of CRD type, add/remove of CRD object
as well as for generating pkg/clientset, informers and listers for each CRD type
Example (run in a separate terminal):
./main crd add [crd-name]
./main crd remove [crd-name]

./main crd generate-client [crd-name]
./main crd delelete-client [crd-name]

./main object add [crd-name] [object-name]
./main object remove [crd-name] [object-name]

the above commands use kubectl output and display results using stdout and stderr fd

How to run:
start monitoring in a separate terminal with
go run operator-v1/controllers/typ0crd/main.go
* or extend functionality to use operator-v1/controllers/main.go file to run more than one types monitoring at once 

* Project is to be extended to implement monitoring of more than one CRDs and their addition/deletion and their objects 

