// This file contains examples of scenarios implementation using
// the SDK for channels management.

const system = require('./system/client');

const client = system.Client('http://localhost:8080');

// Scenario 1: Display available channels.
client
  .listBalancers()
  .then(list => {
    console.log('=== Scenario 1 ===');
    console.log('Balancers:');
    console.log(list);
  })
  .catch(e => {
    console.log(`Problem listing balancers: ${e.message}`);
  });

// Scenario 2: Create new channel.
client
  .updateWorkingStatus(1, true)
  .then(resp => {
    console.log('=== Scenario 2 ===');
    console.log('Update working working status of a machine:', resp);
    return client
      .listBalancers()
      .then(list => console.log('Current Balancers:', list));
  })
  .catch(e => {
    console.log(`Problem updating status of a machine: ${e.message}`);
  });
