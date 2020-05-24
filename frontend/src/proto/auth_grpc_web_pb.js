/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./auth_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AuthClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AuthPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

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
 *   !proto.proto.LoginRequest,
 *   !proto.proto.LoginResponse>}
 */
const methodDescriptor_Auth_Login = new grpc.web.MethodDescriptor(
  '/proto.Auth/Login',
  grpc.web.MethodType.UNARY,
  proto.proto.LoginRequest,
  proto.proto.LoginResponse,
  /**
   * @param {!proto.proto.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.LoginResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.LoginRequest,
 *   !proto.proto.LoginResponse>}
 */
const methodInfo_Auth_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.LoginResponse,
  /**
   * @param {!proto.proto.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.proto.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AuthClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.Auth/Login',
      request,
      metadata || {},
      methodDescriptor_Auth_Login,
      callback);
};


/**
 * @param {!proto.proto.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.LoginResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.AuthPromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.Auth/Login',
      request,
      metadata || {},
      methodDescriptor_Auth_Login);
};


module.exports = proto.proto;

