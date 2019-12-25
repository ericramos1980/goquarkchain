#!/bin/bash


step=60*60*24  # 24h
touch /tmp/status.txt
for (( i = 0; i < 1000000; i=(i+1) )); do
	git pull
	go build
	chmod +x ./stop.sh && ./stop.h
	DATE=`curl https://qkcmainnet-go.s3.amazonaws.com/data/LATEST` >> /tmp/status.txt
	curl https://qkcmainnet-go.s3.amazonaws.com/data/`curl https://qkcmainnet-go.s3.amazonaws.com/data/LATEST`.tar.gz --output data.tar.gz
	tar xvfz data.tar.gz
	rm -rf qkc-data && mkdir -p ./qkc-data/
	rm -rf data.tar.gz && mv mainnet qkc-data/

	# checkdb
	chmod +x cluster
	./cluster --cluster_config ../../mainnet/singularity/cluster_config_template.json --service S0>> S0.log 2>&1 &
	./cluster --cluster_config ../../mainnet/singularity/cluster_config_template.json --service S1>> S1.log 2>&1 &
	./cluster --cluster_config ../../mainnet/singularity/cluster_config_template.json --service S2>> S2.log 2>&1 &
	./cluster --cluster_config ../../mainnet/singularity/cluster_config_template.json --service S3>> S3.log 2>&1 &
	sleep 3
	./cluster --cluster_config ../../mainnet/singularity/cluster_config_template.json  --check_db
	echo $DATE >> /tmp/status.txt

	sleep $step
done
