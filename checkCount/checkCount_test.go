package main

import (
	"context"
	"testing"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
	"google.golang.org/grpc"
)

func BenchmarkRunCheck(b *testing.B) {
	conn, err := grpc.DialContext(context.Background(), ReadRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	if err != nil {
		panic(err)
	}

	cli := acl.NewCheckServiceClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RunCheck(cli)
	}
}
