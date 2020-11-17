package tree
func Tree(list []Org, parentID string) *OrgTree {
	orgTree := &OrgTree{}
	for _, value := range list {
		if value.ParentID == parentID {
			orgTree.Org = value
			orgTree.Data = append(orgTree.Data, *Tree(list, value.OrgID))
		}
	}
	return orgTree
}

// 组织模型
type Org struct {
	// 父组织id
	ParentID string `json:"parentID,omitempty"`
	// 组织ID
	OrgID string `json:"orgID,omitempty"`
	// 组织名称
	OrgName string `json:"orgName,omitempty"`
	// 组织描述
	OrgDescription string `json:"OrgDescription,omitempty"`
}


// 组织数据列表
type ListOrg struct {
	ListData []Org `json:"list"`
}

// 组织树
type OrgTree struct {
	Org  Org       `json:"org,omitempty"`
	Data []OrgTree `json:"data,omitempty"`
}

// 组织树列表
type ListOrgTree struct {
	ListOrgTree []OrgTree `json:"list,omitempty"`
}

