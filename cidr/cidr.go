package cidr

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	ipv4Reg = regexp.MustCompile(`^((2([0-4][0-9]|5[0-5]))|(1?[0-9]{1,2}))(\.((2([0-4][0-9]|5[0-5]))|(1?[0-9]{1,2}))){3}(/((3[0-2])|([1-2]?[0-9])))?$`)
)

// 10.0.0.1
// 10.0.0.1/8
func IsCIDR(n string) bool {
	return ipv4Reg.MatchString(n)
}

type Subnet struct {
	Net  uint32 // net num
	Mask uint32 // net mask
}

func (sn *Subnet) Include(ip string) bool {
	return sn.Net == IP2u32(ip)&sn.Mask
}

func CIDR2Subnet(n string) Subnet {
	var mask uint32
	ps := strings.Split(n, "/")
	if len(ps) > 1 {
		m, _ := strconv.Atoi(ps[1])
		mask = (^uint32(0)) << uint32(m)
	} else {
		mask = ^uint32(0)
	}

	return Subnet{
		Net:  IP2u32(ps[0]) & mask,
		Mask: mask,
	}
}

func IP2u32(ip string) uint32 {
	ps := strings.Split(ip, ".")
	p0, _ := strconv.Atoi(ps[0])
	p1, _ := strconv.Atoi(ps[1])
	p2, _ := strconv.Atoi(ps[2])
	p3, _ := strconv.Atoi(ps[3])
	return uint32(p0)<<24 | uint32(p1)<<16 | uint32(p2)<<8 | uint32(p3)
}
