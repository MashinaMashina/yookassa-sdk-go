package notification

import (
	"net/netip"
)

var YooKassaIPs = []string{
	"185.71.76.0/27",
	"185.71.77.0/27",
	"77.75.153.0/25",
	"77.75.156.11",
	"77.75.156.35",
	"77.75.154.128/25",
	"2a02:5180::/32",
}

type Authenticity struct {
	networks []netip.Prefix
}

func InitAuthenticity(ips []string) (*Authenticity, error) {
	a := &Authenticity{
		networks: make([]netip.Prefix, 0, len(ips)),
	}

	for _, ip := range ips {
		network, err := netip.ParsePrefix(ip)
		if err != nil {
			return nil, err
		}

		a.networks = append(a.networks, network)
	}

	return a, nil
}

func (r *Authenticity) Allowed(userIP string) (bool, error) {
	ip, err := netip.ParseAddr(userIP)
	if err != nil {
		return false, err
	}

	for _, network := range r.networks {
		if network.Contains(ip) {
			return true, nil
		}
	}

	return false, nil
}
