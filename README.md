# create 
`aws --endpoint-url=http://localhost:4566 secretsmanager create-secret --name MySecret --secret-string '{"username":"myuser","password":"mypassword"}'`

# update
`aws --endpoint-url=http://localhost:4566 secretsmanager put-secret-value --secret-id MySecret --secret-string file://secrets.json`




protoc -I . \
  -I ./google/api \
  --go_out=./gen/go \
  --go-grpc_out=./gen/go \
  --grpc-gateway_out=./gen/go \
  --openapiv2_out=./gen/swagger \
  example.proto



protoc -I . \
  -I ./google/api \
  --openapiv2_out=./gen/swagger \
  example.proto
