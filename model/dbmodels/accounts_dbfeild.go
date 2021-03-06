//生成的文件建议不要改动,详见mysql-struct-maker.go ParseColumn方法源码生成格式 
package dbmodels 

type Accounts struct {
	AccountID uint64 `sql:"AccountID"` // 数据库注释:账户id 
 	LoginName string `sql:"LoginName"` // 数据库注释:登陆账号 
 	LoginPwd string `sql:"LoginPwd"` // 数据库注释:登陆密码 
 	RegisterTime string `sql:"RegisterTime"` // 数据库注释:注册时间 
 	RegisterIp string `sql:"RegisterIp"` // 数据库注释:注册ip 
 	LoginTime *string `sql:"LoginTime"` // 数据库注释:登陆时间 
 	LoginIp *string `sql:"LoginIp"` // 数据库注释:登陆ip 
 	IsBan uint32 `sql:"IsBan"` // 数据库注释:是否被屏蔽 
 	IsGM uint8 `sql:"IsGM"` // 数据库注释:是否是gm 
 	Phone string `sql:"Phone"` // 数据库注释:电话号码 
 	RegMacAddr string `sql:"RegMacAddr"` // 数据库注释:注册的mac地址 
 	LoginMacAddr *string `sql:"LoginMacAddr"` // 数据库注释:登录时mac地址 
 }
