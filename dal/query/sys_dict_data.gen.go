// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"golang-example/dal/model"
)

func newSysDictDatum(db *gorm.DB, opts ...gen.DOOption) sysDictDatum {
	_sysDictDatum := sysDictDatum{}

	_sysDictDatum.sysDictDatumDo.UseDB(db, opts...)
	_sysDictDatum.sysDictDatumDo.UseModel(&model.SysDictDatum{})

	tableName := _sysDictDatum.sysDictDatumDo.TableName()
	_sysDictDatum.ALL = field.NewAsterisk(tableName)
	_sysDictDatum.DictCode = field.NewInt64(tableName, "dict_code")
	_sysDictDatum.DictSort = field.NewInt32(tableName, "dict_sort")
	_sysDictDatum.DictLabel = field.NewString(tableName, "dict_label")
	_sysDictDatum.DictValue = field.NewString(tableName, "dict_value")
	_sysDictDatum.DictType = field.NewString(tableName, "dict_type")
	_sysDictDatum.CSSClass = field.NewString(tableName, "css_class")
	_sysDictDatum.ListClass = field.NewString(tableName, "list_class")
	_sysDictDatum.IsDefault = field.NewString(tableName, "is_default")
	_sysDictDatum.Status = field.NewString(tableName, "status")
	_sysDictDatum.CreateBy = field.NewString(tableName, "create_by")
	_sysDictDatum.CreateTime = field.NewTime(tableName, "create_time")
	_sysDictDatum.UpdateBy = field.NewString(tableName, "update_by")
	_sysDictDatum.UpdateTime = field.NewTime(tableName, "update_time")
	_sysDictDatum.Remark = field.NewString(tableName, "remark")

	_sysDictDatum.fillFieldMap()

	return _sysDictDatum
}

// sysDictDatum 字典数据表
type sysDictDatum struct {
	sysDictDatumDo

	ALL        field.Asterisk
	DictCode   field.Int64  // 字典编码
	DictSort   field.Int32  // 字典排序
	DictLabel  field.String // 字典标签
	DictValue  field.String // 字典键值
	DictType   field.String // 字典类型
	CSSClass   field.String // 样式属性（其他样式扩展）
	ListClass  field.String // 表格回显样式
	IsDefault  field.String // 是否默认（Y是 N否）
	Status     field.String // 状态（0正常 1停用）
	CreateBy   field.String // 创建者
	CreateTime field.Time   // 创建时间
	UpdateBy   field.String // 更新者
	UpdateTime field.Time   // 更新时间
	Remark     field.String // 备注

	fieldMap map[string]field.Expr
}

func (s sysDictDatum) Table(newTableName string) *sysDictDatum {
	s.sysDictDatumDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictDatum) As(alias string) *sysDictDatum {
	s.sysDictDatumDo.DO = *(s.sysDictDatumDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictDatum) updateTableName(table string) *sysDictDatum {
	s.ALL = field.NewAsterisk(table)
	s.DictCode = field.NewInt64(table, "dict_code")
	s.DictSort = field.NewInt32(table, "dict_sort")
	s.DictLabel = field.NewString(table, "dict_label")
	s.DictValue = field.NewString(table, "dict_value")
	s.DictType = field.NewString(table, "dict_type")
	s.CSSClass = field.NewString(table, "css_class")
	s.ListClass = field.NewString(table, "list_class")
	s.IsDefault = field.NewString(table, "is_default")
	s.Status = field.NewString(table, "status")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysDictDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictDatum) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["dict_code"] = s.DictCode
	s.fieldMap["dict_sort"] = s.DictSort
	s.fieldMap["dict_label"] = s.DictLabel
	s.fieldMap["dict_value"] = s.DictValue
	s.fieldMap["dict_type"] = s.DictType
	s.fieldMap["css_class"] = s.CSSClass
	s.fieldMap["list_class"] = s.ListClass
	s.fieldMap["is_default"] = s.IsDefault
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
}

func (s sysDictDatum) clone(db *gorm.DB) sysDictDatum {
	s.sysDictDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictDatum) replaceDB(db *gorm.DB) sysDictDatum {
	s.sysDictDatumDo.ReplaceDB(db)
	return s
}

type sysDictDatumDo struct{ gen.DO }

type ISysDictDatumDo interface {
	gen.SubQuery
	Debug() ISysDictDatumDo
	WithContext(ctx context.Context) ISysDictDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysDictDatumDo
	WriteDB() ISysDictDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysDictDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysDictDatumDo
	Not(conds ...gen.Condition) ISysDictDatumDo
	Or(conds ...gen.Condition) ISysDictDatumDo
	Select(conds ...field.Expr) ISysDictDatumDo
	Where(conds ...gen.Condition) ISysDictDatumDo
	Order(conds ...field.Expr) ISysDictDatumDo
	Distinct(cols ...field.Expr) ISysDictDatumDo
	Omit(cols ...field.Expr) ISysDictDatumDo
	Join(table schema.Tabler, on ...field.Expr) ISysDictDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysDictDatumDo
	Group(cols ...field.Expr) ISysDictDatumDo
	Having(conds ...gen.Condition) ISysDictDatumDo
	Limit(limit int) ISysDictDatumDo
	Offset(offset int) ISysDictDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictDatumDo
	Unscoped() ISysDictDatumDo
	Create(values ...*model.SysDictDatum) error
	CreateInBatches(values []*model.SysDictDatum, batchSize int) error
	Save(values ...*model.SysDictDatum) error
	First() (*model.SysDictDatum, error)
	Take() (*model.SysDictDatum, error)
	Last() (*model.SysDictDatum, error)
	Find() ([]*model.SysDictDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysDictDatum, err error)
	FindInBatches(result *[]*model.SysDictDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysDictDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysDictDatumDo
	Assign(attrs ...field.AssignExpr) ISysDictDatumDo
	Joins(fields ...field.RelationField) ISysDictDatumDo
	Preload(fields ...field.RelationField) ISysDictDatumDo
	FirstOrInit() (*model.SysDictDatum, error)
	FirstOrCreate() (*model.SysDictDatum, error)
	FindByPage(offset int, limit int) (result []*model.SysDictDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysDictDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysDictDatumDo) Debug() ISysDictDatumDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictDatumDo) WithContext(ctx context.Context) ISysDictDatumDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictDatumDo) ReadDB() ISysDictDatumDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictDatumDo) WriteDB() ISysDictDatumDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictDatumDo) Session(config *gorm.Session) ISysDictDatumDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictDatumDo) Clauses(conds ...clause.Expression) ISysDictDatumDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictDatumDo) Returning(value interface{}, columns ...string) ISysDictDatumDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictDatumDo) Not(conds ...gen.Condition) ISysDictDatumDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictDatumDo) Or(conds ...gen.Condition) ISysDictDatumDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictDatumDo) Select(conds ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictDatumDo) Where(conds ...gen.Condition) ISysDictDatumDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictDatumDo) Order(conds ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictDatumDo) Distinct(cols ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictDatumDo) Omit(cols ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictDatumDo) Join(table schema.Tabler, on ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictDatumDo) Group(cols ...field.Expr) ISysDictDatumDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictDatumDo) Having(conds ...gen.Condition) ISysDictDatumDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictDatumDo) Limit(limit int) ISysDictDatumDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictDatumDo) Offset(offset int) ISysDictDatumDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictDatumDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictDatumDo) Unscoped() ISysDictDatumDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictDatumDo) Create(values ...*model.SysDictDatum) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictDatumDo) CreateInBatches(values []*model.SysDictDatum, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictDatumDo) Save(values ...*model.SysDictDatum) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictDatumDo) First() (*model.SysDictDatum, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictDatum), nil
	}
}

func (s sysDictDatumDo) Take() (*model.SysDictDatum, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictDatum), nil
	}
}

func (s sysDictDatumDo) Last() (*model.SysDictDatum, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictDatum), nil
	}
}

func (s sysDictDatumDo) Find() ([]*model.SysDictDatum, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysDictDatum), err
}

func (s sysDictDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysDictDatum, err error) {
	buf := make([]*model.SysDictDatum, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictDatumDo) FindInBatches(result *[]*model.SysDictDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictDatumDo) Attrs(attrs ...field.AssignExpr) ISysDictDatumDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictDatumDo) Assign(attrs ...field.AssignExpr) ISysDictDatumDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictDatumDo) Joins(fields ...field.RelationField) ISysDictDatumDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictDatumDo) Preload(fields ...field.RelationField) ISysDictDatumDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictDatumDo) FirstOrInit() (*model.SysDictDatum, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictDatum), nil
	}
}

func (s sysDictDatumDo) FirstOrCreate() (*model.SysDictDatum, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictDatum), nil
	}
}

func (s sysDictDatumDo) FindByPage(offset int, limit int) (result []*model.SysDictDatum, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysDictDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictDatumDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictDatumDo) Delete(models ...*model.SysDictDatum) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictDatumDo) withDO(do gen.Dao) *sysDictDatumDo {
	s.DO = *do.(*gen.DO)
	return s
}
