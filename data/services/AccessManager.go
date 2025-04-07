package services

import (
	"citadel-api/data/model"
	"citadel-api/data/storage"
	"citadel-api/http/security"
	"citadel-api/utils/container"
	"errors"
)

type User struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name"`
	Role string `json:"role"`
}

const UserBlock = "user"
const RoleAdmin = "admin"
const RoleUser = "user"

func init() {
	container.Add[AccessManagerInterface](newAccessManager())
}

type AccessManagerInterface interface {
	Get(uuid string) (*model.Block, error)
	GetFromJWT(token string) (*User, error)
	Create(name string, role string) (string, error)
}

type accessManager struct {
	jwtManager      *security.JWTManager[string]
	blockRepository storage.BlockRepositoryInterface
}

func newAccessManager() AccessManagerInterface {
	return &accessManager{
		jwtManager:      security.NewUnlimitedJWTManager[string]([]byte("123456")),
		blockRepository: storage.NewBlockRepository(),
	}
}

func (a *accessManager) Get(uuid string) (*model.Block, error) {
	block, err := a.blockRepository.Get(uuid)
	if err != nil {
		return nil, err
	}
	if block.Type != UserBlock {
		return nil, nil
	}
	return block, nil
}

func (a *accessManager) GetFromJWT(token string) (*User, error) {
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
	if block == nil {
		return nil, errors.New("nil block")
	}
	if block.Type != UserBlock {
		return nil, nil
	}
	properties := block.Properties.(map[string]any)
	user := User{
		block.Id,
		properties["name"].(string),
		properties["role"].(string)}
	return &user, nil
}

func (a *accessManager) Create(name string, role string) (string, error) {
	block := &model.Block{
		Type:       UserBlock,
		Content:    &name,
		Properties: User{Name: name, Role: role},
	}
	_, err := a.blockRepository.Create(block)
	if err != nil {
		return "", err
	}

	return a.jwtManager.Generate(block.Id)

}
