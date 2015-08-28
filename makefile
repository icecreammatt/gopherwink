winkIP := 192.168.1.11
winkUser := root

build:
	GOOS=linux GOARCH=arm GOARM=5 /usr/local/go/bin/go build .

deploy:
	scp gopherwink ${winkUser}@${winkIP}:/root/gopherwink
