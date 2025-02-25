/**
 * @fileoverview gRPC-Web generated client stub for greet
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.greet = require('./greet_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.greet.GreeterClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.greet.GreeterPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.greet.GreetRequest,
 *   !proto.greet.GreetReply>}
 */
const methodDescriptor_Greeter_Greet = new grpc.web.MethodDescriptor(
  '/greet.Greeter/Greet',
  grpc.web.MethodType.UNARY,
  proto.greet.GreetRequest,
  proto.greet.GreetReply,
  /**
   * @param {!proto.greet.GreetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.greet.GreetReply.deserializeBinary
);


/**
 * @param {!proto.greet.GreetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.greet.GreetReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.greet.GreetReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.greet.GreeterClient.prototype.greet =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/greet.Greeter/Greet',
      request,
      metadata || {},
      methodDescriptor_Greeter_Greet,
      callback);
};


/**
 * @param {!proto.greet.GreetRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.greet.GreetReply>}
 *     Promise that resolves to the response
 */
proto.greet.GreeterPromiseClient.prototype.greet =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/greet.Greeter/Greet',
      request,
      metadata || {},
      methodDescriptor_Greeter_Greet);
};


module.exports = proto.greet;

