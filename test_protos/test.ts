/* eslint-disable */
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "MPSaas.AccountServer";

export interface QueryCommandKeyInfoReq {
  /** 内部口令 */
  commandKey: string;
}

export interface QueryCommandKeyInfoResp {
  /** 有效截止时间戳 */
  validateTime: number;
}

export interface QueryCommandStatusInfoReq {
}

export interface QueryCommandStatusInfoResp {
  /** 功能开关状态 */
  isSwitch: boolean;
}

export interface UpdateUserPhoneInfoReq {
  /** 微信加密以后的手机号码，该字段适用于2021年7月以后的微信版本 */
  codePhone: string;
  /** 包括敏感数据在内的完整用户信息的加密数据，以下四个字段适用于2021年7月之前的微信版本 */
  encryptedData: string;
  /** 加密算法的初始向量 */
  iv: string;
  /** 登录时用户token */
  token: string;
  /** 小程序appid */
  appid: string;
}

export interface UpdateUserPhoneInfoResp {
}

export interface CheckUserInfoReq {
}

export interface CheckUserInfoResp {
}

export interface QueryUserBaseInfoReq {
}

export interface UpdateUserBaseInfoReq {
  /** 用户信息 */
  userInfo: UserInfo | undefined;
}

export interface UserInfo {
  /** 用户昵称 */
  nickName: string;
  /** 用户头像图片 */
  avatarUrl: string;
  /** 用户性别,0：未知，1：男性，2：女性 */
  gender: number;
  /** 用户所在国家 */
  country: string;
  /** 用户所在省份 */
  province: string;
  /** 用户所在城市 */
  city: string;
  /** 手机号码 */
  phone: string;
}

export interface UpdateUserBaseInfoResp {
}

/** 参考：LoginWithThirdPartyReq in https://git.woa.com/MedicalCommon/goproto/blob/master/MediAccount/AccountToken.proto */
export interface QueryOuterWeComUserTokenReq {
  /** 手机号码/第三方ID */
  id: string;
  /** CorpID是企业号的标识，每个企业号拥有一个唯一的CorpID */
  corpId: string;
  /** 若为空使用header-bin中的platform, 1: H5 2: APP ios */
  platform: number;
}

/** 参考：LoginWithThirdPartyResp in https://git.woa.com/MedicalCommon/goproto/blob/master/MediAccount/AccountToken.proto */
export interface QueryOuterWeComUserTokenResp {
  /** 登录成功后的token */
  token: string;
  /** token有效时间，单位秒 */
  expiration: number;
  /** 登录后的用户ID */
  uin: number;
  /** 错误码，0表示正确 */
  errCode: number;
  /** 错误信息 */
  errMsg: string;
}

function createBaseQueryCommandKeyInfoReq(): QueryCommandKeyInfoReq {
  return { commandKey: "" };
}

export const QueryCommandKeyInfoReq = {
  encode(message: QueryCommandKeyInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.commandKey !== "") {
      writer.uint32(10).string(message.commandKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCommandKeyInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCommandKeyInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commandKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCommandKeyInfoReq {
    return { commandKey: isSet(object.commandKey) ? String(object.commandKey) : "" };
  },

  toJSON(message: QueryCommandKeyInfoReq): unknown {
    const obj: any = {};
    message.commandKey !== undefined && (obj.commandKey = message.commandKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCommandKeyInfoReq>, I>>(object: I): QueryCommandKeyInfoReq {
    const message = createBaseQueryCommandKeyInfoReq();
    message.commandKey = object.commandKey ?? "";
    return message;
  },
};

function createBaseQueryCommandKeyInfoResp(): QueryCommandKeyInfoResp {
  return { validateTime: 0 };
}

export const QueryCommandKeyInfoResp = {
  encode(message: QueryCommandKeyInfoResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validateTime !== 0) {
      writer.uint32(8).int64(message.validateTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCommandKeyInfoResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCommandKeyInfoResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validateTime = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCommandKeyInfoResp {
    return { validateTime: isSet(object.validateTime) ? Number(object.validateTime) : 0 };
  },

  toJSON(message: QueryCommandKeyInfoResp): unknown {
    const obj: any = {};
    message.validateTime !== undefined && (obj.validateTime = Math.round(message.validateTime));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCommandKeyInfoResp>, I>>(object: I): QueryCommandKeyInfoResp {
    const message = createBaseQueryCommandKeyInfoResp();
    message.validateTime = object.validateTime ?? 0;
    return message;
  },
};

function createBaseQueryCommandStatusInfoReq(): QueryCommandStatusInfoReq {
  return {};
}

export const QueryCommandStatusInfoReq = {
  encode(_: QueryCommandStatusInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCommandStatusInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCommandStatusInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryCommandStatusInfoReq {
    return {};
  },

  toJSON(_: QueryCommandStatusInfoReq): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCommandStatusInfoReq>, I>>(_: I): QueryCommandStatusInfoReq {
    const message = createBaseQueryCommandStatusInfoReq();
    return message;
  },
};

function createBaseQueryCommandStatusInfoResp(): QueryCommandStatusInfoResp {
  return { isSwitch: false };
}

export const QueryCommandStatusInfoResp = {
  encode(message: QueryCommandStatusInfoResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.isSwitch === true) {
      writer.uint32(8).bool(message.isSwitch);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCommandStatusInfoResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCommandStatusInfoResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.isSwitch = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCommandStatusInfoResp {
    return { isSwitch: isSet(object.isSwitch) ? Boolean(object.isSwitch) : false };
  },

  toJSON(message: QueryCommandStatusInfoResp): unknown {
    const obj: any = {};
    message.isSwitch !== undefined && (obj.isSwitch = message.isSwitch);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCommandStatusInfoResp>, I>>(object: I): QueryCommandStatusInfoResp {
    const message = createBaseQueryCommandStatusInfoResp();
    message.isSwitch = object.isSwitch ?? false;
    return message;
  },
};

function createBaseUpdateUserPhoneInfoReq(): UpdateUserPhoneInfoReq {
  return { codePhone: "", encryptedData: "", iv: "", token: "", appid: "" };
}

export const UpdateUserPhoneInfoReq = {
  encode(message: UpdateUserPhoneInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.codePhone !== "") {
      writer.uint32(10).string(message.codePhone);
    }
    if (message.encryptedData !== "") {
      writer.uint32(18).string(message.encryptedData);
    }
    if (message.iv !== "") {
      writer.uint32(26).string(message.iv);
    }
    if (message.token !== "") {
      writer.uint32(34).string(message.token);
    }
    if (message.appid !== "") {
      writer.uint32(42).string(message.appid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserPhoneInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserPhoneInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.codePhone = reader.string();
          break;
        case 2:
          message.encryptedData = reader.string();
          break;
        case 3:
          message.iv = reader.string();
          break;
        case 4:
          message.token = reader.string();
          break;
        case 5:
          message.appid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserPhoneInfoReq {
    return {
      codePhone: isSet(object.codePhone) ? String(object.codePhone) : "",
      encryptedData: isSet(object.encryptedData) ? String(object.encryptedData) : "",
      iv: isSet(object.iv) ? String(object.iv) : "",
      token: isSet(object.token) ? String(object.token) : "",
      appid: isSet(object.appid) ? String(object.appid) : "",
    };
  },

  toJSON(message: UpdateUserPhoneInfoReq): unknown {
    const obj: any = {};
    message.codePhone !== undefined && (obj.codePhone = message.codePhone);
    message.encryptedData !== undefined && (obj.encryptedData = message.encryptedData);
    message.iv !== undefined && (obj.iv = message.iv);
    message.token !== undefined && (obj.token = message.token);
    message.appid !== undefined && (obj.appid = message.appid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserPhoneInfoReq>, I>>(object: I): UpdateUserPhoneInfoReq {
    const message = createBaseUpdateUserPhoneInfoReq();
    message.codePhone = object.codePhone ?? "";
    message.encryptedData = object.encryptedData ?? "";
    message.iv = object.iv ?? "";
    message.token = object.token ?? "";
    message.appid = object.appid ?? "";
    return message;
  },
};

function createBaseUpdateUserPhoneInfoResp(): UpdateUserPhoneInfoResp {
  return {};
}

export const UpdateUserPhoneInfoResp = {
  encode(_: UpdateUserPhoneInfoResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserPhoneInfoResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserPhoneInfoResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): UpdateUserPhoneInfoResp {
    return {};
  },

  toJSON(_: UpdateUserPhoneInfoResp): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserPhoneInfoResp>, I>>(_: I): UpdateUserPhoneInfoResp {
    const message = createBaseUpdateUserPhoneInfoResp();
    return message;
  },
};

function createBaseCheckUserInfoReq(): CheckUserInfoReq {
  return {};
}

export const CheckUserInfoReq = {
  encode(_: CheckUserInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CheckUserInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCheckUserInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): CheckUserInfoReq {
    return {};
  },

  toJSON(_: CheckUserInfoReq): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CheckUserInfoReq>, I>>(_: I): CheckUserInfoReq {
    const message = createBaseCheckUserInfoReq();
    return message;
  },
};

function createBaseCheckUserInfoResp(): CheckUserInfoResp {
  return {};
}

export const CheckUserInfoResp = {
  encode(_: CheckUserInfoResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CheckUserInfoResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCheckUserInfoResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): CheckUserInfoResp {
    return {};
  },

  toJSON(_: CheckUserInfoResp): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CheckUserInfoResp>, I>>(_: I): CheckUserInfoResp {
    const message = createBaseCheckUserInfoResp();
    return message;
  },
};

function createBaseQueryUserBaseInfoReq(): QueryUserBaseInfoReq {
  return {};
}

export const QueryUserBaseInfoReq = {
  encode(_: QueryUserBaseInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserBaseInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserBaseInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryUserBaseInfoReq {
    return {};
  },

  toJSON(_: QueryUserBaseInfoReq): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUserBaseInfoReq>, I>>(_: I): QueryUserBaseInfoReq {
    const message = createBaseQueryUserBaseInfoReq();
    return message;
  },
};

function createBaseUpdateUserBaseInfoReq(): UpdateUserBaseInfoReq {
  return { userInfo: undefined };
}

export const UpdateUserBaseInfoReq = {
  encode(message: UpdateUserBaseInfoReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userInfo !== undefined) {
      UserInfo.encode(message.userInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserBaseInfoReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserBaseInfoReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userInfo = UserInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserBaseInfoReq {
    return { userInfo: isSet(object.userInfo) ? UserInfo.fromJSON(object.userInfo) : undefined };
  },

  toJSON(message: UpdateUserBaseInfoReq): unknown {
    const obj: any = {};
    message.userInfo !== undefined && (obj.userInfo = message.userInfo ? UserInfo.toJSON(message.userInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserBaseInfoReq>, I>>(object: I): UpdateUserBaseInfoReq {
    const message = createBaseUpdateUserBaseInfoReq();
    message.userInfo = (object.userInfo !== undefined && object.userInfo !== null)
      ? UserInfo.fromPartial(object.userInfo)
      : undefined;
    return message;
  },
};

function createBaseUserInfo(): UserInfo {
  return { nickName: "", avatarUrl: "", gender: 0, country: "", province: "", city: "", phone: "" };
}

export const UserInfo = {
  encode(message: UserInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): UserInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserInfo();
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

  fromJSON(object: any): UserInfo {
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

  toJSON(message: UserInfo): unknown {
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

  fromPartial<I extends Exact<DeepPartial<UserInfo>, I>>(object: I): UserInfo {
    const message = createBaseUserInfo();
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

function createBaseUpdateUserBaseInfoResp(): UpdateUserBaseInfoResp {
  return {};
}

export const UpdateUserBaseInfoResp = {
  encode(_: UpdateUserBaseInfoResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateUserBaseInfoResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserBaseInfoResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): UpdateUserBaseInfoResp {
    return {};
  },

  toJSON(_: UpdateUserBaseInfoResp): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserBaseInfoResp>, I>>(_: I): UpdateUserBaseInfoResp {
    const message = createBaseUpdateUserBaseInfoResp();
    return message;
  },
};

function createBaseQueryOuterWeComUserTokenReq(): QueryOuterWeComUserTokenReq {
  return { id: "", corpId: "", platform: 0 };
}

export const QueryOuterWeComUserTokenReq = {
  encode(message: QueryOuterWeComUserTokenReq, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.corpId !== "") {
      writer.uint32(18).string(message.corpId);
    }
    if (message.platform !== 0) {
      writer.uint32(24).int32(message.platform);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryOuterWeComUserTokenReq {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOuterWeComUserTokenReq();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.corpId = reader.string();
          break;
        case 3:
          message.platform = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOuterWeComUserTokenReq {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      corpId: isSet(object.corpId) ? String(object.corpId) : "",
      platform: isSet(object.platform) ? Number(object.platform) : 0,
    };
  },

  toJSON(message: QueryOuterWeComUserTokenReq): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.corpId !== undefined && (obj.corpId = message.corpId);
    message.platform !== undefined && (obj.platform = Math.round(message.platform));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryOuterWeComUserTokenReq>, I>>(object: I): QueryOuterWeComUserTokenReq {
    const message = createBaseQueryOuterWeComUserTokenReq();
    message.id = object.id ?? "";
    message.corpId = object.corpId ?? "";
    message.platform = object.platform ?? 0;
    return message;
  },
};

function createBaseQueryOuterWeComUserTokenResp(): QueryOuterWeComUserTokenResp {
  return { token: "", expiration: 0, uin: 0, errCode: 0, errMsg: "" };
}

export const QueryOuterWeComUserTokenResp = {
  encode(message: QueryOuterWeComUserTokenResp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.expiration !== 0) {
      writer.uint32(16).uint32(message.expiration);
    }
    if (message.uin !== 0) {
      writer.uint32(24).int64(message.uin);
    }
    if (message.errCode !== 0) {
      writer.uint32(800).int32(message.errCode);
    }
    if (message.errMsg !== "") {
      writer.uint32(810).string(message.errMsg);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryOuterWeComUserTokenResp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOuterWeComUserTokenResp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.expiration = reader.uint32();
          break;
        case 3:
          message.uin = longToNumber(reader.int64() as Long);
          break;
        case 100:
          message.errCode = reader.int32();
          break;
        case 101:
          message.errMsg = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOuterWeComUserTokenResp {
    return {
      token: isSet(object.token) ? String(object.token) : "",
      expiration: isSet(object.expiration) ? Number(object.expiration) : 0,
      uin: isSet(object.uin) ? Number(object.uin) : 0,
      errCode: isSet(object.errCode) ? Number(object.errCode) : 0,
      errMsg: isSet(object.errMsg) ? String(object.errMsg) : "",
    };
  },

  toJSON(message: QueryOuterWeComUserTokenResp): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.expiration !== undefined && (obj.expiration = Math.round(message.expiration));
    message.uin !== undefined && (obj.uin = Math.round(message.uin));
    message.errCode !== undefined && (obj.errCode = Math.round(message.errCode));
    message.errMsg !== undefined && (obj.errMsg = message.errMsg);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryOuterWeComUserTokenResp>, I>>(object: I): QueryOuterWeComUserTokenResp {
    const message = createBaseQueryOuterWeComUserTokenResp();
    message.token = object.token ?? "";
    message.expiration = object.expiration ?? 0;
    message.uin = object.uin ?? 0;
    message.errCode = object.errCode ?? 0;
    message.errMsg = object.errMsg ?? "";
    return message;
  },
};

export interface AccountService {
  /** 根据biz_userid检验用户是否存在于卫星城账号体系中，若不存在，则前往公共组去拉取该用户信息，并保存至卫星城 */
  CheckUserInfo(request: CheckUserInfoReq): Promise<CheckUserInfoResp>;
  /** 根据biz_userid查询头像、昵称 */
  QueryUserBaseInfo(request: QueryUserBaseInfoReq): Promise<UserInfo>;
  /** 用户更新微信的头像和昵称，卫星城也要相对应的更新 */
  UpdateUserBaseInfo(request: UpdateUserBaseInfoReq): Promise<UpdateUserBaseInfoResp>;
  /** 保存用户手机号码 */
  UpdateUserPhoneInfo(request: UpdateUserPhoneInfoReq): Promise<UpdateUserPhoneInfoResp>;
  /** 功能开关接口 */
  QueryCommandStatusInfo(request: QueryCommandStatusInfoReq): Promise<QueryCommandStatusInfoResp>;
  /** 口令校验接口 */
  QueryCommandKeyInfo(request: QueryCommandKeyInfoReq): Promise<QueryCommandKeyInfoResp>;
  /** 获取企微用户的登录态Token（for 外部WeCom用户, by 白名单） */
  QueryOuterWeComUserToken(request: QueryOuterWeComUserTokenReq): Promise<QueryOuterWeComUserTokenResp>;
}

export class AccountServiceClientImpl implements AccountService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CheckUserInfo = this.CheckUserInfo.bind(this);
    this.QueryUserBaseInfo = this.QueryUserBaseInfo.bind(this);
    this.UpdateUserBaseInfo = this.UpdateUserBaseInfo.bind(this);
    this.UpdateUserPhoneInfo = this.UpdateUserPhoneInfo.bind(this);
    this.QueryCommandStatusInfo = this.QueryCommandStatusInfo.bind(this);
    this.QueryCommandKeyInfo = this.QueryCommandKeyInfo.bind(this);
    this.QueryOuterWeComUserToken = this.QueryOuterWeComUserToken.bind(this);
  }
  CheckUserInfo(request: CheckUserInfoReq): Promise<CheckUserInfoResp> {
    const data = CheckUserInfoReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "CheckUserInfo", data);
    return promise.then((data) => CheckUserInfoResp.decode(new _m0.Reader(data)));
  }

  QueryUserBaseInfo(request: QueryUserBaseInfoReq): Promise<UserInfo> {
    const data = QueryUserBaseInfoReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "QueryUserBaseInfo", data);
    return promise.then((data) => UserInfo.decode(new _m0.Reader(data)));
  }

  UpdateUserBaseInfo(request: UpdateUserBaseInfoReq): Promise<UpdateUserBaseInfoResp> {
    const data = UpdateUserBaseInfoReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "UpdateUserBaseInfo", data);
    return promise.then((data) => UpdateUserBaseInfoResp.decode(new _m0.Reader(data)));
  }

  UpdateUserPhoneInfo(request: UpdateUserPhoneInfoReq): Promise<UpdateUserPhoneInfoResp> {
    const data = UpdateUserPhoneInfoReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "UpdateUserPhoneInfo", data);
    return promise.then((data) => UpdateUserPhoneInfoResp.decode(new _m0.Reader(data)));
  }

  QueryCommandStatusInfo(request: QueryCommandStatusInfoReq): Promise<QueryCommandStatusInfoResp> {
    const data = QueryCommandStatusInfoReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "QueryCommandStatusInfo", data);
    return promise.then((data) => QueryCommandStatusInfoResp.decode(new _m0.Reader(data)));
  }

  QueryCommandKeyInfo(request: QueryCommandKeyInfoReq): Promise<QueryCommandKeyInfoResp> {
    const data = QueryCommandKeyInfoReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "QueryCommandKeyInfo", data);
    return promise.then((data) => QueryCommandKeyInfoResp.decode(new _m0.Reader(data)));
  }

  QueryOuterWeComUserToken(request: QueryOuterWeComUserTokenReq): Promise<QueryOuterWeComUserTokenResp> {
    const data = QueryOuterWeComUserTokenReq.encode(request).finish();
    const promise = this.rpc.request("MPSaas.AccountServer.AccountService", "QueryOuterWeComUserToken", data);
    return promise.then((data) => QueryOuterWeComUserTokenResp.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
