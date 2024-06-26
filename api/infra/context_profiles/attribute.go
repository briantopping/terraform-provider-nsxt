//nolint:revive
package contextprofiles

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra/context_profiles"
	model1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/context_profiles"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	client2 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/orgs/projects/infra/context_profiles"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

type AttributeClientContext utl.ClientContext

func NewAttributesClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *AttributeClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewAttributesClient(connector)

	case utl.Global:
		client = client1.NewAttributesClient(connector)

	case utl.Multitenancy:
		client = client2.NewAttributesClient(connector)

	default:
		return nil
	}
	return &AttributeClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID, VPCID: sessionContext.VPCID}
}

func (c AttributeClientContext) List(attributeKeyParam *string, attributeSourceParam *string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model0.PolicyContextProfileListResult, error) {
	var err error
	var obj model0.PolicyContextProfileListResult

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.AttributesClient)
		obj, err = client.List(attributeKeyParam, attributeSourceParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	case utl.Global:
		client := c.Client.(client1.AttributesClient)
		gmObj, err := client.List(attributeKeyParam, attributeSourceParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.PolicyContextProfileListResultBindingType(), model0.PolicyContextProfileListResultBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.PolicyContextProfileListResult)

	case utl.Multitenancy:
		client := c.Client.(client2.AttributesClient)
		obj, err = client.List(utl.DefaultOrgID, c.ProjectID, attributeKeyParam, attributeSourceParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}
