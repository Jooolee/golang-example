// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGenTable = "gen_table"

// GenTable 代码生成业务表
type GenTable struct {
	TableID        int64     `gorm:"column:table_id;primaryKey;autoIncrement:true;comment:编号" json:"table_id"`              // 编号
	TableName_     string    `gorm:"column:table_name;comment:表名称" json:"table_name"`                                       // 表名称
	TableComment   string    `gorm:"column:table_comment;comment:表描述" json:"table_comment"`                                 // 表描述
	ClassName      string    `gorm:"column:class_name;comment:实体类名称" json:"class_name"`                                     // 实体类名称
	TplCategory    string    `gorm:"column:tpl_category;default:crud;comment:使用的模板（crud单表操作 tree树表操作）" json:"tpl_category"` // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string    `gorm:"column:package_name;comment:生成包路径" json:"package_name"`                                 // 生成包路径
	ModuleName     string    `gorm:"column:module_name;comment:生成模块名" json:"module_name"`                                   // 生成模块名
	BusinessName   string    `gorm:"column:business_name;comment:生成业务名" json:"business_name"`                               // 生成业务名
	FunctionName   string    `gorm:"column:function_name;comment:生成功能名" json:"function_name"`                               // 生成功能名
	FunctionAuthor string    `gorm:"column:function_author;comment:生成功能作者" json:"function_author"`                          // 生成功能作者
	Options        string    `gorm:"column:options;comment:其它生成选项" json:"options"`                                          // 其它生成选项
	CreateBy       string    `gorm:"column:create_by;comment:创建者" json:"create_by"`                                         // 创建者
	CreateTime     time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                                    // 创建时间
	UpdateBy       string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                                         // 更新者
	UpdateTime     time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
	Remark         string    `gorm:"column:remark;comment:备注" json:"remark"`                                                // 备注
}

// TableName GenTable's table name
func (*GenTable) TableName() string {
	return TableNameGenTable
}
