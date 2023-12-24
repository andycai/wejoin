package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/andycai/wejoin/model"
)

func newGroupApplication(db *gorm.DB, opts ...gen.DOOption) groupApplication {
	_groupApplication := groupApplication{}

	_groupApplication.groupApplicationDo.UseDB(db, opts...)
	_groupApplication.groupApplicationDo.UseModel(&model.GroupApplication{})

	tableName := _groupApplication.groupApplicationDo.TableName()
	_groupApplication.ALL = field.NewAsterisk(tableName)
	_groupApplication.ID = field.NewInt32(tableName, "id")
	_groupApplication.UserID = field.NewInt32(tableName, "user_id")
	_groupApplication.GroupID = field.NewInt32(tableName, "group_id")
	_groupApplication.Deleted = field.NewInt32(tableName, "deleted")
	_groupApplication.DeleteAt = field.NewTime(tableName, "delete_at")
	_groupApplication.CreateAt = field.NewTime(tableName, "create_at")

	_groupApplication.fillFieldMap()

	return _groupApplication
}

type groupApplication struct {
	groupApplicationDo

	ALL      field.Asterisk
	ID       field.Int32 // 申请id
	UserID   field.Int32 // 用户id
	GroupID  field.Int32 // 群组id
	Deleted  field.Int32 // 是否删除
	DeleteAt field.Time  // 删除id
	CreateAt field.Time  // 创建时间

	fieldMap map[string]field.Expr
}

func (g groupApplication) Table(newTableName string) *groupApplication {
	g.groupApplicationDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g groupApplication) As(alias string) *groupApplication {
	g.groupApplicationDo.DO = *(g.groupApplicationDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *groupApplication) updateTableName(table string) *groupApplication {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.UserID = field.NewInt32(table, "user_id")
	g.GroupID = field.NewInt32(table, "group_id")
	g.Deleted = field.NewInt32(table, "deleted")
	g.DeleteAt = field.NewTime(table, "delete_at")
	g.CreateAt = field.NewTime(table, "create_at")

	g.fillFieldMap()

	return g
}

func (g *groupApplication) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *groupApplication) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 6)
	g.fieldMap["id"] = g.ID
	g.fieldMap["user_id"] = g.UserID
	g.fieldMap["group_id"] = g.GroupID
	g.fieldMap["deleted"] = g.Deleted
	g.fieldMap["delete_at"] = g.DeleteAt
	g.fieldMap["create_at"] = g.CreateAt
}

func (g groupApplication) clone(db *gorm.DB) groupApplication {
	g.groupApplicationDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g groupApplication) replaceDB(db *gorm.DB) groupApplication {
	g.groupApplicationDo.ReplaceDB(db)
	return g
}

type groupApplicationDo struct{ gen.DO }

type IGroupApplicationDo interface {
	gen.SubQuery
	Debug() IGroupApplicationDo
	WithContext(ctx context.Context) IGroupApplicationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGroupApplicationDo
	WriteDB() IGroupApplicationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGroupApplicationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGroupApplicationDo
	Not(conds ...gen.Condition) IGroupApplicationDo
	Or(conds ...gen.Condition) IGroupApplicationDo
	Select(conds ...field.Expr) IGroupApplicationDo
	Where(conds ...gen.Condition) IGroupApplicationDo
	Order(conds ...field.Expr) IGroupApplicationDo
	Distinct(cols ...field.Expr) IGroupApplicationDo
	Omit(cols ...field.Expr) IGroupApplicationDo
	Join(table schema.Tabler, on ...field.Expr) IGroupApplicationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGroupApplicationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGroupApplicationDo
	Group(cols ...field.Expr) IGroupApplicationDo
	Having(conds ...gen.Condition) IGroupApplicationDo
	Limit(limit int) IGroupApplicationDo
	Offset(offset int) IGroupApplicationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupApplicationDo
	Unscoped() IGroupApplicationDo
	Create(values ...*model.GroupApplication) error
	CreateInBatches(values []*model.GroupApplication, batchSize int) error
	Save(values ...*model.GroupApplication) error
	First() (*model.GroupApplication, error)
	Take() (*model.GroupApplication, error)
	Last() (*model.GroupApplication, error)
	Find() ([]*model.GroupApplication, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupApplication, err error)
	FindInBatches(result *[]*model.GroupApplication, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GroupApplication) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGroupApplicationDo
	Assign(attrs ...field.AssignExpr) IGroupApplicationDo
	Joins(fields ...field.RelationField) IGroupApplicationDo
	Preload(fields ...field.RelationField) IGroupApplicationDo
	FirstOrInit() (*model.GroupApplication, error)
	FirstOrCreate() (*model.GroupApplication, error)
	FindByPage(offset int, limit int) (result []*model.GroupApplication, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGroupApplicationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g groupApplicationDo) Debug() IGroupApplicationDo {
	return g.withDO(g.DO.Debug())
}

func (g groupApplicationDo) WithContext(ctx context.Context) IGroupApplicationDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupApplicationDo) ReadDB() IGroupApplicationDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupApplicationDo) WriteDB() IGroupApplicationDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupApplicationDo) Session(config *gorm.Session) IGroupApplicationDo {
	return g.withDO(g.DO.Session(config))
}

func (g groupApplicationDo) Clauses(conds ...clause.Expression) IGroupApplicationDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupApplicationDo) Returning(value interface{}, columns ...string) IGroupApplicationDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupApplicationDo) Not(conds ...gen.Condition) IGroupApplicationDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupApplicationDo) Or(conds ...gen.Condition) IGroupApplicationDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupApplicationDo) Select(conds ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupApplicationDo) Where(conds ...gen.Condition) IGroupApplicationDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupApplicationDo) Order(conds ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupApplicationDo) Distinct(cols ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupApplicationDo) Omit(cols ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupApplicationDo) Join(table schema.Tabler, on ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupApplicationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupApplicationDo) RightJoin(table schema.Tabler, on ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupApplicationDo) Group(cols ...field.Expr) IGroupApplicationDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupApplicationDo) Having(conds ...gen.Condition) IGroupApplicationDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupApplicationDo) Limit(limit int) IGroupApplicationDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupApplicationDo) Offset(offset int) IGroupApplicationDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupApplicationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupApplicationDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupApplicationDo) Unscoped() IGroupApplicationDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupApplicationDo) Create(values ...*model.GroupApplication) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupApplicationDo) CreateInBatches(values []*model.GroupApplication, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupApplicationDo) Save(values ...*model.GroupApplication) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupApplicationDo) First() (*model.GroupApplication, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApplication), nil
	}
}

func (g groupApplicationDo) Take() (*model.GroupApplication, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApplication), nil
	}
}

func (g groupApplicationDo) Last() (*model.GroupApplication, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApplication), nil
	}
}

func (g groupApplicationDo) Find() ([]*model.GroupApplication, error) {
	result, err := g.DO.Find()
	return result.([]*model.GroupApplication), err
}

func (g groupApplicationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupApplication, err error) {
	buf := make([]*model.GroupApplication, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupApplicationDo) FindInBatches(result *[]*model.GroupApplication, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupApplicationDo) Attrs(attrs ...field.AssignExpr) IGroupApplicationDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupApplicationDo) Assign(attrs ...field.AssignExpr) IGroupApplicationDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupApplicationDo) Joins(fields ...field.RelationField) IGroupApplicationDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupApplicationDo) Preload(fields ...field.RelationField) IGroupApplicationDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupApplicationDo) FirstOrInit() (*model.GroupApplication, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApplication), nil
	}
}

func (g groupApplicationDo) FirstOrCreate() (*model.GroupApplication, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApplication), nil
	}
}

func (g groupApplicationDo) FindByPage(offset int, limit int) (result []*model.GroupApplication, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g groupApplicationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupApplicationDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupApplicationDo) Delete(models ...*model.GroupApplication) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupApplicationDo) withDO(do gen.Dao) *groupApplicationDo {
	g.DO = *do.(*gen.DO)
	return g
}
