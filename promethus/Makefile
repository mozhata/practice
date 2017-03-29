docker_start:
	docker run --rm -p 9090:9090 \
	-v /home/kang/go/src/practice/promethus:/conf \
	-d prom/prometheus \
	-config.file=/conf/prometheus.yml \
	-storage.local.path=/prometheus \
	-web.console.libraries=/usr/share/prometheus/console_libraries \
	-web.console.templates=/usr/share/prometheus/consoles

docker_live:
	docker run -it -p 9090:9090 -v /home/kang/go/src/practice/promethus:/conf prom/prometheus -config.file=/conf/prometheus.yml -storage.local.path=/prometheus -web.console.libraries=/usr/share/prometheus/console_libraries -web.console.templates=/usr/share/prometheus/consoles -alertmanager.url=http://192.168.0.66:9093

alertmanager:
	alertmanager -config.file=./alertmanager.yml &

alertmanager_live:
	alertmanager -config.file=./alertmanager.yml
# start:
# 	prometheus -config.file=promethus.yml

random1:
	random -listen-address=:8080 &
random2:
	random -listen-address=:8081 &
random3:
	random -listen-address=:8082 &

random:
	make random1;make random2;make random3

docker_init:
	make docker_start;make node_exporter;make alertmanager


reload:
	curl -X POST http://localhost:9090/-/reload
