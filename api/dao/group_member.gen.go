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

func newGroupMember(db *gorm.DB, opts ...gen.DOOption) groupMember {
	_groupMember := groupMember{}

	_groupMember.groupMemberDo.UseDB(db, opts...)
	_groupMember.groupMemberDo.UseModel(&model.GroupMember{})

	tableName := _groupMember.groupMemberDo.TableName()
	_groupMember.ALL = field.NewAsterisk(tableName)
	_groupMember.ID = field.NewInt32(tableName, "id")
	_groupMember.GroupID = field.NewInt32(tableName, "group_id")
	_groupMember.UserID = field.NewInt32(tableName, "user_id")
	_groupMember.Scores = field.NewInt32(tableName, "scores")
	_groupMember.Position = field.NewInt32(tableName, "position")
	_groupMember.Alias_ = field.NewString(tableName, "alias")
	_groupMember.EnterAt = field.NewTime(tableName, "enter_at")
	_groupMember.CreateAt = field.NewTime(tableName, "create_at")
	_groupMember.UpdateAt = field.NewTime(tableName, "update_at")
	_groupMember.DeleteAt = field.NewTime(tableName, "delete_at")

	_groupMember.fillFieldMap()

	return _groupMember
}

type groupMember struct {
	groupMemberDo

	ALL      field.Asterisk
	ID       field.Int32  // 成员id
	GroupID  field.Int32  // 群组id
	UserID   field.Int32  // 用户id
	Scores   field.Int32  // 积分
	Position field.Int32  // 群组职位
	Alias_   field.String // 群组中别名
	EnterAt  field.Time   // 进入群组时间
	CreateAt field.Time   // 创建时间
	UpdateAt field.Time   // 更新时间
	DeleteAt field.Time   // 删除时间

	fieldMap map[string]field.Expr
}

func (g groupMember) Table(newTableName string) *groupMember {
	g.groupMemberDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g groupMember) As(alias string) *groupMember {
	g.groupMemberDo.DO = *(g.groupMemberDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *groupMember) updateTableName(table string) *groupMember {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.GroupID = field.NewInt32(table, "group_id")
	g.UserID = field.NewInt32(table, "user_id")
	g.Scores = field.NewInt32(table, "scores")
	g.Position = field.NewInt32(table, "position")
	g.Alias_ = field.NewString(table, "alias")
	g.EnterAt = field.NewTime(table, "enter_at")
	g.CreateAt = field.NewTime(table, "create_at")
	g.UpdateAt = field.NewTime(table, "update_at")
	g.DeleteAt = field.NewTime(table, "delete_at")

	g.fillFieldMap()

	return g
}

func (g *groupMember) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *groupMember) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 10)
	g.fieldMap["id"] = g.ID
	g.fieldMap["group_id"] = g.GroupID
	g.fieldMap["user_id"] = g.UserID
	g.fieldMap["scores"] = g.Scores
	g.fieldMap["position"] = g.Position
	g.fieldMap["alias"] = g.Alias_
	g.fieldMap["enter_at"] = g.EnterAt
	g.fieldMap["create_at"] = g.CreateAt
	g.fieldMap["update_at"] = g.UpdateAt
	g.fieldMap["delete_at"] = g.DeleteAt
}

func (g groupMember) clone(db *gorm.DB) groupMember {
	g.groupMemberDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g groupMember) replaceDB(db *gorm.DB) groupMember {
	g.groupMemberDo.ReplaceDB(db)
	return g
}

type groupMemberDo struct{ gen.DO }

type IGroupMemberDo interface {
	gen.SubQuery
	Debug() IGroupMemberDo
	WithContext(ctx context.Context) IGroupMemberDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGroupMemberDo
	WriteDB() IGroupMemberDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGroupMemberDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGroupMemberDo
	Not(conds ...gen.Condition) IGroupMemberDo
	Or(conds ...gen.Condition) IGroupMemberDo
	Select(conds ...field.Expr) IGroupMemberDo
	Where(conds ...gen.Condition) IGroupMemberDo
	Order(conds ...field.Expr) IGroupMemberDo
	Distinct(cols ...field.Expr) IGroupMemberDo
	Omit(cols ...field.Expr) IGroupMemberDo
	Join(table schema.Tabler, on ...field.Expr) IGroupMemberDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGroupMemberDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGroupMemberDo
	Group(cols ...field.Expr) IGroupMemberDo
	Having(conds ...gen.Condition) IGroupMemberDo
	Limit(limit int) IGroupMemberDo
	Offset(offset int) IGroupMemberDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupMemberDo
	Unscoped() IGroupMemberDo
	Create(values ...*model.GroupMember) error
	CreateInBatches(values []*model.GroupMember, batchSize int) error
	Save(values ...*model.GroupMember) error
	First() (*model.GroupMember, error)
	Take() (*model.GroupMember, error)
	Last() (*model.GroupMember, error)
	Find() ([]*model.GroupMember, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupMember, err error)
	FindInBatches(result *[]*model.GroupMember, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GroupMember) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGroupMemberDo
	Assign(attrs ...field.AssignExpr) IGroupMemberDo
	Joins(fields ...field.RelationField) IGroupMemberDo
	Preload(fields ...field.RelationField) IGroupMemberDo
	FirstOrInit() (*model.GroupMember, error)
	FirstOrCreate() (*model.GroupMember, error)
	FindByPage(offset int, limit int) (result []*model.GroupMember, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGroupMemberDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g groupMemberDo) Debug() IGroupMemberDo {
	return g.withDO(g.DO.Debug())
}

func (g groupMemberDo) WithContext(ctx context.Context) IGroupMemberDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupMemberDo) ReadDB() IGroupMemberDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupMemberDo) WriteDB() IGroupMemberDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupMemberDo) Session(config *gorm.Session) IGroupMemberDo {
	return g.withDO(g.DO.Session(config))
}

func (g groupMemberDo) Clauses(conds ...clause.Expression) IGroupMemberDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupMemberDo) Returning(value interface{}, columns ...string) IGroupMemberDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupMemberDo) Not(conds ...gen.Condition) IGroupMemberDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupMemberDo) Or(conds ...gen.Condition) IGroupMemberDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupMemberDo) Select(conds ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupMemberDo) Where(conds ...gen.Condition) IGroupMemberDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupMemberDo) Order(conds ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupMemberDo) Distinct(cols ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupMemberDo) Omit(cols ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupMemberDo) Join(table schema.Tabler, on ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupMemberDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupMemberDo) RightJoin(table schema.Tabler, on ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupMemberDo) Group(cols ...field.Expr) IGroupMemberDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupMemberDo) Having(conds ...gen.Condition) IGroupMemberDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupMemberDo) Limit(limit int) IGroupMemberDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupMemberDo) Offset(offset int) IGroupMemberDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupMemberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupMemberDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupMemberDo) Unscoped() IGroupMemberDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupMemberDo) Create(values ...*model.GroupMember) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupMemberDo) CreateInBatches(values []*model.GroupMember, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupMemberDo) Save(values ...*model.GroupMember) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupMemberDo) First() (*model.GroupMember, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupMember), nil
	}
}

func (g groupMemberDo) Take() (*model.GroupMember, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupMember), nil
	}
}

func (g groupMemberDo) Last() (*model.GroupMember, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupMember), nil
	}
}

func (g groupMemberDo) Find() ([]*model.GroupMember, error) {
	result, err := g.DO.Find()
	return result.([]*model.GroupMember), err
}

func (g groupMemberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupMember, err error) {
	buf := make([]*model.GroupMember, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupMemberDo) FindInBatches(result *[]*model.GroupMember, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupMemberDo) Attrs(attrs ...field.AssignExpr) IGroupMemberDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupMemberDo) Assign(attrs ...field.AssignExpr) IGroupMemberDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupMemberDo) Joins(fields ...field.RelationField) IGroupMemberDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupMemberDo) Preload(fields ...field.RelationField) IGroupMemberDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupMemberDo) FirstOrInit() (*model.GroupMember, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupMember), nil
	}
}

func (g groupMemberDo) FirstOrCreate() (*model.GroupMember, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupMember), nil
	}
}

func (g groupMemberDo) FindByPage(offset int, limit int) (result []*model.GroupMember, count int64, err error) {
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

func (g groupMemberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupMemberDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupMemberDo) Delete(models ...*model.GroupMember) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupMemberDo) withDO(do gen.Dao) *groupMemberDo {
	g.DO = *do.(*gen.DO)
	return g
}
