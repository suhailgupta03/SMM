go build -buildmode=plugin -o plugins/analytics/analytics.so plugins/analytics/analytics.go
go build -buildmode=plugin -o plugins/integrations/integrations.so plugins/integrations/integrations.go
go run runner.go depchecker.go