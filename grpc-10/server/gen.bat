cd pb && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Model.proto
protoc --go_out=plugins=grpc:../services Order.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Order.proto
cd ..