package fakegeo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var isp = []string{"联通", "电信", "移动"}

var geodata map[string]interface{}

func init() {
	filedog, _ := os.ReadFile("./fakegeo/fakegeo.json")
	jsonstr := string(filedog)
	err := json.Unmarshal([]byte(jsonstr), &geodata)
	if err != nil {
		panic(err)
	}
}

func Fakeip() map[string]string {
	// println("Hello, world!")
	// fmt.Println(geodata)
	rand.Seed(time.Now().UnixNano())
	c := geodata["num"].(float64)
	fmt.Println(c)
	data := geodata["data"].([]interface{})
	// m := data[4]
	m := rand.Intn(int(c))
	// fmt.Println(data[m])
	R := data[m].(map[string]interface{})
	// fmt.Println(R["provice"])
	adress := R["province"].(string) + R["city"].(string) + R["area"].(string)
	Isp := isp[rand.Intn(3)]

	// // res:={"provice":R["province"],"city":R["city"],"area":R["area"],"address":adress,"location":Isp}
	res := map[string]string{"province": R["province"].(string), "city": R["city"].(string), "country": "中国", "area": R["area"].(string), "address": adress, "location": Isp}
	return res
	// fmt.Println(res)
}
