package user

import (
	"encoding/json"
	"ny2/bs"
	"ny2/bs/entity"
	"ny2/utils"
	"ny2/wxapi"
	"strconv"
	"time"
)

func login(avatarUrl string, nickName string, openId string) (bool, string) {
	var s bs.Stu
	s.Us.OpenId = openId
	s.Us.AvatarUrl = avatarUrl
	s.Us.NickName = nickName
	r := s.Insert()
	s.SelectByOpenId()

	type t struct {
		OpenId string `json:"open_id"`
		StuId  int64  `json:"stu_id"`
	}
	var data t
	data.OpenId = s.Us.OpenId
	data.StuId = s.StuId
	if bytes, err := json.Marshal(data); err == nil && r {
		return true, string(bytes)
	}

	return false, utils.EmptyString
}

func getStuMsg(stuIdA int64, stuIdB int64, limit int) (bool, string) {
	var u bs.Stu
	u.StuId = stuIdA
	if ok, data := u.SelectStuMsg(stuIdB, limit); ok {
		go func() {
			for _, d := range data {
				var md entity.Detail
				md.MsgId = d.Id
				md.StuId = stuIdA
				md.IsRead = true
				u.InsertMsgDetail(&md)
			}
		}()

		if bytes, err := json.Marshal(data); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

func sendTxtMessage(senderStuId int64, recipientStuId int64, content string) bool {
	var s bs.Stu
	m := entity.Msg{
		SenderStu:    senderStuId,
		RecipientStu: recipientStuId,
		Content:      content,
		Type:         "txt",
	}
	r := s.InsertMsg(&m)

	s.InsertMsgDetail(&entity.Detail{
		MsgId:  m.Id,
		StuId:  senderStuId,
		IsRead: true,
	})

	return r

}

func isOpenIdExist(openId string) bool {
	var s bs.Stu
	s.Us.OpenId = openId
	return s.IsOpenIdExist()
}

func putInfoByOpenId(openId string, dormId int64, stuNumber string, stuMobile string, room string) bool {
	var u bs.Stu
	u.Us.OpenId = openId
	r := u.SelectByOpenId()
	if !r {
		return false
	}
	u.DormId = dormId
	u.StuNumber = stuNumber
	u.Us.Mobile = stuMobile
	u.DormRoom = room
	return u.UpdateById()
}

func getSimpleInfoByStuId(stuId int64) (bool, string) {
	var u bs.Stu
	u.StuId = stuId
	r := u.SelectByStuId()
	if r == false {
		return false, utils.EmptyString
	}
	return getSimpleInfoByOpenId(u.Us.OpenId)
}

func sendOrder(orderType int, stuId int64, price string, endTime time.Time, comment string, templateId string) bool {
	var o entity.Order
	var s bs.Stu
	s.StuId = stuId
	r := s.SelectByStuId()
	if r == false {
		return false
	}

	o.Price = price
	o.FinishTime = &endTime
	o.Comment = comment
	o.TemplateId = templateId
	o.AvatarUrl = s.Us.AvatarUrl
	typeStr := strconv.Itoa(orderType)
	o.Type = typeStr

	r = s.InsertOrder(&o)
	return r

}

func hasNewUnreadMsg(openId string) bool {
	var u bs.Stu
	u.Us.OpenId = openId
	r := u.SelectByOpenId()
	if !r {
		return false
	}
	return u.HasUnReadMsg()
}

func getUnreadNewestMsg(openId string) (bool, string) {
	var u bs.Stu
	u.Us.OpenId = openId
	r := u.SelectByOpenId()
	if !r {
		return false, utils.EmptyString
	}
	if ok, data := u.SelectNewestUnreadMsg(); ok {

		//加入新的字段
		type t struct {
			entity.Msg
			NickName     string `json:"nick_name"`
			SenderAvatar string `json:"sender_avatar"`
		}
		var tmp []t
		for _, v := range data {
			var u bs.Stu
			u.StuId = v.SenderStu
			u.SelectByStuId()

			var t1 t
			t1.RecipientStu = v.RecipientStu
			t1.Id = v.Id
			t1.Type = v.Type
			t1.Content = v.Content
			t1.SenderStu = v.SenderStu
			t1.CreateTime = v.CreateTime
			t1.SenderAvatar = u.Us.AvatarUrl
			t1.NickName = u.Us.NickName

			tmp = append(tmp, t1)
		}

		if bytes, err := json.Marshal(tmp); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

func getSimpleInfoByOpenId(openId string) (bool, string) {
	var u bs.Stu
	u.Us.OpenId = openId
	r := u.SelectByOpenId()
	if !r {
		return false, utils.EmptyString
	}
	type t struct {
		DormId    int64  `json:"dorm_id"`
		Mobile    string `json:"mobile"`
		Room      string `json:"room"`
		SchoolId  int64  `json:"school_id"`
		StuId     int64  `json:"stu_id"`
		StuNumber string `json:"stu_number"`
		NickName  string `json:"nick_name"`
		AvatarUrl string `json:"avatar_url"`
	}
	var tmp t
	tmp.DormId = u.DormId
	tmp.StuNumber = u.StuNumber
	tmp.Mobile = u.Us.Mobile
	tmp.Room = u.DormRoom
	tmp.SchoolId = u.Dm.SchoolId
	tmp.StuId = u.StuId
	tmp.NickName = u.Us.NickName
	tmp.AvatarUrl = u.Us.AvatarUrl

	if bytes, err := json.Marshal(tmp); err == nil {
		return true, string(bytes)
	}
	return false, utils.EmptyString
}

// 获取与stu相关的订单数量
func getStuOrderSize(stuId int64) (bool, string) {
	var u bs.Stu
	u.StuId = stuId
	if ok, size := u.SelectOrderLength(); ok {
		if bytes, err := json.Marshal(size); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString

}

// 获取与stu相关的订单详细信息
func getStuOrder(stuId int64, limit int64, offset int64) (bool, string) {
	var u bs.Stu
	u.StuId = stuId
	ok, data := u.SelectOrderByStuId(limit, offset)
	if ok {
		if bytes, err := json.Marshal(data); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

// 获取与stu相关的订单简略信息
func getStuPreOrder(stuId int64, limit int64, offset int64) (bool, string) {
	var u bs.Stu
	u.StuId = stuId
	ok, data := u.SelectOrderByStuId(limit, offset)

	type t struct {
		OrderId    int64      `json:"order_id"`
		AvatarUrl  string     `json:"avatar_url"`
		FinishTime *time.Time `json:"finish_time"`
		Price      string     `json:"price"`
		Type       string     `json:"type"`
		StuId      int64      `json:"stu_id"`
		NickName   string     `json:"nick_name"`
		Dorm       string     `json:"dorm"`
		Cancel     bool       `json:"cancel"`
		RecvStu    int64      `json:"recv_stu"`
	}
	var tmp []t
	for _, v := range data {
		var k t
		k.FinishTime = v.FinishTime
		k.Price = v.Price
		k.Type = v.Type
		k.AvatarUrl = v.AvatarUrl
		k.OrderId = v.Id
		if v.StuId != 0 {
			k.StuId = v.StuId
		}
		k.Cancel = v.Cancel
		k.RecvStu = v.RecvStu

		var u bs.Stu
		u.StuId = k.StuId
		u.SelectByStuId()
		k.NickName = u.Us.NickName

		var d bs.Dorm
		d.Id = v.DormId
		d.SelectById()

		k.Dorm = d.DormName

		tmp = append(tmp, k)
	}
	if ok {
		if bytes, err := json.Marshal(tmp); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

// 获取订单的详细情况

func getOrderDetail(orderId int64) (bool, string) {
	var o entity.Order
	o.Id = orderId
	r := o.SelectById()

	type t struct {
		entity.Order
		DormName string `json:"dorm_name"`
	}
	var tmp t
	var dm bs.Dorm
	dm.Id = o.DormId
	dm.SelectById()
	tmp.DormName = dm.DormName
	tmp.Order = o

	if r {
		if bytes, err := json.Marshal(tmp); err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

// 修改订单的接收者
func setOrderRecv(orderId int64, stuId int64) bool {
	var o entity.Order
	o.Id = orderId
	r := o.SelectById()
	if r {
		o.RecvStu = stuId
		r = o.Update()

		//发送通知
		var u bs.Stu
		u.StuId = o.StuId
		u.SelectByStuId()
		wxapi.SendOrderNotify(u.Us.OpenId, o.TemplateId, u.Dm.DormName, orderId, o.Comment)
		return r
	}
	return false
}

// 取消订单
func cancelOrder(orderId int64, stuId int64) bool {
	var o entity.Order
	o.Id = orderId
	r := o.SelectById()
	if !r {
		return false
	}
	o.Cancel = true
	return o.Update()
}
