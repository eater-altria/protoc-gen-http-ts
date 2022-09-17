/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "PackageName";

export interface GetUserInfoResp {
  nickName: string;
  avatarUrl: string;
  gender: number;
  country: string;
  province: string;
  city: string;
  phone: string;
}

export interface GetUserInfoReq {
  userId: string;
}

function createBaseGetUserInfoResp(): GetUserInfoResp {
  return { nickName: "", avatarUrl: "", gender: 0, country: "", province: "", city: "", phone: "" };
}

export const GetUserInfoResp = {
  encode(message: GetUserInfoResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.nickName !== "") {
      writer.uint32(10).string(message.nickName);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(18).string(message.avatarUrl);
    }
    if (message.gender !== 0) {
      writer.uint32(24).int32(message.gender);
    }
    if (message.country !== "") {
      writer.uint32(34).string(message.country);
    }
    if (message.province !== "") {
      writer.uint32(42).string(message.province);
    }
    if (message.city !== "") {
      writer.uint32(50).string(message.city);
    }
    if (message.phone !== "") {
      writer.uint32(58).string(message.phone);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserInfoResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserInfoResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nickName = reader.string();
          break;
        case 2:
          message.avatarUrl = reader.string();
          break;
        case 3:
          message.gender = reader.int32();
          break;
        case 4:
          message.country = reader.string();
          break;
        case 5:
          message.province = reader.string();
          break;
        case 6:
          message.city = reader.string();
          break;
        case 7:
          message.phone = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserInfoResp {
    return {
      nickName: isSet(object.nickName) ? String(object.nickName) : "",
      avatarUrl: isSet(object.avatarUrl) ? String(object.avatarUrl) : "",
      gender: isSet(object.gender) ? Number(object.gender) : 0,
      country: isSet(object.country) ? String(object.country) : "",
      province: isSet(object.province) ? String(object.province) : "",
      city: isSet(object.city) ? String(object.city) : "",
      phone: isSet(object.phone) ? String(object.phone) : "",
    };
  },

  toJSON(message: GetUserInfoResp): unknown {
    const obj: any = {};
    message.nickName !== undefined && (obj.nickName = message.nickName);
    message.avatarUrl !== undefined && (obj.avatarUrl = message.avatarUrl);
    message.gender !== undefined && (obj.gender = Math.round(message.gender));
    message.country !== undefined && (obj.country = message.country);
    message.province !== undefined && (obj.province = message.province);
    message.city !== undefined && (obj.city = message.city);
    message.phone !== undefined && (obj.phone = message.phone);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserInfoResp>, I>>(object: I): GetUserInfoResp {
    const message = createBaseGetUserInfoResp();
    message.nickName = object.nickName ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    message.gender = object.gender ?? 0;
    message.country = object.country ?? "";
    message.province = object.province ?? "";
    message.city = object.city ?? "";
    message.phone = object.phone ?? "";
    return message;
  },
};

function createBaseGetUserInfoReq(): GetUserInfoReq {
  return { userId: "" };
}

export const GetUserInfoReq = {
  encode(message: GetUserInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserInfoReq {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: GetUserInfoReq): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserInfoReq>, I>>(object: I): GetUserInfoReq {
    const message = createBaseGetUserInfoReq();
    message.userId = object.userId ?? "";
    return message;
  },
};

export interface UserService {
  GetUserInfo(request: GetUserInfoReq): Promise<GetUserInfoResp>;
}

export class UserServiceClientImpl implements UserService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetUserInfo = this.GetUserInfo.bind(this);
  }
  GetUserInfo(request: GetUserInfoReq): Promise<GetUserInfoResp> {
    const data = GetUserInfoReq.encode(request).finish();
    const promise = this.rpc.request("PackageName.UserService", "GetUserInfo", data);
    return promise.then((data) => GetUserInfoResp.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
