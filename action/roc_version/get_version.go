package roc_version

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/serializer"
	"os"
)

type GetVersion struct {
}

func (v *GetVersion) Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface) {
	return os.Getenv("APP_VERSION"), nil
}
