winkUser := root

build:
	GOOS=linux GOARCH=arm GOARM=5 go build .

.PHONY: deploy install release

debug:
	GOOS=linux GOARCH=arm GOARM=5 go build .
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwinkdebug
	scp settings.ini ${winkUser}@$(WINK_IP_ADDRESS):/root/settings.ini
	scp client.key ${winkUser}@$(WINK_IP_ADDRESS):/root/client.key
	scp client.pem ${winkUser}@$(WINK_IP_ADDRESS):/root/client.pem

deploy: build
	-ssh ${winkUser}@$(WINK_IP_ADDRESS) "/etc/init.d/S63gopherwink stop"
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwink
	scp settings.ini ${winkUser}@$(WINK_IP_ADDRESS):/root/settings.ini

install: build
	scp ./S63gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/etc/init.d/S63gopherwink
	ssh ${winkUser}@$(WINK_IP_ADDRESS) "chmod 755 /etc/init.d/S63gopherwink"
	-ssh ${winkUser}@$(WINK_IP_ADDRESS) "/etc/init.d/S63gopherwink stop"
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwink
	scp settings.ini ${winkUser}@$(WINK_IP_ADDRESS):/root/settings.ini

release:
	rm -rf release
	mkdir release
	cp ./gopherwink release/
	cp ./frontend/dist/index.html release/
	cp ./frontend/dist/assets/main.js release/
	cp ./S63gopherwink release/
	tar cfvz release.tgz ./release
	zip -r release.zip ./release

