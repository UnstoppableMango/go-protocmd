package codec

import (
	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
	"github.com/unstoppablemango/godec/proto"
)

var (
	FromJsonRequest   = proto.NewJson[*cmdv1alpha2.FromJsonRequest]()
	FromJsonResponse  = proto.NewJson[*cmdv1alpha2.FromJsonResponse]()
	FromYamlRequest   = proto.NewJson[*cmdv1alpha2.FromYamlRequest]()
	FromYamlResponse  = proto.NewJson[*cmdv1alpha2.FromYamlResponse]()
	FromProtoRequest  = proto.NewJson[*cmdv1alpha2.FromProtoRequest]()
	FromProtoResponse = proto.NewJson[*cmdv1alpha2.FromProtoResponse]()
)
