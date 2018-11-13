# Availability checker

## How to run

```shell
$ docker run \
    --rm \
    -e SMTP_HOST="" \
    -e SMTP_PORT="" \
    -e SMTP_USER="" \
    -e SMTP_PASSWORD="" \
    -e CRON="0 */10 * * * *"
    --volume ./config.json:/var/lib/availability-checker/config.json \ 
    budry/availability-checker
```

## Configuration

In `config.json`

```json
{
  "sites": [
    {"url": "google.com", "code": 200, "title": "Google"}
  ]
}
```