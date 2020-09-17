/**
 * @fileoverview gRPC-Web generated client stub for zbay
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var user_pb = require('./user_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.zbay = require('./message_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.zbay.MessagesClient =
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
proto.zbay.MessagesPromiseClient =
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
 *   !proto.zbay.Message,
 *   !proto.zbay.Message>}
 */
const methodDescriptor_Messages_Add = new grpc.web.MethodDescriptor(
  '/zbay.Messages/Add',
  grpc.web.MethodType.UNARY,
  proto.zbay.Message,
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Message,
 *   !proto.zbay.Message>}
 */
const methodInfo_Messages_Add = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @param {!proto.zbay.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.Message)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.add =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Messages/Add',
      request,
      metadata || {},
      methodDescriptor_Messages_Add,
      callback);
};


/**
 * @param {!proto.zbay.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.Message>}
 *     A native promise that resolves to the response
 */
proto.zbay.MessagesPromiseClient.prototype.add =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Messages/Add',
      request,
      metadata || {},
      methodDescriptor_Messages_Add);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Message,
 *   !proto.zbay.Message>}
 */
const methodDescriptor_Messages_List = new grpc.web.MethodDescriptor(
  '/zbay.Messages/List',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.zbay.Message,
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Message,
 *   !proto.zbay.Message>}
 */
const methodInfo_Messages_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @param {!proto.zbay.Message} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.list =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/List',
      request,
      metadata || {},
      methodDescriptor_Messages_List);
};


/**
 * @param {!proto.zbay.Message} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesPromiseClient.prototype.list =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/List',
      request,
      metadata || {},
      methodDescriptor_Messages_List);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.User,
 *   !proto.zbay.Message>}
 */
const methodDescriptor_Messages_GroupBy = new grpc.web.MethodDescriptor(
  '/zbay.Messages/GroupBy',
  grpc.web.MethodType.SERVER_STREAMING,
  user_pb.User,
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.User,
 *   !proto.zbay.Message>}
 */
const methodInfo_Messages_GroupBy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @param {!proto.zbay.User} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.groupBy =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/GroupBy',
      request,
      metadata || {},
      methodDescriptor_Messages_GroupBy);
};


/**
 * @param {!proto.zbay.User} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesPromiseClient.prototype.groupBy =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/GroupBy',
      request,
      metadata || {},
      methodDescriptor_Messages_GroupBy);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Message,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_Messages_Send = new grpc.web.MethodDescriptor(
  '/zbay.Messages/Send',
  grpc.web.MethodType.UNARY,
  proto.zbay.Message,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.zbay.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Message,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Messages_Send = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.zbay.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.zbay.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Messages/Send',
      request,
      metadata || {},
      methodDescriptor_Messages_Send,
      callback);
};


/**
 * @param {!proto.zbay.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.zbay.MessagesPromiseClient.prototype.send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Messages/Send',
      request,
      metadata || {},
      methodDescriptor_Messages_Send);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.User,
 *   !proto.zbay.Message>}
 */
const methodDescriptor_Messages_Receive = new grpc.web.MethodDescriptor(
  '/zbay.Messages/Receive',
  grpc.web.MethodType.SERVER_STREAMING,
  user_pb.User,
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.User,
 *   !proto.zbay.Message>}
 */
const methodInfo_Messages_Receive = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Message,
  /**
   * @param {!proto.zbay.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Message.deserializeBinary
);


/**
 * @param {!proto.zbay.User} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.receive =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/Receive',
      request,
      metadata || {},
      methodDescriptor_Messages_Receive);
};


/**
 * @param {!proto.zbay.User} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Message>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesPromiseClient.prototype.receive =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/Receive',
      request,
      metadata || {},
      methodDescriptor_Messages_Receive);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Topic,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_Messages_Publish = new grpc.web.MethodDescriptor(
  '/zbay.Messages/Publish',
  grpc.web.MethodType.UNARY,
  proto.zbay.Topic,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.zbay.Topic} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Topic,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Messages_Publish = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.zbay.Topic} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.zbay.Topic} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.publish =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Messages/Publish',
      request,
      metadata || {},
      methodDescriptor_Messages_Publish,
      callback);
};


/**
 * @param {!proto.zbay.Topic} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.zbay.MessagesPromiseClient.prototype.publish =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Messages/Publish',
      request,
      metadata || {},
      methodDescriptor_Messages_Publish);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Topic,
 *   !proto.zbay.Topic>}
 */
const methodDescriptor_Messages_Subscribe = new grpc.web.MethodDescriptor(
  '/zbay.Messages/Subscribe',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.zbay.Topic,
  proto.zbay.Topic,
  /**
   * @param {!proto.zbay.Topic} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Topic.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Topic,
 *   !proto.zbay.Topic>}
 */
const methodInfo_Messages_Subscribe = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Topic,
  /**
   * @param {!proto.zbay.Topic} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Topic.deserializeBinary
);


/**
 * @param {!proto.zbay.Topic} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Topic>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesClient.prototype.subscribe =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/Subscribe',
      request,
      metadata || {},
      methodDescriptor_Messages_Subscribe);
};


/**
 * @param {!proto.zbay.Topic} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Topic>}
 *     The XHR Node Readable Stream
 */
proto.zbay.MessagesPromiseClient.prototype.subscribe =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Messages/Subscribe',
      request,
      metadata || {},
      methodDescriptor_Messages_Subscribe);
};


module.exports = proto.zbay;

