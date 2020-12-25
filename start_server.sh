#!/bin/bash
if [[ -f `pwd`/API/main ]]; then
	echo "noob"
else
	go build -o api API/main.go
fi

./api &
php -S localhost:8000 -t Website/ 
