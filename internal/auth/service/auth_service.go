package service

import (
	auth "chat/api/v1"
	"chat/internal/auth/model"
	"chat/internal/auth/repository"
	myemail "chat/internal/auth/service/email"
	"chat/internal/auth/service/password"
	"context"
	"fmt"
	"math/rand"
	"strconv"
)

type AuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
}

func (*AuthServiceServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	var userrepo repository.UserRepository
	user, err := userrepo.FindByUsername(ctx, in.Username)
	if err != nil {
		return &auth.LoginResponse{Err: "该用户名尚未注册"}, err
	}
	var hashhandler password.Argon2Hasher
	if hashhandler.Compare(user.PasswordHash, in.Password) == false {
		return &auth.LoginResponse{
			Err: "用户名或密码错误",
		}, fmt.Errorf("wrong password or username")
	}
	return &auth.LoginResponse{
		Err: "",
	}, nil
}

func (*AuthServiceServer) Register(ctx context.Context, in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	code := rand.Intn(900000) + 100000
	fmt.Println(code)
	op := model.EmailOptions{
		MailHost: "smtp.163.com",
		MailPort: 25,
		MailUser: "lhmchat@163.com",
		MailPass: "Zhang.0803",
		MailTo:   in.Email,
		Subject:  "lhmchat验证码",
		Body:     "你的验证码是:" + strconv.Itoa(code),
	}
	if ok, err := repository.Savekv(ctx, "verifycode:send:"+in.Email, 60, 1); !ok {
		return &auth.RegisterResponse{Err: "请求频繁，请稍后再试"}, err
	}
	if ok, err := repository.Savekv(ctx, "verifycode:data:"+in.Email, 300, code); !ok {
		return &auth.RegisterResponse{Err: "请求频繁，请稍后再试"}, err
	}
	if err := myemail.SendMail(&op); err != nil {
		return &auth.RegisterResponse{Err: "未知错误，请稍后重试"}, err
	}
	return &auth.RegisterResponse{Err: ""}, nil
}
