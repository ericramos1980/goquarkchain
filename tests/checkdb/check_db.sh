#!/bin/bash
git pull
go build
CLUSTER_DIR="/go/src/github.com/QuarkChain/goquarkchain/cmd/cluster"
chmod +x ./stop.sh && ./stop.sh
NOW=`date +%Y-%m-%d.%H:%M:%S`
DATE=`curl https://qkcmainnet-go.s3.amazonaws.com/data/LATEST`
echo "current time" $NOW  "remote time" $DATE >> /tmp/status.txt
rm -rf data.tar.gz
curl https://qkcmainnet-go.s3.amazonaws.com/data/`curl https://qkcmainnet-go.s3.amazonaws.com/data/LATEST`.tar.gz --output data.tar.gz
tar xvfz data.tar.gz
rm -rf $CLUSTER_DIR/qkc-data && mkdir -p $CLUSTER_DIR/qkc-data/
rm -rf data.tar.gz && mv mainnet $CLUSTER_DIR/qkc-data

# checkdb
chmod +x $CLUSTER_DIR/cluster
rm -rf *.log
$CLUSTER_DIR/cluster --cluster_config /go/src/github.com/QuarkChain/goquarkchain/mainnet/singularity/cluster_config_template.json --service S0>> S0.log 2>&1 &
$CLUSTER_DIR/cluster --cluster_config /go/src/github.com/QuarkChain/goquarkchain/mainnet/singularity/cluster_config_template.json --service S1>> S1.log 2>&1 &
$CLUSTER_DIR/cluster --cluster_config /go/src/github.com/QuarkChain/goquarkchain/mainnet/singularity/cluster_config_template.json --service S2>> S2.log 2>&1 &
$CLUSTER_DIR/cluster --cluster_config /go/src/github.com/QuarkChain/goquarkchain/mainnet/singularity/cluster_config_template.json --service S3>> S3.log 2>&1 &
sleep 3
echo "ready to checkdb `date +%Y-%m-%d.%H:%M:%S`" >> /tmp/status.txt
$CLUSTER_DIR/cluster --cluster_config /go/src/github.com/QuarkChain/goquarkchain/mainnet/singularity/cluster_config_template.json  --check_db >> cc.log 2>&1
echo "end to checkdb `date +%Y-%m-%d.%H:%M:%S`" >> /tmp/status.txt

ee=`cat /go/src/github.com/QuarkChain/goquarkchain/cmd/cluster/cc.log | grep ERROR`
if [ -z "$ee" ]; then
	echo "no error"
else
	echo $ee >> /tmp/status.txt
	echo "!=" >> /tmp/status.txt
	./stop.sh
	exit 0
fi

if [ "${NOW:0:10}" = "${DATE:0:10}" ];then
	echo "=" >> /tmp/status.txt
else
	echo "!=" >> /tmp/status.txt
fi
./stop.sh
