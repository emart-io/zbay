/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.zbay.Address', null, global);
goog.exportSymbol('proto.zbay.Certification', null, global);
goog.exportSymbol('proto.zbay.User', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.zbay.User = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.zbay.User, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.zbay.User.displayName = 'proto.zbay.User';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.zbay.Certification = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.zbay.Certification, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.zbay.Certification.displayName = 'proto.zbay.Certification';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.zbay.Address = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.zbay.Address, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.zbay.Address.displayName = 'proto.zbay.Address';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.zbay.User.prototype.toObject = function(opt_includeInstance) {
  return proto.zbay.User.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.zbay.User} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zbay.User.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    password: jspb.Message.getFieldWithDefault(msg, 3, ""),
    telephone: jspb.Message.getFieldWithDefault(msg, 4, ""),
    icon: jspb.Message.getFieldWithDefault(msg, 5, ""),
    signature: jspb.Message.getFieldWithDefault(msg, 6, ""),
    cert: (f = msg.getCert()) && proto.zbay.Certification.toObject(includeInstance, f),
    annotationsMap: (f = msg.getAnnotationsMap()) ? f.toObject(includeInstance, undefined) : [],
    created: (f = msg.getCreated()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.zbay.User}
 */
proto.zbay.User.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.zbay.User;
  return proto.zbay.User.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.zbay.User} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.zbay.User}
 */
proto.zbay.User.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setTelephone(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setIcon(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setSignature(value);
      break;
    case 9:
      var value = new proto.zbay.Certification;
      reader.readMessage(value,proto.zbay.Certification.deserializeBinaryFromReader);
      msg.setCert(value);
      break;
    case 7:
      var value = msg.getAnnotationsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 8:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreated(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.zbay.User.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.zbay.User.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.zbay.User} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zbay.User.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getTelephone();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getIcon();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getSignature();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getCert();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      proto.zbay.Certification.serializeBinaryToWriter
    );
  }
  f = message.getAnnotationsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(7, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreated();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


Object.defineProperty(proto.zbay.User.prototype, "id", {
  set: function(value) {
    this.setId(value);
  },
  get: function() {
    return this.getId();
  },
});


/**
 * optional string id = 1;
 * @return {string}
 */
proto.zbay.User.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.zbay.User.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


Object.defineProperty(proto.zbay.User.prototype, "name", {
  set: function(value) {
    this.setName(value);
  },
  get: function() {
    return this.getName();
  },
});


/**
 * optional string name = 2;
 * @return {string}
 */
proto.zbay.User.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.zbay.User.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


Object.defineProperty(proto.zbay.User.prototype, "password", {
  set: function(value) {
    this.setPassword(value);
  },
  get: function() {
    return this.getPassword();
  },
});


/**
 * optional string password = 3;
 * @return {string}
 */
proto.zbay.User.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.zbay.User.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


Object.defineProperty(proto.zbay.User.prototype, "telephone", {
  set: function(value) {
    this.setTelephone(value);
  },
  get: function() {
    return this.getTelephone();
  },
});


/**
 * optional string telephone = 4;
 * @return {string}
 */
proto.zbay.User.prototype.getTelephone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.zbay.User.prototype.setTelephone = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


Object.defineProperty(proto.zbay.User.prototype, "icon", {
  set: function(value) {
    this.setIcon(value);
  },
  get: function() {
    return this.getIcon();
  },
});


/**
 * optional string icon = 5;
 * @return {string}
 */
proto.zbay.User.prototype.getIcon = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.zbay.User.prototype.setIcon = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


Object.defineProperty(proto.zbay.User.prototype, "signature", {
  set: function(value) {
    this.setSignature(value);
  },
  get: function() {
    return this.getSignature();
  },
});


/**
 * optional string signature = 6;
 * @return {string}
 */
proto.zbay.User.prototype.getSignature = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.zbay.User.prototype.setSignature = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


Object.defineProperty(proto.zbay.User.prototype, "cert", {
  set: function(value) {
    this.setCert(value);
  },
  get: function() {
    return this.getCert();
  },
});


/**
 * optional Certification cert = 9;
 * @return {?proto.zbay.Certification}
 */
proto.zbay.User.prototype.getCert = function() {
  return /** @type{?proto.zbay.Certification} */ (
    jspb.Message.getWrapperField(this, proto.zbay.Certification, 9));
};


/** @param {?proto.zbay.Certification|undefined} value */
proto.zbay.User.prototype.setCert = function(value) {
  jspb.Message.setWrapperField(this, 9, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.zbay.User.prototype.clearCert = function() {
  this.setCert(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.zbay.User.prototype.hasCert = function() {
  return jspb.Message.getField(this, 9) != null;
};


Object.defineProperty(proto.zbay.User.prototype, "annotationsMap", {
  set: function(value) {
    this.setAnnotationsMap(value);
  },
  get: function() {
    return this.getAnnotationsMap();
  },
});


/**
 * map<string, string> annotations = 7;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.zbay.User.prototype.getAnnotationsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 7, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 */
proto.zbay.User.prototype.clearAnnotationsMap = function() {
  this.getAnnotationsMap().clear();
};


Object.defineProperty(proto.zbay.User.prototype, "created", {
  set: function(value) {
    this.setCreated(value);
  },
  get: function() {
    return this.getCreated();
  },
});


/**
 * optional google.protobuf.Timestamp created = 8;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.zbay.User.prototype.getCreated = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 8));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.zbay.User.prototype.setCreated = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.zbay.User.prototype.clearCreated = function() {
  this.setCreated(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.zbay.User.prototype.hasCreated = function() {
  return jspb.Message.getField(this, 8) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.zbay.Certification.prototype.toObject = function(opt_includeInstance) {
  return proto.zbay.Certification.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.zbay.Certification} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zbay.Certification.toObject = function(includeInstance, msg) {
  var f, obj = {
    fullName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    idCardNo: jspb.Message.getFieldWithDefault(msg, 2, ""),
    idCardFront: jspb.Message.getFieldWithDefault(msg, 3, ""),
    idCardBack: jspb.Message.getFieldWithDefault(msg, 4, ""),
    livePhoto: jspb.Message.getFieldWithDefault(msg, 5, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.zbay.Certification}
 */
proto.zbay.Certification.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.zbay.Certification;
  return proto.zbay.Certification.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.zbay.Certification} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.zbay.Certification}
 */
proto.zbay.Certification.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFullName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setIdCardNo(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setIdCardFront(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setIdCardBack(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setLivePhoto(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.zbay.Certification.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.zbay.Certification.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.zbay.Certification} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zbay.Certification.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFullName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getIdCardNo();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getIdCardFront();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getIdCardBack();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLivePhoto();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


Object.defineProperty(proto.zbay.Certification.prototype, "fullName", {
  set: function(value) {
    this.setFullName(value);
  },
  get: function() {
    return this.getFullName();
  },
});


/**
 * optional string full_name = 1;
 * @return {string}
 */
proto.zbay.Certification.prototype.getFullName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.zbay.Certification.prototype.setFullName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


Object.defineProperty(proto.zbay.Certification.prototype, "idCardNo", {
  set: function(value) {
    this.setIdCardNo(value);
  },
  get: function() {
    return this.getIdCardNo();
  },
});


/**
 * optional string id_card_no = 2;
 * @return {string}
 */
proto.zbay.Certification.prototype.getIdCardNo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.zbay.Certification.prototype.setIdCardNo = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


Object.defineProperty(proto.zbay.Certification.prototype, "idCardFront", {
  set: function(value) {
    this.setIdCardFront(value);
  },
  get: function() {
    return this.getIdCardFront();
  },
});


/**
 * optional string id_card_front = 3;
 * @return {string}
 */
proto.zbay.Certification.prototype.getIdCardFront = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.zbay.Certification.prototype.setIdCardFront = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


Object.defineProperty(proto.zbay.Certification.prototype, "idCardBack", {
  set: function(value) {
    this.setIdCardBack(value);
  },
  get: function() {
    return this.getIdCardBack();
  },
});


/**
 * optional string id_card_back = 4;
 * @return {string}
 */
proto.zbay.Certification.prototype.getIdCardBack = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.zbay.Certification.prototype.setIdCardBack = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


Object.defineProperty(proto.zbay.Certification.prototype, "livePhoto", {
  set: function(value) {
    this.setLivePhoto(value);
  },
  get: function() {
    return this.getLivePhoto();
  },
});


/**
 * optional string live_photo = 5;
 * @return {string}
 */
proto.zbay.Certification.prototype.getLivePhoto = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.zbay.Certification.prototype.setLivePhoto = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.zbay.Address.prototype.toObject = function(opt_includeInstance) {
  return proto.zbay.Address.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.zbay.Address} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zbay.Address.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    userId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    contact: jspb.Message.getFieldWithDefault(msg, 3, ""),
    telephone: jspb.Message.getFieldWithDefault(msg, 4, ""),
    location: jspb.Message.getFieldWithDefault(msg, 5, ""),
    pb_default: jspb.Message.getBooleanFieldWithDefault(msg, 8, false),
    annotationsMap: (f = msg.getAnnotationsMap()) ? f.toObject(includeInstance, undefined) : [],
    created: (f = msg.getCreated()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.zbay.Address}
 */
proto.zbay.Address.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.zbay.Address;
  return proto.zbay.Address.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.zbay.Address} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.zbay.Address}
 */
proto.zbay.Address.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setContact(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setTelephone(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocation(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDefault(value);
      break;
    case 6:
      var value = msg.getAnnotationsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 7:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreated(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.zbay.Address.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.zbay.Address.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.zbay.Address} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zbay.Address.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getContact();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getTelephone();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLocation();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getDefault();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getAnnotationsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreated();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


Object.defineProperty(proto.zbay.Address.prototype, "id", {
  set: function(value) {
    this.setId(value);
  },
  get: function() {
    return this.getId();
  },
});


/**
 * optional string id = 1;
 * @return {string}
 */
proto.zbay.Address.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.zbay.Address.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


Object.defineProperty(proto.zbay.Address.prototype, "userId", {
  set: function(value) {
    this.setUserId(value);
  },
  get: function() {
    return this.getUserId();
  },
});


/**
 * optional string userId = 2;
 * @return {string}
 */
proto.zbay.Address.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.zbay.Address.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


Object.defineProperty(proto.zbay.Address.prototype, "contact", {
  set: function(value) {
    this.setContact(value);
  },
  get: function() {
    return this.getContact();
  },
});


/**
 * optional string contact = 3;
 * @return {string}
 */
proto.zbay.Address.prototype.getContact = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.zbay.Address.prototype.setContact = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


Object.defineProperty(proto.zbay.Address.prototype, "telephone", {
  set: function(value) {
    this.setTelephone(value);
  },
  get: function() {
    return this.getTelephone();
  },
});


/**
 * optional string telephone = 4;
 * @return {string}
 */
proto.zbay.Address.prototype.getTelephone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.zbay.Address.prototype.setTelephone = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


Object.defineProperty(proto.zbay.Address.prototype, "location", {
  set: function(value) {
    this.setLocation(value);
  },
  get: function() {
    return this.getLocation();
  },
});


/**
 * optional string location = 5;
 * @return {string}
 */
proto.zbay.Address.prototype.getLocation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.zbay.Address.prototype.setLocation = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


Object.defineProperty(proto.zbay.Address.prototype, "pb_default", {
  set: function(value) {
    this.setDefault(value);
  },
  get: function() {
    return this.getDefault();
  },
});


/**
 * optional bool default = 8;
 * @return {boolean}
 */
proto.zbay.Address.prototype.getDefault = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/** @param {boolean} value */
proto.zbay.Address.prototype.setDefault = function(value) {
  jspb.Message.setProto3BooleanField(this, 8, value);
};


Object.defineProperty(proto.zbay.Address.prototype, "annotationsMap", {
  set: function(value) {
    this.setAnnotationsMap(value);
  },
  get: function() {
    return this.getAnnotationsMap();
  },
});


/**
 * map<string, string> annotations = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.zbay.Address.prototype.getAnnotationsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 */
proto.zbay.Address.prototype.clearAnnotationsMap = function() {
  this.getAnnotationsMap().clear();
};


Object.defineProperty(proto.zbay.Address.prototype, "created", {
  set: function(value) {
    this.setCreated(value);
  },
  get: function() {
    return this.getCreated();
  },
});


/**
 * optional google.protobuf.Timestamp created = 7;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.zbay.Address.prototype.getCreated = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 7));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.zbay.Address.prototype.setCreated = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.zbay.Address.prototype.clearCreated = function() {
  this.setCreated(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.zbay.Address.prototype.hasCreated = function() {
  return jspb.Message.getField(this, 7) != null;
};


goog.object.extend(exports, proto.zbay);
