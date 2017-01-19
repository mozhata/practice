docker_start:
	docker run -p 9090:9090 -v /home/kang/go/src/practice/promethus/promethus.yml:/etc/prometheus/prometheus.yml -d prom/prometheus

live:
		docker run -p 9090:9090 --rm -v /home/kang/go/src/practice/promethus/promethus.yml:/etc/prometheus/prometheus.yml prom/prometheus

start:
	prometheus -config.file=promethus.yml
