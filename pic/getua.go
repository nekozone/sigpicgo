package pic

import (
	"fmt"

	"github.com/mileusna/useragent"
)

func Getua(s string) string {
	ua := useragent.Parse(s)
	if ua.Device == "" {
		ua.Device = "PC"
	}
	res := fmt.Sprintf("浏览器: %s %s OS: %s %s 设备: %s", ua.Name, ua.Version, ua.OS, ua.OSVersion, ua.Device)
	return res
}
