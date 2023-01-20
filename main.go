package main

import (
	"go-crud-gin-gorm/initializers"
	"go-crud-gin-gorm/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := routers.SetupRouter()
	r.Run()
}

// go mod init
// go get github.com/joho/godotenv
// go get -u github.com/gin-gonic/gin
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres

// go run migrations/migrations.go

//criar classe de acesso ao banco
//gerenciar todos os erros com o novo objeto
