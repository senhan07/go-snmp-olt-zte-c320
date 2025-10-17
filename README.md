# ZTE C320 OLT SNMP Exporter for Prometheus

This application is a dedicated Prometheus exporter that collects metrics from a ZTE C320 OLT (Optical Line Terminal) via SNMP. It discovers Optical Network Units (ONUs) across configured boards and PONs, and exposes their operational data as Prometheus metrics.

## Features

-   **ONU Discovery**: Automatically scans a configurable range of boards and PONs to find active ONUs.
-   **Rich Metrics**: Exposes a wide range of metrics, including signal strength (Rx/Tx power), uptime, and status.
-   **Prometheus Integration**: Provides a standard `/metrics` endpoint for Prometheus to scrape.
-   **Dockerized**: Comes with a `Dockerfile` for easy deployment.

## Configuration

The exporter is configured through environment variables. The following variables are available:

| Environment Variable        | Description                                                                 | Default |
| --------------------------- | --------------------------------------------------------------------------- | ------- |
| `APP_ENV`                   | Set the application environment (`dev` or `prod`).                          | `dev`     |
| `PROMETHEUS_BOARD_MIN`      | The starting board ID to scan for ONUs.                                     | `1`     |
| `PROMETHEUS_BOARD_MAX`      | The ending board ID to scan for ONUs.                                       | `2`     |
| `PROMETHEUS_PON_MIN`        | The starting PON ID to scan for ONUs.                                       | `1`     |
| `PROMETHEUS_PON_MAX`        | The ending PON ID to scan for ONUs.                                         | `16`    |
| `SNMP_IP`                   | The IP address of the ZTE C320 OLT.                                         |         |
| `SNMP_PORT`                 | The SNMP port of the OLT.                                                   | `161`   |
| `SNMP_COMMUNITY`            | The SNMP community string.                                                  |         |
| `REDIS_HOST`                | The hostname of the Redis server.                                           |         |
| `REDIS_PORT`                | The port of the Redis server.                                               | `6379`  |
| `REDIS_PASSWORD`            | The password for the Redis server.                                          |         |
| `REDIS_DB`                  | The Redis database to use.                                                  | `0`     |


## Exposed Metrics

The following Prometheus metrics are exposed on the `/metrics` endpoint:

| Metric                          | Type  | Labels                                                                                                                              | Description                                                                 |
| ------------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `zte_onu_info`                  | Gauge | `board`, `pon`, `onu_id`, `name`, `serial_number`, `onu_type`, `description`, `ip_address`, `offline_reason`, `status`               | A constant value of `1` with labels providing detailed information about the ONU. |
| `zte_onu_rx_power_dbm`          | Gauge | `board`, `pon`, `onu_id`                                                                                                            | The received optical power of the ONU in dBm.                               |
| `zte_onu_tx_power_dbm`          | Gauge | `board`, `pon`, `onu_id`                                                                                                            | The transmitted optical power of the ONU in dBm.                            |
| `zte_onu_uptime_seconds`        | Gauge | `board`, `pon`, `onu_id`                                                                                                            | The uptime of the ONU in seconds.                                           |
| `zte_onu_last_down_duration_seconds` | Gauge | `board`, `pon`, `onu_id`                                                                                                            | The duration of the last downtime for the ONU in seconds.                   |
| `zte_onu_last_online_timestamp` | Gauge | `board`, `pon`, `onu_id`                                                                                                            | The timestamp of the last time the ONU was online.                          |
| `zte_onu_last_offline_timestamp`| Gauge | `board`, `pon`, `onu_id`                                                                                                            | The timestamp of the last time the ONU was offline.                         |
| `zte_onu_gpon_optical_distance_meters` | Gauge | `board`, `pon`, `onu_id`                                                                                                            | The optical distance of the ONU from the OLT in meters.                     |

**Note**: Power metrics (`zte_onu_rx_power_dbm` and `zte_onu_tx_power_dbm`) are only reported for ONUs with a status of `Online`.

## How to Run

### Using Docker

The recommended way to run the exporter is with Docker. A `Dockerfile` is provided to build the image.

1.  **Build the Docker image:**

    ```sh
    docker build -t zte-olt-exporter .
    ```

2.  **Run the Docker container:**

    Replace the placeholder values with your OLT and Redis connection details.

    ```sh
    docker run -d \
      -p 8081:8081 \
      --name zte-olt-exporter \
      -e APP_ENV=prod \
      -e SNMP_IP=<your_olt_ip> \
      -e SNMP_COMMUNITY=<your_snmp_community> \
      -e REDIS_HOST=<your_redis_host> \
      -e REDIS_PASSWORD=<your_redis_password> \
      zte-olt-exporter
    ```

3.  **Verify the exporter is running:**

    You can check the logs of the container to ensure it started correctly:

    ```sh
    docker logs zte-olt-exporter
    ```

    You should also be able to access the metrics at `http://localhost:8081/metrics`.