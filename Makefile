server:
	cd server && go build main.go

client:
	cd client && go build main.go

generate:
	cd soundservice && protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative *.proto