package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
	"google.golang.org/grpc"

	"github.com/alecthomas/kong"
)

const ReadRemote = "127.0.0.1:4466"
const WriteRemote = "127.0.0.1:4467"

func WriteOrganization(layout OrgLayout) error {
	acls := []*acl.RelationTupleDelta{}

	for i := 0; i < layout.OrgOwners; i++ {
		acls = append(acls, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "organizations",
				Object:    layout.OrgName,
				Relation:  "owner",
				Subject: &acl.Subject{
					Ref: &acl.Subject_Id{
						Id: fmt.Sprintf("%sOrgOwner%d", layout.OrgName, i),
					},
				},
			},
		})
	}

	for i := 0; i < layout.OrgMembers; i++ {
		acls = append(acls, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "organizations",
				Object:    layout.OrgName,
				Relation:  "member",
				Subject: &acl.Subject{
					Ref: &acl.Subject_Id{
						Id: fmt.Sprintf("%sOrgMember%d", layout.OrgName, i),
					},
				},
			},
		})
	}

	for i := 0; i < layout.Groups; i++ {
		acls = append(acls, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "groups",
				Object:    fmt.Sprintf("%sGroup%d", layout.OrgName, i),
				Relation:  "owner",
				Subject: &acl.Subject{
					Ref: &acl.Subject_Set{
						Set: &acl.SubjectSet{
							Namespace: "organizations",
							Object:    layout.OrgName,
							Relation:  "owner",
						},
					},
				},
			},
		})

		for j := 0; j < layout.GroupMembers; j++ {
			acls = append(acls, &acl.RelationTupleDelta{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "groups",
					Object:    fmt.Sprintf("%sGroup%d", layout.OrgName, i),
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: fmt.Sprintf("%sGroupMember%d-%d", layout.OrgName, i, j),
						},
					},
				},
			})
		}
	}

	for i := 0; i < layout.IncidentTypes; i++ {
		acls = append(acls, &acl.RelationTupleDelta{
			Action: acl.RelationTupleDelta_INSERT,
			RelationTuple: &acl.RelationTuple{
				Namespace: "incidents",
				Object:    fmt.Sprintf("%sIncidentType%d", layout.OrgName, i),
				Relation:  "owner",
				Subject: &acl.Subject{
					Ref: &acl.Subject_Set{
						Set: &acl.SubjectSet{
							Namespace: "organizations",
							Object:    layout.OrgName,
							Relation:  "owner",
						},
					},
				},
			},
		})

		for j := 0; j < layout.Groups; j++ {
			acls = append(acls, &acl.RelationTupleDelta{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "incidents",
					Object:    fmt.Sprintf("%sIncidentType%d", layout.OrgName, i),
					Relation:  "member",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "groups",
								Object:    fmt.Sprintf("%sGroup%d", layout.OrgName, j),
								Relation:  "member",
							},
						},
					},
				},
			})
		}
	}

	wconn, err := grpc.DialContext(context.Background(), WriteRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	if err != nil {
		panic(err)
	}

	cli := acl.NewWriteServiceClient(wconn)
	_, err = cli.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: acls,
	})

	return err
}

func RunCheck(layout OrgLayout, rconn *grpc.ClientConn) (bool, error) {

	cli := acl.NewCheckServiceClient(rconn)

	group := rand.Intn(layout.Groups)
	member := rand.Intn(layout.GroupMembers)
	incidentType := rand.Intn(layout.IncidentTypes)

	start := time.Now()
	cr := &acl.CheckRequest{
		Subject: &acl.Subject{
			Ref: &acl.Subject_Id{Id: fmt.Sprintf("%sGroupMember%d-%d", layout.OrgName, group, member)},
		},
		Relation:  "member",
		Namespace: "incidents",
		Object:    fmt.Sprintf("%sIncidentType%d", layout.OrgName, incidentType),
	}
	resp, err := cli.Check(context.Background(), cr)
	if err != nil {
		return false, err
	}
	fmt.Printf("checking: %+v - resp: %v took :%s\n", cr, resp.Allowed, time.Since(start))
	return resp.Allowed, nil
}

type Context struct {
	Debug bool
}

type OrgLayout struct {
	OrgName       string
	OrgOwners     int
	OrgMembers    int
	Groups        int
	GroupMembers  int
	IncidentTypes int
}

func (r *OrgLayout) Run(concurrency int) error {
	// write
	fmt.Printf("in orglayout RUN: %+v\n", r)
	err := WriteOrganization(*r)
	if err != nil {
		return err
	}

	// concurrency
	rand.Seed(time.Now().Unix())

	wg := sync.WaitGroup{}
	fmt.Println("concurrency", concurrency)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		rconn, err := grpc.DialContext(context.Background(), ReadRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
		if err != nil {
			panic(err)
		}

		fmt.Println("connection created ", i)

		go func(rconn *grpc.ClientConn) {
			defer wg.Done()

			for {
				allowed, err := RunCheck(*r, rconn)
				if err != nil {
					panic(err)
				}
				if !allowed {
					panic("now allowed when we should be!!!")
				}
			}
		}(rconn)
	}

	wg.Wait()

	return nil
}

func main() {
	var cli struct {
		Concurrency int `help:"how many"`

		Ol OrgLayout `cmd:"" help:"orglayout"`
	}

	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(cli.Concurrency)
	ctx.FatalIfErrorf(err)
}

// func main() {

// 	rconn, err := grpc.DialContext(context.Background(), ReadRemote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("this will run until something breaks (ctrl-c to exit): enter to continue")
// 	fmt.Scan()

// 	// do this until things break.
// 	for i := 1; ; i++ {
// 		start := time.Now()
// 		orgname := fmt.Sprintf("org%d", i)
// 		cli := acl.NewWriteServiceClient(wconn)
// 		err = WriteOrg(orgname, cli)
// 		if err != nil {
// 			panic(err)
// 		}

// 		writeDelta := time.Since(start)

// 		read := time.Now()
// 		allowed, err := RunCheck(orgname, acl.NewCheckServiceClient(rconn))
// 		if err != nil {
// 			panic(err)
// 		}
// 		if !allowed {
// 			panic("incorrect response")
// 		}
// 		readDelta := time.Since(read)
// 		fmt.Printf("orgcreated: %s\t writeDelta: %s, readDelat: %s\n", orgname, writeDelta, readDelta)
// 	}

// }
