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

var messages = require('./gcd_pb');
var services = require('./gcd_grpc_pb');


const PORT_GRPC = process.env.PORT_GRPC || 3000,
    PORT_API = process.env.PORT_API || 3031,
    GCD_SERVICE_NAME = process.env.GCD_SERVICE_NAME || 'localhost',
    HOST_NAME = process.env.HOST_NAME || 'localhost',
    TLS_ENABLE = process.env.TLS_ENABLE || "false";

console.info(`PORT_GRPC = ${PORT_GRPC}`);
console.info(`PORT_API = ${PORT_API}`);
console.info(`GCD_SERVICE_NAME = ${GCD_SERVICE_NAME}`);
console.info(`HOST_NAME = ${HOST_NAME}`);
console.info(`TLS_ENABLE = ${TLS_ENABLE}`);

var CA_CERTS = './certs/ca/rootCA.crt';
var ROOT_CERTS = './certs/client/client.crt';
var ROOT_KEY = './certs/client/client.key';

var grpc = require('grpc');

const cacert = fs.readFileSync(CA_CERTS),
    cert = fs.readFileSync(ROOT_CERTS),
    key = fs.readFileSync(ROOT_KEY);

const options = {
    'grpc.ssl_target_name_override': HOST_NAME,
    'grpc.default_authority': HOST_NAME
};

const creds = grpc.credentials.createSsl(cacert, key, cert, options);

function Generate(attributes) {

    var client, request;

    if (TLS_ENABLE === "true") {
        client = new services.GCDServiceClient(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, creds);
    } else {
        client = new services.GCDServiceClient(`${GCD_SERVICE_NAME}:${PORT_GRPC}`, grpc.credentials.createInsecure());
    }

    request = new messages.GenerateRequest();
    request.setAttributes(attributes);

    client.generate(request, function(err, response) {
        if(err){
            console.log(err);
            return err;
        }
        console.log(`Greeting:, ${response}`);
        return response;
    });
}

app.use(bodyParser.urlencoded({
    extended: true
}));

app.use(bodyParser.json());

app.post('/generate', function (req, res) {
    let attr = req.body.attributes ? req.body.attributes : null;
    let arg = new Buffer.from(attr.toString());
    console.log(arg);
    Generate(arg);
    res.send('Success!');
});

app.listen(PORT_API, function () {
    console.log(`Example app listening on port ${PORT_API}!`);
});


