package model

// OltConfig struct is a struct that represent the OLT configuration
type OltConfig struct {
	BaseOID                   string
	OnuIDNameOID              string
	OnuTypeOID                string
	OnuSerialNumberOID        string
	OnuRxPowerOID             string
	OnuTxPowerOID             string
	OnuStatusOID              string
	OnuPhaseStateOID          string
	OnuIPAddressOID           string
	OnuDescriptionOID         string
	OnuLastOnlineOID          string
	OnuLastOfflineOID         string
	OnuLastOfflineReasonOID   string
	OnuGponOpticalDistanceOID string
}

// ONUInfo struct is a struct that represent the ONU information
type ONUInfo struct {
	ID   string `json:"onu_id"`
	Name string `json:"name"`
}

// ONUInfoPerBoard struct is a struct that represent the ONU information per board
type ONUInfoPerBoard struct {
	Board        int    `json:"board"`
	PON          int    `json:"pon"`
	ID           int    `json:"onu_id"`
	Name         string `json:"name"`
	OnuType      string `json:"onu_type"`
	SerialNumber string `json:"serial_number"`
	RXPower      string `json:"rx_power"`
	Status       string `json:"status"`
}

// ONUCustomerInfo struct is a struct that represent the detailed ONU information for customer
type ONUCustomerInfo struct {
	Board                int    `json:"board"`
	PON                  int    `json:"pon"`
	ID                   int    `json:"onu_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	OnuType              string `json:"onu_type"`
	SerialNumber         string `json:"serial_number"`
	RXPower              string `json:"rx_power"`
	TXPower              string `json:"tx_power"`
	Status               string `json:"status"`
	IPAddress            string `json:"ip_address"`
	LastOnline           string `json:"last_online"`
	LastOffline          string `json:"last_offline"`
	Uptime               string `json:"uptime"`
	LastDownTimeDuration string `json:"last_down_time_duration"`
	LastOfflineReason    string `json:"offline_reason"`
	GponOpticalDistance  string `json:"gpon_optical_distance"`
	PhaseState           string `json:"phase_state"`
}

// OnuID struct is a struct that represent the ONU ID
type OnuID struct {
	Board int `json:"board"`
	PON   int `json:"pon"`
	ID    int `json:"onu_id"`
}

// OnuOnlyID struct is a struct that represent only the ONU ID without board and PON
type OnuOnlyID struct {
	ID int `json:"onu_id"`
}

// SNMPWalkTask struct is a struct that represent the SNMP walk task
type SNMPWalkTask struct {
	BaseOID   string
	TargetOID string
	BoardID   int
	PON       int
}

// OnuSerialNumber struct is a struct that represent the ONU serial number
type OnuSerialNumber struct {
	Board        int    `json:"board"`
	PON          int    `json:"pon"`
	ID           int    `json:"onu_id"`
	SerialNumber string `json:"serial_number"`
}

// PaginationResult struct is a struct that represent the pagination result
type PaginationResult struct {
	OnuInformationList []ONUInfoPerBoard
	Count              int
}
