package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	o1 := Org{
		ParentID: "",
		OrgID:    "1",
		OrgName: "zs",
	}
	o2 := Org{
		ParentID: "1",
		OrgID:    "2",
		OrgName: "zs",
	}
	o3 := Org{
		ParentID: "2",
		OrgID:    "3",
		OrgName: "zs",
	}
	list := new(ListOrg)
	list.ListData = append(list.ListData, o1)
	list.ListData = append(list.ListData, o2)
	list.ListData = append(list.ListData, o3)
	tree := Tree(list.ListData, "")
	t.Log(tree)
}