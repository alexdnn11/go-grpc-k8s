// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var gcd_pb = require('./gcd_pb.js');

function serialize_pb_GenerateRequest(arg) {
  if (!(arg instanceof gcd_pb.GenerateRequest)) {
    throw new Error('Expected argument of type pb.GenerateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_GenerateRequest(buffer_arg) {
  return gcd_pb.GenerateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_GenerateResponse(arg) {
  if (!(arg instanceof gcd_pb.GenerateResponse)) {
    throw new Error('Expected argument of type pb.GenerateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_GenerateResponse(buffer_arg) {
  return gcd_pb.GenerateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_VerifyRequest(arg) {
  if (!(arg instanceof gcd_pb.VerifyRequest)) {
    throw new Error('Expected argument of type pb.VerifyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_VerifyRequest(buffer_arg) {
  return gcd_pb.VerifyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pb_VerifyResponse(arg) {
  if (!(arg instanceof gcd_pb.VerifyResponse)) {
    throw new Error('Expected argument of type pb.VerifyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pb_VerifyResponse(buffer_arg) {
  return gcd_pb.VerifyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var GCDServiceService = exports.GCDServiceService = {
  generate: {
    path: '/pb.GCDService/Generate',
    requestStream: false,
    responseStream: false,
    requestType: gcd_pb.GenerateRequest,
    responseType: gcd_pb.GenerateResponse,
    requestSerialize: serialize_pb_GenerateRequest,
    requestDeserialize: deserialize_pb_GenerateRequest,
    responseSerialize: serialize_pb_GenerateResponse,
    responseDeserialize: deserialize_pb_GenerateResponse,
  },
  verify: {
    path: '/pb.GCDService/Verify',
    requestStream: false,
    responseStream: false,
    requestType: gcd_pb.VerifyRequest,
    responseType: gcd_pb.VerifyResponse,
    requestSerialize: serialize_pb_VerifyRequest,
    requestDeserialize: deserialize_pb_VerifyRequest,
    responseSerialize: serialize_pb_VerifyResponse,
    responseDeserialize: deserialize_pb_VerifyResponse,
  },
};

exports.GCDServiceClient = grpc.makeGenericClientConstructor(GCDServiceService);
