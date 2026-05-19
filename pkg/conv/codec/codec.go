package codec

import (
	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/godec/proto"
)

var (
	ArgsRequest  = proto.NewJson[*cmdv1alpha1.ArgsRequest]()
	ArgsResponse = proto.NewJson[*cmdv1alpha1.ArgsResponse]()
)
