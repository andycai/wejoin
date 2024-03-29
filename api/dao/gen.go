package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                = new(Query)
	Activity         *activity
	ActivityUser     *activityUser
	Group            *group
	GroupApplication *groupApplication
	GroupMember      *groupMember
	User             *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Activity = &Q.Activity
	ActivityUser = &Q.ActivityUser
	Group = &Q.Group
	GroupApplication = &Q.GroupApplication
	GroupMember = &Q.GroupMember
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:               db,
		Activity:         newActivity(db, opts...),
		ActivityUser:     newActivityUser(db, opts...),
		Group:            newGroup(db, opts...),
		GroupApplication: newGroupApplication(db, opts...),
		GroupMember:      newGroupMember(db, opts...),
		User:             newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Activity         activity
	ActivityUser     activityUser
	Group            group
	GroupApplication groupApplication
	GroupMember      groupMember
	User             user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		Activity:         q.Activity.clone(db),
		ActivityUser:     q.ActivityUser.clone(db),
		Group:            q.Group.clone(db),
		GroupApplication: q.GroupApplication.clone(db),
		GroupMember:      q.GroupMember.clone(db),
		User:             q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		Activity:         q.Activity.replaceDB(db),
		ActivityUser:     q.ActivityUser.replaceDB(db),
		Group:            q.Group.replaceDB(db),
		GroupApplication: q.GroupApplication.replaceDB(db),
		GroupMember:      q.GroupMember.replaceDB(db),
		User:             q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Activity         IActivityDo
	ActivityUser     IActivityUserDo
	Group            IGroupDo
	GroupApplication IGroupApplicationDo
	GroupMember      IGroupMemberDo
	User             IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Activity:         q.Activity.WithContext(ctx),
		ActivityUser:     q.ActivityUser.WithContext(ctx),
		Group:            q.Group.WithContext(ctx),
		GroupApplication: q.GroupApplication.WithContext(ctx),
		GroupMember:      q.GroupMember.WithContext(ctx),
		User:             q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
