winkUser := root

build:
	GOOS=linux GOARCH=arm GOARM=5 /usr/local/go/bin/go build .

.PHONY: deploy install release

deploy:
	-ssh ${winkUser}@$(WINK_IP_ADDRESS) "/etc/init.d/S63gopherwink stop"
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwink

install:
	scp ./S63gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/etc/init.d/S63gopherwink
	ssh ${winkUser}@$(WINK_IP_ADDRESS) "chmod 755 /etc/init.d/S63gopherwink"
	-ssh ${winkUser}@$(WINK_IP_ADDRESS) "/etc/init.d/S63gopherwink stop"
	scp gopherwink ${winkUser}@$(WINK_IP_ADDRESS):/root/gopherwink

release:
	rm -rf release
	mkdir release
	cp ./gopherwink release/
	cp ./frontend/dist/index.html release/
	cp ./frontend/dist/assets/main.js release/
	cp ./S63gopherwink release/
	tar cfvz release.tgz ./release
	zip -r release.zip ./release

