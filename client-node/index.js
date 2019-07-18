/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
var express = require('express');
var app = express();
var bodyParser = require('body-parser');

var fs = require("fs");


const PORT_GRPC = process.env.PORT_GRPC || 3000,
    PORT_API = process.env.PORT_API || 3030,
    GCD_SERVICE_NAME = process.env.GCD_SERVICE_NAME || 'gcd.example.com',
    HOST_NAME = process.env.HOST_NAME || 'api-node.example.com';

console.info(`PORT_GRPC = ${PORT_GRPC}`);
console.info(`PORT_API = ${PORT_API}`);
console.info(`GCD_SERVICE_NAME = ${GCD_SERVICE_NAME}`);
console.info(`HOST_NAME = ${HOST_NAME}`);

var PROTO_PATH = './pb/gcd.proto';
var ROOT_CERTS = './certs/server.crt';
var ROOT_KEY = './certs/server.key';

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });

var gcd_proto = grpc.loadPackageDefinition(packageDefinition).pb;

const cacert = null,
    cert = fs.readFileSync(ROOT_CERTS),
    key = fs.readFileSync(ROOT_KEY);

const options = {
    'grpc.ssl_target_name_override' : HOST_NAME,
    'grpc.default_authority': HOST_NAME
};

const creds = grpc.credentials.createSsl(cacert, key, cert, options);

function Generate(attributes) {

    var client = new gcd_proto.GCDService(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, creds);

    client.Generate({Attributes: attributes}, function (err, response) {
        console.log(`Greeting:, ${response}`);
    });
}

app.use(bodyParser());

app.post('/generate', function (req, res) {
    let attributes = req.body.attributes ? req.body.attributes : null;
    Generate(attributes);
    res.send('Success!');
});

app.listen(PORT_API, function () {
    console.log(`Example app listening on port ${PORT_API}!`);
});


