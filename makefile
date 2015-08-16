build:
	GOOS=linux GOARCH=arm GOARM=5 /usr/local/go/bin/go build .

deploy:
	scp light-control-server root@192.168.1.11:light-control-go1.5
