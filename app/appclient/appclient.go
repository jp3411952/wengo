/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
登录服
*/

package appclient

import (
	"sync"
	"time"
	"github.com/zjytra/wengo/app/appdata"
	"github.com/zjytra/wengo/app/netmsgsys"
	"github.com/zjytra/wengo/cmdconst"
	"github.com/zjytra/wengo/cmdconst/cmdaccount"
	"github.com/zjytra/wengo/csvdata"
	"github.com/zjytra/wengo/dispatch"
	"github.com/zjytra/wengo/model"
	"github.com/zjytra/wengo/network"
	"github.com/zjytra/wengo/protobuf/pb/account_proto"
	"github.com/zjytra/wengo/timersys"
	"github.com/zjytra/wengo/xlog"
	"github.com/zjytra/wengo/xutil/osutil"
)

type AppClient struct {
	NetWorkInfo *model.AppNetWorkModel // 服务器网络信息
	netmsgsys   *netmsgsys.NetMsgSys
	conns       sync.Map
	tcpclient   *network.TCPClient
	sndHeartTimer uint32 //定时起
	dispSys   *dispatch.DispatchSys
	span       time.Duration
}

// 程序启动
func (this *AppClient) OnStart() {
	this.OnInit()
	this.dispSys = dispatch.NewDispatchSys()
	this.dispSys.SetNetObserver(this) //模拟客户端
	// dispSys.SetServiceNet(this) // 模拟服务器连接
	this.tcpclient = network.NewTCPClient(this.dispSys, appdata.NetConf, appdata.WorkPool)
	this.tcpclient.Start()
	this.span = time.Second * 5
	//for i:= 0; i < 3; i++ {
	//	time.Sleep(time.Second * 10)
	//	go func() {
	//		clent := network.NewTCPClient(this.dispSys, appdata.NetConf, appdata.WorkPool)
	//		clent.Start()
	//		// 连接成功发送登陆命令
	//		time.Sleep(time.Second)
	//		this.TestRegisterAccount(clent.GetTcpConn())
	//		//账号登录
	//		//this.TestLoginAccount(clent.GetTcpConn())
	//	}()
	//}
	
	//this.sndHeartTimer = timersys.NewWheelTimer(time.Second * 30,this.TestTimer,this.dispSys)
}

func (this *AppClient) TestRegisterAccount(conn network.Conner) {
	reqCreateAccount := &account_proto.CL_LS_ReqRegisterAccoutMsg{
		Username:   "zjy082",
		Password:   "jp3411952",
		ClientType: model.ClientType_Test,
		MacAddr:    osutil.GetUpMacAddr(),
		Version:    1001,
	}
	var erro error
	if conn != nil {
		erro = conn.WritePBMsg(cmdconst.Main_Account, cmdaccount.Sub_C_LS_RegisterAccount, reqCreateAccount)
	}else {
		erro = this.tcpclient.WritePBMsg(cmdconst.Main_Account, cmdaccount.Sub_C_LS_RegisterAccount, reqCreateAccount)
	}
	if erro != nil {
		xlog.ErrorLogNoInScene("TestRegisterAccount write erro %v ", erro)
	}
}
func (this *AppClient) TestLoginAccount(conn network.Conner) {
	reqMsg := &account_proto.CL_LS_ReqLoginMsg{
		Username:   "zty111uuy",
		Password:   "jp3411952",
		ClientType: model.ClientType_Test,
		MacAddr:    osutil.GetUpMacAddr(),
		Version:    1001,
	}
	var erro error
	if conn != nil {
		erro = conn.WritePBMsg(cmdconst.Main_Account, cmdaccount.Sub_C_LS_LoginAccount, reqMsg)
	}else {
		erro = this.tcpclient.WritePBMsg(cmdconst.Main_Account, cmdaccount.Sub_C_LS_LoginAccount, reqMsg)
	}
	if erro != nil {
		xlog.ErrorLogNoInScene("TestLoginAccount write erro %v ", erro)
	}
}


// 发送心跳给世界服
func (this *AppClient) TestTimerTestTimer() {
	defer xlog.RecoverToLog(func() {
		timersys.StopTimer(this.sndHeartTimer)
	})
	//账号登录
	this.TestLoginAccount(nil)
	//this.TestRegisterAccount()

}
// 初始化
func (this *AppClient) OnInit() bool {
	csvdata.LoadCommonCsvData()
	this.netmsgsys = netmsgsys.NewMsgHandler()
	this.RegisterServerMsg()
	return true
}

// 程序运行
func (this *AppClient) OnUpdate() bool {
	// xlog.DebugLog("","run LoginApp")
	
	return true
}

// 关闭
func (this *AppClient) OnRelease() {
	this.tcpclient.Close()
	timersys.Release()
}

func (this *AppClient) OnServiceLink(conn network.Conner) error {
	return nil
}

func (this *AppClient) OnServiceClose(conn network.Conner) error {
	xlog.DebugLogNoInScene( "AppClient OnServiceClose %v", conn.RemoteAddr())
	return nil
}

func (this *AppClient) OnServiceMsg(msgdata *network.MsgData) error {
	return this.netmsgsys.OnNetWorkMsgHandle(msgdata)
}


//客户端连接
func (this *AppClient)OnNetWorkConnect(conn network.Conner) error{
	xlog.DebugLogNoInScene( "OnNetWorkConnect %v", conn.RemoteAddr())
	//time.Sleep(this.span)
	////this.span += time.Second * 5
	////if 	this.span > time.Second * 20 {
	////	this.span = time.Second
	////}
	////this.TestLoginAccount(conn)
	this.TestRegisterAccount(conn)
	return nil
}

func (this *AppClient) SendCreateAccount() error {
	// 连接成功发送登陆命令
	reqCreateAccount := &account_proto.CL_LS_ReqRegisterAccoutMsg{
		Username:   "zty111uuy",
		Password:   "jp3411952",
		ClientType: model.ClientType_Test,
		MacAddr:    osutil.GetUpMacAddr(),
		Version: 1001,
	}
	erro := this.tcpclient.WritePBMsg(cmdconst.Main_Account, cmdaccount.Sub_C_LS_RegisterAccount, reqCreateAccount)
	if erro != nil {
		xlog.ErrorLogNoInScene("OnNetWorkConnect write erro %v ", erro.Error())
	}
	return nil
}

//客户端关闭连接
func (this *AppClient)OnNetWorkClose(conn network.Conner) error{
	xlog.DebugLogNoInScene("远端关闭了连接%v ", conn.RemoteAddr())
	return  nil
}

//读取客戶端发来的消息
func (this *AppClient)OnNetWorkRead(msgdata *network.MsgData) error{
	return	this.netmsgsys.OnNetWorkMsgHandle(msgdata)
}

// 注册服务器 的消息
func (this *AppClient)RegisterServerMsg(){
	this.netmsgsys.RegisterMsgHandle(cmdconst.Main_Account, cmdaccount.Sub_LS_C_RegisterAccount, OnRegisterAccountHanlder)
	this.netmsgsys.RegisterMsgHandle(cmdconst.Main_Account, cmdaccount.Sub_LS_C_LoginAccount, OnRespnLoginAccountHanlder)
}

