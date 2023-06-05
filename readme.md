## Be sure you have installed latest version protocol buffer compiler to compile properly

### To compile proto files run below command in terminal

protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=pb  --grpc-gateway_opt=paths=source_relative  proto/*.proto

