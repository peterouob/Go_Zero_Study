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

### 封裝response避免每次修改都重複編輯
- 生成模板文件 `goctl template init`
- 更改模板文件會再生成api文黨時進行變更
```tpl
{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		//	{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
		// }
		{{if .HasResp}}response.Response(r,w,resp,err){{else}}reponse.Response(r,w,nil,err){{end}}
```
- `goctl api go -api user.api -dir .`
```api
//入餐，一定要大寫
type LoginRequest {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoResponse {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
}

service users{
	//試圖函數
	@handler login
	post /api/users/login (LoginRequest) returns (string )

	@handler userinfo
	get /api/users/userinfo  returns (UserInfoResponse)
}

// goctl api go -api user.api -dir .
```

```go
resp, err := l.Login(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
		response.Response(r, w, resp, err)
```
### 為api新增路由前綴
```api
//入餐，一定要大寫
type LoginRequest {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoResponse {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
}

@server (
	prefix : /api/users
)

service users{
	//試圖函數
	@handler login
	post /login (LoginRequest) returns (string )

	@handler userinfo
	get /userinfo  returns (UserInfoResponse)
}

// goctl api go -api user.api -dir .
```

### 為api新增JWT
```api
//入餐，一定要大寫
type LoginRequest {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoResponse {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
}

@server (
	prefix : /api/users
)

service users{
	//試圖函數
	@handler login
	post /login (LoginRequest) returns (string )

}

@server (
	prefix : /api/users
	jwt : Auth //固定寫法
)

service users{

	@handler userinfo
	get /userinfo  returns (UserInfoResponse)
}

// goctl api go -api user.api -dir .
```

### 添加生成jwt操作
#### Auth中的JWT到etc資料夾裡面的users.yaml配置
#### 對測試api的工具新增Bearer Token請求頭輸入token資料來驗證和解析
```go
package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayload struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

func GetToken(user JwtPayload, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{user, jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour + time.Duration(expires)))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

func ParseToken(tokenStr string, accessSecret string, expires int64) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

```

### 修改logic/loginlogic.go以獲得jwt
```go
func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	//獲取配置文件
	auth := l.svcCtx.Config.Auth
	token, err := jwt.GetToken(jwt.JwtPayload{
		UserId:   1,
		Username: "peter",
		Role:     1,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}
```
### 修改logic/userinfologic.go以解析jwt並獲得值
```go

func (l *UserinfoLogic) Userinfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	//獲取token值
	userid := l.ctx.Value("userId").(json.Number)

	//遇到不知道的類型錯誤檢查方式
	//fmt.Printf("%v %T",userid,userid)
	uuid, _ := userid.Int64()
	
	username := l.ctx.Value("username").(string)
	return &types.UserInfoResponse{
		UserId:   uint(uuid),
		Username: username,
	}, nil
}
```
### 修改api_jwt底下的users.go完成驗證jwt失敗返回訊息
```go
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err)
	httpx.WriteJson(w, http.StatusOK, response.Body{10087, nil, "認證失敗"})
}

```