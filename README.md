# order-alert-system

# Admin Frontend
The admin frontend is a reactjs websocket app. To build run:
``` bash
cd frontend/
REACT_APP_BACKEND_URL="" REACT_APP_BASIC_AUTH_USER="" REACT_APP_BASIC_AUTH_PASS="" npm run build
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

# Postgres
We use Postgres as our backend for storing user and order data.

The following command set up a local docker postgres instance for development:
``` bash
DB_PASSWORD="somePassword"
docker run --name yumyums-local -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${DB_PASSWORD}  -p 5432:5432 -d postgres
docker exec -ti yumyums-local psql -U postgres -c 'create database "customer_orders";'
```

# Mock
A websocket mock server can easily be spun up for client side UI development.

## Usage
Run:
``` bash
cd mock/
npm i
npm run start
```

All custom business logic lives in `mock/src/api/services`.

## Install
The following commands generates the mock server from scratch:
``` bash
npm install -g @asyncapi/generator
ag docs/order-events.yaml @asyncapi/nodejs-ws-template -o mock -p server=dev
```

# Deploy
We use heroku to deploy our app:

``` bash
REACT_APP_BASIC_AUTH_USER=
REACT_APP_BASIC_AUTH_PASS=
REACT_APP_BACKEND_URL=
DATABASE_URL=
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

### Deploy
After env variables are set, the app can be deployed via:
``` bash
flyctl deploy
```