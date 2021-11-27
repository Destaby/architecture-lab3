const request = require('request');

const getClient = baseUrl => ({
  get: path => {
    return new Promise((resolve, reject) => {
      request(`${baseUrl}${path}`, { json: true }, (err, res, body) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(body);
      });
    });
  },
  put: async (path, data) => {
    return new Promise((resolve, reject) => {
      request(
        `${baseUrl}${path}`,
        { json: true, method: 'PUT', body: data },
        (err, res, body) => {
          if (err) {
            reject(err);
            return;
          }
          resolve(body);
        }
      );
    });
  },
});

module.exports = { getClient };
