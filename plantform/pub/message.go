package pub

type ReceiveMsgType string

type ReplyMsgType string

const (
	ReceiveMsgTypeText       ReceiveMsgType = "text"
	ReceiveMsgTypeImage      ReceiveMsgType = "image"
	ReceiveMsgTypeVoice      ReceiveMsgType = "voice"
	ReceiveMsgTypeVideo      ReceiveMsgType = "video"
	ReceiveMsgTypeShortVideo ReceiveMsgType = "shortvideo"
	ReceiveMsgTypeLocation   ReceiveMsgType = "location"
	ReceiveMsgTypeLink       ReceiveMsgType = "link"

	ReplyMsgTypeText  ReplyMsgType = "text"
	ReplyMsgTypeImage ReplyMsgType = "image"
	ReplyMsgTypeVoice ReplyMsgType = "voice"
	ReplyMsgTypeVideo ReplyMsgType = "video"
	ReplyMsgTypeMusic ReplyMsgType = "music"
	ReplyMsgTypeNews  ReplyMsgType = "news"
)

type CDATA struct {
	Text string `xml:",cdata"`
}

type ReceiveMsg struct {
	ToUserName   string         `xml:"ToUserName"`
	FromUserName string         `xml:"FromUserName"`
	CreateTime   uint64         `xml:"CreateTime"`
	MsgType      ReceiveMsgType `xml:"MsgType"`
	MsgId        uint64         `xml:"MsgId"`
	MsgDataId    string         `xml:"MsgDataId"`
	Idx          string         `xml:"Idx"`
	Content      string         `xml:"Content"`
	PicUrl       string         `xml:"PicUrl"`
	MediaId      string         `xml:"MediaId"`
	Format       string         `xml:"Format"`
	ThumbMediaId string         `xml:"ThumbMediaId"`
	Description  string         `xml:"Description"`
	Url          string         `xml:"Url"`
	LocationX    string         `xml:"Location_X"`
	LocationY    string         `xml:"Location_Y"`
	Label        string         `xml:"Label"`
}

type MediaId struct {
	MediaId CDATA `xml:"MediaId"`
}

type ReplyMsg struct {
	ToUserName   CDATA        `xml:"ToUserName"`
	FromUserName CDATA        `xml:"FromUserName"`
	CreateTime   uint64       `xml:"CreateTime"`
	MsgType      ReplyMsgType `xml:"MsgType"`
	Content      string       `xml:"Content"`
	Image        []MediaId    `xml:"Image"`
	Voice        []MediaId    `xml:"Voice"`
}

//
//type Text struct {
//	Message
//	Content string `xml:"Content"`
//}
//
//type Image struct {
//	Message
//	PicUrl  string `xml:"PicUrl"`
//	MediaId string `xml:"MediaId"`
//}
//
//type Voice struct {
//	MediaId string `xml:"MediaId"`
//	Format  string `xml:"Format"`
//}
//
//type Video struct {
//	MediaId      string `xml:"MediaId"`
//	ThumbMediaId string `xml:"ThumbMediaId"`
//}
//
//type ShortVideo struct {
//	MediaId      string `xml:"MediaId"`
//	ThumbMediaId string `xml:"ThumbMediaId"`
//}
//
//type Link struct {
//	Description string `xml:"Location_x"`
//	LocationY   string `xml:"Location_Y"`
//	MsgType     uint64 `xml:"Scale"`
//	Label       string `xml:"Label"`
//}
