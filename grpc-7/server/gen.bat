cd pb && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Model.proto
cd ..