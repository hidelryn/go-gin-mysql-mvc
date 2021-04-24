package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go_sql_study2/db"
	"github.com/go_sql_study2/routes"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string { // 커스텀 로그 (아파치에서 출력하는 형식)
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	routes.Route(router)
	server := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env 로딩 할 수 업다")
	}
	db.MySQLConnect(0, os.Getenv("MY_SQL_USER"))
	cpu := runtime.NumCPU()
	fmt.Println("CPU 개수: ", cpu)
	runtime.GOMAXPROCS(cpu)
	fmt.Println("서버 동작!")
	server.ListenAndServe()
}
