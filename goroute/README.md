编译docker

```shell script
docker build . -t goroute
```

限制为0.5核1G运行容器

```shell script
docker run -d -p 7080:7080 --cpus 0.5 --memory 1g  --memory-swap 1g  goroute 
```