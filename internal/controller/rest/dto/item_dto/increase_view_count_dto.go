package item_dto

type IncreaseViewCountReq struct {
	Id int64 `json:"item_id" binding:"required"` // 조회 수를 증가시킬 글의 ID
}

type IncreaseViewCountRes struct {
	ItemViewCount int64 `json:"item_view_count"` // 증가된 글의 조회 수
}
