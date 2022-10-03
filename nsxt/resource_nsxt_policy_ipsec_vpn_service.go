/* Copyright © 2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	t0_locale_service "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_0s/locale_services"
	t1_locale_service "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/tier_1s/locale_services"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

var IPSecVpnServiceIkeLogLevelTypes = []string{
	model.IPSecVpnService_IKE_LOG_LEVEL_DEBUG,
	model.IPSecVpnService_IKE_LOG_LEVEL_INFO,
	model.IPSecVpnService_IKE_LOG_LEVEL_WARN,
	model.IPSecVpnService_IKE_LOG_LEVEL_ERROR,
	model.IPSecVpnService_IKE_LOG_LEVEL_EMERGENCY,
}

func resourceNsxtPolicyIPSecVpnService() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtPolicyIPSecVpnServiceCreate,
		Read:   resourceNsxtPolicyIPSecVpnServiceRead,
		Update: resourceNsxtPolicyIPSecVpnServiceUpdate,
		Delete: resourceNsxtPolicyIPSecVpnServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"nsx_id":              getNsxIDSchema(),
			"path":                getPathSchema(),
			"display_name":        getDisplayNameSchema(),
			"description":         getDescriptionSchema(),
			"revision":            getRevisionSchema(),
			"tag":                 getTagsSchema(),
			"locale_service_path": getPolicyPathSchema(true, false, "Polciy path for the locale service."),
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Enable/Disable IPSec VPN service.",
				Optional:    true,
				Default:     true,
			},
			"ha_sync": {
				Type:        schema.TypeBool,
				Description: "Enable/Disable IPSec VPN service HA state sync.",
				Optional:    true,
				Default:     true,
			},
			"ike_log_level": {
				Type:         schema.TypeString,
				Description:  "Log level for internet key exchange (IKE).",
				ValidateFunc: validation.StringInSlice(IPSecVpnServiceIkeLogLevelTypes, false),
				Optional:     true,
				Default:      model.IPSecVpnService_IKE_LOG_LEVEL_INFO,
			},
			"bypass_rule": getIPSecVPNRulesSchema(),
		},
	}
}

func getNsxtPolicyIPSecVpnServiceByID(connector *client.RestConnector, gwID string, isT0 bool, localeServiceID string, serviceID string, isGlobalManager bool) (model.IPSecVpnService, error) {
	if isT0 {
		client := t0_locale_service.NewIpsecVpnServicesClient(connector)
		return client.Get(gwID, localeServiceID, serviceID)
	}
	client := t1_locale_service.NewIpsecVpnServicesClient(connector)
	return client.Get(gwID, localeServiceID, serviceID)
}

func patchNsxtPolicyIPSecVpnService(connector *client.RestConnector, gwID string, localeServiceID string, ipSecVpnService model.IPSecVpnService, isT0 bool) error {
	id := *ipSecVpnService.Id
	if isT0 {
		client := t0_locale_service.NewIpsecVpnServicesClient(connector)
		return client.Patch(gwID, localeServiceID, id, ipSecVpnService)
	}
	client := t1_locale_service.NewIpsecVpnServicesClient(connector)
	return client.Patch(gwID, localeServiceID, id, ipSecVpnService)
}

func updateNsxtPolicyIPSecVpnService(connector *client.RestConnector, gwID string, localeServiceID string, ipSecVpnService model.IPSecVpnService, isT0 bool) error {
	id := *ipSecVpnService.Id
	if isT0 {
		client := t0_locale_service.NewIpsecVpnServicesClient(connector)
		_, err := client.Update(gwID, localeServiceID, id, ipSecVpnService)
		return err
	}
	client := t1_locale_service.NewIpsecVpnServicesClient(connector)
	_, err := client.Update(gwID, localeServiceID, id, ipSecVpnService)
	return err
}

func deleteNsxtPolicyIPSecVpnService(connector *client.RestConnector, gwID string, localeServiceID string, isT0 bool, id string) error {
	if isT0 {
		client := t0_locale_service.NewIpsecVpnServicesClient(connector)
		return client.Delete(gwID, localeServiceID, id)
	}
	client := t1_locale_service.NewIpsecVpnServicesClient(connector)
	return client.Delete(gwID, localeServiceID, id)
}

func setBypassRuleInSchema(d *schema.ResourceData, bypassRules []model.IPSecVpnRule) {
	var ruleList []map[string]interface{}
	for _, rule := range bypassRules {
		elem := make(map[string]interface{})
		var srcList []string
		for _, src := range rule.Sources {
			srcList = append(srcList, *src.Subnet)
		}
		var destList []string
		for _, dest := range rule.Destinations {
			destList = append(destList, *dest.Subnet)
		}
		elem["sources"] = srcList
		elem["destinations"] = destList
		elem["action"] = rule.Action
		ruleList = append(ruleList, elem)
	}
	err := d.Set("bypass_rule", ruleList)
	if err != nil {
		log.Printf("[WARNING] Failed to set bypass_rule in schema: %v", err)
	}
}

func getIPSecVPNBypassRulesFromSchema(d *schema.ResourceData) []model.IPSecVpnRule {
	rules := d.Get("bypass_rule")
	if rules != nil {
		rules := rules.([]interface{})
		var ruleList []model.IPSecVpnRule
		for _, rule := range rules {
			data := rule.(map[string]interface{})
			action := data["action"].(string)
			sourceRanges := interface2StringList(data["sources"].(*schema.Set).List())
			destinationRanges := interface2StringList(data["destinations"].(*schema.Set).List())
			/// Source Subnets
			sourceIPSecVpnSubnetList := make([]model.IPSecVpnSubnet, 0)
			if len(sourceRanges) > 0 {
				for _, element := range sourceRanges {
					subnet := element
					ipSecVpnSubnet := model.IPSecVpnSubnet{
						Subnet: &subnet,
					}
					sourceIPSecVpnSubnetList = append(sourceIPSecVpnSubnetList, ipSecVpnSubnet)
				}
			}
			/// Destination Subnets
			destinationIPSecVpnSubnetList := make([]model.IPSecVpnSubnet, 0)
			if len(destinationRanges) > 0 {
				for _, element := range destinationRanges {
					subnet := element
					ipSecVpnSubnet := model.IPSecVpnSubnet{
						Subnet: &subnet,
					}
					destinationIPSecVpnSubnetList = append(destinationIPSecVpnSubnetList, ipSecVpnSubnet)
				}
			}
			ruleID := data["nsx_id"].(string)
			if ruleID == "" {
				ruleID = newUUID()
			}
			elem := model.IPSecVpnRule{
				Action:       &action,
				Sources:      sourceIPSecVpnSubnetList,
				Destinations: destinationIPSecVpnSubnetList,
				UniqueId:     &ruleID,
				Id:           &ruleID,
			}
			ruleList = append(ruleList, elem)
		}
		return ruleList
	}
	return nil
}

func resourceNsxtPolicyIPSecVpnServiceRead(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining IPSecVpnService ID")
	}
	localeServicePath := d.Get("locale_service_path").(string)
	isT0, gwID, localeServiceID, err := parseLocaleServicePolicyPath(localeServicePath)
	if err != nil {
		return err
	}
	obj, err := getNsxtPolicyIPSecVpnServiceByID(connector, gwID, isT0, localeServiceID, id, isPolicyGlobalManager(m))
	if err != nil {
		return handleReadError(d, "IPSecVpnService", id, err)
	}
	d.Set("display_name", obj.DisplayName)
	d.Set("description", obj.Description)
	setPolicyTagsInSchema(d, obj.Tags)
	d.Set("nsx_id", id)
	d.Set("path", obj.Path)
	d.Set("revision", obj.Revision)
	d.Set("enabled", obj.Enabled)
	d.Set("ha_sync", obj.HaSync)
	setBypassRuleInSchema(d, obj.BypassRules)
	if obj.IkeLogLevel != nil {
		d.Set("ike_log_level", obj.IkeLogLevel)
	}
	return nil
}

func resourceNsxtPolicyIPSecVpnServiceCreate(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)
	localeServicePath := d.Get("locale_service_path").(string)
	isT0, gwID, localeServiceID, err := parseLocaleServicePolicyPath(localeServicePath)
	if err != nil {
		return err
	}
	isGlobalManager := isPolicyGlobalManager(m)
	id := d.Get("nsx_id").(string)
	if id == "" {
		id = newUUID()
	} else {
		_, err := getNsxtPolicyIPSecVpnServiceByID(connector, gwID, isT0, localeServiceID, id, isGlobalManager)
		if err == nil {
			return fmt.Errorf("IPSecVpnService with nsx_id '%s' already exists", id)
		} else if !isNotFoundError(err) {
			return err
		}
	}

	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	haSync := d.Get("ha_sync").(bool)
	rules := getIPSecVPNBypassRulesFromSchema(d)
	tags := getPolicyTagsFromSchema(d)

	ipSecVpnService := model.IPSecVpnService{
		Id:          &id,
		DisplayName: &displayName,
		Description: &description,
		Tags:        tags,
		Enabled:     &enabled,
		HaSync:      &haSync,
		BypassRules: rules,
	}

	ikeLogLevel := d.Get("ike_log_level").(string)
	if ikeLogLevel != "" {
		ipSecVpnService.IkeLogLevel = &ikeLogLevel
	}

	err = patchNsxtPolicyIPSecVpnService(connector, gwID, localeServiceID, ipSecVpnService, isT0)
	if err != nil {
		return handleCreateError("IPSecVpnService", id, err)
	}
	d.SetId(id)
	d.Set("nsx_id", id)
	return resourceNsxtPolicyIPSecVpnServiceRead(d, m)
}

func resourceNsxtPolicyIPSecVpnServiceUpdate(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining IPSec VPN Service ID")
	}
	localeServicePath := d.Get("locale_service_path").(string)
	isT0, gwID, localeServiceID, err := parseLocaleServicePolicyPath(localeServicePath)
	if err != nil {
		return err
	}

	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	haSync := d.Get("ha_sync").(bool)
	rules := getIPSecVPNBypassRulesFromSchema(d)
	tags := getPolicyTagsFromSchema(d)
	revision := int64(d.Get("revision").(int))
	ipSecVpnService := model.IPSecVpnService{
		Id:          &id,
		DisplayName: &displayName,
		Description: &description,
		Tags:        tags,
		Enabled:     &enabled,
		HaSync:      &haSync,
		BypassRules: rules,
		Revision:    &revision,
	}

	ikeLogLevel := d.Get("ike_log_level").(string)
	if ikeLogLevel != "" {
		ipSecVpnService.IkeLogLevel = &ikeLogLevel
	}

	log.Printf("[INFO] Updating IPSecVpnService with ID %s", id)
	err = updateNsxtPolicyIPSecVpnService(connector, gwID, localeServiceID, ipSecVpnService, isT0)
	if err != nil {
		return handleUpdateError("IPSecVpnService", id, err)
	}
	d.Set("nsx_id", id)
	return resourceNsxtPolicyIPSecVpnServiceRead(d, m)
}

func resourceNsxtPolicyIPSecVpnServiceDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	if id == "" {
		return fmt.Errorf("Error obtaining IPSec VPN Service ID")
	}

	localeServicePath := d.Get("locale_service_path").(string)
	isT0, gwID, localeServiceID, err := parseLocaleServicePolicyPath(localeServicePath)
	if err != nil {
		return err
	}

	err = deleteNsxtPolicyIPSecVpnService(getPolicyConnector(m), gwID, localeServiceID, isT0, id)
	if err != nil {
		return handleDeleteError("IPSecVpnService", id, err)
	}
	return nil
}
