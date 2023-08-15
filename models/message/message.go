package message

import "gorm.io/gorm"

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
