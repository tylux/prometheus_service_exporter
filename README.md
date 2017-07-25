# Prometheus Service Exporter

Prometheus Service Exporter is meant to be a way to monitor if a Systemd service is active or not.

Useful for times when there isn't a full exporter for the service and you need to monitor if it is up or not.

# Setup

  - Download latest AMD64 compiled binary from realeases tab and place under /usr/local/bin/prometheus_service_exporter OR if you have a local Golang environment setup run ```Make linux``` to compile it yourself
  - Install the prometheus-service-exporter.service file under /etc/systemd/system/prometheus-service-exporter.service
  - Edit prometheus-service-exporter.service file to monitor your service, if you want to monitor multiple, use a comma-seperated list "defaults to sshd.service"
  - Run ```systemctl daemon-reload``` and ```systemctl start prometheus-service-exporter.service```

# Test and Example output
Run ```curl localhost:9199/metrics```
```sh
# HELP service_up Is the service active
# TYPE service_up gauge
service_up{service="apache2"} 0
service_up{service="kibana"} 1
service_up{service="nginx"} 1
service_up{service="sshd"} 1
```
