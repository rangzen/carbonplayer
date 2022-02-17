# Carbon Player - Santorini - REST API

Demo of a Santorini REST API server.

## Classic HTTP Server

```shell
go run main.go
```

then open `localhost:10842`

## Netlify and AWS Lambda

* Configure a new site with **Functions directory** to `examples/santorini-rest-api/netlify/functions`.
* The client can call `https://your-site-name.netlify.app/.netlify/functions/nextPlay`.

Remember that normal Functions (not Background ones) are limited to 10 seconds of execution.
Play with the `max_plies` configuration.