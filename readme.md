# protoc-gen-http-ts

> `protoc-gen-http-ts` transforms your `.proto` files into out-of-the-box http method By Typescript!

Welcome any issues. If you have some idea, please leave me an issue!

# Overview

More and More backend services use micro-service arch, all the services are micro-service which communicate by gRPC. When we need to request a certain service, we expose it as http interface. If your team use this arch, the plugin may be helpful to you.

In this arch, HTTP interfaces and gRPC methods are one-to-one correspondence. So it makes generate TS source code which issue http requests by proto file possible.

If we input a proto file just like it:

```protobuf
syntax = "proto3";

package PackageName;
option go_package = "github.com/group/resp/TestService";

import "sub_dir/user.proto";



message SearchRequest {
    string query = 1;
    int32 page_number = 2;  // Which page number do we want?
    int32 result_per_page = 3;  // Number of results to return per page.
}

message ResultItem {
    string title = 1;
    string url = 2;
    string time = 3;
}

message SearchResponse {
    repeated ResultItem results = 1;
    bool isEnd = 2;
    int32 count = 3;
}

service SearchService{
    rpc SearchByKeyword(SearchRequest) returns (SearchResponse){}
    rpc GetUserInfo(GetUserInfoReq) returns (GetUserInfoResp) {}
}
```

We will get TS code:

```typescript
/* eslint-disable */
import { GetUserInfoReq, GetUserInfoResp } from "./sub_dir/user";
import { SearchRequest, SearchResponse } from "./test";

export type GeneralRequest = <TReq, TResp>(
  cmd: string,
  payload: TReq,
  options?: any
) => Promise<TResp>;

export class GeneralClass {
  generalRequestMethod: GeneralRequest;
  constructor(generalRequestMethod: any) {
    this.generalRequestMethod = generalRequestMethod as GeneralRequest;
  }
}

export class SearchService extends GeneralClass {
  /**
   * 按关键词搜索
   */
  searchByKeyword(
    payload: SearchRequest,
    options?: any
  ): Promise<SearchResponse> {
    return new Promise((resolve, reject) => {
      this.generalRequestMethod<SearchRequest, SearchResponse>(
        "searchByKeyword",
        payload,
        options
      )
        .then((res) => {
          resolve(res);
        })
        .catch((error) => {
          reject(error);
        });
    });
  }
  /**
   * 获取用户信息
   */
  getUserInfo(
    payload: GetUserInfoReq,
    options?: any
  ): Promise<GetUserInfoResp> {
    return new Promise((resolve, reject) => {
      this.generalRequestMethod<GetUserInfoReq, GetUserInfoResp>(
        "getUserInfo",
        payload,
        options
      )
        .then((res) => {
          resolve(res);
        })
        .catch((error) => {
          reject(error);
        });
    });
  }
}
```

You can import a certen service class, and instantiate to an object to use all the api methods without any other operation. The constructor need your common http request method in your project. The request method often has very complex logic and hooks which can not describe with some complie options and this is why I did not choose to implement request method myself.

The output code is strongly-typed. Instead of implement compilation of types, I opted to use another plugin to compile types：[stephenh/ts-proto: An idiomatic protobuf generator for TypeScript (github.com)](https://github.com/stephenh/ts-proto)

# How to use

## As a NPM package

If you don't want to compile it yourself, you can use the npm package.

```sh
npm install @eater-altria/protoc-gen-http-ts --save
```

It will download corresponding binary program at `./node_modules/@eater-altria/protoc-gen-http-ts/protoc-gen-http-ts`

Then, you can use it! But now it only support Linux, MacOS and WSL, because I use `wget` to download the binary program. I will resolve this problem quickly!

## Compile it yourself

- Firstly, you need install Golang SDK and the version is higher than 1.18. You can refer to Golang's official website: [The Go Programming Language](https://go.dev/).

- Install dependencies：

  ```sh
  $ go get
  ```

- Compile code:

  ```sh
  $ go build main.go
  ```

- Use it!

  When the compilation is done, you will get a executable file called `main`. Now you can compile your proto files.：

  ```sh
  protoc \
  --proto_path=./test_protos \
  --plugin=protoc-gen-http=./main \
  --plugin=protoc-gen-ts=/usr/local/bin/protoc-gen-ts_proto \
  --http_out=./out \
  --ts_out=./out \
  ./test_protos/*.proto
  ```

  You need install `ts-proto` though to provide interface definition.

- import it to your project

  A rpc service correspond a service class. Import your class ,and provide a common http request method like this:

  ```typescript
  export type GeneralRequest = <TReq, TResp>(
    TReq,
    cmd: string,
    options?: any
  ) => Promise<TResp>;
  ```

  Then, you can request the corresponding interface!

# Compile options

- with`--http_opt=nameCase=xxx`, you can get service methods which has specific name style. There are values supported:
  - camel: camelCase
  - pascal: PascalCase
  - snake: snake_case
