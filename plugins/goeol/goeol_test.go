package main

import (
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ExtractVersionFromMod(t *testing.T) {
	modFile := `module cuddly-eureka-

go 1.20

require (
	github.com/aws/aws-sdk-go-v2 v1.17.8
	github.com/aws/aws-sdk-go-v2/config v1.18.21
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.9
	github.com/google/go-github/v50 v50.1.0
	github.com/hashicorp/go-version v1.6.0
	github.com/stretchr/testify v1.8.2
	golang.org/x/oauth2 v0.6.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.13.20 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.26 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.9 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
`

	version := extractVersionFromMod(&modFile)
	assert.Equal(t, "1.20", *version)

	version = extractVersionFromMod(nil)
	assert.Nil(t, version)

	modFile = `module github.com/foo/authorization-api

go 1.19

require (
	// When updating a version for Knetic/govaluate, you should not use the auto-inserted v3.0.0+incompatible version, it will not work for new versions of Golang.
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/brianvoe/gofakeit/v6 v6.18.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/newrelic/go-agent/v3 v3.18.1
	github.com/newrelic/go-agent/v3/integrations/nrgrpc v1.3.1
	github.com/newrelic/go-agent/v3/integrations/nrlogrus v1.0.1
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.0
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require github.com/jackc/pgx/v5 v5.2.0

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/puddle/v2 v2.1.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sync v0.0.0-20220923202941-7f9b1623fab7 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.8 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)`

	version = extractVersionFromMod(&modFile)
	assert.Equal(t, "1.19", *version)
}

func TestGoEOL_Check(t *testing.T) {
	app := initialize.GetAppConstants()
	g := GoEOL{}
	mvalue := g.Check(app.Test.Repo.Node)
	assert.Equal(t, types.MaturityValue2, mvalue)
}

func TestGoEOL_Meta(t *testing.T) {
	g := GoEOL{}
	meta := g.Meta()
	assert.Equal(t, types.MaturityTypeDependency, meta.Type)
	assert.Equal(t, "Not EOL: GO", meta.Name)
}
