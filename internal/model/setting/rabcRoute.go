package setting

type RouteItem struct {
	Module string      `json:"module" description:"模块"`
	Path   string      `json:"path" description:"地址"`
	Method string      `json:"method" description:"请求方式"`
	IsAuth bool        `json:"is_auth" description:"是否验证"`
	Func   interface{} `json:"func" description:"路由方法"`
}
