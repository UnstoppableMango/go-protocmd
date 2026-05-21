package cli

import (
	"fmt"

	cliv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/cli/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/cli/proto"
)

func Parse(req *cliv1alpha1.ParseRequest) (*cliv1alpha1.ParseResponse, error) {
	var err error
	res := &cliv1alpha1.ParseResponse_builder{}
	switch in := req.WhichInput(); in {
	case cliv1alpha1.ParseRequest_Proto_case:
		res.Utility, err = proto.Parse(req.GetProto())
	default:
		return nil, fmt.Errorf("unsupported input: %s", in)
	}

	if err != nil {
		return nil, err
	}
	return res.Build(), nil
}
