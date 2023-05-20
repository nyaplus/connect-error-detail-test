module github.com/2yanpath/connect-error-detail-test

go 1.19

require (
	github.com/bufbuild/connect-go v1.7.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/bufbuild/connect-grpcreflect-go v1.0.1-0.20230515231402-7290a64c9d5a
	golang.org/x/net v0.10.0
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
)

require golang.org/x/text v0.9.0 // indirect

// add
replace github.com/bufbuild/connect-grpcreflect-go => github.com/2yanpath/connect-grpcreflect-go v0.0.0-20230520063326-4a4e4e6c3b8d
