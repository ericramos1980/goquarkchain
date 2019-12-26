#!/bin/bash
touch /tmp/status.txt
git pull
go build
chmod +x ./stop.sh && ./stop.sh
NOW=`date +%Y-%m-%d.%H:%M:%S`
DATE=`curl https://qkcmainnet-go.s3.amazonaws.com/data/LATEST`
echo "current time" $NOW  "remote time" $DATE >> /tmp/status.txt
rm -rf data.tar.gz
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
./cluster --cluster_config ../../mainnet/singularity/cluster_config_template.json  --check_db --check_db_rblock_from=1000
if [ "${NOW:0:10}" = "${DATE:0:10}" ];then
	echo "=" >> /tmp/status.txt
else
	echo "!=" >> /tmp/status.txt
fi
./stop.sh
