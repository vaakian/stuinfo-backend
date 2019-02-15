package main
import(
    // _"database/sql"
    "fmt"
    // _"strconv"
    // _"github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
	// _"net/http"
	routs "./src"
)
func main(){
	gin.SetMode(gin.ReleaseMode)
	// store := sessions.NewCoo
	router := gin.Default()
	const serverInfo string = "0.0.0.0:8080"
	fmt.Printf("running at: %s\n", serverInfo)
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	go router.GET("/", routs.Index)
	go router.GET("/query", routs.Getimg)
	go router.GET("/listnames", routs.Listnames)
	router.Run(serverInfo)
}