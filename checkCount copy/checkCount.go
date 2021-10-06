package main

import (
	"context"
	"fmt"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
	"google.golang.org/grpc"
)

const ReadRemote = "127.0.0.1:4466"
const WriteRemote = "127.0.0.1:4467"

func RunCheck(cli acl.CheckServiceClient) (bool, error) {
	resp, err := cli.Check(context.Background(), &acl.CheckRequest{
		Subject: &acl.Subject{
			Ref: &acl.Subject_Id{Id: "1"},
		},
		Relation:  "owner",
		Namespace: "organizations",
		Object:    "lyonsorg",
	})
	if err != nil {
		return false, err
	}
	return resp.Allowed, nil
}

func WriteOrg(orgName string, cli acl.WriteServiceClient) error {
	_, err := cli.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "organizations",
					Object:    orgName,
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: "1",
						},
					},
				},
			},
		},
	})

	return err
}

func main() {
	// write a typical organization
	wconn, err := grpc.DialContext(context.Background(), WriteRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	if err != nil {
		panic(err)
	}

	cli := acl.NewWriteServiceClient(wconn)
	err = WriteOrg("lyonsorg", cli)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.DialContext(context.Background(), ReadRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	if err != nil {
		panic(err)
	}

	allowed, err := RunCheck(acl.NewCheckServiceClient(conn))

	if err != nil {
		panic(err)
	}
	fmt.Println("allowed:", allowed)
}
