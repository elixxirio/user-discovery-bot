module gitlab.com/elixxir/user-discovery-bot

go 1.13

require (
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.9.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/nyaruka/phonenumbers v1.0.68 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/viper v1.7.1
	gitlab.com/elixxir/client v1.5.1-0.20210401161618-d4b92a84c3f6
	gitlab.com/elixxir/comms v0.0.4-0.20210401161030-7ace84f51ba1
	gitlab.com/elixxir/crypto v0.0.7-0.20210401160850-96cbf25fc59e
	gitlab.com/elixxir/primitives v0.0.3-0.20210401160752-2fe779c9fb2a
	gitlab.com/xx_network/comms v0.0.4-0.20210401160731-7b8890cdd8ad
	gitlab.com/xx_network/crypto v0.0.5-0.20210401160648-4f06cace9123
	gitlab.com/xx_network/primitives v0.0.4-0.20210331161816-ed23858bdb93
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c // indirect
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
	google.golang.org/genproto v0.0.0-20210401141331-865547bb08e2 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.27.1
