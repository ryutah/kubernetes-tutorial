#!/bin/sh

kubectl create configmap my-config --from-literal=key1=value1 --from-literal=key2=value2
