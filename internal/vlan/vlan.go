package vlan

import (
	vlanproto "code-challenge/protos"
	"strings"
)

var vlans []*vlanproto.VLAN

func Save(id, vlan string) {
	if len(vlans) == 0 {
		vlans = append(vlans, &vlanproto.VLAN{Id: id, Vlan: vlan})
		return
	}

	for i := 0; i < len(vlans); i++ {
		if strings.Compare(id, vlans[i].Id) <= 0 {
			vlans = append(vlans, nil)
			copy(vlans[i+1:], vlans[i:])
			vlans[i] = &vlanproto.VLAN{Id: id, Vlan: vlan}
			return
		}
	}

	vlans = append(vlans, &vlanproto.VLAN{Id: id, Vlan: vlan})
}

func VLANs() []*vlanproto.VLAN {
	return vlans
}
