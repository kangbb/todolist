docker pull registry-vpc.cn-shenzhen.aliyuncs.com/selfmysql/master
docker pull registry-vpc.cn-shenzhen.aliyuncs.com/selfmysql/slave
docker-compose up -d
id=$(docker images --filter "dangling=true" -q --no-trunc)
if [ $id ]; then docker rmi $id; fi