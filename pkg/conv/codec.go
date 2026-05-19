package conv

import (
	"io"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/conv/codec"
)

func DecodeArgs(r io.Reader) (*cmdv1alpha1.ArgsRequest, error) {
	var req cmdv1alpha1.ArgsRequest
	d := codec.ArgsRequest.NewDecoder(r)
	if err := d.Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
