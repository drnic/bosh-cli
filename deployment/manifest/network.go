package manifest

import (
	bmproperty "github.com/cloudfoundry/bosh-micro-cli/common/property"
)

type NetworkType string

func (n NetworkType) String() string {
	return string(n)
}

const (
	Dynamic NetworkType = "dynamic"
	Manual  NetworkType = "manual"
	VIP     NetworkType = "vip"
)

type Network struct {
	Name            string
	Type            NetworkType
	CloudProperties bmproperty.Map
	IP              string
	Netmask         string
	Gateway         string
	DNS             []string
}

// Interface returns a property map representing a generic network interface.
// Expected Keys: ip, type, cloud properties.
// Optional Keys: netmask, gateway, dns
func (n Network) Interface() bmproperty.Map {
	iface := bmproperty.Map{
		"type":             n.Type.String(),
		"ip":               n.IP,
		"cloud_properties": n.CloudProperties,
	}

	if n.Netmask != "" {
		iface["netmask"] = n.Netmask
	}

	if n.Gateway != "" {
		iface["gateway"] = n.Gateway
	}

	if len(n.DNS) > 0 {
		iface["dns"] = n.DNS
	}

	return iface
}