package services

import (
	"citadel-api/data/model"
	"citadel-api/data/storage"
	"citadel-api/http/security"
	"errors"
)

const UserBlock = "user"
const RoleAdmin = "admin"
const RoleUser = "user"

type AccessManager struct {
	jwtManager      *security.JWTManager[string]
	blockRepository storage.BlockRepositoryInterface
}

func NewAccessManager() *AccessManager {
	return &AccessManager{
		jwtManager:      security.NewUnlimitedJWTManager[string]([]byte("123456")),
		blockRepository: storage.NewBlockRepository(),
	}
}

func (a *AccessManager) Get(uuid string) (*model.Block, error) {
	block, err := a.blockRepository.Get(uuid)
	if err != nil {
		return nil, err
	}
	if block.Type != UserBlock {
		return nil, nil
	}
	return block, nil
}

func (a *AccessManager) GetFromJWT(token string) (*model.Block, error) {
	id, err := a.jwtManager.Decode(token)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, errors.New("nil token")
	}
	block, err := a.blockRepository.Get(*id)
	if err != nil {
		return nil, err
	}
	if block.Type != UserBlock {
		return nil, nil
	}
	return block, nil
}

func (a *AccessManager) Create(name string, role string) (string, error) {
	block := &model.Block{
		Type:       UserBlock,
		Content:    &name,
		Properties: map[string]any{"role": role, "name": name},
	}
	_, err := a.blockRepository.Create(block)
	if err != nil {
		return "", err
	}

	return a.jwtManager.Generate(block.Id)

}
