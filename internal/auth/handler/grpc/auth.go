package grpc

import auth "chat/api/v1"

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
}
