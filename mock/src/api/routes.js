const util = require('util');
const { Router } = require('express');
const { pathParser } = require('../lib/path');
const { yellow } = require('../lib/colors');
const { subOrdersMessage } = require('./services/orders');
const router = Router();
module.exports = router;
router.ws('/orders', async (ws, req) => {
  const path = pathParser(req.path);
  console.log(`${yellow(path)} client connected.`);
  await subOrdersMessage(ws);
});
