FROM docker/whalesay:latest

RUN apt-get --ignore-missing -y update && apt-get install -y fortune

CMD /usr/games/fortune -a | cowsay