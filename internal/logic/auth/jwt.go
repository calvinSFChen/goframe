package auth

import (
	"context"
	"fmt"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe/api/v1/backend/setting/rabc"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/utility/cerrors"
	"goframe/utility/response"
	"goframe/utility/utils"
	"time"
)

type sAuth struct {
	AuthJwt *jwt.GfJWTMiddleware
}

func NewAuthJwt() *sAuth {
	return &sAuth{}
}

var AuthsService = &sAuth{
	AuthJwt: &jwt.GfJWTMiddleware{},
}

func (s *sAuth) AuthJwtService() *jwt.GfJWTMiddleware {
	return AuthsService.AuthJwt
}

// 初始化
func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		//用户的领域名称，必传
		Realm: "asimov",
		// 签名算法
		SigningAlgorithm: "HS256",
		// 签名密钥
		Key: []byte("asimov123"),
		// 时效
		Timeout: time.Minute * 60 * 24,
		// 	token过期后，可凭借旧token获取新token的刷新时间
		MaxRefresh: time.Minute * 24 * 5,
		// 身份验证的key值
		IdentityKey: "id",
		//token检索模式，用于提取token-> Authorization
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// token在请求头时的名称，默认值为Bearer.客户端在header中传入"Authorization":"token xxxxxx"
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,           // 测试或服务器在其他时区可设置该属
		Authenticator:   LoginAuthenticator, // 根据登录信息对用户进行身份验证的回调函数
		Unauthorized:    LoginUnauthorized,  // 处理不进行授权的逻辑
		IdentityHandler: IdentityHandler,    // 解析并设置用户身份信息，并设置身份信息至每次请求中
		PayloadFunc:     PayloadFunc,        // 登录期间的回调的函数
	})
	AuthsService = &sAuth{
		AuthJwt: auth,
	}
}

// PayloadFunc 向web添加额外的有效负载数据。
func PayloadFunc(data interface{}) jwt.MapClaims {
	fmt.Printf("PayloadFunc:  %+v\n", data)
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler 标识
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[AuthsService.AuthJwt.IdentityKey]
}

// LoginUnauthorized 验证不通过
func LoginUnauthorized(ctx context.Context, code int, message string) {
	fmt.Println("验证不通过")
	r := g.RequestFromCtx(ctx)
	codes := gcode.WithCode(cerrors.CodeExpire, nil)
	response.JsonExit(r, codes.Code(), codes.Message(), nil)
	r.ExitAll()
}

// LoginAuthenticator 登录验证
func LoginAuthenticator(ctx context.Context) (interface{}, error) {
	fmt.Println("登录验证")
	var (
		r     = g.RequestFromCtx(ctx)
		input rabc.AdminLoginReq

		rabcAdmin entity.MuRabcAdmin
	)
	err := r.Parse(&input)
	if err != nil {
		return nil, gerror.New(utils.T(ctx, "参数异常"))
	}

	err = dao.MuRabcAdmin.Ctx(ctx).
		WhereOr("username=?", input.Username).
		WhereOr("phone=?", input.Username).
		WhereOr("email=?", input.Username).Scan(&rabcAdmin)
	if err != nil {
		return nil, err
	}

	if !utils.EncryptionVerify(rabcAdmin.Password, input.Password) {
		return nil, gerror.New(utils.T(ctx, "账号或密码有误"))
	}

	return utils.StructToMap(rabcAdmin), nil
}