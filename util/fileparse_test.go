package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersionFromRequirementsTxt(t *testing.T) {
	fileData := `amqp==2.6.0
				billiard==3.6.4.0
				boto==2.49.0
				boto3==1.24.47
				celery==4.4.0
				cffi==1.13.2
				chardet==2.3.0
				codecov==2.0.15
				Django==3.2.15
				coverage==6.4.3`

	version := GetVersionFromRequirementsTxt(fileData, "django")
	assert.Equal(t, "3.2.15", *version)

	version = GetVersionFromRequirementsTxt(fileData, "boto")
	assert.Equal(t, "2.49.0", *version)

	version = GetVersionFromRequirementsTxt(fileData, "foobar")
	assert.Nil(t, version)

	version = GetVersionFromRequirementsTxt("", "django")
	assert.Nil(t, version)
}

func TestGetVersionFromPackageJSON(t *testing.T) {
	pjson := PackageJson{
		"name":    "FooBar",
		"version": "0.0.1",
		"private": true,
		"scripts": map[string]interface{}{
			"start":       "npx react-native start",
			"start:reset": "npx react-native start --reset-cache",
			"test":        "jest",
			"test:watch":  "jest --watch",
			"lint":        "eslint .",
			"android":     "npx react-native run-android",
			"ios":         "npx react-native run-ios",
		},
		"dependencies": map[string]interface{}{
			"react":  "17.0.1",
			"moment": "^2.29.2",
			"axios":  "^0.21.1",
		},
	}

	version := GetVersionFromPackageJSON(pjson, "react")
	assert.NotNil(t, version)
	assert.Equal(t, "17.0.1", *version)

	version = GetVersionFromPackageJSON(pjson, "axios")
	assert.NotNil(t, version)
	assert.Equal(t, "0.21.1", *version)

	version = GetVersionFromPackageJSON(pjson, "foobar")
	assert.Nil(t, version)

	version = GetVersionFromPackageJSON(PackageJson{}, "react")
	assert.Nil(t, version)
}

func TestParseDockerFileFromCommand(t *testing.T) {
	dockerFile := `FROM python:3.10.2-slim
					ENV PYTHONDONTWRITEBYTECODE=1
					ENV PYTHONUNBUFFERED=1
					COPY requirements.txt /
					RUN pip install -r /requirements.txt
					COPY . /src/
					WORKDIR /src/integrations
					EXPOSE 8000
					ENTRYPOINT ["/src/entrypoint.sh"]`

	commands := ParseDockerFileFromCommand(dockerFile)
	assert.Len(t, commands, 1)
	assert.Equal(t, "python", *commands[0].Image)
	assert.Equal(t, "3.10.2-slim", *commands[0].Tag)
	assert.Nil(t, commands[0].As)
	assert.Nil(t, commands[0].Platform)
	assert.Nil(t, commands[0].Digest)

	dockerFile = `FROM golang:1.19-alpine AS build

		WORKDIR /src/
		COPY . /src/
		RUN GO111MODULE=on CGO_ENABLED=0 go build -mod=readonly -o out/bin/client ./cmd/client && \
			GO111MODULE=on CGO_ENABLED=0 go build -mod=readonly -o out/bin/gateway ./cmd/gateway && \
			GO111MODULE=on CGO_ENABLED=0 go build -mod=readonly -o out/bin/server ./cmd/server
		
		
		FROM alpine
		
		RUN mkdir -p app
		WORKDIR /app
		COPY --from=build /src/package.json /src/config/config.toml /src/out/bin/server /src/out/bin/client /src/out/bin/gateway /app/
		COPY --from=build /src/internal/migrations /app/internal/migrations
		RUN chmod +x config.toml package.json && \
			apk --no-cache add curl supervisor tzdata
		
		COPY supervisord.conf /etc/supervisord.conf
		
		RUN \
			# migrate
			MIGRATE_VERSION=v4.15.2 && \
			curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz && \
			mv migrate /usr/bin/migrate && \
			chmod +x /usr/bin/migrate && \
			# enable gRPC health probe
			GRPC_HEALTH_PROBE_VERSION=v0.4.11 && \
			curl -L https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 --output /usr/local/bin/grpc_health_probe && \
			chmod +x /usr/local/bin/grpc_health_probe
		
		EXPOSE 50051
		EXPOSE 9051
		
		CMD ["/usr/bin/supervisord"]`

	commands = ParseDockerFileFromCommand(dockerFile)
	assert.Len(t, commands, 2)
	assert.Equal(t, "golang", *commands[0].Image)
	assert.Equal(t, "alpine", *commands[1].Image)
	assert.Equal(t, "1.19-alpine", *commands[0].Tag)
}
