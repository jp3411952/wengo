/*
创建时间: 2020/08/2020/8/25
作者: Administrator
功能介绍:

*/
package datacenter

import (
	"errors"
	"fmt"
	"github.com/zjytra/wengo/app/datacenter/dcdbmodel"
	"github.com/zjytra/wengo/cmdconst"
	"github.com/zjytra/wengo/cmdconst/cmdaccount"
	"github.com/zjytra/wengo/dbsys"
	"github.com/zjytra/wengo/dispatch"
	"github.com/zjytra/wengo/model/dbmodels"
	"github.com/zjytra/wengo/msgcode"
	"github.com/zjytra/wengo/protobuf/pb/account_proto"
	"github.com/zjytra/wengo/xlog"
	"github.com/zjytra/wengo/xutil/timeutil"
	"time"
)

type AccountRegister struct {
	DBParams    *dispatch.DBEventParam
	accountNum  uint8              //数据库返回的账号数量
	accountData *dbmodels.Accounts //数据库查询出的数据
}

func (this *AccountRegister) SetDBEventParam(param *dispatch.DBEventParam) {
	this.DBParams = param
}

//数据库线程调用
func (this *AccountRegister) ExcouteQueryFun() error {
	regAccout, ok := this.DBParams.ReqParam.(*dcdbmodel.DB_Req_CreateAccount) //创建
	if !ok {
		return errors.New("AccountRegister ExcouteQueryFun 解析对象dcdbmodel.DB_Req_CreateAccount失败")
	}
	//保证账号唯一性
	dbsys.AccountMutex.Lock()
	//这里执行查询
	row, erro := dbsys.GameDB.Query("SELECT * FROM accounts where LoginName = ?; ", regAccout.Username)
	if erro != nil {
		this.ErroReturn()
		return erro
	}
	this.accountData = new(dbmodels.Accounts)
	dbsys.RowToStruct(row, this.accountData) //解析row
	row.Close()
	if this.accountData.AccountID == 0 { //没有数据
		err := this.InsertAccountInDB(regAccout)
		if err != nil {
			return err
		}
	} else {
		this.DBParams.DBRestCode = msgcode.AccountCode_IsExsist
	}
	//释放锁
	dbsys.AccountMutex.Unlock()
	//向逻辑线程投递
	if this.DBParams.CbDispSys == nil {
		xlog.DebugLogNoInScene("AccountRegister.ExcouteQueryFun 调度事件=nil 或者回调 = nil 就不向逻辑队列投递事件了")
		//出错了要回收数据
		dbsys.PDBParamPool.Recycle(this.DBParams)
		PAccountPool.Recycle(this)
		return nil
	}
	//进入调度队列
	onEventErro := this.DBParams.CbDispSys.PostCustomDBOperateOneRow(this)
	if onEventErro != nil {
		xlog.ErrorLogNoInScene("投递查询事件 %v", onEventErro)
	}
	return onEventErro
}

func (this *AccountRegister) InsertAccountInDB(regAccout *dcdbmodel.DB_Req_CreateAccount) error {
	
	//账号创建过多直接返回不走后面的逻辑
	if !this.CheckAccountNumIsMore(regAccout) {
		return nil
	}
	//获取当前时间
	currentT := time.Now()
	nowTimeStr := timeutil.GetTimeALLStr(currentT)
	result, erro := dbsys.GameDB.Excute("INSERT INTO accounts(LoginName,LoginPwd,RegisterTime,RegisterIp,Phone,RegMacAddr) VALUES (?,?,?,?,?,?)",
		regAccout.Username,
		regAccout.Password,
		nowTimeStr,
		regAccout.ClientIp,
		regAccout.PhoneNum,
		regAccout.MacAddr)
	
	if erro != nil {
		this.ErroReturn()
		return erro
	}
	lastID, erro := result.LastInsertId()
	if erro != nil {
		xlog.ErrorLogNoInScene("LastInsertId erro = %v", erro)
		this.ErroReturn()
		return erro
	}
	xlog.ErrorLogNoInScene("创建账号ID = %v", lastID)
	logSql := fmt.Sprintf("INSERT INTO accounts_register_record_%s(AccountID,LoginName,LoginPwd,RegisterTime,RegisterIp,Phone,RegMacAddr) VALUES (?,?,?,?,?,?,?)",
		timeutil.GetYearMonthFromatStr(currentT))
	
	_, erro = dbsys.LogDB.Excute(logSql,
		lastID,
		regAccout.Username,
		regAccout.Password,
		nowTimeStr,
		regAccout.ClientIp,
		regAccout.PhoneNum,
		regAccout.MacAddr)
	
	if erro != nil {
		xlog.ErrorLogNoInScene("创建账号记录日志失败= %v", erro)
	}
	//再执行查询
	row, erro := dbsys.GameDB.Query("SELECT * FROM accounts where LoginName =? AND LoginPwd =?; ", regAccout.Username, regAccout.Password)
	if erro != nil {
		this.ErroReturn()
		return erro
	}
	//解析row
	dbsys.RowToStruct(row, this.accountData)
	row.Close()
	this.DBParams.DBRestCode = msgcode.AccountCode_Succeed //创建账号成功
	return nil
}

//查看机器码的账号是否过多
func (this *AccountRegister) CheckAccountNumIsMore(regAccout *dcdbmodel.DB_Req_CreateAccount) bool {
	//查看机器码创建账号数量
	rows, erro := dbsys.GameDB.Query("SELECT count(1) as num FROM accounts where RegMacAddr = ?; ", regAccout.MacAddr)
	if erro != nil {
		this.ErroReturn()
		return false
	}
	for rows.Next() {
		erro := rows.Scan(&this.accountNum)
		if erro != nil {
			xlog.ErrorLogNoInScene("扫描账号错误 erro = %v", erro)
			this.ErroReturn()
			return false
		}
	}
	rows.Close()
	if this.accountNum > 9 { //最多10个
		this.DBParams.DBRestCode = msgcode.AccountCode_MACAccountNumIsMore //创建的账号过多
		return false
	}
	return true
}

//错误返回
func (this *AccountRegister) ErroReturn() {
	//释放锁
	dbsys.AccountMutex.Unlock()
	//出错了要回收数据
	dbsys.PDBParamPool.Recycle(this.DBParams)
	PAccountPool.Recycle(this)
}

//逻辑线程调用
func (this *AccountRegister) OnQueryCB() error {
	if this.accountData == nil {
		return nil
	}
	restMsg := &account_proto.DC_LS_RespnRegisterAccoutMsg{
		ClientConnID: this.DBParams.ClientConnID,
		RestCode:     this.DBParams.DBRestCode,
		Username:     this.accountData.LoginName,
	}
	erro := PdataCenter.tcpserver.WritePBMsgByConnID(this.DBParams.ServerConnID,cmdconst.Main_Account,cmdaccount.Sub_DC_LS_RegisterAccount, restMsg)
	if this.accountData != nil {
		if this.accountNum > 0 {
			//设置账号数量
			PaccountMgr.SetMacCreateAccountNum(this.accountData.RegMacAddr, this.accountNum)
		}
		//将账号保存在内存中
		if PaccountMgr.GetAccountByUserName(this.accountData.LoginName) == nil && this.accountData.AccountID > 0 {
			paccount := PaccountMgr.AddAccunts(this.accountData)
			if paccount != nil {
				now := timeutil.NowAddDate(0,0,7) //设置7天后过期
				paccount.SetExprationTime(now.Unix())
			}
			xlog.DebugLogNoInScene("OnQueryOneRowCB accout = %v", this.accountData)
		}
	}
	//最后回收数据
	dbsys.PDBParamPool.Recycle(this.DBParams)
	PAccountPool.Recycle(this)
	return erro
}
func (this *AccountRegister) Reset() {
	this.DBParams.Reset()
	this.accountNum = 0
	this.accountData = nil
}
