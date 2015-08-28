winkIP := 192.168.1.11
winkUser := root

build:
	GOOS=linux GOARCH=arm GOARM=5 /usr/local/go/bin/go build .

.PHONY: deploy install

deploy:
	scp gopherwink ${winkUser}@${winkIP}:/root/gopherwink

install:
	scp gopherwink ${winkUser}@${winkIP}:/root/gopherwink
	scp ./S63gopherwink ${winkUser}@${winkIP}:/etc/init.d/S63gopherwink
	ssh ${winkUser}@${winkIP} "chmod 755 /etc/init.d/S63gopherwink"
