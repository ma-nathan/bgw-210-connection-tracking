#!/bin/bash

while [ 1 ] ; do
	/bgw210 >> /var/log/bgw210.log 2>&1 &
	sleep 3600
done

