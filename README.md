# GopherWink

This is server that runs on [Wink Hub](http://www.wink.com/products/wink-hub/) written in Go  

Currently the only supported devices are lights. This is a very early work in progress. This is not intended to be a replacement for OpenHAB. I just thought this would be fun to create.

## Current Features

* Add new lights
* List lights
* Turn lights on and off
* Adjust brightness

## Setup

1. [Root Winkhub]()
2. `scp light-control-server root@<WINKHUB_IP>:`
3. `ssh root<WINKHUB_IP> "./light-control-server &"`
4. Visit `http://WINKHUB_IP:3000/` in the browser.

## Future Features

* Light sleep timer
* Support for door sensors
* Clock control (Daylight based settings)

## TODO
* Need to fix error:

```
Error: unexpected end of JSON input
value body: {"id":1,"active":true,"value":255}
value body:
Error: unexpected end of JSON input
value body: {"id":2,"active":true,"value":252}
```

* Add install script to place files into the proper locations
* Add service as part of init scripts so it does not have to manually be started over SSH
* Add ability to remove existing lights

## License
