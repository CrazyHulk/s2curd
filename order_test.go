package main

import (
	"context"
	"time"
)

// Order 支付订单
type Order struct {
	ID             int64     // 订单号
	RootOrderID    int64     // 母单号，自动扣费订单对应的签约订单号
	UID            int64     // 用户 ID
	PayType        int32     // 支付类型，1、普通；2、IAP
	PayAmount      int32     // 实际支付金额
	PayMsgContent  string    // 支付返回的额外信息，json格式字符串，{"payCounponAmount":0,"payBpAmount":9}
	OriginalAmount int32     // 原始支付金额
	Platform       int32     // 平台
	ProductID      int32     // 商品标识
	ProductAmount  int32     // 商品总数量（包含奖励数量）
	PayStatus      int32     // 支付状态：1、支付成功
	PayTime        time.Time // 支付时间
	CreateTime     time.Time // 创建时间
	MTime          time.Time
	ExpireTime     time.Time // 过期时间
	ExtInfo        ExtInfo   // 业务额外信息

	// 以下信息均为支付平台回调数据，用于统计和对账
	TxID          string // 支付平台订单号
	PayDeviceType int32  // 支付设备渠道类型，1 pc 2 webapp 3 app 4 jsapi 5 server 6 小程序支付 7 聚合二维码支付
	PayChannel    int32  // 支付渠道，alipay、wechat, iap、花呗

	version int32 // 当前 row 的版本号，注意业务不要写该字段
}

type ExtInfo struct {
	IsFirst             bool   `json:"is_first,omitempty"`               // 首次购买该档位标记
	IsRoot              bool   `json:"is_root,omitempty"`                // 签约订单标记
	ContractNo          string `json:"contract_no,omitempty"`            // 签约合同号
	IsWithhold          bool   `json:"is_withhold,omitempty"`            // 是否代扣
	PayFailReason       string `json:"fail_reason,omitempty"`            // 支付失败原因
	CannelReason        string `json:"cannel_reason,omitempty"`          // 取消原因
	BizFlag             int32  `json:"biz_flag,omitempty"`               // 业务
	IsMainSite          bool   `json:"is_main_site,omitempty"`           // 是否来自主站
	WithholdProductID   int32  `json:"withhold_product_id,omitempty"`    // 代扣的product_id
	RefundLevel         int32  `json:"refund_level,omitempty"`           // 退款等级 0 普通 1 普通恶意
	RefundAmount        int32  `json:"refund_amount,omitempty"`          // 回收的数量(已发放)
	RefundUsedAmount    int32  `json:"refund_used_amount,omitempty"`     // 已使用的数量
	RefundNotSendAmount int32  `json:"refund_not_send_amount,omitempty"` // 回收的数量(未发放)
	RefundCardSetType   int32  `json:"refund_card_set_type,omitempty"`   // 卡面类型
	BuySendAmount       int32  `json:"BuySendAmount,omitempty"`          // 买赠份数
	RecieverUID         int64  `json:"RecieverUID,omitempty"`            // 收卡人UID
	ExtraAmount         int32  `json:"extra_amount,omitempty"`           // 额外赠送数量
	ActID               int32  `json:"act_id,omitempty"`                 // 活动ID

	PaySuccessTime time.Time `json:"pay_success_time,omitempty"` // 支付回调时间
	RefundTime     time.Time `json:"refund_time,omitempty"`      // 退款回调时间

	IsRefuel       bool      `json:"is_refuel,omitempty"`        // 是否是加油包
	RefuelSendTime time.Time `json:"refuel_send_time,omitempty"` // 是否是加油包
	RefuelDeadline time.Time `json:"refuel_deadline,omitempty"`  // 加油包结束时间
	RefuelAmount   int32     `json:"refuel_amount,omitempty"`    // 加油包数量
	RefuelSilverID int32     `json:"refuel_silver_id,omitempty"` // 加油包附加的通用券 id
	RefuelCardID   int32     `json:"refuel_card_id,omitempty"`   // 加油包附加的
}
