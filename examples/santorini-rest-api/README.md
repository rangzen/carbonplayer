# Carbon Player - Santorini REST API

## Classic HTTP Server

```shell
go run main.go
```

then open `localhost:10842`

## Netlify

* Configure a new site with **Functions directory** to `examples/santorini-rest-api/netlify/functions`.
* The client can call `https://your-site-name.netlify.app/.netlify/functions/nextPlay`.