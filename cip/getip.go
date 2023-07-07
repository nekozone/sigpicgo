package cip

import (
	"context"
	"database/sql"
	"encoding/json"
	"sigpicgo/fakegeo"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func GetIp_c(rds *redis.Client, db *sql.DB, ip string) map[string]interface{} {
	if !IsIp(ip) {
		// fmt.Println(ip)
		return nil
	}
	qidx := "r" + Ip2index(ip)
	res, err := rds.Get(ctx, qidx).Result()
	if err == redis.Nil {
		a := GetAllIP(db, ip)
		josa, err := json.Marshal(a)
		if err != nil {
			panic(err)
		}
		rds.Set(ctx, qidx, josa, time.Hour*2)
		return a
	} else if err != nil {
		panic(err)
	} else {
		var a map[string]interface{}
		err := json.Unmarshal([]byte(res), &a)
		if err != nil {
			panic(err)
		}
		return a
	}

}

func GetFIp_c(rds *redis.Client, ip string) map[string]string {
	if !IsIp(ip) {
		return nil
	}
	qidx := "f" + Ip2index(ip)
	res, err := rds.Get(ctx, qidx).Result()
	if err == redis.Nil {
		a := fakegeo.Fakeip()
		josa, err := json.Marshal(a)
		if err != nil {
			panic(err)
		}
		rds.Set(ctx, qidx, josa, time.Hour*2)
		return a
	} else if err != nil {
		panic(err)
	} else {
		var a map[string]string
		err := json.Unmarshal([]byte(res), &a)
		if err != nil {
			panic(err)
		}
		return a
	}
}
