package tree

func Tree(list []Org, parentID string) OrgTree {
	var orgTree OrgTree
	for _, value := range list {
		if value.ParentID == parentID {
			orgTree.Org = value
			orgTree.Data = append(orgTree.Data, Tree(list, value.OrgID))
		}
	}
	return orgTree
}

// 组织模型
type Org struct {
	// 父组织id
	ParentID string `json:"parentID"`
	// 组织ID
	OrgID string `json:"orgID"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织描述
	OrgDescription string `json:"OrgDescription"`
	// 创建时间
	CreateAt int64 `json:"createAt"`
	// 最近更新时间
	LastUpdateAt int64 `json:"lastUpdateAt"`
}

// 创建请求数据
type CreateOrg struct {
	// 父组织id
	ParentID string `json:"parentID,omitempty" validate:"@string[0,]"`
	// 组织名称
	OrgName string `json:"orgName" validate:"@string[1,]"`
	// 组织描述
	OrgDescription string `json:"orgDescription,omitempty" validate:"@string[0,]"`
}

// 更新请求数据模型
type UpdateOrg struct {
	// 父组织id
	ParentID string `json:"parentID,omitempty" validate:"@string[0,]"`
	// 组织ID
	OrgID string `json:"orgID" validate:"@string[1,]"`
	// 组织名称
	OrgName string `json:"orgName" validate:"@string[1,]"`
	// 组织描述
	OrgDescription string `json:"orgDescription,omitempty" validate:"@string[0,]"`
}

// 组织数据列表
type ListOrg struct {
	ListData []Org `json:"list"`
}

// 组织树
type OrgTree struct {
	Org  Org       `json:"org,omitempty"`
	Data []OrgTree `json:"data"`
}

// 组织树列表
type ListOrgTree struct {
	ListOrgTree []OrgTree `json:"list"`
}

