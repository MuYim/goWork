package main

import (
	"ginEssential/common"
	"ginEssential/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = routes.CollectRoute(r)
	r.Run(":8081")
}
