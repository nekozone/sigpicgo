package cip

import (
	"fmt"
	"math/big"
	"net"
	"strings"
)

func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

func InetAtoN4(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}
func InetAtoN6(ipd string) string {
	ip := net.ParseIP(ipd)
	d := fmt.Sprintf("x'%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x'",
		ip[0], ip[1], ip[2], ip[3], ip[4], ip[5], ip[6], ip[7], ip[8], ip[9], ip[10], ip[11], ip[12], ip[13], ip[14], ip[15])
	return d
}

func Ip42index(ip string) string {
	nip := strings.ReplaceAll(ip, ".", "-")
	return fmt.Sprintf("ip_%s", nip)
}

func Ip62index(ip string) string {
	ipd := net.ParseIP(ip)
	return fmt.Sprintf("i6p_%02x%02x%02x%02x%02x%02x%02x%02x", ipd[0], ipd[1], ipd[2], ipd[3], ipd[4], ipd[5], ipd[6], ipd[7])
}

func Ip2index(ip string) string {
	if IsIPv6(ip) {
		return Ip62index(ip)
	} else {
		return Ip42index(ip)
	}
}

func IsIp(ip string) bool {
	return net.ParseIP(ip) != nil
}
