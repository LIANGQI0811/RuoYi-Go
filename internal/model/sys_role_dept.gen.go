// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRoleDept = "sys_role_dept"

// SysRoleDept mapped from table <sys_role_dept>
type SysRoleDept struct {
	RoleID int64 `gorm:"column:role_id;primaryKey;comment:角色ID" json:"role_id"` // 角色ID
	DeptID int64 `gorm:"column:dept_id;primaryKey;comment:部门ID" json:"dept_id"` // 部门ID
}

// TableName SysRoleDept's table name
func (*SysRoleDept) TableName() string {
	return TableNameSysRoleDept
}
