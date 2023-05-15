package setting

type MenuListItem struct {
	MenuListItemEntity
	Children []MenuListItem `json:"children"`
}

type MenuListItemEntity struct {
	Id         int    `json:"id"       description:"图标"`
	Path       string `json:"path"       description:"图标"`
	UniqueAuth string `json:"unique_auth"      description:"名称"`
	Title      string `json:"title"      description:"名称"`
	Icon       string `json:"icon"       description:"排序"`
	Header     int    `json:"header"     description:"是否隐藏菜单 0隐藏 1显示"`
	IsHeader   string `json:"is_header"       description:"路径"`
}
