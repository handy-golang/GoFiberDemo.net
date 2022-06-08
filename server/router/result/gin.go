package result

import "github.com/EasyGolang/goTools/mRes"

var (
	Not                 = mRes.Response(0, "Not")     // 无响应
	OK                  = mRes.Response(1, "Succeed") // 通用成功
	LoginSucceed        = mRes.Response(2, "登录成功")
	RegisterSucceed     = mRes.Response(3, "用户数据录入成功")
	EditPasswordSucceed = mRes.Response(4, "修改成功")
	SendCodeSucceed     = mRes.Response(5, "验证码已发送")
	UploadSucceed       = mRes.Response(6, "上传成功")
	EditUserSucceed     = mRes.Response(7, "修改成功")
	CreateHunterSucceed = mRes.Response(8, "创建成功")
	CreateHunterShell   = mRes.Response(9, "脚本已生成")

	Fail           = mRes.Response(-1, "Fail") // 通用错误
	ErrToken       = mRes.Response(-2, "Token验证失败")
	ErrApiAuth     = mRes.Response(-3, "授权验证失败")
	ErrCodeExpired = mRes.Response(-4, "验证码已过期")
	ErrPassword    = mRes.Response(-6, "密码错误")
	ErrHz          = mRes.Response(-7, "请求太频繁")
	ErrName        = mRes.Response(-8, "5-10位小写字母或数字")
	ErrDB          = mRes.Response(-9, "数据库出错")
	ErrRole        = mRes.Response(-10, "角色不正确!")
	ErrUserNote    = mRes.Response(-11, "请输入角色描述")
	ErrRoleAuth    = mRes.Response(-12, "当前用户无权限")
	ErrUserNot     = mRes.Response(-13, "该账号不存在")
	ErrUserYet     = mRes.Response(-14, "该账号已存在")
	ErrRmUser      = mRes.Response(-15, "注销账户失败")
	// -15
	// -16
	ErrUpload       = mRes.Response(-17, "上传失败")
	ErrEditUser     = mRes.Response(-18, "用户信息修改失败")
	ErrCreateHunter = mRes.Response(-19, "Hunter创建失败")
)
