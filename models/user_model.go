package models

type User struct {
	Id        uint32 `json:"id"`
	NickName  string `json:"nickname" binding:"required"`
	Create_at uint64 `json:"create_at"`
	Update_at uint64 `json:"update_at"`
}
