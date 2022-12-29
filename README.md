# order-alert-system

# SMS
Sms notifications are sent via the twilio api: https://www.twilio.com

# Postgres
We use Postgre as our backend for storing user and order data.

The following command set up a local docker postgre instance for development.
``` bash
DB_PASSWORD="somePassword"
docker run --name yumyums-local -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${DB_PASSWORD}  -p 5432:5432 -d postgres
docker exec -ti yumyums-local psql -U postgres -c 'create database "customer_orders";'
docker exec -ti yumyums-local psql -U postgres -c 'grant all privileges ON DATABASE "customer_orders" TO postgres;'
docker exec -ti yumyums-local psql -U postgres -c 'create role "customer-orders";'
docker exec -ti yumyums-local psql -U postgres -c 'grant usage on schema public to "customer-orders";'
docker exec -ti yumyums-local psql -U postgres -c 'grant all privileges ON DATABASE "customer_orders" TO "customer-orders";'
```

# Mock
To start:
``` bash
cd mock/
npm i
npm run start
```

All custom business logic lives in `mock/src/api/services`.

The following commands generates the mock server from scratch:
``` bash
npm install -g @asyncapi/generator
ag docs/order-events.yaml @asyncapi/nodejs-ws-template -o mock -p server=dev
```