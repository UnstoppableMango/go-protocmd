package conv_test

import (
	_ "embed"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/conv"
)

//go:embed testdata/map.json
var mapJSON []byte

var _ = Describe("Args", func() {
	It("empty message", func() {
		msg := (&cmdv1alpha1.Process_builder{}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(BeEmpty())
	})

	It("string field", func() {
		msg := (&cmdv1alpha1.Process_builder{Path: new("/bin/ls")}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("path", "/bin/ls"))
	})

	It("bool field true emits flag name", func() {
		msg := (&cmdv1alpha1.Process_builder{Terminal: new(true)}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("terminal"))
	})

	It("bool field false emits nothing", func() {
		msg := (&cmdv1alpha1.Process_builder{Terminal: new(false)}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(BeEmpty())
	})

	It("int32 field", func() {
		msg := (&cmdv1alpha1.RunResponse_builder{ExitCode: proto.Int32(1)}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("exit_code", "1"))
	})

	It("uint32 field", func() {
		msg := (&cmdv1alpha1.User_builder{Uid: proto.Uint32(1000)}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("uid", "1000"))
	})

	It("enum field emits enum number", func() {
		msg := (&cmdv1alpha1.File_builder{Mode: cmdv1alpha1.OpenMode_OPEN_MODE_READ.Enum()}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("mode", "1"))
	})

	It("bytes field", func() {
		data := []byte("hello")
		msg := (&cmdv1alpha1.RunResponse_builder{Stdout: data}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("stdout", fmt.Sprint(data)))
	})

	It("nested message recurses into fields", func() {
		user := (&cmdv1alpha1.User_builder{Uid: proto.Uint32(1000)}).Build()
		msg := (&cmdv1alpha1.Process_builder{User: user}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("uid", "1000"))
	})

	It("multiple fields", func() {
		msg := (&cmdv1alpha1.User_builder{
			Uid: proto.Uint32(1000),
			Gid: proto.Uint32(1000),
		}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(ConsistOf("uid", "1000", "gid", "1000"))
	})

	It("repeated field emits name+value per element", func() {
		msg := (&cmdv1alpha1.Process_builder{Args: []string{"foo", "bar"}}).Build()
		args, err := conv.Args(msg.ProtoReflect())
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("args", "foo", "args", "bar"))
	})

	It("map field emits name+key+value per entry", func() {
		var fdp descriptorpb.FileDescriptorProto
		Expect(protojson.Unmarshal(mapJSON, &fdp)).To(Succeed())
		fd, err := protodesc.NewFile(&fdp, nil)
		Expect(err).NotTo(HaveOccurred())

		md := fd.Messages().ByName("TestMsg")
		msg := dynamicpb.NewMessage(md)
		fld := md.Fields().ByName("labels")
		msg.Mutable(fld).Map().Set(
			protoreflect.ValueOfString("mykey").MapKey(),
			protoreflect.ValueOfString("myval"),
		)

		args, err := conv.Args(msg)
		Expect(err).NotTo(HaveOccurred())
		Expect(args).To(HaveExactElements("labels", "mykey", "myval"))
	})
})
