```shell
go mod init baas
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
cobra-cli init
go run main.go
go build
./baas
```