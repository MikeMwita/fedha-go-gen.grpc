#!/usr/bin/env bash

#Generate db proto to sdk/go-proto-gen

protoc --proto_path=./db --go_out=sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/dbserver.proto
protoc --proto_path=./db --go_out=sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/balance.proto
protoc --proto_path=./db --go_out=sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/expense.proto
protoc --proto_path=./db --go_out=sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/income.proto
protoc --proto_path=./db --go_out=sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/monthly_summary.proto
protoc --proto_path=./db --go_out=sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/User.proto



#generate expense proto to sdk/go-proto-gen
protoc --proto_path=./expense --go_out=sdk/go-proto-gen/expense --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/expense --go-grpc_opt=paths=source_relative expense/expense_server.proto
protoc --proto_path=./expense --go_out=sdk/go-proto-gen/expense --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/expense --go-grpc_opt=paths=source_relative expense/balance.proto

protoc --proto_path=./expense --go_out=sdk/go-proto-gen/expense --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/expense --go-grpc_opt=paths=source_relative expense/expense.proto

protoc --proto_path=./expense --go_out=sdk/go-proto-gen/expense --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/expense --go-grpc_opt=paths=source_relative expense/Income.proto

protoc --proto_path=./expense --go_out=sdk/go-proto-gen/expense --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/expense --go-grpc_opt=paths=source_relative expense/monthly_summary.proto
protoc --proto_path=./expense --go_out=sdk/go-proto-gen/expense --go_opt=paths=source_relative --go-grpc_out=sdk/go-proto-gen/expense --go-grpc_opt=paths=source_relative expense/User.proto


