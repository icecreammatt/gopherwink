#!/bin/bash
# Send message to Slack channel

channelid=$SLACK_CHANNEL_ID
SLACK_TOKEN=$SLACK_API_TOKEN
message=${*:1}

echo "Sending message [ $message ]"

postresp=$(curl -sSLk -X POST \
--data-urlencode "token=${SLACK_TOKEN}" \
--data-urlencode "channel=${channelid}" \
--data-urlencode "text=${message}" \
--data-urlencode "as_user=1" \
    https://slack.com/api/chat.postMessage)

echo $postresp
