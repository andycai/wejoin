package constant

const (
	Success                          = 0
	SucGroupApply                    = 200
	SucGroupApprove                  = 201
	SucGroupRefuse                   = 202
	SucGroupUpdateName               = 203
	SucGroupUpdateNotice             = 204
	SucGroupUpdateAddr               = 205
	SucGroupQuit                     = 206
	SucGroupRemove                   = 207
	SucGroupFire                     = 208
	SucGroupTransfer                 = 209
	SucGroupPromote                  = 210
	SucGroupCancel                   = 211
	ErrParam                         = 10000
	ErrData                          = 10001
	ErrOp                            = 10002
	ErrTwoPasswordNotMatch           = 10003
	ErrUserAuth                      = 10100
	ErrUserRegister                  = 10101
	ErrUserData                      = 10102
	ErrUserUpdateData                = 10103
	ErrUserNotFound                  = 10104
	ErrUserEmailFormat               = 10105
	ErrGroupManagerLimit             = 10200
	ErrGroupGetData                  = 10201
	ErrGroupApprove                  = 10202
	ErrGroupUpdateOp                 = 10203
	ErrGroupNonManager               = 10204
	ErrGroupNonOwner                 = 10205
	ErrGroupPromote                  = 10206
	ErrGroupRemove                   = 10207
	ErrGroupTransfer                 = 10208
	ErrGroupNonMember                = 10209
	ErrGroupNotFound                 = 10210
	ErrGroupApplicationListNotFound  = 10211
	ErrGroupApplyTwice               = 10212
	ErrGroupApply                    = 10213
	ErrGroupRefuse                   = 10214
	ErrGroupUpdateName               = 10215
	ErrGroupUpdateNotice             = 10216
	ErrGroupUpdateAddr               = 10217
	ErrGroupQuit                     = 10218
	ErrGroupFire                     = 10219
	ErrGroupCancel                   = 10220
	ErrActivityGetData               = 10300
	ErrActivityCannotApplyNotInGroup = 10301
	ErrActivityUpdate                = 10302
	ErrActivityNonPlanner            = 10303
	ErrActivityCreate                = 10304
	ErrActivityFee                   = 10305
	ErrActivityOverQuota             = 10306
	ErrActivityNotEnough             = 10307
	ErrActivityRemove                = 10308
	ErrActivityNotDoing              = 10309
	ErrActivityCannotCancel          = 10310
	ErrActivityHasBegun              = 10311
	ErrActivityNotFound              = 10312
	ErrActivityApply                 = 10313
)