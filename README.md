# Golang App with Nginx Proxy

minimal setup for golang web app with nginx proxy (for demo only)


### Usage

```sh
% docker compose up -d
```


```sh
% curl http://localhost:8080

Host => "localhost"
RemoteAddr => "192.168.48.3:33522"

"X-Forwarded-For" => ["192.168.48.1"]
"Connection" => ["close"]
"User-Agent" => ["curl/7.78.0"]
"Accept" => ["*/*"]
"X-Real-Ip" => ["192.168.48.1"]
"X-Forwarded-Proto" => ["http"]
```
