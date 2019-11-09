#!/usr/bin/env bash

IP="192.168.99.100"
PORT="8081"

curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
    --data '{"schema": "{\"type\": \"string\"}"}' \
    http://${IP}:${PORT}/subjects/test1-key/versions

curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
    --data '{"schema": "{\"type\": \"string\"}"}' \
     http://${IP}:${PORT}/subjects/test2-value/versions

curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
    --data '{"schema": "{\"type\": \"int\"}"}' \
    http://${IP}:${PORT}/subjects/test3-key/versions

curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
    --data '{"schema": "{\"type\": \"long\"}"}' \
    http://${IP}:${PORT}/subjects/test3-key/versions

curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
    --data '{ "schema": "{ \"type\": \"record\", \"name\": \"Person\", \"namespace\": \"com.thomas.lamirault\", \"fields\": [ { \"name\": \"firstName\", \"type\": \"string\" }, { \"name\": \"lastName\", \"type\": \"string\" }, { \"name\": \"birthDate\", \"type\": \"long\" } ]}" }' \
    http://${IP}:${PORT}/subjects/test3-value/versions