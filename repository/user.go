package repository

import (
	"backend-server/types"
	"backend-server/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		// userMap 초기화
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser)
	return nil
}

func (u *UserRepository) Update(userName string, newAge int64) error {
	isExisted := false

	for _, user := range u.userMap {
		if user.Name == userName {
			user.Age = newAge
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(userName string) error {
	isExisted := false

	for index, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...) // 해당값만 빼고 다시 할당

			isExisted = true
			continue
		}
	}

	if !isExisted {
		// 에러코드 작성
		//return errors.New("삭제오류!") // 별로 좋지 않음
		//return nil
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
