package users_lnk

type CreateGroupUserLnkParams struct {
	GuUserID  int32  `json:"user_id"`
	GuGroupID int32  `json:"group_id"`
	GName     string `json:"group_name"`
}

type CreateGroupUserLnkParamsRes struct {
	GuUserID  int32 `json:"user_id"`
	GuGroupID int32 `json:"group_id"`
}

type GroupsUsersLnk struct {
	GuID      int32 `json:"id"`
	GuUserID  int32 `json:"user_id"`
	GuGroupID int32 `json:"group_id"`
}
