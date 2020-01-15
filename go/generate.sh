protoc -I=../declarations/casbin --micro_out=./casbinpb/ --go_out=./casbinpb/ ../declarations/casbin/casbin.proto

protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/billing.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/cardpay.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/grpc.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/paylink.proto

protoc -I=../declarations/currencies-service --micro_out=./currenciespb/ --go_out=./currenciespb/ ../declarations/currencies-service/currencies.proto

protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/cardpay.proto
protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/entity.proto
protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/repository.proto
protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/xsolla.proto

protoc -I=../declarations/postmark-sender --micro_out=./postmarkpb/ --go_out=./postmarkpb/ ../declarations/postmark-sender/postmark.proto

protoc -I=../declarations/reporter --micro_out=./reporterpb/ --go_out=./reporterpb/ ../declarations/reporter/reporter.proto

protoc -I=../declarations/tax-service --micro_out=./taxpb/ --go_out=./taxpb/ ../declarations/tax-service/tax_service.proto

protoc -I=../declarations/document-signer --micro_out=./document_signerpb/ --go_out=./document_signerpb/ ../declarations/document-signer/signer.proto

echo "INJECTING TAGS"

protoc-go-inject-tag -input=./billingpb/billing.pb.go -XXX_skip=bson,json,structure,validate
protoc-go-inject-tag -input=./billingpb/cardpay.pb.go -XXX_skip=bson,json,structure,validate
protoc-go-inject-tag -input=./billingpb/grpc.pb.go -XXX_skip=bson,json,structure,validate
protoc-go-inject-tag -input=./billingpb/paylink.pb.go -XXX_skip=bson,json,structure,validate

protoc-go-inject-tag -input=./currenciespb/currencies.pb.go -XXX_skip=bson,json,structure,validate

protoc-go-inject-tag -input=./document_signerpb/signer.pb.go -XXX_skip=bson,json,structure,validate

protoc-go-inject-tag -input=./postmarkpb/postmark.pb.go -XXX_skip=bson,json,structure,validate

protoc-go-inject-tag -input=./recurringpb/cardpay.pb.go -XXX_skip=bson,json,structure,validate
protoc-go-inject-tag -input=./recurringpb/entity.pb.go -XXX_skip=bson,json,structure,validate
protoc-go-inject-tag -input=./recurringpb/repository.pb.go -XXX_skip=bson,json,structure,validate
protoc-go-inject-tag -input=./recurringpb/xsolla.pb.go -XXX_skip=bson,json,structure,validate

protoc-go-inject-tag -input=./reporterpb/reporter.pb.go -XXX_skip=bson,json,structure,validate

protoc-go-inject-tag -input=./taxpb/tax_service.pb.go -XXX_skip=bson,json,structure,validate

echo "GENERATING MOCKS"

for d in */ ; do
    mockery -recursive=true -all -dir=./"$d" -output=./"$d"/mocks
done