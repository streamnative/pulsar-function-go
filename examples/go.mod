module github.com/apache/pulsar/pulsar-function-go/examples

go 1.21.5

require (
	github.com/streamnative/pulsar-function-go v0.0.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/streamnative/pulsar-function-go => ../

replace github.com/streamnative/pulsar-function-go/pf => ../pf
