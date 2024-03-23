package postclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/sgokul961/echo-hub-notification-svc/pkg/config"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/pb"
	postclientinterface "github.com/sgokul961/echo-hub-notification-svc/pkg/postClient/postClientInterface"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PostServiceCLient struct {
	Client pb.PostServiceClient
}

func NewPostServiceClient(c config.Config) postclientinterface.PostServiceClient {
	cc, err := grpc.Dial(c.PostHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}
	return &PostServiceCLient{Client: pb.NewPostServiceClient(cc)}

}
func (p *PostServiceCLient) GetUserId(post_id int64) (int64, error) {
	req := &pb.GetUserIdRequest{
		PostId: post_id,
	}
	res, err := p.Client.GetUserId(context.Background(), req)
	if err != nil {
		return res.UserId, errors.New("error in getting post service method")
	}
	return res.UserId, nil

}
