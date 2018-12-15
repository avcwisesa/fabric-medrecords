'use strict';

var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');

// Membuat fabric client
var fabric_client = new Fabric_Client();

// Setup Fabric network
var channel = fabric_client.newChannel('mychannel');
var peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(peer);

// Init path
var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
console.log('Store path:'+store_path);
var tx_id = null;

// Membuat key-value store sesuai dengan apa yang didefinisikan di fabric-client/config/default.json 'key-value-store' setting (belum buat)
Fabric_Client.newDefaultKeyValueStore({ path: store_path
}).then((state_store) => {
    // TODO
}).then((user_from_store) => {
    // TODO
}).then((results) => {
    // TODO
}).then((results) => {
    // TODO
}).catch((err) => {
    console.error('Failed to invoke successfully :: ' + err);
})