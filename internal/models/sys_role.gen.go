// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	RoleID            int64     `gorm:"column:role_id;primaryKey;comment:角色ID" json:"role_id"`                                      // 角色ID
	RoleName          string    `gorm:"column:role_name;not null;comment:角色名称" json:"role_name"`                                    // 角色名称
	RoleKey           string    `gorm:"column:role_key;not null;comment:角色权限字符串" json:"role_key"`                                   // 角色权限字符串
	RoleSort          int32     `gorm:"column:role_sort;not null;comment:显示顺序" json:"role_sort"`                                    // 显示顺序
	DataScope         string    `gorm:"column:data_scope;comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）" json:"data_scope"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly int16     `gorm:"column:menu_check_strictly;comment:菜单树选择项是否关联显示" json:"menu_check_strictly"`                 // 菜单树选择项是否关联显示
	DeptCheckStrictly int16     `gorm:"column:dept_check_strictly;comment:部门树选择项是否关联显示" json:"dept_check_strictly"`                 // 部门树选择项是否关联显示
	Status            string    `gorm:"column:status;not null;comment:角色状态（0正常 1停用）" json:"status"`                                 // 角色状态（0正常 1停用）
	DelFlag           string    `gorm:"column:del_flag;comment:删除标志（0代表存在 2代表删除）" json:"del_flag"`                                  // 删除标志（0代表存在 2代表删除）
	CreateBy          string    `gorm:"column:create_by;comment:创建者" json:"create_by"`                                              // 创建者
	CreateTime        time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                                         // 创建时间
	UpdateBy          string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                                              // 更新者
	UpdateTime        time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                         // 更新时间
	Remark            string    `gorm:"column:remark;comment:备注" json:"remark"`                                                     // 备注
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
