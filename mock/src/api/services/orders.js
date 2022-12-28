const service = module.exports = {};

const dummy = require('dummy-json');

/**
 * Client can receive chat messages.
 * @param {object} ws WebSocket connection.
 */
service.subOrdersMessage = async (ws) => {
    (function myLoop () {
      setTimeout(() => {
        try {
          ws.send(generateResponse());
          myLoop();
        } catch (err) {
          return;
        }
      }, 4000);
    }());

    const myHelpers = {
      status() {
        // Use randomArrayItem() to ensure the seeded random number generator is used
        return dummy.utils.randomArrayItem(['open', 'awaiting_pickup']);
      }
    };

    function generateResponse() {
      const template = `[
      {{#repeat 0 3}}
      {"orderId": "{{int 0 10}}","phoneNumber": "{{phone "+1 (xxx) xxx-xxxx"}}","status": "{{status}}"}
      {{/repeat}}
    ]`;
      return dummy.parse(template, { helpers: myHelpers });
    }
};
