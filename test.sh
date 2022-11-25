#!/bin/bash

for((i=101;i<=150;i++))
do
title="test-${i}"
echo $title

curl --request POST \
  --url https://api.pagerduty.com/incidents \
  --header 'Accept: application/vnd.pagerduty+json;version=2' \
  --header 'Authorization: Token token=u+YhtuzQa11j1-8da-Ag' \
  --header 'Content-Type: application/json' \
  --header 'From: ved@turbot.com' \
  --data "{
  \"incident\": {
    \"type\": \"incident\",
    \"title\": \"${title}\",
    \"service\": {
      \"id\": \"PQZITI6\",
      \"type\": \"service_reference\"
    },
    \"priority\": {
      \"id\": \"P53ZZH5\",
      \"type\": \"priority_reference\"
    },
    \"urgency\": \"high\",
    \"body\": {
      \"type\": \"incident_body\",
      \"details\": \"A disk is getting full on this machine. You should investigate what is causing the disk to fill, and ensure that there is an automated process in place for ensuring data is rotated (eg. logs should have logrotate around them). If data is expected to stay on this disk forever, you should start planning to scale up to a larger disk.\"
    },
    \"escalation_policy\": {
      \"id\": \"P2LOCGZ\",
      \"type\": \"escalation_policy_reference\"
    }
  }
}"
done