package api

import (
	"net"
	"net/http"
	"strings"
)

// IPExtractor provides ability to extract clients ip address from the http.Request
type IPExtractor struct {
	privatesubnets []*net.IPNet
}

// NewIPExtractorT instantiates new extractor
func NewIPExtractorT() IPExtractor {
	return IPExtractor{
		[]*net.IPNet{
			getsubnet("127.0.0.1/8"),
			getsubnet("10.0.0.0/8"),
			getsubnet("172.16.0.0/12"),
			getsubnet("192.168.0.0/16"),
			getsubnet("169.254.0.0/16"),
			getsubnet("::1/128"),
			getsubnet("fc00::/7"),
			getsubnet("fe80::/10"),
		},
	}
}

func (extractor *IPExtractor) insideprivate(ipaddr net.IP) bool {
	for i := range extractor.privatesubnets {
		if extractor.privatesubnets[i].Contains(ipaddr) {
			return true
		}
	}
	return false
}

func getsubnet(cidr string) *net.IPNet {
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}
	return subnet
}

// Extract real ipaddress form http.Request's "X-Forwarded-For", "X-Real-Ip" headers
// Falls back to the http.Request's RemoteAddr in case not being behind reverse proxy
func (extractor *IPExtractor) Extract(r *http.Request) (string, error) {
	for _, header := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		ipss := strings.Split(r.Header.Get(header), ",")
		for i := len(ipss) - 1; i >= 0; i-- {
			ips := strings.TrimSpace(ipss[i])
			ip := net.ParseIP(ips)
			if ip == nil || !ip.IsGlobalUnicast() ||
				extractor.insideprivate(ip) {
				continue
			}
			return ips, nil
		}
	}
	ips, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ips, nil
}
