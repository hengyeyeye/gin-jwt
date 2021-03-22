package main

import (
	"gin-jwt/common"
	"gin-jwt/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = router.CollertRouter(r)
	panic(r.Run())
}
