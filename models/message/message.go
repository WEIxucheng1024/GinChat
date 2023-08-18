package message

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net"
	"net/http"
	"sync"
)

// 消息
type Message struct {
	gorm.Model
	FromId      string // 发送者
	TargetId    string // 接受者
	Type        string // 发送类型，比如私聊、群聊、广播啥的
	MessageType int    // 消息类型，比如纯文本、图片、音频等
	Context     string // 消息内容
	Pic         string // 图片
	Url         string // 附件的url等
	Desc        string // 描述
	Amounr      int    // 其他数字统计等
}

func (this *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets []interface{}
}

// 映射关系
var ClientMap = make(map[string]*Node)

// 读写锁
var reLocker sync.RWMutex

// 需要发送者UUID、接受者UUID、消息类型、消息内容、发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 1.获取参数，并鉴权
	query := request.URL.Query()
	userUUID := query.Get("userUUID")
	//token := query.Get("token")
	//msgType := query.Get("type")
	targetUUID := query.Get("targetUUID")
	//context := query.Get("context")

	// 鉴权，校验token
	/*
		这部分是鉴权，目前使用的是网上的测试工具，没办法传UUID和token
		var isvalida bool
		loginLog, err := auth.FindToken(userUUID, token)
		if err != nil {
			fmt.Println(err)
			return
		}
		if loginLog.UserUUID == "" {
			isvalida = false
		} else {
			isvalida = true
		}
	*/

	conn, err := (&websocket.Upgrader{
		// token校验
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2.获取连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte),
		GroupSets: make([]interface{}, 0),
	}

	// 3.获取用户关系

	// 4.userUUID跟node绑定并且加锁
	reLocker.RLock()
	ClientMap[userUUID] = node
	reLocker.RUnlock()

	// 5.完成发送逻辑
	go sendProc(node)
	// 6.完成接受逻辑
	go recvProc(node)

	sendMsg(targetUUID, []byte("首次进入测试！targetUUID"))
	sendMsg(userUUID, []byte("首次进入测试！fromUUID"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("监听DataQueue，数据为：", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<<", string(data))
	}
}

var udpSendChan = make(chan []byte)

func broadMsg(data []byte) {
	fmt.Println("写入udpSendChan，数据为：", string(data))
	udpSendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成upd数据发送协成
func udpSendProc() {
	// 这里配置发送目标的IP和端口,正常情况下应该是从上下文获取到IP，这里IPv4zero是本机IP
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()

	for {
		select {
		case data := <-udpSendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("监听udpSendChan，数据为：", string(data))
		}
	}
}

// 完成upd数据接收协成
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑
func dispatch(data []byte) {
	msg := Message{}
	json.Unmarshal(data, &msg)
	msg.Type = "privateChat"

	switch msg.Type {
	// 私聊
	case "privateChat":
		sendMsg(msg.TargetId, data)
	// 群聊
	case "groupChat":
		sendGroupMsg()
	// 广播
	case "All":
		sendAllMsg()
	}
}

func sendMsg(argetID string, msg []byte) {
	reLocker.RLock()
	node, ok := ClientMap[argetID]
	reLocker.RUnlock()
	if ok {
		fmt.Println("写入DataQueue，数据为：", string(msg))
		node.DataQueue <- msg
	}
}
func sendGroupMsg() {

}
func sendAllMsg() {

}
