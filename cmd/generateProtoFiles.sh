#!/bin/bash
protoc --go_out=plugins=grpc:. ./FileService/proto/*.proto
protoc --go_out=plugins=grpc:. ./MailService/proto/*.proto
protoc --go_out=plugins=grpc:. ./SmtpService/proto/*.proto
protoc --go_out=plugins=grpc:. ./UserService/proto/*.proto
