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
const express = require('express');
const app = express();
const bodyParser = require('body-parser');
const protoLoader = require('@grpc/proto-loader');
const fs = require("fs");

const PORT_GRPC = process.env.PORT_GRPC || 3000,
    PORT_API = process.env.PORT_API || 3030,
    GCD_SERVICE_NAME = process.env.GCD_SERVICE_NAME || 'localhost',
    HOST_NAME = process.env.HOST_NAME || 'localhost',
    TLS_ENABLE = process.env.TLS_ENABLE || "false",
    DEBUG_MODE = process.env.DEBUG_MODE || "true";

let RP = './';

console.info(`PORT_GRPC = ${PORT_GRPC}`);
console.info(`PORT_API = ${PORT_API}`);
console.info(`GCD_SERVICE_NAME = ${GCD_SERVICE_NAME}`);
console.info(`HOST_NAME = ${HOST_NAME}`);
console.info(`TLS_ENABLE = ${TLS_ENABLE}`);
console.info(`DEBUG_MODE = ${DEBUG_MODE}`);

if (DEBUG_MODE === "true") {
    RP = '../';
}

let CA_CERTS = RP + 'certs/ca/rootCA.crt';
let ROOT_CERTS = RP + 'certs/client/client.crt';
let ROOT_KEY = RP + 'certs/client/client.key';
let PROTO_PATH = RP + 'pb/idemix.proto';

let grpc = require('grpc');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const packageDescriptor = grpc.loadPackageDefinition(packageDefinition);
const gcd_proto = packageDescriptor.pb;

const cacert = fs.readFileSync(CA_CERTS),
    cert = fs.readFileSync(ROOT_CERTS),
    key = fs.readFileSync(ROOT_KEY);

const options = {
    'grpc.ssl_target_name_override': HOST_NAME,
    'grpc.default_authority': HOST_NAME
};

const creds = grpc.credentials.createSsl(cacert, key, cert, options);

app.use(bodyParser.urlencoded({
    extended: true
}));

app.use(bodyParser.json());

app.post('/generate', (req, res) => {

    let client;
    let result;

    let attrObj = req.body.attributes ? req.body.attributes : null;
    let attrBytes = new Buffer.from(JSON.stringify(attrObj));

    if (TLS_ENABLE === "true") {
        client = new gcd_proto.Idemix(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, creds);
    } else {
        client = new gcd_proto.Idemix(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, grpc.credentials.createInsecure());
    }

    client.Generate({attributes: attrBytes}, (err, resGCD) => {
        if (err == null) {
            result = resGCD.result.toString();
        }else{
            result = `Error: ${err}`;
        }
        res.send({result: result});
    });

});

app.post('/verify', (req, res) => {

    let client;

    let proofObj = req.body.proof ? req.body.proof : null;
    let proofBytes = new Buffer.from(JSON.stringify(proofObj));

    if (TLS_ENABLE === "true") {
        client = new gcd_proto.Idemix(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, creds);
    } else {
        client = new gcd_proto.Idemix(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, grpc.credentials.createInsecure());
    }

    client.Verify({proof: proofBytes}, (err, resGCD) => {
        if (err == null) {
            result = resGCD.result.toString();
        }else{
            result = `Error: ${err}`;
        }
        res.send({result: result});
    });

});

app.listen(PORT_API, function () {
    console.log(`Example app listening on port ${PORT_API}!`);
});


