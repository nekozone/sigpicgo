package cip

import (
	"database/sql"
)

func GetIP(db *sql.DB, dip string) (string, string) {
	if !IsIp(dip) {
		return "", ""
	}
	iptype := IsIPv6(dip)
	if iptype {
		ipnum := InetAtoN6(dip)
		w := db.QueryRow("select address,location from ipv6_range_info where " + ipnum + " between ip_start_num and ip_end_num")
		var address, location string
		w.Scan(&address, &location)
		return address, location
	} else {
		ipnum := InetAtoN4(dip)
		qu, err := db.Prepare("select address,location from iprange_info where ? between ip_start_num and ip_end_num")
		if err != nil {
			panic(err)
		}
		w := qu.QueryRow(ipnum)
		var address, location string
		w.Scan(&address, &location)
		return address, location
	}
}

func GetAllIP(db *sql.DB, dip string) map[string]interface{} {
	if !IsIp(dip) {
		return map[string]interface{}{"country": "", "province": "", "city": "", "area": "", "address": "", "location": ""}
	}
	iptype := IsIPv6(dip)
	if iptype {
		ipnum := InetAtoN6(dip)
		w := db.QueryRow("select country,province,city,area,address,location from ipv6_range_info where " + ipnum + " between ip_start_num and ip_end_num")
		var country, province, city, area, address, location string
		w.Scan(&country, &province, &city, &area, &address, &location)
		if location == " CZ88.NET" {
			location = ""
		}
		return map[string]interface{}{"country": country, "province": province, "city": city, "area": area, "address": address, "location": location}
	} else {
		ipnum := InetAtoN4(dip)
		// qu, err := db.Prepare("select country,province,area,address,location from iprange_info where ? between ip_start_num and ip_end_num")
		// if err != nil {
		// 	panic(err)
		// }
		// w := qu.QueryRow(ipnum)
		// var country, province, city, area, address, location string
		qu, err := db.Prepare("select country,province,city,area,address,location from iprange_info where ? between ip_start_num and ip_end_num")
		if err != nil {
			panic(err)
		}
		w := qu.QueryRow(ipnum)
		var country, province, city, area, address, location string
		w.Scan(&country, &province, &city, &area, &address, &location)
		if location == " CZ88.NET" {
			location = ""
		}
		return map[string]interface{}{"country": country, "province": province, "city": city, "area": area, "address": address, "location": location}
	}
}
