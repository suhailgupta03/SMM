go test ./... -coverpkg=./... -coverprofile coverage.txt

# Print overall and modulewise test coverage
go tool cover -func=coverage.txt

# View the coverage report in browser
go tool cover -html=coverage.txt