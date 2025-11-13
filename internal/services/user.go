package services

import (
	"GoGin/internal/model"
	"GoGin/internal/repository"
	"GoGin/internal/util/jwt_util"
	"errors"
	"strings"
)

type UserService struct {
	UserRepo repository.UserRepository
	jwtUtil  jwt_util.Util
}

func NewUserService(userRepo repository.UserRepository, jwtUtil jwt_util.Util) *UserService {
	return &UserService{
		UserRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *UserService) Register(req *model.RegisterRequest) (*model.User, error) {
	//密码时候否符合格式
	var flagPassword bool
	for i := 0; i < len(req.Password); i++ {
		if !((req.Password[i] >= 'a' && req.Password[i] <= 'z') || (req.Password[i] >= '0' && req.Password[i] <= '9') || (req.Password[i] >= 'A' && req.Password[i] <= 'Z')) {
			flagPassword = true
		}
	}
	if flagPassword {
		return nil, errors.New("password format Error")
	}

	//邮箱是否符合格式
	if !strings.Contains(req.Email, "@") {
		return nil, errors.New("email format Error")
	}

	//创建用户
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	//传入数据库
	if err := s.UserRepo.AddUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(loginKey, password string) (string, *model.User, error) {
	//判断是邮箱登录还是用户名登录
	var user *model.User
	var at, point bool
	for i := 0; i < len(loginKey); i++ {
		if loginKey[i] == '@' {
			at = true
		}
		if loginKey[i] == '.' {
			point = true
		}
	}
	if at && point { // 邮箱登录
		userByEmail, err := s.UserRepo.SelectByEmail(loginKey)
		if err != nil {
			return "", nil, err
		}
		user = userByEmail
	} else { // 用户名登录
		userByUsername, err := s.UserRepo.SelectByUsername(loginKey)
		if err != nil {
			return "", nil, err
		}
		user = userByUsername
	}

	//检查用户是否存在
	if flag := s.UserRepo.Exists(user.Username, user.Email); !flag {
		return "", nil, errors.New("username Not Exist")
	}

	//检验密码正确性
	if password != user.Password {
		return "", nil, errors.New("password error")
	}

	//返回token
	token, err := s.jwtUtil.GenerateToken(user.UserID, user.Username)
	if err != nil {
		return "", nil, errors.New("token Error")
	}

	return token, user, nil
}
