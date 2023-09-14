package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"gin/global/email"
	"gin/global/errors"
	"gin/global/jwt"
	"gin/global/password"
	"gin/global/validator"
	"gin/internal/domain/entity"
	"gin/internal/domain/repository"
	"gin/internal/domain/service"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gofrs/uuid"
)

type UserService struct {
	Repository repository.UserRepositoryImpl
}

func NewUserService(ur repository.UserRepositoryImpl) service.UserServiceImpl {
	return &UserService{
		Repository: ur,
	}
}

func (us *UserService) Register(req *entity.UserRegister, ctx context.Context) (*entity.User, error) {
	err := validateRequestRegister(req)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	uuid := uuid.Must(uuid.NewV4())
	user := &entity.User{
		ID:       uuid,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Username: req.Username,
		Phone:    req.Phone,
	}

	user, err = us.Repository.Create(user, ctx)
	if err != nil {
		return nil, err
	}

	go sendEmail(req, uuid)

	return user, nil
}

func validateRequestRegister(req *entity.UserRegister) error {
	switch {
	case req.Name == "":
		return errors.ErrNameRequired
	case req.Email == "" || !validator.ValidateEmail(req.Email):
		return errors.ErrInvalidEmail
	case req.Password == "" || !validator.ValidatePassword(req.Password):
		return errors.ErrInvalidPassword
	case req.Username == "":
		return errors.ErrUsernameRequired
	case req.Phone == "" || !validator.ValidatePhone(req.Phone):
		return errors.ErrInvalidPhoneNumber
	}
	return nil
}

func sendEmail(req *entity.UserRegister, uuid uuid.UUID) error {
	mailer := email.NewMailClient()
	mailer.SetSubject("Email Verification")
	mailer.SetReciever(req.Email)
	mailer.SetSender(os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetBodyHTML(req.Username, fmt.Sprintf("%s/%s", os.Getenv("URL_VERIFY"), uuid.String()))
	if err := mailer.SendMail(); err != nil {
		return err
	}
	return nil
}

func (us *UserService) VerifyAccount(id uuid.UUID, ctx context.Context) (*entity.User, error) {
	user, err := us.Repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}

	if user.Status {
		return nil, errors.ErrAccountAlreadyVerified
	}

	user.Status = true
	user, err = us.Repository.Update(user, ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Login(req *entity.UserLogin, ctx context.Context) (*entity.ResponseLogin, error) {
	var res *entity.ResponseLogin

	user, err := us.Repository.FindByEmail(req.Email, ctx)
	if err != nil {
		return res, err
	}

	if !user.Status {
		return res, errors.ErrAccountNotVerified
	}

	if err := password.ComparePassword(req.Password, user.Password); err != nil {
		return res, errors.ErrInvalidPassword
	}

	jwt, err := jwt.EncodeToken(user)
	res = &entity.ResponseLogin{
		User:  user,
		Token: jwt,
	}
	return res, err
}

func (us *UserService) UploadPhoto(file *multipart.FileHeader, id uuid.UUID, ctx context.Context) (*entity.User, error) {
	user, err := us.Repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}

	name := user.Name + "-" + user.ID.String()
	link, err := uploadPhoto(file, name)
	if err != nil {
		return nil, err
	}

	user.Photo = link
	user, err = us.Repository.Update(user, ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func uploadPhoto(file *multipart.FileHeader, name string) (string, error) {
	supClient := supabasestorageuploader.New(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_KEY"),
		os.Getenv("SUPABASE_BUCKET"),
	)

	link, err := supClient.Upload(file)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (us *UserService) UpdateUser(req *entity.UserUpdate, ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := us.Repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}

	user, err = validateRequestUpdate(req, user)
	if err != nil {
		return nil, err
	}

	user, err = us.Repository.Update(user, ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validateRequestUpdate(req *entity.UserUpdate, user *entity.User) (*entity.User, error) {
	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" {
		if !validator.ValidateEmail(req.Email) {
			return nil, errors.ErrInvalidEmail
		}
		user.Email = req.Email
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.Phone != "" {
		if !validator.ValidatePhone(req.Phone) {
			return nil, errors.ErrInvalidPhoneNumber
		}
		user.Phone = req.Phone
	}

	if req.Password != "" {
		if !validator.ValidatePassword(req.Password) {
			return nil, errors.ErrInvalidPassword
		}
		hashedPassword, err := password.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	return user, nil
}
