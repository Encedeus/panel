package validate

import (
	"net"
	"net/netip"
	"slices"
)

/*func IsIpAddress(addr string) bool {
	fmt.Println(addr)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf(""), 1*time.Second)
	fmt.Println("test")
	if err != nil {
		return false
	}
	_ = conn.Close()

	return true
}*/

func IsIpv4Address(addr string) bool {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return false
	}

	if !ip.Is4() {
		return false
	}

	return true
}

func IsIpv6Address(addr string) bool {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return false
	}

	if !ip.Is6() {
		return false
	}

	return true
}

func IsDomain(fqdn string, addr string) bool {
	addrs, err := net.LookupHost(fqdn)
	if err != nil {
		return false
	}
	if !slices.Contains(addrs, addr) {
		return false
	}

	return true
}
