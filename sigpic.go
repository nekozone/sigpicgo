package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"image"
	"math/rand"
	"net/http"
	"os"
	"sigpicgo/cip"
	"sigpicgo/fakegeo"
	"sigpicgo/global"
	"sigpicgo/pic"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Backpiclist []string `json:"backpiclist"`
	Ptext       string   `json:"ptext"`
	Fontpath    string   `json:"fontpath"`
	Token       string   `json:"token"`
	Dbpath      string   `json:"dbpath"`
	Redis       string   `json:"redis"`
	Maxnum      int      `json:"maxnum"`
	Iplim       int      `json:"iplim"`
}

// var Piclist []image.Image
// var Piclistnum = 0
// var Font *truetype.Font

type userinfo struct {
	Ip    string `json:"ip"`
	Token string `json:"token"`
}

var config Config
var ctx = context.Background()

func init() {
	configfile, _ := os.ReadFile("config.json")
	// jsonstr := string(configfile)
	err := json.Unmarshal(configfile, &config)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(config.Backpiclist); i++ {
		input, err := os.Open(config.Backpiclist[i])
		if err != nil {
			panic(err)
		}
		defer input.Close()
		img, _, err := image.Decode(input)
		if err != nil {
			panic(err)
		}
		global.Piclist = append(global.Piclist, img)
	}
	global.Piclistnum = len(global.Piclist)
	fontfile, err := os.ReadFile(config.Fontpath)
	if err != nil {
		panic(err)
	}
	// defer fontfile.Close()
	global.Font, err = freetype.ParseFont(fontfile)
	if err != nil {
		panic(err)
	}
	// ctx := context.Background()

}

func main() {
	rc, err := redis.ParseURL(config.Redis)
	if err != nil {
		panic(err)
	}
	rcl := redis.NewClient(rc)
	fmt.Println(config.Dbpath)
	db, err := sql.Open("sqlite3", config.Dbpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		qm := c.Query("token")
		q := c.ClientIP()
		if qm == config.Token {
			res := cip.GetIp_c(rcl, db, q)
			c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
		} else {
			if !Numlim(rcl) {
				res := cip.GetFIp_c(rcl, q)
				c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
			} else {
				res := cip.GetIp_c(rcl, db, q)
				c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
			}
		}
	})
	r.GET("/ip/:ip", func(c *gin.Context) {
		qm := c.Query("token")
		q := c.Params.ByName("ip")
		if qm == config.Token {
			res := cip.GetIp_c(rcl, db, c.Params.ByName("ip"))
			c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
		} else {
			if !Numlim(rcl) {
				res := cip.GetFIp_c(rcl, q)
				c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
			} else {
				res := cip.GetIp_c(rcl, db, c.Params.ByName("ip"))
				c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
			}
		}
	})

	r.GET("/fakeip/:ip", func(c *gin.Context) {
		q := c.Params.ByName("ip")
		if !cip.IsIp(q) {
			c.JSON(http.StatusBadGateway, gin.H{"err": "ip address is not valid"})
			return
		}
		qidex := cip.Ip2index(q)
		uij, err := rcl.Get(ctx, qidex).Result()
		if err == redis.Nil {
			a := fakegeo.Fakeip()
			jsoa, err := json.Marshal(a)
			if err != nil {
				panic(err)
			}
			rcl.Set(ctx, qidex, jsoa, time.Hour*2)
			fmt.Println(jsoa)
			c.JSON(http.StatusOK, gin.H{"ip": q, "res": a})
		} else if err != nil {
			panic(err)
		} else {
			var a map[string]string
			err := json.Unmarshal([]byte(uij), &a)
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{"ip": q, "res": a})
		}

	})
	r.GET("/fakeip/", func(c *gin.Context) {
		q := c.ClientIP()
		res := cip.GetFIp_c(rcl, q)
		c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
	})

	r.POST("/ip/", func(c *gin.Context) {
		var u userinfo
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		q := u.Ip
		if u.Token != config.Token {
			if !Numlim(rcl) {
				res := cip.GetFIp_c(rcl, q)
				c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
			} else {
				res := cip.GetIp_c(rcl, db, c.Params.ByName("ip"))
				c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
			}
		} else {
			res := cip.GetIp_c(rcl, db, c.Params.ByName("ip"))
			c.JSON(http.StatusOK, gin.H{"ip": q, "res": res})
		}

	})

	r.GET("/img.jpg", func(c *gin.Context) {
		q := c.ClientIP()
		ua := c.Request.UserAgent()
		usinfo := pic.Getua(ua)
		res := cip.GetIp_c(rcl, db, q)
		refer := c.Request.Referer()
		number := rcl.Incr(ctx, "ddog").Val()
		ntime := time.Now().Format("2006-01-02 15:04:05")
		if refer == "" {
			refer = "neko.red"
		}
		text := [6]string{"忽有狂徒夜磨刀，帝星飘摇荧惑高。", fmt.Sprintf("IP地址: %s  %s", q, res["location"]), refer, usinfo, fmt.Sprintf("No. %d  %s", number, ntime), "@dogcraft neko.red"}
		rawpic := pic.Genpic(text)
		c.Data(http.StatusOK, "image/jpeg", rawpic)
	})

	r.GET("/r/", func(c *gin.Context) {
		picurl := rcl.SRandMember(ctx, "piclist").Val()
		if picurl == "" {
			picurl = "https://neko.red/img.jpg"
		}
		c.Redirect(http.StatusFound, picurl)
	})

	// r.Run(fmt.Sprint(config.Address) + ":" + fmt.Sprint(config.Port))
	r.Run(":8080")

}
