IMAGE_NAME=bloxroute/bdn-protobuf:v3.19.3
LOCAL_IMAGE_NAME=bx-proto-gen

# Paths and files
PROTO_PATH_IN=.
PROTO_PATH_OUT=./go/mevrelaypb
PROTO_FILES=$$(ls *.proto)

# Proto generation command
define PROTO_GEN_CMD
docker run -v $(CURDIR):/go/protobuf --platform linux/amd64 $(1) protoc \
--proto_path=$(2) \
--go_out=$(3) \
--go_opt=paths=source_relative \
--go-grpc_out=$(3) \
--go-grpc_opt=paths=source_relative \
$(4)
endef

.PHONY: genproto
.SILENT: genproto
genproto:
	mkdir -p $(PROTO_PATH_OUT)
	$(call PROTO_GEN_CMD,$(IMAGE_NAME),$(PROTO_PATH_IN),$(PROTO_PATH_OUT),$(PROTO_FILES))

.PHONY: clearproto
.SILENT: clearproto
clearproto:
	rm -f $(PROTO_PATH_OUT)/*.pb.go
