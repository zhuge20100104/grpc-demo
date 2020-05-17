cd pb && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services --validate_out=lang=go:../services Model.proto
protoc --go_out=plugins=grpc:../services Order.proto

protoc --go_out=plugins=grpc:../services User.proto

protoc --grpc-gateway_out=logtostderr=true:../services User.proto
protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Order.proto
cd ..