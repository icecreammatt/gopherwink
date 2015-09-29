#!/bin/sh
# Monitor messages for zwave security devices
tail -Fn0 /var/log/messages | while read line ; do
    echo "$line" | grep "Received zwave General event UPDATE callback for node"
    if [ $? = 0 ]
    then
    # Actions
    res=$(echo "Logging: $line" | grep -Eo "node [0-9]")
    message="`date` zwave device triggered: $res"
    /root/monitor/slack-notify.sh ${message}
    fi
done
