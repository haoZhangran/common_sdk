protoc --go_out=plugins=grpc:. .\user_service.proto

proto文件放在idls下，注意路径
proto生成的文件在protoc_gen_code下，cd到目标目录下执行protoc --go_out=plugins=grpc:. .\user_service.proto

scripts下放的是脚本，暂时没有用到

sql下面存的是sql文件和对应的go结构体