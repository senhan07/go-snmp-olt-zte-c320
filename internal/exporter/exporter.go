package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// OnuInfoGauge provides basic information about the ONU.
	OnuInfoGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_info",
			Help: "Information about the ZTE ONU device.",
		},
		[]string{"board", "pon", "onu_id", "name", "serial_number", "onu_type", "description", "ip_address", "offline_reason", "status"},
	)

	// OnuRxPowerGauge shows the received optical power of the ONU.
	OnuRxPowerGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_rx_power_dbm",
			Help: "The received optical power of the ONU in dBm.",
		},
		[]string{"board", "pon", "onu_id"},
	)

	// OnuTxPowerGauge shows the transmitted optical power of the ONU.
	OnuTxPowerGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_tx_power_dbm",
			Help: "The transmitted optical power of the ONU in dBm.",
		},
		[]string{"board", "pon", "onu_id"},
	)

	// OnuUptimeGauge shows the uptime of the ONU in seconds.
	OnuUptimeGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_uptime_seconds",
			Help: "The uptime of the ONU in seconds.",
		},
		[]string{"board", "pon", "onu_id"},
	)

	// OnuLastDownDurationGauge shows the duration of the last downtime in seconds.
	OnuLastDownDurationGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_last_down_duration_seconds",
			Help: "The duration of the last downtime in seconds.",
		},
		[]string{"board", "pon", "onu_id"},
	)

	// OnuLastOnlineGauge shows the last online timestamp as a Unix epoch.
	OnuLastOnlineGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_last_online_timestamp_seconds",
			Help: "The last online timestamp of the ONU as a Unix epoch.",
		},
		[]string{"board", "pon", "onu_id"},
	)

	// OnuLastOfflineGauge shows the last offline timestamp as a Unix epoch.
	OnuLastOfflineGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_last_offline_timestamp_seconds",
			Help: "The last offline timestamp of the ONU as a Unix epoch.",
		},
		[]string{"board", "pon", "onu_id"},
	)

	// OnuGponOpticalDistanceGauge shows the GPON optical distance in meters.
	OnuGponOpticalDistanceGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zte_onu_gpon_optical_distance_meters",
			Help: "The GPON optical distance to the ONU in meters.",
		},
		[]string{"board", "pon", "onu_id"},
	)
)