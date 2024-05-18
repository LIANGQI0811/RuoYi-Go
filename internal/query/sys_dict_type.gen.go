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
)

func newSysDictType(db *gorm.DB, opts ...gen.DOOption) sysDictType {
	_sysDictType := sysDictType{}

	_sysDictType.sysDictTypeDo.UseDB(db, opts...)
	_sysDictType.sysDictTypeDo.UseModel(&model.SysDictType{})

	tableName := _sysDictType.sysDictTypeDo.TableName()
	_sysDictType.ALL = field.NewAsterisk(tableName)
	_sysDictType.DictID = field.NewInt64(tableName, "dict_id")
	_sysDictType.DictName = field.NewString(tableName, "dict_name")
	_sysDictType.DictType = field.NewString(tableName, "dict_type")
	_sysDictType.Status = field.NewString(tableName, "status")
	_sysDictType.CreateBy = field.NewString(tableName, "create_by")
	_sysDictType.CreateTime = field.NewTime(tableName, "create_time")
	_sysDictType.UpdateBy = field.NewString(tableName, "update_by")
	_sysDictType.UpdateTime = field.NewTime(tableName, "update_time")
	_sysDictType.Remark = field.NewString(tableName, "remark")

	_sysDictType.fillFieldMap()

	return _sysDictType
}

type sysDictType struct {
	sysDictTypeDo sysDictTypeDo

	ALL        field.Asterisk
	DictID     field.Int64  // 字典主键
	DictName   field.String // 字典名称
	DictType   field.String // 字典类型
	Status     field.String // 状态（0正常 1停用）
	CreateBy   field.String // 创建者
	CreateTime field.Time   // 创建时间
	UpdateBy   field.String // 更新者
	UpdateTime field.Time   // 更新时间
	Remark     field.String // 备注

	fieldMap map[string]field.Expr
}

func (s sysDictType) Table(newTableName string) *sysDictType {
	s.sysDictTypeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictType) As(alias string) *sysDictType {
	s.sysDictTypeDo.DO = *(s.sysDictTypeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictType) updateTableName(table string) *sysDictType {
	s.ALL = field.NewAsterisk(table)
	s.DictID = field.NewInt64(table, "dict_id")
	s.DictName = field.NewString(table, "dict_name")
	s.DictType = field.NewString(table, "dict_type")
	s.Status = field.NewString(table, "status")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysDictType) WithContext(ctx context.Context) *sysDictTypeDo {
	return s.sysDictTypeDo.WithContext(ctx)
}

func (s sysDictType) TableName() string { return s.sysDictTypeDo.TableName() }

func (s sysDictType) Alias() string { return s.sysDictTypeDo.Alias() }

func (s sysDictType) Columns(cols ...field.Expr) gen.Columns { return s.sysDictTypeDo.Columns(cols...) }

func (s *sysDictType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictType) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["dict_id"] = s.DictID
	s.fieldMap["dict_name"] = s.DictName
	s.fieldMap["dict_type"] = s.DictType
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
}

func (s sysDictType) clone(db *gorm.DB) sysDictType {
	s.sysDictTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictType) replaceDB(db *gorm.DB) sysDictType {
	s.sysDictTypeDo.ReplaceDB(db)
	return s
}

type sysDictTypeDo struct{ gen.DO }

func (s sysDictTypeDo) Debug() *sysDictTypeDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictTypeDo) WithContext(ctx context.Context) *sysDictTypeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictTypeDo) ReadDB() *sysDictTypeDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictTypeDo) WriteDB() *sysDictTypeDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictTypeDo) Session(config *gorm.Session) *sysDictTypeDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictTypeDo) Clauses(conds ...clause.Expression) *sysDictTypeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictTypeDo) Returning(value interface{}, columns ...string) *sysDictTypeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictTypeDo) Not(conds ...gen.Condition) *sysDictTypeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictTypeDo) Or(conds ...gen.Condition) *sysDictTypeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictTypeDo) Select(conds ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictTypeDo) Where(conds ...gen.Condition) *sysDictTypeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictTypeDo) Order(conds ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictTypeDo) Distinct(cols ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictTypeDo) Omit(cols ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictTypeDo) Join(table schema.Tabler, on ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictTypeDo) Group(cols ...field.Expr) *sysDictTypeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictTypeDo) Having(conds ...gen.Condition) *sysDictTypeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictTypeDo) Limit(limit int) *sysDictTypeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictTypeDo) Offset(offset int) *sysDictTypeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysDictTypeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictTypeDo) Unscoped() *sysDictTypeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictTypeDo) Create(values ...*model.SysDictType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictTypeDo) CreateInBatches(values []*model.SysDictType, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictTypeDo) Save(values ...*model.SysDictType) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictTypeDo) First() (*model.SysDictType, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) Take() (*model.SysDictType, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) Last() (*model.SysDictType, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) Find() ([]*model.SysDictType, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysDictType), err
}

func (s sysDictTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysDictType, err error) {
	buf := make([]*model.SysDictType, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictTypeDo) FindInBatches(result *[]*model.SysDictType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictTypeDo) Attrs(attrs ...field.AssignExpr) *sysDictTypeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictTypeDo) Assign(attrs ...field.AssignExpr) *sysDictTypeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictTypeDo) Joins(fields ...field.RelationField) *sysDictTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictTypeDo) Preload(fields ...field.RelationField) *sysDictTypeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictTypeDo) FirstOrInit() (*model.SysDictType, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) FirstOrCreate() (*model.SysDictType, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDictType), nil
	}
}

func (s sysDictTypeDo) FindByPage(offset int, limit int) (result []*model.SysDictType, count int64, err error) {
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

func (s sysDictTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictTypeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictTypeDo) Delete(models ...*model.SysDictType) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictTypeDo) withDO(do gen.Dao) *sysDictTypeDo {
	s.DO = *do.(*gen.DO)
	return s
}
