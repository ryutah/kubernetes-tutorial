#!/bin/sh

kubectl get secret my-password -o json | jq -r ".data.password" | base64 -D
