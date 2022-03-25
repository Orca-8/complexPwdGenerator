package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	number = "0123456789"
	letter = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	symbol = ",.<>/?!@#$%^&*()_-+~;{}[]"
)

func main() {
	route := gin.Default()
	route.Static("/static", "./static")
	route.LoadHTMLGlob("static/**/*")
	route.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	route.GET("/create", func(ctx *gin.Context) {
		go func(ctx *gin.Context) {
			length, _ := strconv.ParseInt(ctx.Query("length"), 10, 64)
			isMix, _ := strconv.ParseBool(ctx.Query("isMix"))
			rand.Seed(time.Now().Unix())
			symbols := make([]string, 0)
			symbols = append(symbols, strings.Split(number, "")...)
			symbols = append(symbols, strings.Split(letter, "")...)
			if isMix {
				// 带特殊符号
				symbols = append(symbols, strings.Split(symbol, "")...)
			}
			num := len(symbols)
			var pwd = ""
			for i := int64(0); i < length; i++ {
				pwd = fmt.Sprint(pwd, symbols[rand.Intn(num)])
			}
			ctx.JSON(http.StatusOK, gin.H{
				"pwd": pwd,
			})
		}(ctx)
	})
	route.Run(":52989")
}
