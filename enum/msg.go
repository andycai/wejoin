package enum

var codeText = map[int]string{
	Success:                          "成功！",
	SucGroupApply:                    "申请加入群组成功！",
	SucGroupApprove:                  "成功通过申请！",
	SucGroupRefuse:                   "成功拒绝申请！",
	SucGroupUpdateName:               "更新群组名称成功！",
	SucGroupUpdateNotice:             "更新群组公告成功！",
	SucGroupUpdateAddr:               "更新群组地址成功！",
	SucGroupQuit:                     "退出群组成功！",
	SucGroupFire:                     "踢出群组成功！",
	SucGroupTransfer:                 "转让群主成功！",
	SucGroupPromote:                  "提升群管理员成功！",
	ErrParam:                         "参数错误！",
	ErrData:                          "数据错误！",
	ErrOp:                            "操作失败！",
	ErrTwoPasswordNotMatch:           "输入的两次密码不一致！",
	ErrUserAuth:                      "登录验证失败，请重新登录！",
	ErrUserRegister:                  "注册失败！",
	ErrUserData:                      "获取用户数据失败！",
	ErrUserUpdateData:                "更新用户数据失败！",
	ErrUserNotFound:                  "用户不存在！",
	ErrUserEmailFormat:               "邮件地址格式错误！",
	ErrGroupManagerLimit:             "副群主数量超过限制，不能再委任！",
	ErrGroupGetData:                  "获取群数据失败！",
	ErrGroupApprove:                  "入群审批失败！",
	ErrGroupUpdateOp:                 "更新群信息失败！",
	ErrGroupNonManager:               "不是群管理员，没权限操作！",
	ErrGroupNonOwner:                 "不是群主，没权限操作！",
	ErrGroupPromote:                  "提升群管理员失败！",
	ErrGroupRemove:                   "删除群失败！",
	ErrGroupTransfer:                 "转让群主失败！",
	ErrGroupNonMember:                "不是群成员，没有权限操作！",
	ErrGroupNotFound:                 "群组不存在！",
	ErrGroupApplicationListNotFound:  "获取群组申请列表数据失败！",
	ErrGroupApplyTwice:               "已经申请过次群组！",
	ErrGroupApply:                    "申请加入群组失败！",
	ErrGroupRefuse:                   "拒绝加入群组失败！",
	ErrGroupUpdateName:               "更新群组名称失败！",
	ErrGroupUpdateNotice:             "更新群组公告失败！",
	ErrGroupUpdateAddr:               "更新群组地址失败！",
	ErrGroupQuit:                     "退出群组失败！",
	ErrGroupFire:                     "踢出群组失败！",
	ErrGroupCancel:                   "取消申请群组失败！",
	ErrActivityGetData:               "获取活动数据失败！",
	ErrActivityCannotApplyNotInGroup: "你不是群组成员不能报名或取消报名群组活动！",
	ErrActivityUpdate:                "更新活动信息失败！",
	ErrActivityNonPlanner:            "你不是活动发起人，没有权限操作！",
	ErrActivityCreate:                "创建新活动失败！",
	ErrActivityFee:                   "选择活动前结算的活动，必须要填写费用！",
	ErrActivityOverQuota:             "报名候补的数量超出限制，请稍后再报名！",
	ErrActivityNotEnough:             "取消报名的数量不正确！",
	ErrActivityRemove:                "移除报名失败！",
	ErrActivityNotDoing:              "活动已经结束，不能再操作！",
	ErrActivityCannotCancel:          "取消报名的时间已过，不能取消报名！",
	ErrActivityHasBegun:              "活动已经开始，不能报名！",
	ErrActivityNotFound:              "活动不存在！",
	ErrActivityUserCreate:            "参与活动失败！",
}

func CodeText(code int) string {
	return codeText[code]
}
