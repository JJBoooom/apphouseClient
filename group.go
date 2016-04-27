package client

type UserGroup struct {
	Id         int    `json:"_id"`
	GroupName  string `json:"group_name"`
	Namespace  string `json:"namespace"`
	CreateTime int64  `json:"create_time"`
	Desc       string `json:"desc"`
}

type UserGroupCollection struct {
	Groups []UserGroup `json:"group"`
}

type UserGroupOperations interface {
	List(info interface{}) (*UserGroupCollection, error)
	Create(info interface{}) error
	Delete(info interface{}) error
	Info(tag interface{}) (*UserGroup, error)
	Update(existirng *UserGroup, updates interface{}) error
}

type UserGroupClient struct {
	apphouseClient *ApphouseClient
}

func newUserGroupClient(apphouseClient *ApphouseClient) *UserGroupClient {
	return &UserGroupClient{
		apphouseClient: apphouseClient,
	}
}

func (c *UserGroupClient) List(info interface{}) (*UserGroupCollection, error) {
	return &UserGroupCollection{}, nil
}

func (c *UserGroupClient) Create(info interface{}) error {
	return nil
}

func (c *UserGroupClient) Delete(info interface{}) error {
	return nil
}

func (c *UserGroupClient) Info(info interface{}) (*UserGroup, error) {
	return &UserGroup{}, nil
}

func (c *UserGroupClient) Update(existing *UserGroup, info interface{}) error {
	return nil
}
