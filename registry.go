package internalcmd

import (
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"reflect"
)

var globalRegistry *registry

func init() {
	globalRegistry = &registry{
		uri2Type: make(map[string]reflect.Type),
		type2URI: make(map[reflect.Type]string),
		creators: make(map[string]func() interface{}),
	}
	globalRegistry.Register(
		&internal_command.CmdCheckup{},
	)
}

type registry struct {
	uri2Type map[string]reflect.Type
	type2URI map[reflect.Type]string
	creators map[string]func() interface{}
}

func (r *registry) Register(mpList ...interface{}) {
	for _, mp := range mpList {
		uri, rt := mustURIResolve(mp)
		r.uri2Type[uri] = rt
		r.type2URI[rt] = uri
		r.creators[uri] = func() interface{} { return reflect.New(rt.Elem()).Interface() }
	}
}

func (r *registry) NewObject(uri string) (obj interface{}, created bool) {
	if cf, ok := r.creators[uri]; ok {
		obj = cf()
		created = true
	}
	return
}

func mustURIResolve(pkt interface{}) (uri string, pType reflect.Type) {
	pType = reflect.TypeOf(pkt)
	if protoMsg, ok := pkt.(proto.Message); ok {
		// use proto message name if is a proto message
		uri = proto.MessageName(protoMsg)
	} else {
		// otherwise just get type name
		uri = pType.Elem().Name()
	}

	return
}
