# GopherWink

This is controller for the [Wink Hub](http://www.wink.com/products/wink-hub/) written in Go  

Currently the only supported devices are lights

## Light Features

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

* Sleep timer
* Clock control (Daylight based settings)

## License
