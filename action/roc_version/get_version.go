package roc_version

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/serializer"
)

type GetVersion struct {
}

func (v *GetVersion) Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface) {
	return "v1.0.0", nil
}
