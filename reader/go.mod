module gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader

go 1.17

require (
	gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common v1.2.7
	gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities v1.2.7
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/google/go-cmp v0.4.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
)

replace gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common => ../common

replace gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities => ../entities

replace gerrit.o-ran-sc.org/r/ric-plt/sdlgo => gerrit.o-ran-sc.org/r/ric-plt/sdlgo.git v0.8.0
