docker rmi registry-vpc.cn-shenzhen.aliyuncs.com/selfmysql/master
docker rmi registry-vpc.cn-shenzhen.aliyuncs.com/selfmysql/slave
docker pull registry-vpc.cn-shenzhen.aliyuncs.com/selfmysql/master
docker pull registry-vpc.cn-shenzhen.aliyuncs.com/selfmysql/slave
docker-compose up -d
docker ps