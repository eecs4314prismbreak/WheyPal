module github.com/eecs4314prismbreak/WheyPal

go 1.15

replace github.com/eecs4314prismbreak/WheyPal/user => ./user

replace github.com/eecs4314prismbreak/WheyPal/auth => ./auth

replace github.com/eecs4314prismbreak/WheyPal/database => ./database

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/eecs4314prismbreak/WheyPal/auth v1.0.0
	github.com/eecs4314prismbreak/WheyPal/user v1.0.0
	github.com/eecs4314prismbreak/WheyPal/database v1.0.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/jackc/pgx/v4 v4.9.2
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lib/pq v1.8.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/ugorji/go v1.1.8 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect

)
