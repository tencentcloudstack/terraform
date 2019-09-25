package tencentcloud

/*
 all gate way types
 https://cloud.tencent.com/document/api/215/15824#Route
*/
const GATE_WAY_TYPE_CVM = "CVM"
const GATE_WAY_TYPE_VPN = "VPN"
const GATE_WAY_TYPE_DIRECTCONNECT = "DIRECTCONNECT"
const GATE_WAY_TYPE_SSLVPN = "SSLVPN"
const GATE_WAY_TYPE_NAT = "NAT"
const GATE_WAY_TYPE_NORMAL_CVM = "NORMAL_CVM"
const GATE_WAY_TYPE_EIP = "EIP"
const GATE_WAY_TYPE_CCN = "CCN"

var ALL_GATE_WAY_TYPES = []string{GATE_WAY_TYPE_CVM,
	GATE_WAY_TYPE_VPN,
	GATE_WAY_TYPE_DIRECTCONNECT,
	GATE_WAY_TYPE_SSLVPN,
	GATE_WAY_TYPE_NAT,
	GATE_WAY_TYPE_NORMAL_CVM,
	GATE_WAY_TYPE_EIP,
	GATE_WAY_TYPE_CCN,
}

/*
EIP
*/
const (
	EIP_STATUS_CREATING  = "CREATING"
	EIP_STATUS_BINDING   = "BINDING"
	EIP_STATUS_BIND      = "BIND"
	EIP_STATUS_UNBINDING = "UNBINDING"
	EIP_STATUS_UNBIND    = "UNBIND"
	EIP_STATUS_OFFLINING = "OFFLINING"
	EIP_STATUS_BIND_ENI  = "BIND_ENI"
)
