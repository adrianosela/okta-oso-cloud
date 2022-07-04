package main

import (
	"fmt"
	oso "github.com/osohq/go-oso-cloud"
	"log"
	"os"
)

type OktaUser struct{ id string }
type OktaGroup struct{ id string }
type FeatureFlag struct{ id string }

func (ou OktaUser) Id() string    { return ou.id }
func (og OktaGroup) Id() string   { return og.id }
func (ff FeatureFlag) Id() string { return ff.id }

func (ou OktaUser) Type() string    { return "OktaUser" }
func (og OktaGroup) Type() string   { return "OktaGroup" }
func (ff FeatureFlag) Type() string { return "FeatureFlag" }

func groupListToBulkFact(userID string, groupIDs ...string) []oso.BulkFact {
	bf := []oso.BulkFact{}
	for _, groupID := range groupIDs {
		bf = append(bf, oso.BulkFact{
			Predicate: "has_group",
			Args:      []oso.Instance{OktaUser{id: userID}, OktaGroup{id: groupID}},
		})
	}
	return bf
}

func main() {
	// mock data
	oktaUserID := "larry"
	oktaGroupIDs := []string{"Infrastructure Engineering", "Everyone", "Internal Tools"}

	osoClient := oso.NewClient("https://cloud.osohq.com", os.Getenv("OSO_AUTH"))

	allowed, err := osoClient.AuthorizeWithContext(OktaUser{id: oktaUserID}, "read", FeatureFlag{id: "1"}, groupListToBulkFact(oktaUserID, oktaGroupIDs...))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(allowed)
}
