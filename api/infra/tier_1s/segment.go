//nolint:revive
package tier1s

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra/tier_1s"
	model1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_1s"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	client2 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/orgs/projects/infra/tier_1s"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

type SegmentClientContext utl.ClientContext

func NewSegmentsClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *SegmentClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewSegmentsClient(connector)

	case utl.Global:
		client = client1.NewSegmentsClient(connector)

	case utl.Multitenancy:
		client = client2.NewSegmentsClient(connector)

	default:
		return nil
	}
	return &SegmentClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID, VPCID: sessionContext.VPCID}
}

func (c SegmentClientContext) Get(tier1IdParam string, segmentIdParam string) (model0.Segment, error) {
	var obj model0.Segment
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.SegmentsClient)
		obj, err = client.Get(tier1IdParam, segmentIdParam)
		if err != nil {
			return obj, err
		}

	case utl.Global:
		client := c.Client.(client1.SegmentsClient)
		gmObj, err1 := client.Get(tier1IdParam, segmentIdParam)
		if err1 != nil {
			return obj, err1
		}
		var rawObj interface{}
		rawObj, err = utl.ConvertModelBindingType(gmObj, model1.SegmentBindingType(), model0.SegmentBindingType())
		obj = rawObj.(model0.Segment)

	case utl.Multitenancy:
		client := c.Client.(client2.SegmentsClient)
		obj, err = client.Get(utl.DefaultOrgID, c.ProjectID, tier1IdParam, segmentIdParam)
		if err != nil {
			return obj, err
		}

	default:
		return obj, errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c SegmentClientContext) Patch(tier1IdParam string, segmentIdParam string, segmentParam model0.Segment) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.SegmentsClient)
		err = client.Patch(tier1IdParam, segmentIdParam, segmentParam)

	case utl.Global:
		client := c.Client.(client1.SegmentsClient)
		gmObj, err1 := utl.ConvertModelBindingType(segmentParam, model0.SegmentBindingType(), model1.SegmentBindingType())
		if err1 != nil {
			return err1
		}
		err = client.Patch(tier1IdParam, segmentIdParam, gmObj.(model1.Segment))

	case utl.Multitenancy:
		client := c.Client.(client2.SegmentsClient)
		err = client.Patch(utl.DefaultOrgID, c.ProjectID, tier1IdParam, segmentIdParam, segmentParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c SegmentClientContext) Update(tier1IdParam string, segmentIdParam string, segmentParam model0.Segment) (model0.Segment, error) {
	var err error
	var obj model0.Segment

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.SegmentsClient)
		obj, err = client.Update(tier1IdParam, segmentIdParam, segmentParam)

	case utl.Global:
		client := c.Client.(client1.SegmentsClient)
		gmObj, err := utl.ConvertModelBindingType(segmentParam, model0.SegmentBindingType(), model1.SegmentBindingType())
		if err != nil {
			return obj, err
		}
		gmObj, err = client.Update(tier1IdParam, segmentIdParam, gmObj.(model1.Segment))
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.SegmentBindingType(), model0.SegmentBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.Segment)

	case utl.Multitenancy:
		client := c.Client.(client2.SegmentsClient)
		obj, err = client.Update(utl.DefaultOrgID, c.ProjectID, tier1IdParam, segmentIdParam, segmentParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c SegmentClientContext) Delete(tier1IdParam string, segmentIdParam string) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.SegmentsClient)
		err = client.Delete(tier1IdParam, segmentIdParam)

	case utl.Global:
		client := c.Client.(client1.SegmentsClient)
		err = client.Delete(tier1IdParam, segmentIdParam)

	case utl.Multitenancy:
		client := c.Client.(client2.SegmentsClient)
		err = client.Delete(utl.DefaultOrgID, c.ProjectID, tier1IdParam, segmentIdParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}

func (c SegmentClientContext) List(tier1IdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, segmentTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model0.SegmentListResult, error) {
	var err error
	var obj model0.SegmentListResult

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.SegmentsClient)
		obj, err = client.List(tier1IdParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, segmentTypeParam, sortAscendingParam, sortByParam)

	case utl.Global:
		client := c.Client.(client1.SegmentsClient)
		gmObj, err := client.List(tier1IdParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, segmentTypeParam, sortAscendingParam, sortByParam)
		if err != nil {
			return obj, err
		}
		obj1, err1 := utl.ConvertModelBindingType(gmObj, model1.SegmentListResultBindingType(), model0.SegmentListResultBindingType())
		if err1 != nil {
			return obj, err1
		}
		obj = obj1.(model0.SegmentListResult)

	case utl.Multitenancy:
		client := c.Client.(client2.SegmentsClient)
		obj, err = client.List(utl.DefaultOrgID, c.ProjectID, tier1IdParam, cursorParam, includeMarkForDeleteObjectsParam, includedFieldsParam, pageSizeParam, segmentTypeParam, sortAscendingParam, sortByParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return obj, err
}
