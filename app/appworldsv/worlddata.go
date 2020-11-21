/*
创建时间: 2020/5/17
作者: zjy
功能介绍:

*/

package appworldsv

import (
	"wengo/app/datacenter/dcmodel"
	"wengo/xlog"
)

var (
	ServerInfoS map[uint32]*dcmodel.SeverInfoModel //使用connID作为key 查找得
	ServerKind  map[int32][]uint32
	LinkServer  map[int32]uint32    //避免同一个服务器重复连接
)

func NewData() {
	ServerInfoS = make(map[uint32]*dcmodel.SeverInfoModel)
	ServerKind = make(map[int32][]uint32)
	LinkServer = make(map[int32]uint32)
}

//添加服务器信息
func AddServerInfo(info *dcmodel.SeverInfoModel) bool {
	if info == nil {
		xlog.DebugLogNoInScene("AddServerInfo erro")
		return false
	}
	_,hasLink := LinkServer[info.AppId]
	if hasLink {
		xlog.DebugLogNoInScene("serverId%v 已经在中心服注册",info.AppId)
		return false
	}
	//服务器已经注册
	if _,ok := ServerInfoS[info.ConnID];ok{
		xlog.DebugLogNoInScene("ClientConnID %v 已经在中心服注册",info.ConnID)
		return false
	}
	
	// 将同一类的服务器放在一起
	if connIDs, ok := ServerKind[info.AppKind]; ok { //已经存在同类
		var isFind bool
		for _, conid := range connIDs {
			if conid == info.ConnID {
				isFind = true
				//已经放在列表里面
				xlog.DebugLogNoInScene("serverId %v 已经在类型列表中",info.AppId)
				break
			}
		}
		//没找到才添加
		if !isFind {
			connIDs = append(connIDs, info.ConnID )
		}
	} else {
		var connIDs []uint32
		connIDs = append(connIDs, info.ConnID )
		ServerKind[info.AppKind] = connIDs
	}
	LinkServer[info.AppId] = info.ConnID
	ServerInfoS[info.ConnID] = info
	xlog.DebugLogNoInScene("AppId %v  info.ConnID = %v注册成功",info.AppId,info.ConnID)
	return  true
}


// 移除某个连接
func RemoveServerInfo(connID uint32) bool {
	
	//服务器已经注册
	pServerInfo,ok := ServerInfoS[connID];
	if !ok {
		xlog.DebugLogNoInScene("未找到 connID = %v 的服务器",connID)
		return false
	}
	var appid  = pServerInfo.AppId  //服务器id
	var appkind = pServerInfo.AppKind //服务器类型
	pServerInfo = nil //移除变量

	_,hasLink := LinkServer[appid]
	if hasLink {
		delete(LinkServer,appid)
		xlog.DebugLogNoInScene("移除 appid = %v 的服务器",appid)
	}else {
		xlog.ErrorLogNoInScene("未找到 appid = %v 的服务器",appid)
	}
	var isFind bool
	if 	connIDs,ok := ServerKind[appkind]; ok {
		var svlen = len(connIDs)
		for i := 0 ; i < svlen ; i++ {
			if  connIDs[i] == connID {
				isFind = true
				xlog.DebugLogNoInScene("在同类中移除连接 connID = %v 的服务器",connID)
				connIDs = append(connIDs[:i],connIDs[i+1:]...) //移除找到连接
				break
			}
		}
	}
	if !isFind {
		xlog.DebugLogNoInScene("在同类中服务器中未找到 connID = %v 的服务器",connID)
	}
	delete(ServerInfoS,connID)
	xlog.DebugLogNoInScene("移除连接 appid =%v connID = %v 的服务器",appid ,connID)
	return true
}

func ClearAllServerData()  {
	ServerInfoS = nil
	for _,v := range ServerKind  {
		v = v[:0:0] //清空切片
	}
	ServerKind = nil
	LinkServer = nil
}