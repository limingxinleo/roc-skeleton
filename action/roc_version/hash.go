package roc_version

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/serializer"
)

type Hash struct {
}

type HashRequest struct {
	User UserDTO
}

func (h *HashRequest) UnmarshalJSON(bytes []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	if err := json.Unmarshal(raw[0], &h.User); err != nil {
		return err
	}
	return nil
}

type UserDTO struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type HashDTO struct {
	Version string `json:"version"`
}

func (h *Hash) Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface) {
	//id := packet.GetId()
	body := packet.GetBody()

	req := &formatter.JsonRPCRequest[*HashRequest, any]{}

	if err := serializer.UnSerialize(body, req); err != nil {
		return nil, exception.NewDefaultException(err.Error())
	}

	user := req.Data.User

	return &HashDTO{Version: h.toHash(&user)}, nil
}

func (h *Hash) toHash(user *UserDTO) string {
	data := []byte(string(user.Id) + user.Name + user.Email)
	hash := md5.Sum(data)
	return "v1.0.0@" + hex.EncodeToString(hash[:])
}
