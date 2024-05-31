// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameSysConfig = "sys_config"

// SysConfig mapped from table <sys_config>
type SysConfig struct {
	ConfigID    int32     `gorm:"column:config_id;primaryKey;comment:参数主键" json:"config_id"` // 参数主键
	ConfigName  string    `gorm:"column:config_name;comment:参数名称" json:"config_name"`        // 参数名称
	ConfigKey   string    `gorm:"column:config_key;comment:参数键名" json:"config_key"`          // 参数键名
	ConfigValue string    `gorm:"column:config_value;comment:参数键值" json:"config_value"`      // 参数键值
	ConfigType  string    `gorm:"column:config_type;comment:系统内置（Y是 N否）" json:"config_type"` // 系统内置（Y是 N否）
	CreateBy    string    `gorm:"column:create_by;comment:创建者" json:"create_by"`             // 创建者
	CreateTime  time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`        // 创建时间
	UpdateBy    string    `gorm:"column:update_by;comment:更新者" json:"update_by"`             // 更新者
	UpdateTime  time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`        // 更新时间
	Remark      string    `gorm:"column:remark;comment:备注" json:"remark"`                    // 备注
}

// TableName SysConfig's table name
func (*SysConfig) TableName() string {
	return TableNameSysConfig
}