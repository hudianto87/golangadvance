package middleware

import (
	"belajargolangpart2/session11-user-crud-grpc-gateway-cache/config"
	"context"
	"encoding/base64"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		publicMethods := []string{
			"/proto.user_service.v1.UserService/GetUsers",
			"/proto.user_service.v1.UserService/GetUserByID",
		}

		for _, method := range publicMethods {
			if info.FullMethod == method {
				return handler(ctx, req)
			}
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Meta data not provide")
		}

		authHeader, ok := md["authorization"]
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization header is missing")
		}

		if !strings.HasPrefix(authHeader[0], "Basic ") {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid authoritasion header scheme")
		}

		decoded, err := base64.StdEncoding.DecodeString(authHeader[0][6:])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
		}

		creds := strings.SplitN(string(decoded), ":", 2)
		if len(creds) != 2 {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
		}

		username, password := creds[0], creds[1]

		if username != config.AuthBasicUsername || password != config.AuthBasicPassword {
			return nil, status.Errorf(codes.Unauthenticated, "invalid user and password")
		}

		return handler(ctx, req)
	}
}
