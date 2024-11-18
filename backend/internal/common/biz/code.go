package biz

const Ok = 200

var (
	DBError         = NewError(10000, "数据库错误")
	AlreadyRegister = NewError(10100, "用户已注册")
)