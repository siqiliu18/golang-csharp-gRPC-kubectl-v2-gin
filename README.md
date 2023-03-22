# golang-csharp-gRPC-kubectl-v2-gin

### qh_data_service
* go gRPC service - mock DB server
* kubectl create -f deployment

### merge_service
* csharp gRPC client and server in one service
* kubectl create -f deployment

### gfp_service
* go gRPC client service using gin as RestAPI handler
* kubectl create -f deployment

## Examples:
* http://localhost:30008/gfp/template/pptx
* 30008 is a NodePort
* "/gfp/:a/:b" is the path and params
