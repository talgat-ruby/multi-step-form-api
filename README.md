# multi-step-form-api

multi step form api for information of planets.

## Requirements

Install [docker desktop](https://www.docker.com/products/docker-desktop/)

Just run build once. If it stacks, restart docker desktop.

```shell
docker build --no-cache -t multi-step-form/api .
```

Can run as many times as you want.

```shell
docker run -d -p 8081:80 multi-step-form/api
```

Unfortunately no swagger. But you can find description below.

## API

Valid **plans** are:

- arcade
- advanced
- pro

Valid **periods** are:

- monthly
- yearly

`POST` `/forms`

**Response**

Sample Success Response for

```bash
curl -X POST http://localhost:8081/api/v1/forms \
      -H 'Content-Type: application/json' \
       -d '{"name": "Talgat", "email": "t@example.com", "phone": "+770875195", "plan": "arcade", "period": "yearly", "addOns": {"onlineService": true, "largerStorage": false}}'
```

**DON'T FORGET ABOUT `'Content-Type: application/json` HEADER!!!**

Request Body

```json
{
  "name": "Talga", // required
  "email": "ta@eample.com", // must be valid email
  "phone": "+770875195", // must be valid phone, starts with +. e164 format
  "plan": "arcade", // one of: arcade, advanced, pro
  "period": "yearly", // one of: monthly, yearly
  "addOns": { // required, can be empty object
    "onlineService": true,
    "largerStorage": false
  }
}
```

Sample Error Response for

```bash
curl -X POST http://localhost:8081/api/v1/forms \
      -H 'Content-Type: application/json' \
       -d '{"name": "Talga", "phone": "+770875195", "plan": "arcade"}'
```

```json
{
  "error": {
    "email": "email is invalid",
    "period": "period is invalid",
    "addOns": "add-ons are required"
  }
}
```

## License

MIT
