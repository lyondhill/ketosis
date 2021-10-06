package main

import (
	"context"
	"fmt"
	"time"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
	"google.golang.org/grpc"
)

const ReadRemote = "127.0.0.1:4466"
const WriteRemote = "127.0.0.1:4467"

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
							Id: orgName + "orgPerson",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "organizations",
					Object:    orgName,
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "orgPerson2",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "organizations",
					Object:    orgName,
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "organizations",
								Object:    orgName,
								Relation:  "owner",
							},
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "organizations",
					Object:    orgName,
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "orgperson3",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group1",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "organizations",
								Object:    orgName,
								Relation:  "member",
							},
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group1",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "groupOwner1",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group1",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "groupOwner2",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group1",
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "groupMember1",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group2",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "organizations",
								Object:    orgName,
								Relation:  "member",
							},
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group2",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "groupOwner3",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group2",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "groupOwner4",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    orgName + "group2",
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: orgName + "groupMember2",
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "incidents",
					Object:    orgName + "incidnetType1",
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "organizations",
								Object:    orgName,
								Relation:  "member",
							},
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "incidents",
					Object:    orgName + "incidnetType1",
					Relation:  "owner",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "groups",
								Object:    orgName + "group1",
								Relation:  "member",
							},
						},
					},
				},
			},
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "incidents",
					Object:    orgName + "incidnetType1",
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "groups",
								Object:    orgName + "group2",
								Relation:  "member",
							},
						},
					},
				},
			},
		},
	})

	return err
}

func RunCheck(orgName string, cli acl.CheckServiceClient) (bool, error) {
	resp, err := cli.Check(context.Background(), &acl.CheckRequest{
		Subject: &acl.Subject{
			Ref: &acl.Subject_Id{Id: orgName + "groupMember1"},
		},
		Relation:  "owner",
		Namespace: "incidents",
		Object:    orgName + "incidnetType1",
	})
	if err != nil {
		return false, err
	}
	return resp.Allowed, nil
}

func main() {
	// write a typical organization
	wconn, err := grpc.DialContext(context.Background(), WriteRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	if err != nil {
		panic(err)
	}

	rconn, err := grpc.DialContext(context.Background(), ReadRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	if err != nil {
		panic(err)
	}

	fmt.Printf("this will run until something breaks (ctrl-c to exit): enter to continue")
	fmt.Scan()

	// do this until things break.
	for i := 1; ; i++ {
		start := time.Now()
		orgname := fmt.Sprintf("org%d", i)
		cli := acl.NewWriteServiceClient(wconn)
		err = WriteOrg(orgname, cli)
		if err != nil {
			panic(err)
		}

		writeDelta := time.Since(start)

		read := time.Now()
		allowed, err := RunCheck(orgname, acl.NewCheckServiceClient(rconn))
		if err != nil {
			panic(err)
		}
		if !allowed {
			panic("incorrect response")
		}
		readDelta := time.Since(read)
		fmt.Printf("orgcreated: %s\t writeDelta: %s, readDelat: %s\n", orgname, writeDelta, readDelta)
	}

}
