module gitlab.com/elixxir/user-discovery-bot

go 1.13

require (
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.9.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.9.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/viper v1.7.1
	github.com/ugorji/go v1.1.4 // indirect
	github.com/xordataexchange/crypt v0.0.3-0.20170626215501-b2862e3d0a77 // indirect
	gitlab.com/elixxir/client v1.5.1-0.20210520232413-2a3894c2b690
	gitlab.com/elixxir/comms v0.0.4-0.20210520231539-ca1bbeb6e3ec
	gitlab.com/elixxir/crypto v0.0.7-0.20210520231341-cc91d0be28ae
	gitlab.com/elixxir/primitives v0.0.3-0.20210520220650-16cb34e6b7e3
	gitlab.com/xx_network/comms v0.0.4-0.20210517205649-06ddfa8d2a75
	gitlab.com/xx_network/crypto v0.0.5-0.20210517205543-4ae99cbb9063
	gitlab.com/xx_network/primitives v0.0.4-0.20210517202253-c7b4bd0087ea
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	golang.org/x/sys v0.0.0-20210426230700-d19ff857e887 // indirect
	google.golang.org/genproto v0.0.0-20210427215850-f767ed18ee4d // indirect
	google.golang.org/grpc v1.37.0 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.27.1
