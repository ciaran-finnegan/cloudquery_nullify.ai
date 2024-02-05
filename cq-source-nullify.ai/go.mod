module github.com/org/cq-source-nullify.ai

go 1.21

require (
	github.com/apache/arrow/go/v14 a526ba697d4e
	github.com/cloudquery/plugin-pb-go v1.13.4
	github.com/cloudquery/plugin-sdk/v4 v4.17.1
	github.com/rs/zerolog v1.32.0
)

replace github.com/apache/arrow/go/v14 => github.com/cloudquery/arrow/go/v14 cd3d4114faa0

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
