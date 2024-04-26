سامانه نوبت دهی آنلاین

این پروژه یک سامانه نوبت دهی آنلاین است که به کاربران امکان می‌دهد به راحتی و سریع‌تر از طریق اینترنت نوبت‌های خود را دریافت کنند. این سامانه بهبودی عمده در فرآیند نوبت دهی و ارائه خدمات به مشتریان برای کسب و کارها و سازمان‌های مختلف را فراهم می‌کند.

### Create Network
```shell
docker network create appointment_network
```

### Run Frontend

```Shell
cd frontend/appointment-scheduler
docker build . -t front
docker run --network appointment_network --name frontend -d front
```

> TODO: some conflicts exists on package and should be solve
> because of that we use from `--legacy-peer-deps` option in `Dockerfile`.

### Run Backend

```shell
cd backend
docker build . -f src/the-wall/Dockerfile -t walltest:test1
docker run --network appointment_network --name the_wall -d walltest:test1
docker container ls
docker logs -f the_wall
docker stop the_wall
```

### Run Gateway
```shell
cd gateway/nginx/
docker build . -t gateway:v1
docker run -p 80:80 --network appointment_network --name gateway -d gateway:v1
```
