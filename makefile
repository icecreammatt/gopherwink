winkUser := root

build:
	GOOS=linux GOARCH=arm GOARM=5 /usr/local/go/bin/go build .

.PHONY: deploy install

deploy:
	-ssh ${winkUser}@$(WINK_IP_ADDRESS) "/etc/init.d/S63gopherwink stop"
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwink

install:
	scp ./S63gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/etc/init.d/S63gopherwink
	ssh ${winkUser}@$(WINK_IP_ADDRESS) "chmod 755 /etc/init.d/S63gopherwink"
	-ssh ${winkUser}@$(WINK_IP_ADDRESS) "/etc/init.d/S63gopherwink stop"
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwink
