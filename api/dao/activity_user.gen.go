package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/andycai/axe-fiber/model"
)

func newActivityUser(db *gorm.DB, opts ...gen.DOOption) activityUser {
	_activityUser := activityUser{}

	_activityUser.activityUserDo.UseDB(db, opts...)
	_activityUser.activityUserDo.UseModel(&model.ActivityUser{})

	tableName := _activityUser.activityUserDo.TableName()
	_activityUser.ALL = field.NewAsterisk(tableName)
	_activityUser.ID = field.NewInt32(tableName, "id")
	_activityUser.ActivityID = field.NewInt32(tableName, "activity_id")
	_activityUser.UserID = field.NewInt32(tableName, "user_id")
	_activityUser.Alias_ = field.NewString(tableName, "alias")
	_activityUser.IsFriend = field.NewInt32(tableName, "is_friend")
	_activityUser.CreateAt = field.NewTime(tableName, "create_at")
	_activityUser.UpdateAt = field.NewTime(tableName, "update_at")
	_activityUser.DeleteAt = field.NewTime(tableName, "delete_at")

	_activityUser.fillFieldMap()

	return _activityUser
}

type activityUser struct {
	activityUserDo

	ALL        field.Asterisk
	ID         field.Int32  // 活动报名id
	ActivityID field.Int32  // 活动id
	UserID     field.Int32  // 报名用户id
	Alias_     field.String // 报名昵称
	IsFriend   field.Int32  // 是否朋友
	CreateAt   field.Time   // 创建时间
	UpdateAt   field.Time   // 修改时间
	DeleteAt   field.Time   // 删除时间

	fieldMap map[string]field.Expr
}

func (a activityUser) Table(newTableName string) *activityUser {
	a.activityUserDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activityUser) As(alias string) *activityUser {
	a.activityUserDo.DO = *(a.activityUserDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activityUser) updateTableName(table string) *activityUser {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.ActivityID = field.NewInt32(table, "activity_id")
	a.UserID = field.NewInt32(table, "user_id")
	a.Alias_ = field.NewString(table, "alias")
	a.IsFriend = field.NewInt32(table, "is_friend")
	a.CreateAt = field.NewTime(table, "create_at")
	a.UpdateAt = field.NewTime(table, "update_at")
	a.DeleteAt = field.NewTime(table, "delete_at")

	a.fillFieldMap()

	return a
}

func (a *activityUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activityUser) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["activity_id"] = a.ActivityID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["alias"] = a.Alias_
	a.fieldMap["is_friend"] = a.IsFriend
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["delete_at"] = a.DeleteAt
}

func (a activityUser) clone(db *gorm.DB) activityUser {
	a.activityUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activityUser) replaceDB(db *gorm.DB) activityUser {
	a.activityUserDo.ReplaceDB(db)
	return a
}

type activityUserDo struct{ gen.DO }

type IActivityUserDo interface {
	gen.SubQuery
	Debug() IActivityUserDo
	WithContext(ctx context.Context) IActivityUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActivityUserDo
	WriteDB() IActivityUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActivityUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActivityUserDo
	Not(conds ...gen.Condition) IActivityUserDo
	Or(conds ...gen.Condition) IActivityUserDo
	Select(conds ...field.Expr) IActivityUserDo
	Where(conds ...gen.Condition) IActivityUserDo
	Order(conds ...field.Expr) IActivityUserDo
	Distinct(cols ...field.Expr) IActivityUserDo
	Omit(cols ...field.Expr) IActivityUserDo
	Join(table schema.Tabler, on ...field.Expr) IActivityUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActivityUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActivityUserDo
	Group(cols ...field.Expr) IActivityUserDo
	Having(conds ...gen.Condition) IActivityUserDo
	Limit(limit int) IActivityUserDo
	Offset(offset int) IActivityUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityUserDo
	Unscoped() IActivityUserDo
	Create(values ...*model.ActivityUser) error
	CreateInBatches(values []*model.ActivityUser, batchSize int) error
	Save(values ...*model.ActivityUser) error
	First() (*model.ActivityUser, error)
	Take() (*model.ActivityUser, error)
	Last() (*model.ActivityUser, error)
	Find() ([]*model.ActivityUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityUser, err error)
	FindInBatches(result *[]*model.ActivityUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ActivityUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActivityUserDo
	Assign(attrs ...field.AssignExpr) IActivityUserDo
	Joins(fields ...field.RelationField) IActivityUserDo
	Preload(fields ...field.RelationField) IActivityUserDo
	FirstOrInit() (*model.ActivityUser, error)
	FirstOrCreate() (*model.ActivityUser, error)
	FindByPage(offset int, limit int) (result []*model.ActivityUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActivityUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a activityUserDo) Debug() IActivityUserDo {
	return a.withDO(a.DO.Debug())
}

func (a activityUserDo) WithContext(ctx context.Context) IActivityUserDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityUserDo) ReadDB() IActivityUserDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityUserDo) WriteDB() IActivityUserDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityUserDo) Session(config *gorm.Session) IActivityUserDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityUserDo) Clauses(conds ...clause.Expression) IActivityUserDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityUserDo) Returning(value interface{}, columns ...string) IActivityUserDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityUserDo) Not(conds ...gen.Condition) IActivityUserDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityUserDo) Or(conds ...gen.Condition) IActivityUserDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityUserDo) Select(conds ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityUserDo) Where(conds ...gen.Condition) IActivityUserDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityUserDo) Order(conds ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityUserDo) Distinct(cols ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityUserDo) Omit(cols ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityUserDo) Join(table schema.Tabler, on ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityUserDo) Group(cols ...field.Expr) IActivityUserDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityUserDo) Having(conds ...gen.Condition) IActivityUserDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityUserDo) Limit(limit int) IActivityUserDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityUserDo) Offset(offset int) IActivityUserDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityUserDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityUserDo) Unscoped() IActivityUserDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityUserDo) Create(values ...*model.ActivityUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityUserDo) CreateInBatches(values []*model.ActivityUser, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityUserDo) Save(values ...*model.ActivityUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityUserDo) First() (*model.ActivityUser, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityUser), nil
	}
}

func (a activityUserDo) Take() (*model.ActivityUser, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityUser), nil
	}
}

func (a activityUserDo) Last() (*model.ActivityUser, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityUser), nil
	}
}

func (a activityUserDo) Find() ([]*model.ActivityUser, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActivityUser), err
}

func (a activityUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityUser, err error) {
	buf := make([]*model.ActivityUser, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityUserDo) FindInBatches(result *[]*model.ActivityUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityUserDo) Attrs(attrs ...field.AssignExpr) IActivityUserDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityUserDo) Assign(attrs ...field.AssignExpr) IActivityUserDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityUserDo) Joins(fields ...field.RelationField) IActivityUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityUserDo) Preload(fields ...field.RelationField) IActivityUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityUserDo) FirstOrInit() (*model.ActivityUser, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityUser), nil
	}
}

func (a activityUserDo) FirstOrCreate() (*model.ActivityUser, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityUser), nil
	}
}

func (a activityUserDo) FindByPage(offset int, limit int) (result []*model.ActivityUser, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a activityUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityUserDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityUserDo) Delete(models ...*model.ActivityUser) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityUserDo) withDO(do gen.Dao) *activityUserDo {
	a.DO = *do.(*gen.DO)
	return a
}
