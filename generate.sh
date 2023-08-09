#!/bin/bash

protoc -I ./proto/ \
    -I include/googleapis -I include/grpc-gateway \
    --go_out=. --go_opt=module=github.com/bbengfort/notes \
    --go-grpc_out=. --go-grpc_opt=module=github.com/bbengfort/notes \
    --openapiv2_out ./openapiv2 --openapiv2_opt logtostderr=true \
    --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    proto/notes/v1/*.proto
