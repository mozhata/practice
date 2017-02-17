docker_start:
	docker run --rm -p 9090:9090 -v /home/kang/go/src/practice/promethus/promethus.yml:/etc/prometheus/prometheus.yml -d prom/prometheus

live:
		docker run --rm -p 9090:9090 --rm -v /home/kang/go/src/practice/promethus/promethus.yml:/etc/prometheus/prometheus.yml prom/prometheus

start:
	prometheus -config.file=promethus.yml

random1:
	random -listen-address=:8080 &
random2:
	random -listen-address=:8081 &
random3:
	random -listen-address=:8082 &

random:
	make random1;make random2;make random3

docker_init:
	make docker_start;make random;
