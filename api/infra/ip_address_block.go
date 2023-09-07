//nolint:revive
package infra

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/orgs/projects/infra"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

type IpAddressBlockClientContext utl.ClientContext

func NewIpBlocksClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *IpAddressBlockClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewIpBlocksClient(connector)

	case utl.Multitenancy:
		client = client1.NewIpBlocksClient(connector)

	default:
		return nil
	}
	return &IpAddressBlockClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID}
}

func (c IpAddressBlockClientContext) Get(ipBlockIdParam string) (model0.IpAddressBlock, error) {
	var obj model0.IpAddressBlock
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.IpBlocksClient)
		obj, err = client.Get(ipBlockIdParam)
		if err != nil {
			return obj, err
		}

	case utl.Multitenancy:
		client := c.Client.(client1.IpBlocksClient)
		obj, err = client.Get(utl.DefaultOrgID, c.ProjectID, ipBlockIdParam)
		if err != nil {
			return obj, err
		}

	default:
		return obj, errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c IpAddressBlockClientContext) Patch(ipBlockIdParam string, ipAddressBlockParam model0.IpAddressBlock) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.IpBlocksClient)
		err = client.Patch(ipBlockIdParam, ipAddressBlockParam)

	case utl.Multitenancy:
		client := c.Client.(client1.IpBlocksClient)
		err = client.Patch(utl.DefaultOrgID, c.ProjectID, ipBlockIdParam, ipAddressBlockParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c IpAddressBlockClientContext) Update(ipBlockIdParam string, ipAddressBlockParam model0.IpAddressBlock) (model0.IpAddressBlock, error) {
	var err error
	var obj model0.IpAddressBlock

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.IpBlocksClient)
		obj, err = client.Update(ipBlockIdParam, ipAddressBlockParam)

	case utl.Multitenancy:
		client := c.Client.(client1.IpBlocksClient)
		obj, err = client.Update(utl.DefaultOrgID, c.ProjectID, ipBlockIdParam, ipAddressBlockParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c IpAddressBlockClientContext) Delete(ipBlockIdParam string) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.IpBlocksClient)
		err = client.Delete(ipBlockIdParam)

	case utl.Multitenancy:
		client := c.Client.(client1.IpBlocksClient)
		err = client.Delete(utl.DefaultOrgID, c.ProjectID, ipBlockIdParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c IpAddressBlockClientContext) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model0.IpAddressBlockListResult, error) {
	var err error
	var obj model0.IpAddressBlockListResult

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.IpBlocksClient)
		obj, err = client.List(cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	case utl.Multitenancy:
		client := c.Client.(client1.IpBlocksClient)
		obj, err = client.List(utl.DefaultOrgID, c.ProjectID, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, sortAscendingParam, sortByParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}