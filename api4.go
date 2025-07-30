package main

import (
	"os"
	"fmt"
	"log"
	"path"
	"time"
	"runtime"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint is %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	gin.DisableConsoleColor()
	r := gin.New()
	//r.SetTrustedProxies(nil)
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
	r.Use(gin.Recovery())
	//fmt.Println(gin.Mode())
	// r.GET("/ping", func(c *gin.Context) {
 //        c.JSON(200, gin.H{
 //            "message": "pong",
 //        })
 //    })
    r.GET("/tile/:z/:x/:y",tile)
    r.Run(":9971")
}

func tile(c *gin.Context){
    //exPath := getCurrentAbPathByCaller()
    exPath := ""
    //fmt.Println(exPath)
	//c.JSON(200,gin.H{"msg":c.Param("z")+","+c.Param("x")+","+c.Param("y")})
	// mPath := "/out"
	mPath := "out"
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