actor OktaUser {}
actor OktaGroup {}

has_role(entity: OktaUser, role: String, resource: Resource) if
    group matches OktaGroup and
    has_group(entity, group) and
    has_role(group, role, resource);

resource FeatureFlag {
	permissions = ["read", "toggle", "delete"];
	roles = ["contributor", "maintainer", "admin"];

	"read" if "contributor";
	"toggle" if "maintainer";
	"delete" if "admin";

	"maintainer" if "admin";
	"contributor" if "maintainer";
}


## Example CLI calls

## Add data
# oso-cloud tell has_role User:larry admin FeatureFlag:0
# oso-cloud tell has_role OktaGroup:Engineering contributor FeatureFlag:2
# oso-cloud tell has_role OktaGroup:Engineering contributor FeatureFlag:2

## Check it works
# oso-cloud authorize OktaUser:larry read FeatureFlag:2
# oso-cloud list OktaUser:larry read FeatureFlag

## Coming soon: context
# oso-cloud tell has_role 'OktaGroup:Infrastructure Engineering' maintainer FeatureFlag:1
# oso-cloud authorize OktaUser:larry read FeatureFlag:1 -c 'has_group OktaUser:larry OktaGroup:InfrastructureEngineering'
