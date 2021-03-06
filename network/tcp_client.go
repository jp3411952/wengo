package network

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/panjf2000/ants/v2"
	"github.com/zjytra/wengo/csvdata"
	"github.com/zjytra/wengo/model"
	"github.com/zjytra/wengo/xlog"
	"net"
	"sync"
	"time"
)


// 每個对象维持一个连接,连接其他服务器使用
type TCPClient struct {
	sync.Mutex
	Addr            string        // 服务器连接地址
	ConnectInterval time.Duration // 重连时间
	PendingWriteNum int
	AutoReconnect   bool
	netObserver     NetWorkObserver // 客户端事件观察者
	tcpconn         *TcpConn        // 连接对象
	wg              sync.WaitGroup
	closeFlag       *model.AtomicBool
	netConf         *csvdata.Networkconf
	workPool        *ants.Pool // 协程池
	msgParser        *MsgParser // 数据包解析对象 共用一个对象
}

// 创建tcp 客戶端
func NewTCPClient(netObserver NetWorkObserver, netconf *csvdata.Networkconf, pool *ants.Pool) *TCPClient {
	if netconf == nil {
		xlog.WarningLogNoInScene( "server conf is nil")
		return nil
	}
	if netObserver == nil {
		xlog.WarningLogNoInScene( "服务器 消息处理 svNet is nil")
		return nil
	}
	client := new(TCPClient)
	client.netObserver = netObserver
	client.netConf = netconf
	client.workPool = pool
	client.closeFlag = model.NewAtomicBool()
	client.msgParser = NewMsgParser(netconf.Msglen_size, netconf.Max_msglen, netconf.Msg_isencrypt)
	return client
}

func (client *TCPClient) Start() {
	client.init()
	client.wg.Add(1)
	go client.connect() // 开启一个连接
}

func (client *TCPClient) init() {
	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 5 * time.Second
		xlog.DebugLogNoInScene( "invalid ConnectInterval, reset to %v", client.ConnectInterval)
	}
	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 100
		xlog.DebugLogNoInScene( "invalid PendingWriteNum, reset to %v", client.PendingWriteNum)
	}
	client.closeFlag.SetFalse()
	client.AutoReconnect = true
	
	client.Addr = fmt.Sprintf("%s:%s", client.netConf.Out_addr, client.netConf.Out_prot)
	xlog.DebugLogNoInScene( "client.Addr connet %v ", client.Addr)
}

func (client *TCPClient) dial() net.Conn {
	for {
		if client.closeFlag.IsTrue() {
			return nil
		}
		conn, err := net.Dial("tcp", client.Addr)
		if err == nil {
			return conn
		}
		
		xlog.DebugLogNoInScene( "TCPClient dial to %v error: %v", client.Addr, err)
		time.Sleep(client.ConnectInterval)
		continue
	}
}

func (client *TCPClient) connect() {
	defer client.wg.Done()
redial:
	client.doconnect() // 执行连接及读
	// 没有关闭才进行重连
	if client.closeFlag.IsFalse() && client.AutoReconnect {
		time.Sleep(client.ConnectInterval)
		goto redial
	}
}

func (client *TCPClient) doconnect() {
	conn := client.dial()
	if conn == nil {
		return
	}
	if !client.setConn(conn) {
		return
	}
}

// 添加链接信息
func (client *TCPClient) setConn(conn net.Conn) bool {
	if client.closeFlag.IsTrue() {
		conn.Close()
		return false
	}
	tcpconn := newTcpConn(conn, nextID(), client.netObserver, client.netConf, client.workPool,client.msgParser)
	client.tcpconn = tcpconn
	xlog.DebugLogNoInScene( "连接远程 %v 地址成功", conn.RemoteAddr())
	// 连接成功,将阻塞读取数据
	client.ReceiveData(client.tcpconn)
	xlog.DebugLogNoInScene( "TCPClient结束读取")
	return true
}

// 连接中读取数据
func (client *TCPClient) ReceiveData(conn *TcpConn) {
	for {
		err := conn.ReadMsg()
		if err != nil { // 这里读到错误消息,关闭
			xlog.WarningLogNoInScene("read message: %v ", err)
			break // 关闭连接
		}
	}
	// cleanup
	client.closeConn(conn)
}

func (client *TCPClient) closeConn(conn *TcpConn) {
	conn.Close()
	xlog.DebugLogNoInScene( "关闭远程连接")
}

// 写单个消息
func (client *TCPClient) WriteOneMsg(maincmd, subcmd uint16, msg []byte) error {
	if client.closeFlag.IsTrue() {
		return colseErro
	}
	if client.tcpconn == nil {
		return errors.New(fmt.Sprintf("TCPClient WriteOneMsg 未建立连接 %v",client.Addr))
	}
	return client.tcpconn.WriteOneMsg(maincmd, subcmd, msg)
}

// 将消息体构建为[]byte数组，最终要发出去的单包
func (client *TCPClient) GetOneMsgByteArr(maincmd, subcmd uint16, msg []byte) ([]byte, error) {
	if client.closeFlag.IsTrue() {
		return nil, colseErro
	}
	if client.tcpconn == nil {
		return nil,errors.New(fmt.Sprintf("GetOneMsgByteArr未建立连接 %v", client.Addr))
	}
	return client.tcpconn.GetOneMsgByteArr(maincmd, subcmd, msg)
}

// 写单个消息pb实现
func (client *TCPClient) WritePBMsg(maincmd, subcmd uint16, pb proto.Message) error {
	if client.closeFlag.IsTrue() {
		return colseErro
	}
	if client.tcpconn == nil {
		return errors.New(fmt.Sprintf("WritePBMsg未建立连接 %v", client.Addr))
	}
	return client.tcpconn.WritePBMsg(maincmd, subcmd, pb)
}

// 将消息体构建为[]byte数组，最终要发出去的单包 pb实现
func (client *TCPClient) GetPBByteArr(maincmd, subcmd uint16, pb proto.Message) ([]byte, error) {
	if client.closeFlag.IsTrue() {
		return nil, colseErro
	}
	if client.tcpconn == nil {
		return nil,errors.New(fmt.Sprintf("GetPBByteArr未建立连接 %v", client.Addr))
	}
	return client.tcpconn.GetPBByteArr(maincmd, subcmd, pb)
}

// 一起写多个数据包
// 每个包的数据 由GetOneMsgByteArr构建
func (client *TCPClient) WriteMsg(args ...[]byte) error {
	if client.closeFlag.IsTrue() {
		return colseErro
	}
	return client.tcpconn.WriteMsg(args...)
}

// 是否存活 没有存活会被提下线
func (client *TCPClient) IsAlive() bool {
	return client.tcpconn.IsAlive()
}

// 是否关闭
func (client *TCPClient) IsClose() bool {
	return client.tcpconn.IsClose()
}

// 获取连接对象id
func (client *TCPClient) GetConnID() uint32 {
	return client.tcpconn.GetConnID()
}

//关闭服务
func (client *TCPClient) Close() {
	client.closeFlag.SetTrue()
	client.tcpconn.Close()
	client.wg.Wait()
}

//关闭连接
func (client *TCPClient) DoCloseConn() {
	client.tcpconn.Close()
}

func (client *TCPClient) GetTcpConn() *TcpConn  {
	return client.tcpconn
}