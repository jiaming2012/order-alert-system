# order-alert-system

# Admin Frontend
The admin frontend is a reactjs websocket app. To build run:
``` bash
cd frontend/
REACT_APP_BASIC_AUTH_USER="" REACT_APP_BASIC_AUTH_PASS="" npm run build
```

# SMS
Sms notifications are sent via the twilio api: https://www.twilio.com

# Development
## Environment Variables
The following environment variables need to be set:
``` bash 
BASIC_AUTH_PASS=
BASIC_AUTH_USER=
TWILIO_ACCOUNT_SID=
TWILIO_AUTH_TOKEN=
TWILIO_PHONE_NUMBER=
DATABASE_URL=
```

## Postgres
We use Postgres as our backend for storing user and order data.

The following command set up a local docker postgres instance for development:
``` bash
DB_PASSWORD="somePassword"
docker run --name yumyums-postgres-local -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${DB_PASSWORD}  -p 5432:5432 -d postgres
docker exec -ti yumyums-postgres-local psql -U postgres -c 'create database "customer_orders";'
```

## Docker
To run locally:
``` bash
docker run -p 8080:8080 -e BASIC_AUTH_PASS=$BASIC_AUTH_PASS -e BASIC_AUTH_USER=$BASIC_AUTH_USER -e DATABASE_URL="host=postgres user=postgres password=mysecretpass dbname=customer_orders port=5432 sslmode=disable" -e TWILIO_ACCOUNT_SID=$TWILIO_ACCOUNT_SID -e TWILIO_AUTH_TOKEN=$TWILIO_AUTH_TOKEN -e TWILIO_PHONE_NUMBER=$TWILIO_PHONE_NUMBER --link yumyums-postgres-local:postgres yumyums/order-messenger
```

## Mock
A websocket mock server can easily be spun up for client side UI development.

### Usage
Run:
``` bash
cd mock/
npm i
npm run start
```

All custom business logic lives in `mock/src/api/services`.

### Install
The following commands generates the mock server from scratch:
``` bash
npm install -g @asyncapi/generator
ag docs/order-events.yaml @asyncapi/nodejs-ws-template -o mock -p server=dev
```

# Deploy
## Docker
Before deploying to Heroku, we must dockerize our app:
``` bash
docker build -t yumyums/order-messenger .
```

## Heroku
We use heroku to deploy our app. A new version of our app is deployed by pushing to a remote heroku git branch:
``` bash
git push heroku main
```

### Database
Set up a heroku database instance:
``` bash
heroku addons:create -a yumyums-order-messenger heroku-postgresql:mini
```

``` bash
REACT_APP_BASIC_AUTH_USER=
REACT_APP_BASIC_AUTH_PASS=
```
## Backend
### Environment variables
All environment variable were set as fly secrets, using example command:
``` bash
flyctl secrets set BASIC_AUTH_PASS=mysecret
```
The following environment variables will need to be set:
``` bash
BASIC_AUTH_PASS=
BASIC_AUTH_USER=
TWILIO_ACCOUNT_SID=
TWILIO_AUTH_TOKEN=
TWILIO_PHONE_NUMBER=
```