# order-alert-system


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