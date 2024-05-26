package services

import (
	"fmt"

	"api_user_service_booking/config"
	pba "api_user_service_booking/genproto/attraction_proto"
	pbu "api_user_service_booking/genproto/user_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	AttractionService() pba.AttractionServiceClient
}

type serviceManager struct {
	userService       pbu.UserServiceClient
	attractionService pba.AttractionServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) AttractionService() pba.AttractionServiceClient {
	return s.attractionService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// dial to attraction service
	connAttraction, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.AttractionServiceHost, conf.AttractionServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:       pbu.NewUserServiceClient(connUser),
		attractionService: pba.NewAttractionServiceClient(connAttraction),
	}

	return serviceManager, nil
}
