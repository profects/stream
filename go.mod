module github.com/profects/stream

go 1.13

// Lock go-micro
replace github.com/micro/go-micro => github.com/micro/go-micro v1.16.1-0.20191125191034-db03a564fb42

// Lock go-plugins
replace github.com/micro/go-plugins => github.com/micro/go-plugins v1.4.0

// Fix github.com/gogo/protobuf@v0.0.0-20190410021324-65acae22fc9: invalid pseudo-version: revision is shorter than canonical (65acae22fc9d)
replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.3.0

// Fix not enough arguments in call to watch.NewStreamWatcher
replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

require (
	git.profects.com/profects/utils/chunk v0.0.0-20191123103927-bbbb0baf3d8e
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.16.0
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
)
