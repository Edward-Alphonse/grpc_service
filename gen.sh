protoc --go_out=./  --go-grpc_out=./ pb_idl/grpc/*proto


# pb_idl/grpc/*.proto 为pb的输入路径，路径相对于gen.sh所在的文件夹，即go/src/grpc_service
# go_out为pb文件的输出路径，路径相对于gen.sh所在的文件夹，即go/src/grpc_service
# --go-grpc_out为grpc代码的输出路径，路径相对于gen.sh所在的文件夹，即go/src/grpc_service
# pb_gen/grpc路径为 pb文件中配置的go_package