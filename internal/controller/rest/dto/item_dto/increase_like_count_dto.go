package item_dto

type IncreaseLikeCountReq struct {
	Id int64 `json:"item_id" binding:"required"` // 좋아요 수를 증가시킬 글의 ID
}

type IncreaseLikeCountRes struct {
	ItemLikeCount int64 `json:"item_like_count"` // 증가된 글의 종아요 수
}
