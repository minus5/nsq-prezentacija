go build || { echo 'build failed' ; exit 1; }
mkdir -p ./tmp/nsqd
mkdir -p ./tmp/log
consul agent -config-file=consul.json &>./tmp/log/consul.log &
nsqlookupd &>./tmp/log/nsqlookupd.log &
nsqd -lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=127.0.0.1 -data-path=./tmp/nsqd -max-msg-size=10485760 &>./tmp/log/nsqd.log &
nsqadmin -lookupd-http-address=127.0.0.1:4161 &>./tmp/log/nsqadmin.log &
./nsq-prezentacija &
wait
