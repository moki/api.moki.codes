#!/bin/sh
i=0
while [ "$i" != 10000 ]
do
	echo "Looping ... i is set to $i"
	curl --data "email=email-address-$i@moki.codes&name=$i" -sL https://api.moki.codes.localhost/newsletter/subscribers
	((i++))
done
