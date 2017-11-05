#!/bin/sh
# Monitor messages for zwave security devices
tail -Fn0 /var/log/messages | while read line ; do
    echo "$line" | grep "Received zwave General event UPDATE callback for node"
    if [ $? = 0 ]
    then
    res=$(echo "Logging: $line" | grep -Eo "node [0-9]")
    idx=$(echo $res | grep -Eo "[0-9]")

    device=$(sqlite3 /database/apron.db "select masterId, userName from zwaveDevice, masterDevice where nodeId = $idx and deviceId=masterId")
    masterid=`echo $device | cut -d \| -f 1`
    userName=`echo $device | cut -d \| -f 2`

    data=$(aprontest -m $masterid -l)
    newState=$(echo "$data" | grep -Ew "On_Off" | awk '{ print $9 }')

    switchState=""
    if [ $newState = "TRUE" ]
    then
    switchState="OPENED"
    curl 192.168.40.18:8080/sound &
    else
    switchState="CLOSED"
    fi

    message="`date` $userName = $switchState"
    /root/monitor/slack-notify.sh ${message}
    fi
done
