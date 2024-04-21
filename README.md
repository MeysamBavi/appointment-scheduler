سامانه نوبت دهی آنلاین

این پروژه یک سامانه نوبت دهی آنلاین است که به کاربران امکان می‌دهد به راحتی و سریع‌تر از طریق اینترنت نوبت‌های خود را دریافت کنند. این سامانه بهبودی عمده در فرآیند نوبت دهی و ارائه خدمات به مشتریان برای کسب و کارها و سازمان‌های مختلف را فراهم می‌کند.

### Run Frontend

```Shell
cd frontend/appointment-scheduler
docker build . -t front
docker run front
```

> TODO: some conflicts exists on package and should be solve
> because of that we use from `--legacy-peer-deps` option in `Dockerfile`.

### Run Backend

```shell
cd backend
docker build . -f src/the-wall/Dockerfile -t walltest:test1
docker run -d -P walltest:test1
docker container ls
docker logs -f <container-id>
docker stop <container-id>
```
