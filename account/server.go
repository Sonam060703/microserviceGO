//go:generate protoc ./account.proto --go_out=plugins=grpc:./pb
package account

import (
	"context"
	"fmt"
	"net"

	"github.com/Sonam060703/microserviceGO/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*                Req
  	 	CLIENT   ----->   SERVER
                  Res
		CLIENT   <-----   SERVER

*/

// server interect with both service and client making grpc calls
type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	// Instance of Service
	service Service
}

// Invoked in main.go
func ListenGRPC(s Service, port int) error {
	// Listen function of net package to establish tcp connection
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	// It will create a new grpc server serv
	serv := grpc.NewServer()
	// Register the gRPC server with the Service (Account Service)
	pb.RegisterAccountServiceServer(serv, &grpcServer{
		service: s,
	})
	// Enable gRPC reflection for debugging tools like Evans/Postman
	reflection.Register(serv)
	// Start serving and return any error
	return serv.Serve(lis)
}

// pb -> protobuff (pb. things comes from the account.proto ) , s -> Receiver
// Accepting the request
// return the response that Server get from Service after modifying
func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	// invoking Service PostAccount fun to create Account
	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   a.ID,
		Name: a.Name,
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	a, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(
			accounts,
			&pb.Account{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
