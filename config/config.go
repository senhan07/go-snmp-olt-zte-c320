package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Config represents the main application configuration structure
// that contains all sub-configurations for SNMP, Redis, OLT, and individual PON boards.
type Config struct {
	SnmpCfg     SnmpConfig
	RedisCfg    RedisConfig
	OltCfg      OltConfig
	Board1Pon1  Board1Pon1
	Board1Pon2  Board1Pon2
	Board1Pon3  Board1Pon3
	Board1Pon4  Board1Pon4
	Board1Pon5  Board1Pon5
	Board1Pon6  Board1Pon6
	Board1Pon7  Board1Pon7
	Board1Pon8  Board1Pon8
	Board1Pon9  Board1Pon9
	Board1Pon10 Board1Pon10
	Board1Pon11 Board1Pon11
	Board1Pon12 Board1Pon12
	Board1Pon13 Board1Pon13
	Board1Pon14 Board1Pon14
	Board1Pon15 Board1Pon15
	Board1Pon16 Board1Pon16
	Board2Pon1  Board2Pon1
	Board2Pon2  Board2Pon2
	Board2Pon3  Board2Pon3
	Board2Pon4  Board2Pon4
	Board2Pon5  Board2Pon5
	Board2Pon6  Board2Pon6
	Board2Pon7  Board2Pon7
	Board2Pon8  Board2Pon8
	Board2Pon9  Board2Pon9
	Board2Pon10 Board2Pon10
	Board2Pon11 Board2Pon11
	Board2Pon12 Board2Pon12
	Board2Pon13 Board2Pon13
	Board2Pon14 Board2Pon14
	Board2Pon15 Board2Pon15
	Board2Pon16 Board2Pon16
}

// SnmpConfig contains configuration parameters for SNMP connection
// including target IP address, port, and community string.
type SnmpConfig struct {
	IP        string `mapstructure:"ip"` // Target IP address of the SNMP device
	Port      uint16 `mapstructure:"port"`
	Community string `mapstructure:"community"`
}

// RedisConfig contains configuration parameters for Redis connection
// including host, port, authentication, and connection pooling settings.
type RedisConfig struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Password           string `mapstructure:"password"`
	DB                 int    `mapstructure:"db"`
	DefaultDB          int    `mapstructure:"default_db"`
	MinIdleConnections int    `mapstructure:"min_idle_connections"`
	PoolSize           int    `mapstructure:"pool_size"`
	PoolTimeout        int    `mapstructure:"pool_timeout"`
}

// OltConfig contains base OID configurations for OLT device management
// including common OIDs for ONU identification and type mapping.
type OltConfig struct {
	BaseOID1        string `mapstructure:"base_oid_1"`
	BaseOID2        string `mapstructure:"base_oid_2"`
	OnuIDNameAllPon string `mapstructure:"onu_id_name"`
	OnuTypeAllPon   string `mapstructure:"onu_type"`
}

// Board1Pon1 contains OID configurations for Board 1 Port 1 ONU management
// including identifiers, status, power levels, and diagnostic information.
type Board1Pon1 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon2 contains OID configurations for Board 1 Port 2 ONU management.
type Board1Pon2 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon3 contains OID configurations for Board 1 Port 3 ONU management.
type Board1Pon3 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon4 contains OID configurations for Board 1 Port 4 ONU management.
type Board1Pon4 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon5 contains OID configurations for Board 1 Port 5 ONU management.
type Board1Pon5 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon6 contains OID configurations for Board 1 Port 6 ONU management.
type Board1Pon6 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon7 contains OID configurations for Board 1 Port 7 ONU management.
type Board1Pon7 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon8 contains OID configurations for Board 1 Port 8 ONU management.
type Board1Pon8 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon9 contains OID configurations for Board 1 Port 9 ONU management.
type Board1Pon9 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon10 contains OID configurations for Board 1 Port 10 ONU management.
type Board1Pon10 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon11 contains OID configurations for Board 1 Port 11 ONU management.
type Board1Pon11 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon12 contains OID configurations for Board 1 Port 12 ONU management.
type Board1Pon12 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon13 contains OID configurations for Board 1 Port 13 ONU management.
type Board1Pon13 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon14 contains OID configurations for Board 1 Port 14 ONU management.
type Board1Pon14 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon15 contains OID configurations for Board 1 Port 15 ONU management.
type Board1Pon15 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board1Pon16 contains OID configurations for Board 1 Port 16 ONU management.
type Board1Pon16 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon1 contains OID configurations for Board 2 Port 1 ONU management.
type Board2Pon1 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon2 contains OID configurations for Board 2 Port 2 ONU management.
type Board2Pon2 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon3 contains OID configurations for Board 2 Port 3 ONU management.
type Board2Pon3 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon4 contains OID configurations for Board 2 Port 4 ONU management.
type Board2Pon4 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon5 contains OID configurations for Board 2 Port 5 ONU management.
type Board2Pon5 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon6 contains OID configurations for Board 2 Port 6 ONU management.
type Board2Pon6 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon7 contains OID configurations for Board 2 Port 7 ONU management.
type Board2Pon7 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon8 contains OID configurations for Board 2 Port 8 ONU management.
type Board2Pon8 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon9 contains OID configurations for Board 2 Port 9 ONU management.
type Board2Pon9 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon10 contains OID configurations for Board 2 Port 10 ONU management.
type Board2Pon10 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon11 contains OID configurations for Board 2 Port 11 ONU management.
type Board2Pon11 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon12 contains OID configurations for Board 2 Port 12 ONU management.
type Board2Pon12 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon13 contains OID configurations for Board 2 Port 13 ONU management.
type Board2Pon13 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon14 contains OID configurations for Board 2 Port 14 ONU management.
type Board2Pon14 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon15 contains OID configurations for Board 2 Port 15 ONU management.
type Board2Pon15 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// Board2Pon16 contains OID configurations for Board 2 Port 16 ONU management.
type Board2Pon16 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuPhaseStateOID          string `mapstructure:"onu_phase_state"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// LoadConfig file from given path using viper
func LoadConfig(filename string) (*Config, error) {

	// Initialize viper
	v := viper.New()

	// Set config file name
	v.SetConfigName(filename)

	// Set config path in current directory
	v.AddConfigPath(".")

	// Allow environment variables to override config
	v.AutomaticEnv()

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError // Initialize config file not found error
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var cfg Config // Initialize config variable

	// Unmarshal config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
