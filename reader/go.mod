module gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader

go 1.12

require (
	gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common v1.0.17
	gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities v1.0.17
	gerrit.o-ran-sc.org/r/ric-plt/sdlgo v0.2.0
	github.com/golang/protobuf v1.3.1
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
)

replace gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common => ../common

replace gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities => ../entities

replace gerrit.o-ran-sc.org/r/ric-plt/sdlgo => gerrit.o-ran-sc.org/r/ric-plt/sdlgo.git v0.2.0
