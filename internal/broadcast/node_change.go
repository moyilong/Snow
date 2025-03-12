package broadcast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"snow/internal/membership"
	"snow/tool"
)

func initMsg(msg []byte) (MsgAction, int64, []byte) {
	//1 byte的类型，8byte的时间戳，然后是发消息节点的 ip加端口号
	time := tool.BytesToTime(msg[1 : 1+TimeLen])
	return msg[0], time, msg[1+TimeLen:]
}

func NodeChange(msg []byte, ip string, s *Server, conn net.Conn) {
	changeType, _, data := initMsg(msg)
	msg = msg[1:]
	switch changeType {
	case applyJoin:
		applyJoining(s, conn)
	case joinStateSync:
		joinStateSynchronizing(ip, data, s)
	case nodeJoin:
		joinStateSynchronizing(ip, data, s)
	default:
	}
}
func applyJoining(s *Server, conn net.Conn) {
	//接收到消息然后推送
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(s.Member.MetaData)
	if err != nil {
		fmt.Println("GOB Serialization failed:", err)
		return
	}
	fmt.Println(conn.RemoteAddr())
	state := buffer.Bytes()
	data := PackTagToHead(nodeChange, joinStateSync, state)
	replayMessage(conn, s.Config, data)
}

// 接收到同步消息
func joinStateSynchronizing(ip string, msg []byte, s *Server) {
	//同步节点的信息，同步完毕之后请求加入节点
	buffer := bytes.NewBuffer(msg)
	var MetaData map[string]*membership.MetaData
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&MetaData)
	if err != nil {
		fmt.Println("GOB Desialization failed:", err)
		return
	}
	s.Member.InitState(MetaData, s.Config.IPBytes())
	//使用标准消息广播自己需要加入
	s.ColoringMessage(s.Config.IPBytes(), nodeJoin)
}
func PackTagToHead(msgType MsgType, changeType MsgAction, msg []byte) []byte {
	data := make([]byte, len(msg)+TimeLen+TagLen)
	data[0] = msgType
	data[1] = changeType
	timeBytes := tool.TimeBytes()
	copy(data[TagLen:], timeBytes)
	copy(data[TimeLen+TagLen:], msg)
	return data
}

func PackTag(msgType MsgType, changeType MsgAction) []byte {
	data := make([]byte, TimeLen+TagLen)
	data[0] = msgType
	data[1] = changeType
	timeBytes := tool.TimeBytes()
	copy(data[TagLen:], timeBytes)
	return data
}
func nodeJoining(ip []byte) {

}
