## Build
```
docker build -t pressly/screenshot .
```

## Run in server mode
```
docker run -d -p 9222:9222 --cap-add=SYS_ADMIN pressly/screenshot
```

## TODO
1. Twirp server should start on :80
2. The entrypoint binary should start both Chrome
3. ^ starting Chrome will be based on original justinribeiro/chrome-headless image:
```
docker run -d -p 9222:9222 --cap-add=SYS_ADMIN justinribeiro/chrome-headless
```
4. Run in CLI mode (client+server) .. `docker run -d -p 9222:9222 --cap-add=SYS_ADMIN pressly/screenshot -width=800 -height=600 URL`