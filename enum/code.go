package enum

const (
	Success                          = 0
	SucGroupApply                    = 200
	SucGroupApprove                  = 201
	SucGroupRefuse                   = 202
	ErrParam                         = -100
	ErrData                          = -101
	ErrOp                            = -102
	ErrAuth                          = -103
	ErrRegister                      = -104
	ErrUserData                      = -105
	ErrUserUpdateData                = -106
	ErrUserNotFound                  = -107
	ErrGroupManagerLimit             = -200
	ErrGroupGetData                  = -201
	ErrGroupApprove                  = -202
	ErrGroupUpdateOp                 = -203
	ErrGroupNonManager               = -204
	ErrGroupNonOwner                 = -205
	ErrGroupPromote                  = -206
	ErrGroupRemove                   = -207
	ErrGroupTransfer                 = -208
	ErrGroupNonMember                = -209
	ErrGroupNotFound                 = -210
	ErrGroupApplicationListNotFound  = -211
	ErrGroupApplyTwice               = -212
	ErrGroupApply                    = -213
	ErrGroupRefuse                   = -214
	ErrActivityGetData               = -300
	ErrActivityCannotApplyNotInGroup = -301
	ErrActivityUpdate                = -302
	ErrActivityNonPlanner            = -303
	ErrActivityCreate                = -304
	ErrActivityFee                   = -305
	ErrActivityOverQuota             = -306
	ErrActivityNotEnough             = -307
	ErrActivityRemove                = -308
	ErrActivityNotDoing              = -309
	ErrActivityCannotCancel          = -310
	ErrActivityHasBegun              = -311
	ErrActivityNotFound              = -312
)
