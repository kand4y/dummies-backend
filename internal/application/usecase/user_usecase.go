package usecase

import (
	"context"
	"dummies-backend/internal/application/repository"
	"dummies-backend/internal/domain/model"
	"errors"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) GetProfile(ctx context.Context, userID string) (*model.User, error) {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (uc *UserUseCase) UpdateProfile(ctx context.Context, userID, userHandle, userName string) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	user.UserHandle = userHandle
	user.UserName = userName
	return uc.userRepo.Update(ctx, user)
}

func (uc *UserUseCase) DeleteAccount(ctx context.Context, userID string) error {
	return uc.userRepo.SoftDelete(ctx, userID)
}
