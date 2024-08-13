PROTO_FILES = $(shell find ./proto -name "*.proto" -type f)

.PHONY: protoc
protoc:
	rm -rf ./proto_generated/*
	mkdir ./proto_generated/openapi
	protoc -I ./proto --openapiv2_out proto_generated/openapi --openapiv2_opt logtostderr=true \
	--openapiv2_opt disable_default_errors=true \
	--openapiv2_opt allow_merge=true \
	--openapiv2_opt merge_file_name="api_definition.yml" ./proto/*.proto $(PROTO_FILES)

