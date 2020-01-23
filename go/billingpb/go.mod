module github.com/paysuper/paysuper-proto/go/billingpb

go 1.13

require (
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/jinzhu/now v1.1.1
	github.com/micro/go-micro v1.8.0
	github.com/paysuper/paysuper-tools v0.0.0-20200117101901-522574ce4d1c
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.2.1
	go.uber.org/zap v1.13.0
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/micro/go-micro => github.com/micro/go-micro v1.8.0
)
