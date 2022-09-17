import {
  GetUserInfoReq,
  GetUserInfoResp,
}from "./../sub_dir/user"

export type GeneralRequest = <TReq, TResp>(TReq, cmd: string, options?: any) => Promise<TResp>

export class GeneralClass {
  GeneralRequestMethod: GeneralRequest;
  constructor(GeneralRequestMethod: GeneralRequest) {
    this.GeneralRequestMethod = GeneralRequestMethod;
  }
}

export class UserService extends GeneralClass {
  constructor(GeneralRequestMethod: GeneralRequest) {
    super(GeneralRequestMethod)
  }
  GetUserInfo(payload: GetUserInfoReq, options?: any): Promise<GetUserInfoResp> {
    return new Promise((resolve, reject) => {
      this.GeneralRequestMethod<GetUserInfoReq,GetUserInfoResp>(payload, 'GetUserInfo', options).then(res => {
        resolve(res)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

