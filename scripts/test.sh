`go test ./... -v -cover -coverprofile=./test/cover.out`

`go tool cover -html=./test/cover.out`