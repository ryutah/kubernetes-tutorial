#!/bin/sh

kubectl create secret generic my-password --from-literal=password=mysqlpassword
