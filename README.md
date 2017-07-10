# Prometheus Service Exporter

Prometheus Service Exporter is meant to be a way to monitor if a systemd service is active or not.

Useful for times when there isn't a full exporter for the service and you need to monitor if it is up or not.

# Setup

  - Download latest AMD64 compiled binary from realeases tab and place under /usr/local/bin/prometheus_service_exporter OR if you have a local Golang environment setup run ```Make linux``` to compile it yourself
  - Install the prometheus-service-exporter.service file under /etc/systemd/system/prometheus-service-exporter.service
  - Edit prometheus-service-exporter.service file to monitor your service "defaults to sshd.service"
  - Run ```systemctl daemon-reload``` and ```systemctl start prometheus-service-exporter.service```

# Test
Run ```curl localhost:8080/metrics```
```sh
# HELP service_up Check if the Service is up.
# TYPE service_up gauge
service_up 1
```
