package route

import (
	"context"
	"fmt"
	"goframe/internal/consts"
	"goframe/internal/model/system"
	"goframe/utility/redis"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func NewRoute() *sRoute {
	return &sRoute{}
}

type sRoute struct {
}

var (
	commonRouteKey = "common_routes"
	routesKey      = "routes"
	routeArray     []system.RouteItem
	modules        = garray.NewStrArray(true)
	moduleArr      = garray.NewStrArray(true)
	moduleDefault  string
)

func (s *sRoute) AddRoute(routes ...system.RouteItem) {
	for _, v := range routes {
		routeArray = append(routeArray, v)
	}
}
func (s *sRoute) InitModule(module string, isBool bool) {
	moduleArr = moduleArr.Append(module)
	if isBool {
		moduleDefault = moduleArr.Join("/")
	}
}
func (s *sRoute) InitRoute(group *ghttp.RouterGroup, isPrivate bool) {
	for _, value := range routeArray {
		if value.Module == moduleDefault && value.IsAuth == isPrivate {
			methodTmp := gstr.ToLower(value.Method)
			switch methodTmp {
			case gstr.ToLower(consts.MethodGet):
				group.GET(value.Path, value.Func)
			case gstr.ToLower(consts.MethodPost):
				group.POST(value.Path, value.Func)
			default:
				group.ALL(value.Path, value.Func)
			}
		}
	}
	if isPrivate {
		// 重置
		moduleDefault = ""
		moduleArr = garray.NewStrArray(true)
	}
}

func SaveCache(ctx context.Context, module string) {
	for _, value := range routeArray {
		if value.Module == module {
			var key string
			if value.IsAuth {
				key = routesKey
			} else {
				key = commonRouteKey
			}
			key += "_" + gstr.ToLower(module)
			redis.Connect().SAdd(ctx, key, fmt.Sprintf("/%v%v", value.Module, value.Path))
		}
	}
}
func GetRoute(ctx context.Context, module, key string) []string {
	lens, _ := redis.Connect().SCard(ctx, key)
	if lens <= 0 {
		for _, v := range routeArray {
			var keyName string
			if v.IsAuth {
				keyName = routesKey
			} else {
				keyName = commonRouteKey
			}
			keyName += "_" + gstr.ToLower(module)
			redis.Connect().SAdd(ctx, keyName, fmt.Sprintf("/%v%v", v.Module, v.Path))
		}
	}

	// 处理数组
	var routeArr []string
	res, _ := redis.Connect().SMembers(ctx, key)
	for _, v := range res {
		routeArr = append(routeArr, gconv.String(v))
	}
	return routeArr
}
func GetAuthRoute(ctx context.Context) []string {
	var resArr []string
	for _, v := range modules.Slice() {
		tmp := GetRoute(ctx, v, routesKey)
		resArr = append(resArr, tmp...)
	}
	return resArr
}

func GetCommonRoute(ctx context.Context) []string {
	var resArr []string
	for _, v := range modules.Slice() {
		tmp := GetRoute(ctx, v, commonRouteKey)
		resArr = append(resArr, tmp...)
	}
	return resArr
}
