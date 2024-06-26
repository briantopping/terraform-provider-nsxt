//nolint:revive
package infra

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra"
	model1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	client2 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/orgs/projects/infra"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

type GatewayQosProfileClientContext utl.ClientContext

func NewGatewayQosProfilesClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *GatewayQosProfileClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewGatewayQosProfilesClient(connector)

	case utl.Global:
		client = client1.NewGatewayQosProfilesClient(connector)

	case utl.Multitenancy:
		client = client2.NewGatewayQosProfilesClient(connector)

	default:
		return nil
	}
	return &GatewayQosProfileClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID, VPCID: sessionContext.VPCID}
}

func (c GatewayQosProfileClientContext) Get(qosProfileIdParam string) (model0.GatewayQosProfile, error) {
	var obj model0.GatewayQosProfile
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.GatewayQosProfilesClient)
		obj, err = client.Get(qosProfileIdParam)
		if err != nil {
			return obj, err
		}

	case utl.Global:
		client := c.Client.(client1.GatewayQosProfilesClient)
		gmObj, err1 := client.Get(qosProfileIdParam)
		if err1 != nil {
			return obj, err1
		}
		var rawObj interface{}
		rawObj, err = utl.ConvertModelBindingType(gmObj, model1.GatewayQosProfileBindingType(), model0.GatewayQosProfileBindingType())
		obj = rawObj.(model0.GatewayQosProfile)

	case utl.Multitenancy:
		client := c.Client.(client2.GatewayQosProfilesClient)
		obj, err = client.Get(utl.DefaultOrgID, c.ProjectID, qosProfileIdParam)
		if err != nil {
			return obj, err
		}

	default:
		return obj, errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c GatewayQosProfileClientContext) Patch(qosProfileIdParam string, gatewayQosProfileParam model0.GatewayQosProfile, overrideParam *bool) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.GatewayQosProfilesClient)
		err = client.Patch(qosProfileIdParam, gatewayQosProfileParam, overrideParam)

	case utl.Global:
		client := c.Client.(client1.GatewayQosProfilesClient)
		gmObj, err1 := utl.ConvertModelBindingType(gatewayQosProfileParam, model0.GatewayQosProfileBindingType(), model1.GatewayQosProfileBindingType())
		if err1 != nil {
			return err1
		}
		err = client.Patch(qosProfileIdParam, gmObj.(model1.GatewayQosProfile), overrideParam)

	case utl.Multitenancy:
		client := c.Client.(client2.GatewayQosProfilesClient)
		err = client.Patch(utl.DefaultOrgID, c.ProjectID, qosProfileIdParam, gatewayQosProfileParam, overrideParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c GatewayQosProfileClientContext) Update(qosProfileIdParam string, gatewayQosProfileParam model0.GatewayQosProfile, overrideParam *bool) (model0.GatewayQosProfile, error) {
	var err error
	var obj model0.GatewayQosProfile

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.GatewayQosProfilesClient)
		obj, err = client.Update(qosProfileIdParam, gatewayQosProfileParam, overrideParam)

	case utl.Global:
		client := c.Client.(client1.GatewayQosProfilesClient)
		gmObj, err := utl.ConvertModelBindingType(gatewayQosProfileParam, model0.GatewayQosProfileBindingType(), model1.GatewayQosProfileBindingType())
		if err != nil {
			return obj, err
		}
		gmObj, err = client.Update(qosProfileIdParam, gmObj.(model1.GatewayQosProfile), overrideParam)
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.GatewayQosProfileBindingType(), model0.GatewayQosProfileBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.GatewayQosProfile)

	case utl.Multitenancy:
		client := c.Client.(client2.GatewayQosProfilesClient)
		obj, err = client.Update(utl.DefaultOrgID, c.ProjectID, qosProfileIdParam, gatewayQosProfileParam, overrideParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c GatewayQosProfileClientContext) Delete(qosProfileIdParam string, overrideParam *bool) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.GatewayQosProfilesClient)
		err = client.Delete(qosProfileIdParam, overrideParam)

	case utl.Global:
		client := c.Client.(client1.GatewayQosProfilesClient)
		err = client.Delete(qosProfileIdParam, overrideParam)

	case utl.Multitenancy:
		client := c.Client.(client2.GatewayQosProfilesClient)
		err = client.Delete(utl.DefaultOrgID, c.ProjectID, qosProfileIdParam, overrideParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c GatewayQosProfileClientContext) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model0.GatewayQosProfileListResult, error) {
	var err error
	var obj model0.GatewayQosProfileListResult

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.GatewayQosProfilesClient)
		obj, err = client.List(cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	case utl.Global:
		client := c.Client.(client1.GatewayQosProfilesClient)
		gmObj, err := client.List(cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.GatewayQosProfileListResultBindingType(), model0.GatewayQosProfileListResultBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.GatewayQosProfileListResult)

	case utl.Multitenancy:
		client := c.Client.(client2.GatewayQosProfilesClient)
		obj, err = client.List(utl.DefaultOrgID, c.ProjectID, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}
