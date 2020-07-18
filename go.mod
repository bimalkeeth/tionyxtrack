module tionyxtrack

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.24.0

require (
	github.com/RussellLuo/validating v0.0.0-20200703005320-d33b89852d8a // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.2.1
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.14
	github.com/kr/pretty v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.1.16
	github.com/micro/go-micro/v2 v2.9.1-0.20200709170224-318a80f824c2
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.3.0
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.8.1-0.20200706111908-f9feeae399e9
	github.com/micro/go-plugins/transport/nats/v2 v2.3.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.5.1 // indirect
	google.golang.org/protobuf v1.25.0
)

replace github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.9.1-0.20200702162645-b5314829fa7d
