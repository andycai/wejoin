package enum

var codeText = map[int]string{
	Success:                          "成功",
	ErrParam:                         "参数错误",
	ErrData:                          "数据错误",
	ErrOp:                            "操作失败",
	ErrAuth:                          "登录验证失败，请重新登录！",
	ErrRegister:                      "注册失败！",
	ErrUserData:                      "获取用户数据失败！",
	ErrUserUpdateData:                "更新用户数据失败！",
	ErrUserNotFound:                  "用户不存在！",
	ErrGroupManagerLimit:             "副群主数量超过限制，不能再委任！",
	ErrGroupGetData:                  "获取群数据失败！",
	ErrGroupApprove:                  "入群审批失败！",
	ErrGroupUpdateOp:                 "更新群信息失败！",
	ErrGroupNonManager:               "不是群管理员，没权限操作！",
	ErrGroupNonOwner:                 "不是群主，没权限操作！",
	ErrGroupPromote:                  "委任副群主失败！",
	ErrGroupRemove:                   "删除群成员失败！",
	ErrGroupTransfer:                 "转让群主失败！",
	ErrGroupNonMember:                "不是群成员，没有权限操作！",
	ErrGroupNotFound:                 "群组不存在！",
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
}

func CodeText(code int) string {
	return codeText[code]
}
