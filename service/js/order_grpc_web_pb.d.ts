import * as grpcWeb from 'grpc-web';

import * as user_pb from './user_pb';
import * as commodity_pb from './commodity_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import {Order} from './order_pb';

export class OrdersClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  add(
    request: Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Order) => void
  ): grpcWeb.ClientReadableStream<Order>;

  get(
    request: Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Order) => void
  ): grpcWeb.ClientReadableStream<Order>;

  update(
    request: Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: Order) => void
  ): grpcWeb.ClientReadableStream<Order>;

  delete(
    request: Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  listByUser(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<Order>;

  listByOrder(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<Order>;

  signAlipay(
    request: Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.StringValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.StringValue>;

}

export class OrdersPromiseClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  add(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): Promise<Order>;

  get(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): Promise<Order>;

  update(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): Promise<Order>;

  delete(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  listByUser(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<Order>;

  listByOrder(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<Order>;

  signAlipay(
    request: Order,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.StringValue>;

}

