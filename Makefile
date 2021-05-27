all:	build

build:
	@echo Building...
	go get -v .
	strip bgw210
	cp bgw210 files/
	docker build -t bgw .

go:
	@echo Compiling go tool...
	go build -v .

run:
	docker run -d -p 9222:9222 --rm --name bgw --shm-size 2G bgw

kill: stop

stop:
	docker stop bgw

