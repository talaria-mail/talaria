/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./users_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.UserServiceClient =
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
proto.proto.UserServicePromiseClient =
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
 *   !proto.proto.CreateUserRequest,
 *   !proto.proto.CreateUserResponse>}
 */
const methodDescriptor_UserService_Create = new grpc.web.MethodDescriptor(
  '/proto.UserService/Create',
  grpc.web.MethodType.UNARY,
  proto.proto.CreateUserRequest,
  proto.proto.CreateUserResponse,
  /**
   * @param {!proto.proto.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.CreateUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.CreateUserRequest,
 *   !proto.proto.CreateUserResponse>}
 */
const methodInfo_UserService_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.CreateUserResponse,
  /**
   * @param {!proto.proto.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.CreateUserResponse.deserializeBinary
);


/**
 * @param {!proto.proto.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.CreateUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.CreateUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.UserServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.UserService/Create',
      request,
      metadata || {},
      methodDescriptor_UserService_Create,
      callback);
};


/**
 * @param {!proto.proto.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.CreateUserResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.UserServicePromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.UserService/Create',
      request,
      metadata || {},
      methodDescriptor_UserService_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.FetchUserRequest,
 *   !proto.proto.FetchUserResponse>}
 */
const methodDescriptor_UserService_Fetch = new grpc.web.MethodDescriptor(
  '/proto.UserService/Fetch',
  grpc.web.MethodType.UNARY,
  proto.proto.FetchUserRequest,
  proto.proto.FetchUserResponse,
  /**
   * @param {!proto.proto.FetchUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.FetchUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.FetchUserRequest,
 *   !proto.proto.FetchUserResponse>}
 */
const methodInfo_UserService_Fetch = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.FetchUserResponse,
  /**
   * @param {!proto.proto.FetchUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.FetchUserResponse.deserializeBinary
);


/**
 * @param {!proto.proto.FetchUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.FetchUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.FetchUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.UserServiceClient.prototype.fetch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.UserService/Fetch',
      request,
      metadata || {},
      methodDescriptor_UserService_Fetch,
      callback);
};


/**
 * @param {!proto.proto.FetchUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.FetchUserResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.UserServicePromiseClient.prototype.fetch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.UserService/Fetch',
      request,
      metadata || {},
      methodDescriptor_UserService_Fetch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.ListUsersRequest,
 *   !proto.proto.ListUsersResponse>}
 */
const methodDescriptor_UserService_List = new grpc.web.MethodDescriptor(
  '/proto.UserService/List',
  grpc.web.MethodType.UNARY,
  proto.proto.ListUsersRequest,
  proto.proto.ListUsersResponse,
  /**
   * @param {!proto.proto.ListUsersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ListUsersResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.ListUsersRequest,
 *   !proto.proto.ListUsersResponse>}
 */
const methodInfo_UserService_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.ListUsersResponse,
  /**
   * @param {!proto.proto.ListUsersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.ListUsersResponse.deserializeBinary
);


/**
 * @param {!proto.proto.ListUsersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.ListUsersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.ListUsersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.UserServiceClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.UserService/List',
      request,
      metadata || {},
      methodDescriptor_UserService_List,
      callback);
};


/**
 * @param {!proto.proto.ListUsersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.ListUsersResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.UserServicePromiseClient.prototype.list =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.UserService/List',
      request,
      metadata || {},
      methodDescriptor_UserService_List);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.UpdateUserRequest,
 *   !proto.proto.UpdateUserResponse>}
 */
const methodDescriptor_UserService_Update = new grpc.web.MethodDescriptor(
  '/proto.UserService/Update',
  grpc.web.MethodType.UNARY,
  proto.proto.UpdateUserRequest,
  proto.proto.UpdateUserResponse,
  /**
   * @param {!proto.proto.UpdateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.UpdateUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.UpdateUserRequest,
 *   !proto.proto.UpdateUserResponse>}
 */
const methodInfo_UserService_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.UpdateUserResponse,
  /**
   * @param {!proto.proto.UpdateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.UpdateUserResponse.deserializeBinary
);


/**
 * @param {!proto.proto.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.UpdateUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.UpdateUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.UserServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.UserService/Update',
      request,
      metadata || {},
      methodDescriptor_UserService_Update,
      callback);
};


/**
 * @param {!proto.proto.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.UpdateUserResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.UserServicePromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.UserService/Update',
      request,
      metadata || {},
      methodDescriptor_UserService_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.DeleteUserRequest,
 *   !proto.proto.DeleteUserResponse>}
 */
const methodDescriptor_UserService_Delete = new grpc.web.MethodDescriptor(
  '/proto.UserService/Delete',
  grpc.web.MethodType.UNARY,
  proto.proto.DeleteUserRequest,
  proto.proto.DeleteUserResponse,
  /**
   * @param {!proto.proto.DeleteUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.DeleteUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.DeleteUserRequest,
 *   !proto.proto.DeleteUserResponse>}
 */
const methodInfo_UserService_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.DeleteUserResponse,
  /**
   * @param {!proto.proto.DeleteUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.DeleteUserResponse.deserializeBinary
);


/**
 * @param {!proto.proto.DeleteUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.DeleteUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.DeleteUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.UserServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.UserService/Delete',
      request,
      metadata || {},
      methodDescriptor_UserService_Delete,
      callback);
};


/**
 * @param {!proto.proto.DeleteUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.DeleteUserResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.UserServicePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.UserService/Delete',
      request,
      metadata || {},
      methodDescriptor_UserService_Delete);
};


module.exports = proto.proto;

