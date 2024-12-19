package api

type LeaveAccount struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type LeavePeriodLog struct {
	Id      int    `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Comment string `json:"comment,omitempty"`
	Status  int    `json:"status"`
}

type LeavePeriod struct {
	Id                 int              `json:"id,omitempty"`
	OwnerId            int              `json:"ownerId,omitempty"`
	IsConfirmed        bool             `json:"isConfirmed,omitempty"`
	ConfirmationDate   string           `json:"confirmationDate,omitempty"`
	AttachmentId       string           `json:"attachmentId,omitempty"`
	Logs               []LeavePeriodLog `json:"logs,omitempty"`
	Value              string           `json:"value,omitempty"`
	CreationDate       string           `json:"creationDate,omitempty"`
	IsActive           bool             `json:"isActive,omitempty"`
	CancellationDate   string           `json:"cancellationDate,omitempty"`
	CancellationUserId string           `json:"cancellationUserId,omitempty"`
	Comment            string           `json:"comment,omitempty"`
}

type Leave struct {
	Id             int           `json:"id,omitempty"`
	Date           string        `json:"name,omitempty"`
	IsAm           bool          `json:"isAm,omitempty"`
	LeaveAccountId int           `json:"leaveAccountId,omitempty"`
	LeaveAccount   *LeaveAccount `json:"leaveAccount,omitempty"`
	LeavePeriodId  int           `json:"leavePeriodId,omitempty"`
	LeavePeriod    *LeavePeriod  `json:"leavePeriod,omitempty"`
	// ...
}

type ListLeavesRequest struct {
	Date               string `json:"date,omitempty"`
	LeavePeriodOwnerId int    `json:"leavePeriod.ownerId,omitempty"`
	Paging             *Page  `json:"paging,omitempty"`
	LeaveAccountId     int    `json:"leaveAccountId,omitempty"`
}

// type ListUsersResponse ListResponse[User]

// func (c *Client) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
// 	res := &ListUsersResponse{}
// 	err := c.Get(ctx, "/users", req, res)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// type GetUserRequest struct {
// 	UserId int      `json:"-"`
// 	Fields []string `json:"fields,omitempty"`
// }

// type GetUserResponse GetResponse[User]

// func (c *Client) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
// 	res := &GetUserResponse{}
// 	err := c.Get(ctx, fmt.Sprintf("/users/%d", req.UserId), req, res)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
