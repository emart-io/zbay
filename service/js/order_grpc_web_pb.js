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

var commodity_pb = require('./commodity_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.zbay = require('./order_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.zbay.OrdersClient =
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
proto.zbay.OrdersPromiseClient =
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
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodDescriptor_Orders_Add = new grpc.web.MethodDescriptor(
  '/zbay.Orders/Add',
  grpc.web.MethodType.UNARY,
  proto.zbay.Order,
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodInfo_Orders_Add = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.Order)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.add =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Orders/Add',
      request,
      metadata || {},
      methodDescriptor_Orders_Add,
      callback);
};


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.Order>}
 *     A native promise that resolves to the response
 */
proto.zbay.OrdersPromiseClient.prototype.add =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Orders/Add',
      request,
      metadata || {},
      methodDescriptor_Orders_Add);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodDescriptor_Orders_Get = new grpc.web.MethodDescriptor(
  '/zbay.Orders/Get',
  grpc.web.MethodType.UNARY,
  proto.zbay.Order,
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodInfo_Orders_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.Order)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Orders/Get',
      request,
      metadata || {},
      methodDescriptor_Orders_Get,
      callback);
};


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.Order>}
 *     A native promise that resolves to the response
 */
proto.zbay.OrdersPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Orders/Get',
      request,
      metadata || {},
      methodDescriptor_Orders_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodDescriptor_Orders_Update = new grpc.web.MethodDescriptor(
  '/zbay.Orders/Update',
  grpc.web.MethodType.UNARY,
  proto.zbay.Order,
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodInfo_Orders_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.Order)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Orders/Update',
      request,
      metadata || {},
      methodDescriptor_Orders_Update,
      callback);
};


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.Order>}
 *     A native promise that resolves to the response
 */
proto.zbay.OrdersPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Orders/Update',
      request,
      metadata || {},
      methodDescriptor_Orders_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Order,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_Orders_Delete = new grpc.web.MethodDescriptor(
  '/zbay.Orders/Delete',
  grpc.web.MethodType.UNARY,
  proto.zbay.Order,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.zbay.Order} request
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
 *   !proto.zbay.Order,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Orders_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Orders/Delete',
      request,
      metadata || {},
      methodDescriptor_Orders_Delete,
      callback);
};


/**
 * @param {!proto.zbay.Order} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.zbay.OrdersPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Orders/Delete',
      request,
      metadata || {},
      methodDescriptor_Orders_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodDescriptor_Orders_ListByOrder = new grpc.web.MethodDescriptor(
  '/zbay.Orders/ListByOrder',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.zbay.Order,
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Order,
 *   !proto.zbay.Order>}
 */
const methodInfo_Orders_ListByOrder = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.Order} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @param {!proto.zbay.Order} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.listByOrder =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Orders/ListByOrder',
      request,
      metadata || {},
      methodDescriptor_Orders_ListByOrder);
};


/**
 * @param {!proto.zbay.Order} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersPromiseClient.prototype.listByOrder =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Orders/ListByOrder',
      request,
      metadata || {},
      methodDescriptor_Orders_ListByOrder);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.ListQuery,
 *   !proto.zbay.Order>}
 */
const methodDescriptor_Orders_ListForBuyer = new grpc.web.MethodDescriptor(
  '/zbay.Orders/ListForBuyer',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.zbay.ListQuery,
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.ListQuery} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.ListQuery,
 *   !proto.zbay.Order>}
 */
const methodInfo_Orders_ListForBuyer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.ListQuery} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @param {!proto.zbay.ListQuery} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.listForBuyer =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Orders/ListForBuyer',
      request,
      metadata || {},
      methodDescriptor_Orders_ListForBuyer);
};


/**
 * @param {!proto.zbay.ListQuery} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersPromiseClient.prototype.listForBuyer =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Orders/ListForBuyer',
      request,
      metadata || {},
      methodDescriptor_Orders_ListForBuyer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.ListQuery,
 *   !proto.zbay.Order>}
 */
const methodDescriptor_Orders_ListForSeller = new grpc.web.MethodDescriptor(
  '/zbay.Orders/ListForSeller',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.zbay.ListQuery,
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.ListQuery} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.ListQuery,
 *   !proto.zbay.Order>}
 */
const methodInfo_Orders_ListForSeller = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Order,
  /**
   * @param {!proto.zbay.ListQuery} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Order.deserializeBinary
);


/**
 * @param {!proto.zbay.ListQuery} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersClient.prototype.listForSeller =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Orders/ListForSeller',
      request,
      metadata || {},
      methodDescriptor_Orders_ListForSeller);
};


/**
 * @param {!proto.zbay.ListQuery} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Order>}
 *     The XHR Node Readable Stream
 */
proto.zbay.OrdersPromiseClient.prototype.listForSeller =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Orders/ListForSeller',
      request,
      metadata || {},
      methodDescriptor_Orders_ListForSeller);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.zbay.AccountsClient =
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
proto.zbay.AccountsPromiseClient =
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
 *   !proto.zbay.Account,
 *   !proto.zbay.Account>}
 */
const methodDescriptor_Accounts_Add = new grpc.web.MethodDescriptor(
  '/zbay.Accounts/Add',
  grpc.web.MethodType.UNARY,
  proto.zbay.Account,
  proto.zbay.Account,
  /**
   * @param {!proto.zbay.Account} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Account.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Account,
 *   !proto.zbay.Account>}
 */
const methodInfo_Accounts_Add = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Account,
  /**
   * @param {!proto.zbay.Account} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Account.deserializeBinary
);


/**
 * @param {!proto.zbay.Account} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.Account)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Account>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsClient.prototype.add =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Accounts/Add',
      request,
      metadata || {},
      methodDescriptor_Accounts_Add,
      callback);
};


/**
 * @param {!proto.zbay.Account} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.Account>}
 *     A native promise that resolves to the response
 */
proto.zbay.AccountsPromiseClient.prototype.add =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Accounts/Add',
      request,
      metadata || {},
      methodDescriptor_Accounts_Add);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.User,
 *   !proto.zbay.Account>}
 */
const methodDescriptor_Accounts_List = new grpc.web.MethodDescriptor(
  '/zbay.Accounts/List',
  grpc.web.MethodType.SERVER_STREAMING,
  user_pb.User,
  proto.zbay.Account,
  /**
   * @param {!proto.zbay.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Account.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.User,
 *   !proto.zbay.Account>}
 */
const methodInfo_Accounts_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Account,
  /**
   * @param {!proto.zbay.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Account.deserializeBinary
);


/**
 * @param {!proto.zbay.User} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Account>}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsClient.prototype.list =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Accounts/List',
      request,
      metadata || {},
      methodDescriptor_Accounts_List);
};


/**
 * @param {!proto.zbay.User} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Account>}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsPromiseClient.prototype.list =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/zbay.Accounts/List',
      request,
      metadata || {},
      methodDescriptor_Accounts_List);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.Account,
 *   !proto.zbay.Account>}
 */
const methodDescriptor_Accounts_Total = new grpc.web.MethodDescriptor(
  '/zbay.Accounts/Total',
  grpc.web.MethodType.UNARY,
  proto.zbay.Account,
  proto.zbay.Account,
  /**
   * @param {!proto.zbay.Account} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Account.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.Account,
 *   !proto.zbay.Account>}
 */
const methodInfo_Accounts_Total = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.Account,
  /**
   * @param {!proto.zbay.Account} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.Account.deserializeBinary
);


/**
 * @param {!proto.zbay.Account} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.Account)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.Account>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsClient.prototype.total =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Accounts/Total',
      request,
      metadata || {},
      methodDescriptor_Accounts_Total,
      callback);
};


/**
 * @param {!proto.zbay.Account} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.Account>}
 *     A native promise that resolves to the response
 */
proto.zbay.AccountsPromiseClient.prototype.total =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Accounts/Total',
      request,
      metadata || {},
      methodDescriptor_Accounts_Total);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.PayMap,
 *   !proto.zbay.PayMap>}
 */
const methodDescriptor_Accounts_Alipay = new grpc.web.MethodDescriptor(
  '/zbay.Accounts/Alipay',
  grpc.web.MethodType.UNARY,
  proto.zbay.PayMap,
  proto.zbay.PayMap,
  /**
   * @param {!proto.zbay.PayMap} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.PayMap.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.PayMap,
 *   !proto.zbay.PayMap>}
 */
const methodInfo_Accounts_Alipay = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.PayMap,
  /**
   * @param {!proto.zbay.PayMap} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.PayMap.deserializeBinary
);


/**
 * @param {!proto.zbay.PayMap} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.PayMap)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.PayMap>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsClient.prototype.alipay =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Accounts/Alipay',
      request,
      metadata || {},
      methodDescriptor_Accounts_Alipay,
      callback);
};


/**
 * @param {!proto.zbay.PayMap} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.PayMap>}
 *     A native promise that resolves to the response
 */
proto.zbay.AccountsPromiseClient.prototype.alipay =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Accounts/Alipay',
      request,
      metadata || {},
      methodDescriptor_Accounts_Alipay);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.PayMap,
 *   !proto.zbay.PayMap>}
 */
const methodDescriptor_Accounts_WechatHtml5Pay = new grpc.web.MethodDescriptor(
  '/zbay.Accounts/WechatHtml5Pay',
  grpc.web.MethodType.UNARY,
  proto.zbay.PayMap,
  proto.zbay.PayMap,
  /**
   * @param {!proto.zbay.PayMap} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.PayMap.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.PayMap,
 *   !proto.zbay.PayMap>}
 */
const methodInfo_Accounts_WechatHtml5Pay = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.PayMap,
  /**
   * @param {!proto.zbay.PayMap} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.PayMap.deserializeBinary
);


/**
 * @param {!proto.zbay.PayMap} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.PayMap)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.PayMap>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsClient.prototype.wechatHtml5Pay =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Accounts/WechatHtml5Pay',
      request,
      metadata || {},
      methodDescriptor_Accounts_WechatHtml5Pay,
      callback);
};


/**
 * @param {!proto.zbay.PayMap} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.PayMap>}
 *     A native promise that resolves to the response
 */
proto.zbay.AccountsPromiseClient.prototype.wechatHtml5Pay =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Accounts/WechatHtml5Pay',
      request,
      metadata || {},
      methodDescriptor_Accounts_WechatHtml5Pay);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.zbay.PayMap,
 *   !proto.zbay.PayMap>}
 */
const methodDescriptor_Accounts_WechatJSPay = new grpc.web.MethodDescriptor(
  '/zbay.Accounts/WechatJSPay',
  grpc.web.MethodType.UNARY,
  proto.zbay.PayMap,
  proto.zbay.PayMap,
  /**
   * @param {!proto.zbay.PayMap} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.PayMap.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.zbay.PayMap,
 *   !proto.zbay.PayMap>}
 */
const methodInfo_Accounts_WechatJSPay = new grpc.web.AbstractClientBase.MethodInfo(
  proto.zbay.PayMap,
  /**
   * @param {!proto.zbay.PayMap} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.zbay.PayMap.deserializeBinary
);


/**
 * @param {!proto.zbay.PayMap} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.zbay.PayMap)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.zbay.PayMap>|undefined}
 *     The XHR Node Readable Stream
 */
proto.zbay.AccountsClient.prototype.wechatJSPay =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/zbay.Accounts/WechatJSPay',
      request,
      metadata || {},
      methodDescriptor_Accounts_WechatJSPay,
      callback);
};


/**
 * @param {!proto.zbay.PayMap} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.zbay.PayMap>}
 *     A native promise that resolves to the response
 */
proto.zbay.AccountsPromiseClient.prototype.wechatJSPay =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/zbay.Accounts/WechatJSPay',
      request,
      metadata || {},
      methodDescriptor_Accounts_WechatJSPay);
};


module.exports = proto.zbay;

