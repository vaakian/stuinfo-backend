package routers
import (
	"net/http"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)
var db *sql.DB
func init(){
	db,_ = sql.Open("mysql", "alva:zhangguang@tcp(banlan6.cn:3306)/alva")
	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(2000)
	db.Ping()
}
type Stu struct {
    Baominghao string `json:"bmh"`
    Changci string `json:"cc"`
    Img string `json:"img"`
    Kaochang string `json:"kc"`
    Kaoshishijian string `json:"kssj"`
    Xingming string `json:"xm"`
    Zhunkaozheng string `json:"zkzh"`
    Zuowei string `json:"zwh"`
}
type Names struct {
	Value	string 	`json:"value"`
}
func GetInfoByName(xm string) []Stu {
	q := fmt.Sprintf("SELECT bmh, cc, img, kc, kssj, xm, zkzh, zwh FROM students WHERE xm LIKE '%s' LIMIT 50", xm)
	fmt.Printf("query: %s\n", q)
	rows,_ := db.Query(q)
	var data []Stu
	var temp = &Stu{}
	// defer db.Close()
	for rows.Next() {
		rows.Scan(
			&temp.Baominghao,
			&temp.Changci,
			&temp.Img, 
			&temp.Kaochang, 
			&temp.Kaoshishijian, 
			&temp.Xingming, 
			&temp.Zhunkaozheng, 
			&temp.Zuowei,
		)
		data = append(data, *temp)
	}
	return data
}
func ListNames(keyword string) []Names {
	q := fmt.Sprintf("SELECT xm FROM students WHERE xm LIKE '%%%s%%' GROUP BY xm LIMIT 300", keyword)
	fmt.Printf("query: %s\n", q)
	rows,_ := db.Query(q)
	// defer db.Close()
	var data []Names
	var temp = &Names{}
	for rows.Next() {
		rows.Scan(
			&temp.Value,
		)
		data = append(data, *temp)
	}
	return data
}
func Getimg(c *gin.Context) {
	xm := c.Query("xm")
	infoData := GetInfoByName(xm)
	fmt.Printf("姓名是:%s\n", xm)
	c.JSON(http.StatusOK, infoData)
}
func Listnames(c *gin.Context) {
	keyword := c.Query("keyword")
	c.JSON(http.StatusOK, ListNames(keyword))
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

