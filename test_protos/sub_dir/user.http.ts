/* eslint-disable */
import {
  GetUserInfoReq,
  GetUserInfoResp,
} from './../sub_dir/user';

export type GeneralRequest = <TReq, TResp>(cmd: string, payload: TReq, options?: any) => Promise<TResp>;

export class GeneralClass {
  generalRequestMethod: GeneralRequest;
  constructor(generalRequestMethod: any) {
    this.generalRequestMethod = generalRequestMethod as GeneralRequest;
  };
};

export class UserService extends GeneralClass {
  GetUserInfo(payload: GetUserInfoReq, options?: any): Promise<GetUserInfoResp> {
    return new Promise((resolve, reject) => {
      this.generalRequestMethod<GetUserInfoReq, GetUserInfoResp>('GetUserInfo', payload, options).then((res) => {
        resolve(res);
      })
        .catch((error) => {
          reject(error);
        });
    });
  };
};

