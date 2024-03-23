package postclientinterface

type PostServiceClient interface {
	GetUserId(post_id int64) (int64, error)
}
