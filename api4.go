package main

import (
	"os"
	"fmt"
	"path"
	"runtime"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/tile/:z/:x/:y",tile)
    r.Run(":9971")
}

func tile(c *gin.Context){
    exPath := getCurrentAbPathByCaller()
    fmt.Println(exPath)
	//c.JSON(200,gin.H{"msg":c.Param("z")+","+c.Param("x")+","+c.Param("y")})
	mPath := "/out"
	file := fmt.Sprintf("%s/%s/%s/%s.jpg",exPath+mPath,c.Param("z"),c.Param("x"),c.Param("y"))
	fmt.Println(file,isExists(file))
	if !isExists(file){
		// c.File(file)
		file = fmt.Sprintf("%s/%s.jpg",exPath+mPath,"00/00")
	}
	file1, _ := ioutil.ReadFile(file)
	c.Writer.WriteString(string(file1))			
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func isExists(path string) bool {
	_, err := os.Stat(path)
    return  err == nil || os.IsExist(err)
}