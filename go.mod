module git.profects.com/profects/examples/stream

go 1.13

// Fix github.com/go-resty/resty unknown revision
replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

// Fix github.com/gogo/protobuf@v0.0.0-20190410021324-65acae22fc9: invalid pseudo-version: revision is shorter than canonical (65acae22fc9d)
replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d

// Fix not enough arguments in call to watch.NewStreamWatcher
replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

// Fix module github.com/go-log/log@latest (v0.1.0) found, but does not contain package github.com/go-log/log/print
replace github.com/go-log/log => github.com/go-log/log v0.1.1-0.20181211034820-a514cf01a3eb

// Fix ../../../../pkg/mod/github.com/lucas-clemente/quic-go@v0.11.2/internal/handshake/crypto_setup.go:459:42: undefined: qtls.CipherSuite
replace github.com/marten-seemann/qtls => github.com/marten-seemann/qtls v0.3.3

require (
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.1
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.10.0
	golang.org/x/net v0.0.0-20190918130420-a8b05e9114ab
	google.golang.org/grpc v1.23.1
)
