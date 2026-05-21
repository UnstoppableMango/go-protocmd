package proto_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/anypb"

	cliv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/cli/v1alpha1"
	testingv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/testing/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/cli/proto"
)

var _ = Describe("Parser", func() {
	It("should work", func() {
		msg := &testingv1alpha1.Example_builder{
			Str:     new("test-string"),
			Integer: new(int32(420)),
			Opt:     new(true),
		}
		a, err := anypb.New(msg.Build())
		Expect(err).NotTo(HaveOccurred())

		u, err := proto.Parse(a)

		Expect(err).NotTo(HaveOccurred())
		Expect(u.GetOptions()).To(ConsistOf(
			(&cliv1alpha1.Option_builder{
				Name: new("opt"),
			}).Build(),
		))
		Expect(u.GetOperands()).To(ConsistOf("test-string"))
	})
})
