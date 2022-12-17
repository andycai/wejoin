// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/andycai/axe-fiber/v2/model"
)

func newActivity(db *gorm.DB, opts ...gen.DOOption) activity {
	_activity := activity{}

	_activity.activityDo.UseDB(db, opts...)
	_activity.activityDo.UseModel(&model.Activity{})

	tableName := _activity.activityDo.TableName()
	_activity.ALL = field.NewAsterisk(tableName)
	_activity.ID = field.NewInt32(tableName, "id")
	_activity.Planner = field.NewInt32(tableName, "planner")
	_activity.GroupID = field.NewInt32(tableName, "group_id")
	_activity.Kind = field.NewInt32(tableName, "kind")
	_activity.Type = field.NewInt32(tableName, "type")
	_activity.Title = field.NewString(tableName, "title")
	_activity.Remark = field.NewString(tableName, "remark")
	_activity.Quota = field.NewInt32(tableName, "quota")
	_activity.Waiting = field.NewInt32(tableName, "waiting")
	_activity.Status = field.NewInt32(tableName, "status")
	_activity.FeeType = field.NewInt32(tableName, "fee_type")
	_activity.FeeMale = field.NewInt32(tableName, "fee_male")
	_activity.FeeFemale = field.NewInt32(tableName, "fee_female")
	_activity.Addr = field.NewString(tableName, "addr")
	_activity.Ahead = field.NewInt32(tableName, "ahead")
	_activity.BeginAt = field.NewTime(tableName, "begin_at")
	_activity.EndAt = field.NewTime(tableName, "end_at")
	_activity.CreateAt = field.NewTime(tableName, "create_at")
	_activity.UpdateAt = field.NewTime(tableName, "update_at")
	_activity.DeleteAt = field.NewTime(tableName, "delete_at")

	_activity.fillFieldMap()

	return _activity
}

type activity struct {
	activityDo activityDo

	ALL       field.Asterisk
	ID        field.Int32  // 活动ID
	Planner   field.Int32  // 组织者ID
	GroupID   field.Int32  // 群组ID
	Kind      field.Int32  // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type      field.Int32  // 活动类型:1全局保护,2全局公开,3群组
	Title     field.String // 活动标题
	Remark    field.String // 活动描述
	Quota     field.Int32  // 名额
	Waiting   field.Int32  // 候补数量限制
	Status    field.Int32  // 活动状态:1进行中,2正常结算完成,3手动终止
	FeeType   field.Int32  // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale   field.Int32  // 男费用,单位:分
	FeeFemale field.Int32  // 女费用,单位:分
	Addr      field.String // 活动地址
	Ahead     field.Int32  // 可提前取消时间(小时)
	BeginAt   field.Time   // 开始时间
	EndAt     field.Time   // 结束时间
	CreateAt  field.Time   // 创建时间
	UpdateAt  field.Time   // 更新时间
	DeleteAt  field.Time   // 删除时间

	fieldMap map[string]field.Expr
}

func (a activity) Table(newTableName string) *activity {
	a.activityDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activity) As(alias string) *activity {
	a.activityDo.DO = *(a.activityDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activity) updateTableName(table string) *activity {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.Planner = field.NewInt32(table, "planner")
	a.GroupID = field.NewInt32(table, "group_id")
	a.Kind = field.NewInt32(table, "kind")
	a.Type = field.NewInt32(table, "type")
	a.Title = field.NewString(table, "title")
	a.Remark = field.NewString(table, "remark")
	a.Quota = field.NewInt32(table, "quota")
	a.Waiting = field.NewInt32(table, "waiting")
	a.Status = field.NewInt32(table, "status")
	a.FeeType = field.NewInt32(table, "fee_type")
	a.FeeMale = field.NewInt32(table, "fee_male")
	a.FeeFemale = field.NewInt32(table, "fee_female")
	a.Addr = field.NewString(table, "addr")
	a.Ahead = field.NewInt32(table, "ahead")
	a.BeginAt = field.NewTime(table, "begin_at")
	a.EndAt = field.NewTime(table, "end_at")
	a.CreateAt = field.NewTime(table, "create_at")
	a.UpdateAt = field.NewTime(table, "update_at")
	a.DeleteAt = field.NewTime(table, "delete_at")

	a.fillFieldMap()

	return a
}

func (a *activity) WithContext(ctx context.Context) *activityDo { return a.activityDo.WithContext(ctx) }

func (a activity) TableName() string { return a.activityDo.TableName() }

func (a activity) Alias() string { return a.activityDo.Alias() }

func (a *activity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activity) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 20)
	a.fieldMap["id"] = a.ID
	a.fieldMap["planner"] = a.Planner
	a.fieldMap["group_id"] = a.GroupID
	a.fieldMap["kind"] = a.Kind
	a.fieldMap["type"] = a.Type
	a.fieldMap["title"] = a.Title
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["quota"] = a.Quota
	a.fieldMap["waiting"] = a.Waiting
	a.fieldMap["status"] = a.Status
	a.fieldMap["fee_type"] = a.FeeType
	a.fieldMap["fee_male"] = a.FeeMale
	a.fieldMap["fee_female"] = a.FeeFemale
	a.fieldMap["addr"] = a.Addr
	a.fieldMap["ahead"] = a.Ahead
	a.fieldMap["begin_at"] = a.BeginAt
	a.fieldMap["end_at"] = a.EndAt
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["delete_at"] = a.DeleteAt
}

func (a activity) clone(db *gorm.DB) activity {
	a.activityDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activity) replaceDB(db *gorm.DB) activity {
	a.activityDo.ReplaceDB(db)
	return a
}

type activityDo struct{ gen.DO }

func (a activityDo) Debug() *activityDo {
	return a.withDO(a.DO.Debug())
}

func (a activityDo) WithContext(ctx context.Context) *activityDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityDo) ReadDB() *activityDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityDo) WriteDB() *activityDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityDo) Session(config *gorm.Session) *activityDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityDo) Clauses(conds ...clause.Expression) *activityDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityDo) Returning(value interface{}, columns ...string) *activityDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityDo) Not(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityDo) Or(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityDo) Select(conds ...field.Expr) *activityDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityDo) Where(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *activityDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a activityDo) Order(conds ...field.Expr) *activityDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityDo) Distinct(cols ...field.Expr) *activityDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityDo) Omit(cols ...field.Expr) *activityDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityDo) Join(table schema.Tabler, on ...field.Expr) *activityDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityDo) LeftJoin(table schema.Tabler, on ...field.Expr) *activityDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityDo) RightJoin(table schema.Tabler, on ...field.Expr) *activityDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityDo) Group(cols ...field.Expr) *activityDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityDo) Having(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityDo) Limit(limit int) *activityDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityDo) Offset(offset int) *activityDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *activityDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityDo) Unscoped() *activityDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityDo) Create(values ...*model.Activity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityDo) CreateInBatches(values []*model.Activity, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityDo) Save(values ...*model.Activity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityDo) First() (*model.Activity, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Take() (*model.Activity, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Last() (*model.Activity, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Find() ([]*model.Activity, error) {
	result, err := a.DO.Find()
	return result.([]*model.Activity), err
}

func (a activityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Activity, err error) {
	buf := make([]*model.Activity, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityDo) FindInBatches(result *[]*model.Activity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityDo) Attrs(attrs ...field.AssignExpr) *activityDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityDo) Assign(attrs ...field.AssignExpr) *activityDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityDo) Joins(fields ...field.RelationField) *activityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityDo) Preload(fields ...field.RelationField) *activityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityDo) FirstOrInit() (*model.Activity, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) FirstOrCreate() (*model.Activity, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) FindByPage(offset int, limit int) (result []*model.Activity, count int64, err error) {
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

func (a activityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityDo) Delete(models ...*model.Activity) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityDo) withDO(do gen.Dao) *activityDo {
	a.DO = *do.(*gen.DO)
	return a
}