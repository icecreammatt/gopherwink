# GopherWink

This a is server that runs on the [Wink Hub](http://www.wink.com/products/wink-hub/) written in Go.  

Currently the only supported devices are lights. This is a very early work in progress. This is not intended to be a replacement for OpenHAB. I just thought this would be fun to create.

## Current Features
* Add new lights
* List lights
* Turn lights on and off
* Adjust brightness
* Adjust Winkhub LED colors

## GopherWink setup instructions:

### Building from source requirements:

* [Rooted Winkhub](http://localhost:1313/post/hacking-the-winkhub-part-1/) with SSH access
* Go 1.5
* NodeJS
* ReactJS

### Install from source:

* `export WINK_IP_ADDRESS=X.X.X.X` (replace the x's with the Wink IP)
* `git clone https://github.com/icecreammatt/gopherwink`
* `git submodule init && git submodule update`
* `cd frontend && npm install`
* `make deploy`
* `cd ..`
* `make build`
* `make install`
* Visit `http://WINK_IP_ADDRESS/index.html` to access the controls.

### Install release manually

* Download the release from [here](https://github.com/icecreammatt/gopherwink/releases)
* Extract the zip file
* `export WINK_IP_ADDRESS=X.X.X.X` (replace the x's with the Wink IP)
* `scp gopherwink root@$WINK_IP_ADDRESS:/root/gopherwink`
* `scp S63gopherwink root@$WINK_IP_ADDRESS:/etc/init.d/S63gopherwink`
* `scp index.html root@$WINK_IP_ADDRESS:/var/www`
* `ssh root@$WINK_IP_ADDRESS "mkdir /var/www/assets`
* `scp main.js root@$WINK_IP_ADDRESS:/var/www/assets/`
* `ssh root@$WINK_IP_ADDRESS "/etc/init.d/S63gopherwink start"`
* Visit `http://WINK_IP_ADDRESS/index.html` to access the controls.

## Usage notes
* New devices for now need to be connected using `aprontest` or the WinkApp
* Zigbee lights can be added by visiting `http://WINK_IP_ADDRESS:5000/light/search`

## Todo
* Add install script to place files into the proper locations
* Add ability to remove existing lights

## Bugs
* File bugs [here](https://github.com/icecreammatt/gopherwink/issues)
* Need to fix error:

```
Error: unexpected end of JSON input
value body: {"id":1,"active":true,"value":255}
value body:
Error: unexpected end of JSON input
value body: {"id":2,"active":true,"value":252}
```

## Future Plans
* Improved UI to add and remove lights
* Improved UI to rename devices
* Refactor API to be RESTful
* Add Proper support for GoControl Door Window Sensors
* Sleep timer to keep light on for late nights
* Snooze timer to turn on light after x minutes
* Automatic brightness based on the time of day
* TLS Authentication for API

## License

GPLv3
