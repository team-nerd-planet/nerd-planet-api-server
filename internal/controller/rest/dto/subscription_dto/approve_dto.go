package subscription_dto

type ApproveReq struct {
	Token string `form:"token" json:"token" validate:"required"`
}

type ApproveRes struct {
	Ok bool `json:"ok"` // 구독 인증 결과
}
