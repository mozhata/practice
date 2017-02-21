# docker_start:
	# docker run --rm -p 9090:9090 -v /home/kang/go/src/practice/promethus:/data -v /home/kang/go/src/practice/promethus/promethus.yml:/etc/prometheus/prometheus.yml -v /etc/localtime:/etc/localtime -d prom/prometheus

# docker_start:
	# docker run --rm -p 9090:9090 -v /home/kang/go/src/practice/promethus:/data -d prom/prometheus -config.file=/data/prometheus.yml -storage.local.path=/prometheus -web.console.libraries=/usr/share/prometheus/console_libraries -web.console.templates=/usr/share/prometheus/consoles -alertmanager.url=192.168.0.66:9093
docker_start:
	docker run -it -p 9090:9090 -v /home/kang/go/src/practice/promethus:/data prom/prometheus -config.file=/data/prometheus.yml -storage.local.path=/prometheus -web.console.libraries=/usr/share/prometheus/console_libraries -web.console.templates=/usr/share/prometheus/consoles -alertmanager.url=http://192.168.0.66:9093

# live:
# 		docker run --rm -p 9090:9090 --rm -v /home/kang/go/src/practice/promethus/promethus.yml:/etc/prometheus/prometheus.yml prom/prometheus

node_exporter:
	node_exporter &

alertmanager:
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
	make docker_start;make random;make node_exporter


reload:
	curl -X POST http://localhost:9090/-/reload