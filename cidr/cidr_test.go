package cidr

import (
	"testing"
)

func TestCidr(t *testing.T) {
	ns := []string{
		"10.0.0.1",
		"1.0.0.1/23",
		"1.3.3.2/8",
		"192.168.50.12/8",
	}
	cs := make([]Subnet, 0, len(ns))
	for _, n := range ns {
		if !IsCIDR(n) {
			t.Fatal("invalid cidr:", n)
		}
		cs = append(cs, CIDR2Subnet(n))
	}

	for _, c := range cs {
		t.Logf("%x %b\n", c.Net, c.Mask)
	}

	ip := "192.168.50.1"

	t.Log(cs[3].Include(ip))
}
