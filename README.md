# bearbear
demo app running at docker with Kong and Consul integrated

demo steps:
1. start consul  
cd <bearbear_path>/consul  
docker-compose up

2. start bearbear app 
cd <bearbear_path>/  
docker run -p 8080:8080 --network consul_bearnet -d europeli/bearbear:0.0.1

3. start kong  
cd <bearbear_path>/kong-compose  
docker-compose up

4. start konga  
docker run -p 1337:1337 --network consul_bearnet --name konga -e "NODE_ENV=production" pantsel/konga
