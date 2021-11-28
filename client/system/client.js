const { getClient } = require('../common/http');

const Client = baseUrl => {
  const client = getClient(baseUrl);

  return {
    listBalancers: () => client.get('/balancers'),
    updateWorkingStatus: (id, working) =>
      client.put('/machines', { id, working }),
  };
};

module.exports = { Client };
