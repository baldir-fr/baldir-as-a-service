```shell
go mod init baas
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
cobra-cli init
go run main.go
go build
./baas
go install github.com/kyoh86/richgo@latest
go install github.com/mitchellh/gox@latest
```

Makefile
```makefile
XC_OS="linux darwin"
XC_ARCH="amd64"
XC_PARALLEL="2"
BIN="bin"
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go get github.com/kyoh86/richgo")
endif

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif

.PHONY: all build fmt lint test install_deps clean

default: all

all: fmt test build

build: install_deps
	gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BIN)/{{.Dir}}_{{.OS}}_{{.Arch}} \
		;

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: install_deps
	$(info ******************** running tests ********************)
	richgo test -v ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

clean:
	rm -rf $(BIN)
```

```shell
make
```

## baas contact command

Let's add a command to show my contact informations.

```shell
cobra-cli add contact
```

Result

```
contact created at /path/to/baldir-as-a-service
```

Directories and files
```
.  
├── bin/  
├── cmd/  
│   ├── contact.go         <---- This file has been created
│   └── root.go  
├── docs/  
├── LICENSE  
├── Makefile  
├── README.md  
├── go.mod  
├── go.sum  
└── main.go  
```

`cmd/contact.go`

```go  
package cmd  
  
import (  
   "fmt"  
  
   "github.com/spf13/cobra")  
  
// contactCmd represents the contact commandvar contactCmd = &cobra.Command{  
   Use:   "contact",  
   Short: "A brief description of your command",  
   Long: `A longer description that spans multiple lines and likely contains examples  
and usage of using your command. For example:  
  
Cobra is a CLI library for Go that empowers applications.  
This application is a tool to generate the needed files  
to quickly create a Cobra application.`,  
   Run: func(cmd *cobra.Command, args []string) {  
      fmt.Println("contact called")  
   },  
}  
  
func init() {  
   rootCmd.AddCommand(contactCmd)  
  
   // Here you will define your flags and configuration settings.  
  
   // Cobra supports Persistent Flags which will work for this command   // and all subcommands, e.g.:   // contactCmd.PersistentFlags().String("foo", "", "A help for foo")  
   // Cobra supports local flags which will only run when this command   // is called directly, e.g.:   // contactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")}
```


This generated allows to cal the command `baas contact`

Let's try it out.

```shell
go build
./baas contact
```

Result
```
contact called

```

Now, we need to create behaviour separated from the command execution.
This will make it testable in isolation from the command runner.

The behaviour will be invoked from

```go
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("contact called") // Here
	},
```


Now we will move and rename `cmd/contact.go` to `contact/cmd.go`

In `contact/cmd.go` change package name to `contact`

```go
/*
Package contact
...
*/

// ...

package contact

// ...

// Cmd represents the contact command
var Cmd = &cobra.Command{

//...

func init() {
// remove :
// rootCmd.AddCommand(contactCmd)

// ...
```

In `cmd/root.go`

```go
// ...

func init() {

// ...
rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")  
rootCmd.AddCommand(contact.Cmd)

```

Build an verify everything is still all right.

```shell
go build
./baas contact
```

Expected result

```
contact called
```

