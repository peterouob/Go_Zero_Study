# DEMO go_zero

### 新建proto文件
```protobuf
syntax = "proto3";

package user;

option go_package = "./user";

message IdRequest {
  string id = 1;
}

message UserResponse {
  string id = 1;
  string name = 2;
  bool gender = 3;
}

service User{
  rpc getUser(IdRequest) returns(UserResponse);
}

// goctl rpc protoc user/rpc/user.proto --go_out=user/rpc/types --go-grpc_out=user/rpc/types --zrpc_out=user/rpc/
// goctl rpc protoc user.proto --go_out=types --go-grpc_out=types --zrpc_out=.
```
- 轉換proto文件
  - `goctl rpc protoc user/rpc/user.proto --go_out=user/rpc/types --go-grpc_out=user/rpc/types --zrpc_out=user/rpc/`
  - OR
  - `goctl rpc protoc user.proto --go_out=types --go-grpc_out=types --zrpc_out=.`
### 修改proto文件
- 將要改變地方寫在/internal/logic/getuserlogic 裡面
### 使用postman發送服務
- 設定好地址 : grpc://localhost:[port]
- 導入proto 後按next再選擇設定地址旁邊的input

### 新建api文件
```api
type (
	VideoReq {
		Id string `path:"id"`
	}
	VideoRes {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
)
service video {
	@handler getVideo
	get /api/videos/:id (VideoReq) returns (VideoRes)
}

```
- 轉換api文件
    - `goctl api go -api video/api/video.api -dir video/api/`
### 添加grpc服務
    - 修改/internal/config/config.go
```golang
type Config struct {
    rest.RestConf
    UserRpc zrpc.RpcClientConf
}

```

### 完善grpc依賴
    - 修改/internal/svc/servicecontext.go
```golang
type ServiceContext struct {
  Config  config.Config
  UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
  return &ServiceContext{
    Config:  c,
    UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
  }
}

```
### 修改api文件以添加請求服務
    - 將要改變地方寫在/internal/logic/getVideologic 裡面
```go
func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error) {
	// todo: add your logic here and delete this line
	user1, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "12",
	})
	if err != nil {
		panic(err)
	}
	return &types.VideoRes{Id: req.Id, Name: user1.Name}, nil
}
```

## 啟動服務
1. 開啟etcd服務器，終端輸入etcd (在專案根目錄執行)
2. 開啟rpc服務 go run user.go 
3. 開啟api服務 go run api.go
### 如果遇到明明可以請求rpc但api始終報[rpc服務名稱].rpc未開啟
1. 將yaml配置文欓中的地址改為localhost而非etcd
2. 重開電腦