
## build
```
docker build . --tag gitlab_webhook:0.0.1
```

### run
```
docker run -p3000:3000 -e WECOM_TOKEN="" --rm gitlab_webhook:latest
```