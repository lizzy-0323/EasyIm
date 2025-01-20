make build
cd bin/

echo "启动business服务"
nohup ./business-server &

echo "启动connect服务"
nohup ./conn-server &

echo "启动logic服务"
nohup ./logic-server &
