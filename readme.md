# protoc-gen-http-ts

> `protoc-gen-http-ts` transforms your `.proto` files into out-of-the-box http method By Typescript!

(This photo plugin is in a very early stage of development， so **DO NOT USE IT IN YOUR PRODUCTION ENVIRNMENT!!**)

(Welcome any issues. If you have some idea, please leave me an issue!)

[TOC]

# Overview

More and More backend services use micro-service arch, all the services are micro-service which communicate by gRPC. When we need to request a certain service, we expose it as http interface.  If your group use  this arch, the plugin may be helpful you.

In this model, HTTP interfaces and gRPC methods are one-to-one correspondence. So it makes generate TS source code to make http request by proto file possible.

If we input a proto file just like it: 

```protobuf
syntax = "proto3";

package PackageName;
option go_package = "github.com/group/resp/TestService";



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
}
```

We will get TS code:

```typescript
import {
  SearchResponse,
  SearchRequest,
}from "./test"

export type GeneralRequest = <TReq, TResp>(TReq, cmd: string, options?: any) => Promise<TResp>

export class GeneralClass {
  GeneralRequestMethod: GeneralRequest;
  constructor(GeneralRequestMethod: GeneralRequest) {
    this.GeneralRequestMethod = GeneralRequestMethod;
  }
}

export class SearchService extends GeneralClass {
  constructor(GeneralRequestMethod: GeneralRequest) {
    super(GeneralRequestMethod)
  }
  SearchByKeyword(payload: SearchRequest, options?: any): Promise<SearchResponse> {
    return new Promise((resolve, reject) => {
      this.GeneralRequestMethod<SearchRequest,SearchResponse>(payload, 'SearchByKeyword', options).then(res => {
        resolve(res)
      }).catch(error => {
        reject(error)
      })
    })
  }
}


```

You can import a certen service class, and  instantiate an object to use all the api without any other operation. The constructor need your common http request method. The request method often has very complex logic and hooks which can not describe with some complie options and this is why I did not choose to implement request method myself.

The output code is strongly-typed. Instead of implement compilation of types, I opted to use another plugin to compile types：[stephenh/ts-proto: An idiomatic protobuf generator for TypeScript (github.com)](https://github.com/stephenh/ts-proto)





# How to use



Currently I have not uploaded it to any package management platform. So you need to clone this repository and compile it yourself. After I implement the main function, I will try to make it easier to use.

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
  --proto_path=./ \
  --plugin=protoc-gen-http=./main \
  --plugin=protoc-gen-ts=/usr/local/bin/protoc-gen-ts_proto \
  --http_out=./ \
  --ts_out=./ \
  ./test_protos/*.proto
  ```

  You need install `ts-proto` though to provide interface definition.

- import it to your project

  A rpc service correspond a service class. Import your class ,and provide a common http request method like this:

  ```typescript
  export type GeneralRequest = <TReq, TResp>(TReq, cmd: string, options?: any) => Promise<TResp>
  ```

  Then, you can request the corresponding interface!



# Follow-up planning

- Multi-file compile. Now if input multi files and they has dependency relationship, it can't import interface from current file.
- N+1 problem. If we try to resolve the first problem, there will be a N + 1 promble to overcome.
- More compilation options to adapt to different usage scenarios.
- Easier to install.
