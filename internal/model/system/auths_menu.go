package system

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
}

type MenuTreeListItem struct {
	Label    string             `json:"label" description:"标题"`
	Value    string             `json:"value" description:"值"`
	Children []MenuTreeListItem `json:"children" description:"子项"`
}


type MenuEntity struct {
	Id         int    `json:"id"       description:"ID"`
	Path       string `json:"path"       description:"路径"`
	ApiUrl string `json:"api_url"      description:"接口"`
}