package file

import (
	f "github.com/micro-blonde/file"
	pf "github.com/micro-blonde/file/proto/file"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetFileProto[T f.Model](f *f.File[T]) *pf.File {
	r := &pf.File{
		Id:        f.Id,
		Key:       f.Key,
		CreatedAt: f.CreatedAt.UnixMicro(),
		Name:      f.Name,
		Type:      f.Type,
		Path:      f.Path,
	}
	if f.ExpirationTime != nil {
		r.ExpirationTime = timestamppb.New(*f.ExpirationTime)
	}
	if f.AccountId != nil {
		r.AccountId = *f.AccountId
	}
	if f.Meta != nil {
		r.Meta = &pf.Meta{}
	}
	return r
}
