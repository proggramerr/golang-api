#!/bin/bash

while ! nc -z -w 1 db 5432; do
  sleep 1
done

sleep 5


./main