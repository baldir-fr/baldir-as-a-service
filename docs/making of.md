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

### bootstrapping

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

Directories and files

```
.
├── bin/
├── cmd
│   └── root.go
├── contact         <------
│   └── cmd.go      <------
├── docs/
├── LICENSE
├── Makefile
├── README.md
├── baas
├── go.mod
├── go.sum
└── main.go
```

### Testing

Create a subpackage in the `contact` package.
Add a `contact/show.go` and `contact/show_test.go` files.

This submodule will be responsible for producing the output for the `contact` command.
To make it testable it is decoupled from cobra command dependency.

First write a failing test.

`contact/show_test.go` 

```go
package show  
  
import (  
   "github.com/stretchr/testify/assert"  
   "testing")  
  
func TestDo(t *testing.T) {  
   expectedContactInfos := `Hi! I am Marc Bouvier, aspiring Software Craftsman.  
I am always learning to improve my craft and trying new ways to add value to products.  
  
To reach me:  
- e-mail:     mailto:m.bouvier.dev@gmail.com  
- phone:      (+33) 6 66 15 95 38  
- LinkedIn:   https://www.linkedin.com/in/profileofmarcbouvier/  
- Web Site:   https://baldir.fr/`  
   actual, err := Do()  
  
   assert.Nil(t, err)  
   assert.Equal(t, actual, expectedContactInfos)  
}

```

We added new dependencies, so we update module reference and download the imported modules.

```shell
go mod tidy
go mod download
```

Let's run the test.

```shell
go test -v ./...
```

Result: failure as expected because the system under test does not exist yet.

```
# baas/contact/show [baas/contact/show.test]
contact/show/show_test.go:17:17: undefined: Do
?       baas    [no test files]
?       baas/cmd        [no test files]
?       baas/contact    [no test files]
FAIL    baas/contact/show [build failed]
FAIL

```

Let's create the implementation which still does nothing, but compiles.

`contact/show.go`

```go
package show  
  
func Do() (string, error) {  
    return "", nil  
}
```

Let's run the test again. 

```shell
go test -v ./...
```

Now, it fails but for a good reason.

```
?       baas    [no test files]
?       baas/cmd        [no test files]
?       baas/contact    [no test files]
=== RUN   TestDo
    show_test.go:20: 
                Error Trace:    show_test.go:20
                Error:          Not equal: 
                                expected: ""
                                actual  : "Hi! I am Marc Bouvier, aspiring Software Craftsman.\nI am always learning to improve my craft and trying new ways to add value to products.\n\nTo reach me:\n- e-mail:     mailto:m.bouvier.dev@gmail.com\n- phone:      (+33) 6 66 15 95 38\n- LinkedIn:   https://www.linkedin.com/in/profileofmarcbouvier/\n- Web Site:   https://baldir.fr/"
                                
                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1,8 @@
                                +Hi! I am Marc Bouvier, aspiring Software Craftsman.
                                +I am always learning to improve my craft and trying new ways to add value to products.
                                 
                                +To reach me:
                                +- e-mail:     mailto:m.bouvier.dev@gmail.com
                                +- phone:      (+33) 6 66 15 95 38
                                +- LinkedIn:   https://www.linkedin.com/in/profileofmarcbouvier/
                                +- Web Site:   https://baldir.fr/
                Test:           TestDo
--- FAIL: TestDo (0.00s)
FAIL
FAIL    baas/contact/show       0.261s
FAIL
```

Let's write the minimal code to make the test pass.

`contact/show.go`
```go
package show

func Do() (string, error) {
	return `Hi! I am Marc Bouvier, aspiring Software Craftsman.  
I am always learning to improve my craft and trying new ways to add value to products.  
  
To reach me:  
- e-mail:     mailto:m.bouvier.dev@gmail.com  
- phone:      (+33) 6 66 15 95 38  
- LinkedIn:   https://www.linkedin.com/in/profileofmarcbouvier/  
- Web Site:   https://baldir.fr/`, nil
}

```

Let's run the test again.

```
?       baas    [no test files]
?       baas/cmd        [no test files]
?       baas/contact    [no test files]
=== RUN   TestDo
--- PASS: TestDo (0.00s)
PASS
ok      baas/contact/show       0.158s

```

Test passes, we can refactor to split responsibilities inside the `show` submodule.

`contact/show.go`

```go
package show

func Do() (string, error) {
	return plainTextContactDescription(), nil
}

func plainTextContactDescription() string {
    return `Hi! I am Marc Bouvier, aspiring Software Craftsman.  
I am always learning to improve my craft and trying new ways to add value to products.  
  
To reach me:  
- e-mail:     mailto:m.bouvier.dev@gmail.com  
- phone:      (+33) 6 66 15 95 38  
- LinkedIn:   https://www.linkedin.com/in/profileofmarcbouvier/  
- Web Site:   https://baldir.fr/`
}

```

Run the tests again... still passes.

```
?       baas    [no test files]
?       baas/cmd        [no test files]
?       baas/contact    [no test files]
=== RUN   TestDo
--- PASS: TestDo (0.00s)
PASS
ok      baas/contact/show       (cached)
```

Directories and files

```
.
├── bin/
├── cmd
│   └── root.go
├── contact
│   ├── show
│   │   ├── show.go
│   │   └── show_test.go
│   └── cmd.go
├── docs/
├── LICENSE
├── Makefile
├── README.md
├── baas
├── go.mod
├── go.sum
└── main.go
```