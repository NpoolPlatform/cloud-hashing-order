module github.com/NpoolPlatform/cloud-hashing-order

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/cloud-hashing-billing v0.0.0-20211120094336-58e1a1ffa8be
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211122070118-139aac84bc79
	github.com/NpoolPlatform/message v0.0.0-20211123165556-9c8ac6d13305
	github.com/NpoolPlatform/sphinx-coininfo v0.0.0-20211123094225-56e6d9f90828
	github.com/NpoolPlatform/sphinx-service v0.0.0-20211124061425-6ef21f0b762f
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1
	google.golang.org/grpc v1.42.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
