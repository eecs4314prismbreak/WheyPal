module github.com/eecs4314prismbreak/WheyPal

go 1.15

replace github.com/eecs4314prismbreak/WheyPal/user => ./user

replace github.com/eecs4314prismbreak/WheyPal/auth => ./auth

require (
	github.com/eecs4314prismbreak/WheyPal/auth v1.0.0
	github.com/eecs4314prismbreak/WheyPal/user v1.0.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/ugorji/go v1.1.8 // indirect
	google.golang.org/protobuf v1.25.0 // indirect

)
