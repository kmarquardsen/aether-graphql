package store

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	externalRef0 "github.com/onosproject/config-models/modelplugin/aether-2.0.0/aether_2_0_0"
)

// TODO: remove this file dependency entirely and dependency on roc-api

// ModelPluginDevice - a wrapper for the model plugin
type ModelPluginDevice struct {
	Device externalRef0.Device
}

// ToEnterprises converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprises(params ...string) (*types.Enterprises, error) {
	resource := new(types.Enterprises)

	// Property: enterprise []EnterprisesEnterprise
	// Handle []Object
	enterprises := make([]types.EnterprisesEnterprise, 0)
	reflectEnterprisesEnterprise, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterprise", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterprise != nil {
		for _, key := range reflectEnterprisesEnterprise.MapKeys() {
			v := reflectEnterprisesEnterprise.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			enterprise, err := d.ToEnterprisesEnterprise(childParams...)
			if err != nil {
				return nil, err
			}
			enterprises = append(enterprises, *enterprise)
		}
	}
	resource.Enterprise = &enterprises

	return resource, nil
}

// toEnterprisesEnterprise converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterprise(params ...string) (*types.EnterprisesEnterprise, error) {
	resource := new(types.EnterprisesEnterprise)

	// Property: application []EnterprisesEnterpriseApplication
	// Handle []Object
	applications := make([]types.EnterprisesEnterpriseApplication, 0)
	reflectEnterprisesEnterpriseApplication, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplication", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseApplication != nil {
		for _, key := range reflectEnterprisesEnterpriseApplication.MapKeys() {
			v := reflectEnterprisesEnterpriseApplication.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			application, err := d.ToEnterprisesEnterpriseApplication(childParams...)
			if err != nil {
				return nil, err
			}
			applications = append(applications, *application)
		}
	}
	resource.Application = &applications

	// Property: connectivity-service []EnterprisesEnterpriseConnectivityService
	// Handle []Object
	connectivityServices := make([]types.EnterprisesEnterpriseConnectivityService, 0)
	reflectEnterprisesEnterpriseConnectivityService, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseConnectivityService", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseConnectivityService != nil {
		for _, key := range reflectEnterprisesEnterpriseConnectivityService.MapKeys() {
			v := reflectEnterprisesEnterpriseConnectivityService.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			connectivityService, err := d.ToEnterprisesEnterpriseConnectivityService(childParams...)
			if err != nil {
				return nil, err
			}
			connectivityServices = append(connectivityServices, *connectivityService)
		}
	}
	resource.ConnectivityService = &connectivityServices

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: enterprise-id string
	//encoding gNMI attribute to OAPI
	reflectEnterpriseId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseEnterpriseId", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterpriseId != nil {
		attrEnterpriseId := reflectEnterpriseId.Interface().(string)
		resource.EnterpriseId = attrEnterpriseId
	}

	// Property: site []EnterprisesEnterpriseSite
	// Handle []Object
	sites := make([]types.EnterprisesEnterpriseSite, 0)
	reflectEnterprisesEnterpriseSite, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSite", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSite != nil {
		for _, key := range reflectEnterprisesEnterpriseSite.MapKeys() {
			v := reflectEnterprisesEnterpriseSite.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			site, err := d.ToEnterprisesEnterpriseSite(childParams...)
			if err != nil {
				return nil, err
			}
			sites = append(sites, *site)
		}
	}
	resource.Site = &sites

	// Property: template []EnterprisesEnterpriseTemplate
	// Handle []Object
	templates := make([]types.EnterprisesEnterpriseTemplate, 0)
	reflectEnterprisesEnterpriseTemplate, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplate", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseTemplate != nil {
		for _, key := range reflectEnterprisesEnterpriseTemplate.MapKeys() {
			v := reflectEnterprisesEnterpriseTemplate.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			template, err := d.ToEnterprisesEnterpriseTemplate(childParams...)
			if err != nil {
				return nil, err
			}
			templates = append(templates, *template)
		}
	}
	resource.Template = &templates

	// Property: traffic-class []EnterprisesEnterpriseTrafficClass
	// Handle []Object
	trafficClasss := make([]types.EnterprisesEnterpriseTrafficClass, 0)
	reflectEnterprisesEnterpriseTrafficClass, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClass", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseTrafficClass != nil {
		for _, key := range reflectEnterprisesEnterpriseTrafficClass.MapKeys() {
			v := reflectEnterprisesEnterpriseTrafficClass.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			trafficClass, err := d.ToEnterprisesEnterpriseTrafficClass(childParams...)
			if err != nil {
				return nil, err
			}
			trafficClasss = append(trafficClasss, *trafficClass)
		}
	}
	resource.TrafficClass = &trafficClasss

	return resource, nil
}

// toEnterprisesEnterpriseApplication converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseApplication(params ...string) (*types.EnterprisesEnterpriseApplication, error) {
	resource := new(types.EnterprisesEnterpriseApplication)

	// Property: address string
	//encoding gNMI attribute to OAPI
	reflectAddress, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationAddress", params...)
	if err != nil {
		return nil, err
	}
	if reflectAddress != nil {
		attrAddress := reflectAddress.Interface().(string)
		resource.Address = attrAddress
	}

	// Property: application-id string
	//encoding gNMI attribute to OAPI
	reflectApplicationId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationApplicationId", params...)
	if err != nil {
		return nil, err
	}
	if reflectApplicationId != nil {
		attrApplicationId := reflectApplicationId.Interface().(string)
		resource.ApplicationId = attrApplicationId
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: endpoint []EnterprisesEnterpriseApplicationEndpoint
	// Handle []Object
	endpoints := make([]types.EnterprisesEnterpriseApplicationEndpoint, 0)
	reflectEnterprisesEnterpriseApplicationEndpoint, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpoint", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseApplicationEndpoint != nil {
		for _, key := range reflectEnterprisesEnterpriseApplicationEndpoint.MapKeys() {
			v := reflectEnterprisesEnterpriseApplicationEndpoint.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			endpoint, err := d.ToEnterprisesEnterpriseApplicationEndpoint(childParams...)
			if err != nil {
				return nil, err
			}
			endpoints = append(endpoints, *endpoint)
		}
	}
	resource.Endpoint = &endpoints

	return resource, nil
}

// ToEnterprisesEnterpriseConnectivityService converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseConnectivityService(params ...string) (*types.EnterprisesEnterpriseConnectivityService, error) {
	resource := new(types.EnterprisesEnterpriseConnectivityService)

	// Property: connectivity-service string
	//encoding gNMI attribute to OAPI
	reflectConnectivityService, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseConnectivityServiceConnectivityService", params...)
	if err != nil {
		return nil, err
	}
	if reflectConnectivityService != nil {
		attrConnectivityService := reflectConnectivityService.Interface().(string)
		resource.ConnectivityService = attrConnectivityService
	}

	// Property: enabled bool
	//encoding gNMI attribute to OAPI
	reflectEnabled, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseConnectivityServiceEnabled", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnabled != nil {
		boolEnabled := reflectEnabled.Interface().(bool)
		resource.Enabled = &boolEnabled
	}

	return resource, nil
}

// toEnterprisesEnterpriseSite converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSite(params ...string) (*types.EnterprisesEnterpriseSite, error) {
	resource := new(types.EnterprisesEnterpriseSite)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: Device []EnterprisesEnterpriseSiteDevice
	// Handle []Object
	devices := make([]types.EnterprisesEnterpriseSiteDevice, 0)
	reflectEnterprisesEnterpriseSiteDevice, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDevice", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteDevice != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteDevice.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteDevice.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			device, err := d.ToEnterprisesEnterpriseSiteDevice(childParams...)
			if err != nil {
				return nil, err
			}
			devices = append(devices, *device)
		}
	}
	resource.Device = &devices

	// Property: Device-group []EnterprisesEnterpriseSiteDeviceGroup
	// Handle []Object
	deviceGroups := make([]types.EnterprisesEnterpriseSiteDeviceGroup, 0)
	reflectEnterprisesEnterpriseSiteDeviceGroup, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroup", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteDeviceGroup != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteDeviceGroup.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteDeviceGroup.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			deviceGroup, err := d.ToEnterprisesEnterpriseSiteDeviceGroup(childParams...)
			if err != nil {
				return nil, err
			}
			deviceGroups = append(deviceGroups, *deviceGroup)
		}
	}
	resource.DeviceGroup = &deviceGroups

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: imsi-definition EnterprisesEnterpriseSiteImsiDefinition
	//Handle object
	attrImsiDefinition, err := d.ToEnterprisesEnterpriseSiteImsiDefinition(params...)
	if err != nil {
		return nil, err
	}
	resource.ImsiDefinition = attrImsiDefinition

	// Property: ip-domain []EnterprisesEnterpriseSiteIpDomain
	// Handle []Object
	ipDomains := make([]types.EnterprisesEnterpriseSiteIpDomain, 0)
	reflectEnterprisesEnterpriseSiteIpDomain, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomain", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteIpDomain != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteIpDomain.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteIpDomain.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			ipDomain, err := d.ToEnterprisesEnterpriseSiteIpDomain(childParams...)
			if err != nil {
				return nil, err
			}
			ipDomains = append(ipDomains, *ipDomain)
		}
	}
	resource.IpDomain = &ipDomains

	// Property: monitoring EnterprisesEnterpriseSiteMonitoring
	//Handle object
	attrMonitoring, err := d.ToEnterprisesEnterpriseSiteMonitoring(params...)
	if err != nil {
		return nil, err
	}
	resource.Monitoring = attrMonitoring

	// Property: sim-card []EnterprisesEnterpriseSiteSimCard
	// Handle []Object
	simCards := make([]types.EnterprisesEnterpriseSiteSimCard, 0)
	reflectEnterprisesEnterpriseSiteSimCard, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSimCard", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSimCard != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSimCard.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSimCard.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			simCard, err := d.ToEnterprisesEnterpriseSiteSimCard(childParams...)
			if err != nil {
				return nil, err
			}
			simCards = append(simCards, *simCard)
		}
	}
	resource.SimCard = &simCards

	// Property: site-id string
	//encoding gNMI attribute to OAPI
	reflectSiteId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSiteId", params...)
	if err != nil {
		return nil, err
	}
	if reflectSiteId != nil {
		attrSiteId := reflectSiteId.Interface().(string)
		resource.SiteId = attrSiteId
	}

	// Property: slice []EnterprisesEnterpriseSiteSlice
	// Handle []Object
	slices := make([]types.EnterprisesEnterpriseSiteSlice, 0)
	reflectEnterprisesEnterpriseSiteSlice, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlice", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSlice != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSlice.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSlice.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			slice, err := d.ToEnterprisesEnterpriseSiteSlice(childParams...)
			if err != nil {
				return nil, err
			}
			slices = append(slices, *slice)
		}
	}
	resource.Slice = &slices

	// Property: small-cell []EnterprisesEnterpriseSiteSmallCell
	// Handle []Object
	smallCells := make([]types.EnterprisesEnterpriseSiteSmallCell, 0)
	reflectEnterprisesEnterpriseSiteSmallCell, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCell", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSmallCell != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSmallCell.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSmallCell.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			smallCell, err := d.ToEnterprisesEnterpriseSiteSmallCell(childParams...)
			if err != nil {
				return nil, err
			}
			smallCells = append(smallCells, *smallCell)
		}
	}
	resource.SmallCell = &smallCells

	// Property: upf []EnterprisesEnterpriseSiteUpf
	// Handle []Object
	upfs := make([]types.EnterprisesEnterpriseSiteUpf, 0)
	reflectEnterprisesEnterpriseSiteUpf, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpf", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteUpf != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteUpf.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteUpf.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			upf, err := d.ToEnterprisesEnterpriseSiteUpf(childParams...)
			if err != nil {
				return nil, err
			}
			upfs = append(upfs, *upf)
		}
	}
	resource.Upf = &upfs

	return resource, nil
}

// toEnterprisesEnterpriseTemplate converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseTemplate(params ...string) (*types.EnterprisesEnterpriseTemplate, error) {
	resource := new(types.EnterprisesEnterpriseTemplate)

	// Property: default-behavior string
	//encoding gNMI attribute to OAPI
	reflectDefaultBehavior, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateDefaultBehavior", params...)
	if err != nil {
		return nil, err
	}
	if reflectDefaultBehavior != nil {
		attrDefaultBehavior := reflectDefaultBehavior.Interface().(string)
		resource.DefaultBehavior = attrDefaultBehavior
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: mbr EnterprisesEnterpriseTemplateMbr
	//Handle object
	attrMbr, err := d.ToEnterprisesEnterpriseTemplateMbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Mbr = attrMbr

	// Property: sd int32
	//encoding gNMI attribute to OAPI
	reflectSd, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateSd", params...)
	if err != nil {
		return nil, err
	}
	if reflectSd != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Sd, err = utils.ToInt32Ptr(reflectSd); err != nil {
			return nil, err
		}
	}

	// Property: sst int
	//encoding gNMI attribute to OAPI
	reflectSst, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateSst", params...)
	if err != nil {
		return nil, err
	}
	if reflectSst != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Sst, err = utils.ToIntPtr(reflectSst); err != nil {
			return nil, err
		}
	}

	// Property: template-id string
	//encoding gNMI attribute to OAPI
	reflectTemplateId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateTemplateId", params...)
	if err != nil {
		return nil, err
	}
	if reflectTemplateId != nil {
		attrTemplateId := reflectTemplateId.Interface().(string)
		resource.TemplateId = attrTemplateId
	}

	return resource, nil
}

// toEnterprisesEnterpriseTrafficClass converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseTrafficClass(params ...string) (*types.EnterprisesEnterpriseTrafficClass, error) {
	resource := new(types.EnterprisesEnterpriseTrafficClass)

	// Property: arp int
	//encoding gNMI attribute to OAPI
	reflectArp, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassArp", params...)
	if err != nil {
		return nil, err
	}
	if reflectArp != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Arp, err = utils.ToIntPtr(reflectArp); err != nil {
			return nil, err
		}
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: pdb int
	//encoding gNMI attribute to OAPI
	reflectPdb, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassPdb", params...)
	if err != nil {
		return nil, err
	}
	if reflectPdb != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Pdb, err = utils.ToIntPtr(reflectPdb); err != nil {
			return nil, err
		}
	}

	// Property: pelr int
	//encoding gNMI attribute to OAPI
	reflectPelr, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassPelr", params...)
	if err != nil {
		return nil, err
	}
	if reflectPelr != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Pelr, err = utils.ToIntPtr(reflectPelr); err != nil {
			return nil, err
		}
	}

	// Property: qci int
	//encoding gNMI attribute to OAPI
	reflectQci, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassQci", params...)
	if err != nil {
		return nil, err
	}
	if reflectQci != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Qci, err = utils.ToIntPtr(reflectQci); err != nil {
			return nil, err
		}
	}

	// Property: traffic-class-id string
	//encoding gNMI attribute to OAPI
	reflectTrafficClassId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTrafficClassTrafficClassId", params...)
	if err != nil {
		return nil, err
	}
	if reflectTrafficClassId != nil {
		attrTrafficClassId := reflectTrafficClassId.Interface().(string)
		resource.TrafficClassId = attrTrafficClassId
	}

	return resource, nil
}

// toEnterprisesEnterpriseApplicationEndpoint converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseApplicationEndpoint(params ...string) (*types.EnterprisesEnterpriseApplicationEndpoint, error) {
	resource := new(types.EnterprisesEnterpriseApplicationEndpoint)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: endpoint-id string
	//encoding gNMI attribute to OAPI
	reflectEndpointId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointEndpointId", params...)
	if err != nil {
		return nil, err
	}
	if reflectEndpointId != nil {
		attrEndpointId := reflectEndpointId.Interface().(string)
		resource.EndpointId = attrEndpointId
	}

	// Property: mbr EnterprisesEnterpriseApplicationEndpointMbr
	//Handle object
	attrMbr, err := d.ToEnterprisesEnterpriseApplicationEndpointMbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Mbr = attrMbr

	// Property: port-end int
	//encoding gNMI attribute to OAPI
	reflectPortEnd, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointPortEnd", params...)
	if err != nil {
		return nil, err
	}
	if reflectPortEnd != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.PortEnd, err = utils.ToIntPtr(reflectPortEnd); err != nil {
			return nil, err
		}
	}

	// Property: port-start int
	//encoding gNMI attribute to OAPI
	reflectPortStart, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointPortStart", params...)
	if err != nil {
		return nil, err
	}
	if reflectPortStart != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.PortStart, err = utils.ToIntPtr(reflectPortStart); err != nil {
			return nil, err
		}
	}

	// Property: protocol string
	//encoding gNMI attribute to OAPI
	reflectProtocol, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointProtocol", params...)
	if err != nil {
		return nil, err
	}
	if reflectProtocol != nil {
		attrProtocol := reflectProtocol.Interface().(string)
		resource.Protocol = &attrProtocol
	}

	// Property: traffic-class string
	//encoding gNMI attribute to OAPI
	reflectTrafficClass, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointTrafficClass", params...)
	if err != nil {
		return nil, err
	}
	if reflectTrafficClass != nil {
		attrTrafficClass := reflectTrafficClass.Interface().(string)
		resource.TrafficClass = &attrTrafficClass
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDevice converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteDevice(params ...string) (*types.EnterprisesEnterpriseSiteDevice, error) {
	resource := new(types.EnterprisesEnterpriseSiteDevice)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: Device-id string
	//encoding gNMI attribute to OAPI
	reflectDeviceId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceDeviceId", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceId != nil {
		attrDeviceId := reflectDeviceId.Interface().(string)
		resource.DeviceId = attrDeviceId
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: imei string
	//encoding gNMI attribute to OAPI
	reflectImei, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceImei", params...)
	if err != nil {
		return nil, err
	}
	if reflectImei != nil {
		attrImei := reflectImei.Interface().(string)
		resource.Imei = &attrImei
	}

	// Property: sim-card string
	//encoding gNMI attribute to OAPI
	reflectSimCard, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceSimCard", params...)
	if err != nil {
		return nil, err
	}
	if reflectSimCard != nil {
		attrSimCard := reflectSimCard.Interface().(string)
		resource.SimCard = &attrSimCard
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDeviceGroup converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteDeviceGroup(params ...string) (*types.EnterprisesEnterpriseSiteDeviceGroup, error) {
	resource := new(types.EnterprisesEnterpriseSiteDeviceGroup)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: Device []EnterprisesEnterpriseSiteDeviceGroupDevice
	// Handle []Object
	devices := make([]types.EnterprisesEnterpriseSiteDeviceGroupDevice, 0)
	reflectEnterprisesEnterpriseSiteDeviceGroupDevice, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupDevice", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteDeviceGroupDevice != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteDeviceGroupDevice.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteDeviceGroupDevice.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			device, err := d.ToEnterprisesEnterpriseSiteDeviceGroupDevice(childParams...)
			if err != nil {
				return nil, err
			}
			devices = append(devices, *device)
		}
	}
	resource.Device = &devices

	// Property: Device-group-id string
	//encoding gNMI attribute to OAPI
	reflectDeviceGroupId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupDeviceGroupId", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceGroupId != nil {
		attrDeviceGroupId := reflectDeviceGroupId.Interface().(string)
		resource.DeviceGroupId = attrDeviceGroupId
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: ip-domain string
	//encoding gNMI attribute to OAPI
	reflectIpDomain, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupIpDomain", params...)
	if err != nil {
		return nil, err
	}
	if reflectIpDomain != nil {
		attrIpDomain := reflectIpDomain.Interface().(string)
		resource.IpDomain = &attrIpDomain
	}

	// Property: mbr EnterprisesEnterpriseSiteDeviceGroupMbr
	//Handle object
	attrMbr, err := d.ToEnterprisesEnterpriseSiteDeviceGroupMbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Mbr = attrMbr

	// Property: traffic-class string
	//encoding gNMI attribute to OAPI
	reflectTrafficClass, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupTrafficClass", params...)
	if err != nil {
		return nil, err
	}
	if reflectTrafficClass != nil {
		attrTrafficClass := reflectTrafficClass.Interface().(string)
		resource.TrafficClass = attrTrafficClass
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteImsiDefinition converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteImsiDefinition(params ...string) (*types.EnterprisesEnterpriseSiteImsiDefinition, error) {
	resource := new(types.EnterprisesEnterpriseSiteImsiDefinition)

	// Property: enterprise int32
	//encoding gNMI attribute to OAPI
	reflectEnterprise, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteImsiDefinitionEnterprise", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprise != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Enterprise, err = utils.ToInt32(reflectEnterprise); err != nil {
			return nil, err
		}
	}

	// Property: format string
	//encoding gNMI attribute to OAPI
	reflectFormat, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteImsiDefinitionFormat", params...)
	if err != nil {
		return nil, err
	}
	if reflectFormat != nil {
		attrFormat := reflectFormat.Interface().(string)
		resource.Format = attrFormat
	}

	// Property: mcc string
	//encoding gNMI attribute to OAPI
	reflectMcc, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteImsiDefinitionMcc", params...)
	if err != nil {
		return nil, err
	}
	if reflectMcc != nil {
		attrMcc := reflectMcc.Interface().(string)
		resource.Mcc = attrMcc
	}

	// Property: mnc string
	//encoding gNMI attribute to OAPI
	reflectMnc, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteImsiDefinitionMnc", params...)
	if err != nil {
		return nil, err
	}
	if reflectMnc != nil {
		attrMnc := reflectMnc.Interface().(string)
		resource.Mnc = attrMnc
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteIpDomain converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteIpDomain(params ...string) (*types.EnterprisesEnterpriseSiteIpDomain, error) {
	resource := new(types.EnterprisesEnterpriseSiteIpDomain)

	// Property: admin-status string
	//encoding gNMI attribute to OAPI
	reflectAdminStatus, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainAdminStatus", params...)
	if err != nil {
		return nil, err
	}
	if reflectAdminStatus != nil {
		attrAdminStatus := reflectAdminStatus.Interface().(string)
		resource.AdminStatus = &attrAdminStatus
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: dnn string
	//encoding gNMI attribute to OAPI
	reflectDnn, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainDnn", params...)
	if err != nil {
		return nil, err
	}
	if reflectDnn != nil {
		attrDnn := reflectDnn.Interface().(string)
		resource.Dnn = attrDnn
	}

	// Property: dns-primary string
	//encoding gNMI attribute to OAPI
	reflectDnsPrimary, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainDnsPrimary", params...)
	if err != nil {
		return nil, err
	}
	if reflectDnsPrimary != nil {
		attrDnsPrimary := reflectDnsPrimary.Interface().(string)
		resource.DnsPrimary = &attrDnsPrimary
	}

	// Property: dns-secondary string
	//encoding gNMI attribute to OAPI
	reflectDnsSecondary, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainDnsSecondary", params...)
	if err != nil {
		return nil, err
	}
	if reflectDnsSecondary != nil {
		attrDnsSecondary := reflectDnsSecondary.Interface().(string)
		resource.DnsSecondary = &attrDnsSecondary
	}

	// Property: ip-domain-id string
	//encoding gNMI attribute to OAPI
	reflectIpDomainId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainIpDomainId", params...)
	if err != nil {
		return nil, err
	}
	if reflectIpDomainId != nil {
		attrIpDomainId := reflectIpDomainId.Interface().(string)
		resource.IpDomainId = attrIpDomainId
	}

	// Property: mtu int
	//encoding gNMI attribute to OAPI
	reflectMtu, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainMtu", params...)
	if err != nil {
		return nil, err
	}
	if reflectMtu != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Mtu, err = utils.ToIntPtr(reflectMtu); err != nil {
			return nil, err
		}
	}

	// Property: subnet string
	//encoding gNMI attribute to OAPI
	reflectSubnet, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteIpDomainSubnet", params...)
	if err != nil {
		return nil, err
	}
	if reflectSubnet != nil {
		attrSubnet := reflectSubnet.Interface().(string)
		resource.Subnet = attrSubnet
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteMonitoring converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteMonitoring(params ...string) (*types.EnterprisesEnterpriseSiteMonitoring, error) {
	resource := new(types.EnterprisesEnterpriseSiteMonitoring)

	// Property: edge-cluster-prometheus-url string
	//encoding gNMI attribute to OAPI
	reflectEdgeClusterPrometheusUrl, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteMonitoringEdgeClusterPrometheusUrl", params...)
	if err != nil {
		return nil, err
	}
	if reflectEdgeClusterPrometheusUrl != nil {
		attrEdgeClusterPrometheusUrl := reflectEdgeClusterPrometheusUrl.Interface().(string)
		resource.EdgeClusterPrometheusUrl = &attrEdgeClusterPrometheusUrl
	}

	// Property: edge-Device []EnterprisesEnterpriseSiteMonitoringEdgeDevice
	// Handle []Object
	edgeDevices := make([]types.EnterprisesEnterpriseSiteMonitoringEdgeDevice, 0)
	reflectEnterprisesEnterpriseSiteMonitoringEdgeDevice, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteMonitoringEdgeDevice", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteMonitoringEdgeDevice != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteMonitoringEdgeDevice.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteMonitoringEdgeDevice.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			edgeDevice, err := d.ToEnterprisesEnterpriseSiteMonitoringEdgeDevice(childParams...)
			if err != nil {
				return nil, err
			}
			edgeDevices = append(edgeDevices, *edgeDevice)
		}
	}
	resource.EdgeDevice = &edgeDevices

	// Property: edge-monitoring-prometheus-url string
	//encoding gNMI attribute to OAPI
	reflectEdgeMonitoringPrometheusUrl, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteMonitoringEdgeMonitoringPrometheusUrl", params...)
	if err != nil {
		return nil, err
	}
	if reflectEdgeMonitoringPrometheusUrl != nil {
		attrEdgeMonitoringPrometheusUrl := reflectEdgeMonitoringPrometheusUrl.Interface().(string)
		resource.EdgeMonitoringPrometheusUrl = &attrEdgeMonitoringPrometheusUrl
	}

	return resource, nil
}

// ToEnterprisesEnterpriseSiteSimCard converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSimCard(params ...string) (*types.EnterprisesEnterpriseSiteSimCard, error) {
	resource := new(types.EnterprisesEnterpriseSiteSimCard)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSimCardDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSimCardDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: iccid string
	//encoding gNMI attribute to OAPI
	reflectIccid, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSimCardIccid", params...)
	if err != nil {
		return nil, err
	}
	if reflectIccid != nil {
		attrIccid := reflectIccid.Interface().(string)
		resource.Iccid = &attrIccid
	}

	// Property: imsi int64
	//encoding gNMI attribute to OAPI
	reflectImsi, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSimCardImsi", params...)
	if err != nil {
		return nil, err
	}
	if reflectImsi != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Imsi, err = utils.ToInt64Ptr(reflectImsi); err != nil {
			return nil, err
		}
	}

	// Property: sim-id string
	//encoding gNMI attribute to OAPI
	reflectSimId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSimCardSimId", params...)
	if err != nil {
		return nil, err
	}
	if reflectSimId != nil {
		attrSimId := reflectSimId.Interface().(string)
		resource.SimId = attrSimId
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSlice converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSlice(params ...string) (*types.EnterprisesEnterpriseSiteSlice, error) {
	resource := new(types.EnterprisesEnterpriseSiteSlice)

	// Property: default-behavior string
	//encoding gNMI attribute to OAPI
	reflectDefaultBehavior, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceDefaultBehavior", params...)
	if err != nil {
		return nil, err
	}
	if reflectDefaultBehavior != nil {
		attrDefaultBehavior := reflectDefaultBehavior.Interface().(string)
		resource.DefaultBehavior = attrDefaultBehavior
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: Device-group []EnterprisesEnterpriseSiteSliceDeviceGroup
	// Handle []Object
	deviceGroups := make([]types.EnterprisesEnterpriseSiteSliceDeviceGroup, 0)
	reflectEnterprisesEnterpriseSiteSliceDeviceGroup, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceDeviceGroup", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSliceDeviceGroup != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSliceDeviceGroup.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSliceDeviceGroup.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			deviceGroup, err := d.ToEnterprisesEnterpriseSiteSliceDeviceGroup(childParams...)
			if err != nil {
				return nil, err
			}
			deviceGroups = append(deviceGroups, *deviceGroup)
		}
	}
	resource.DeviceGroup = &deviceGroups

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: filter []EnterprisesEnterpriseSiteSliceFilter
	// Handle []Object
	filters := make([]types.EnterprisesEnterpriseSiteSliceFilter, 0)
	reflectEnterprisesEnterpriseSiteSliceFilter, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceFilter", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSliceFilter != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSliceFilter.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSliceFilter.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			filter, err := d.ToEnterprisesEnterpriseSiteSliceFilter(childParams...)
			if err != nil {
				return nil, err
			}
			filters = append(filters, *filter)
		}
	}
	resource.Filter = &filters

	// Property: mbr EnterprisesEnterpriseSiteSliceMbr
	//Handle object
	attrMbr, err := d.ToEnterprisesEnterpriseSiteSliceMbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Mbr = attrMbr

	// Property: priority-traffic-rule []EnterprisesEnterpriseSiteSlicePriorityTrafficRule
	// Handle []Object
	priorityTrafficRules := make([]types.EnterprisesEnterpriseSiteSlicePriorityTrafficRule, 0)
	reflectEnterprisesEnterpriseSiteSlicePriorityTrafficRule, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRule", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSlicePriorityTrafficRule != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSlicePriorityTrafficRule.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSlicePriorityTrafficRule.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			priorityTrafficRule, err := d.ToEnterprisesEnterpriseSiteSlicePriorityTrafficRule(childParams...)
			if err != nil {
				return nil, err
			}
			priorityTrafficRules = append(priorityTrafficRules, *priorityTrafficRule)
		}
	}
	resource.PriorityTrafficRule = &priorityTrafficRules

	// Property: sd int32
	//encoding gNMI attribute to OAPI
	reflectSd, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceSd", params...)
	if err != nil {
		return nil, err
	}
	if reflectSd != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Sd, err = utils.ToInt32(reflectSd); err != nil {
			return nil, err
		}
	}

	// Property: slice-id string
	//encoding gNMI attribute to OAPI
	reflectSliceId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceSliceId", params...)
	if err != nil {
		return nil, err
	}
	if reflectSliceId != nil {
		attrSliceId := reflectSliceId.Interface().(string)
		resource.SliceId = attrSliceId
	}

	// Property: sst int
	//encoding gNMI attribute to OAPI
	reflectSst, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceSst", params...)
	if err != nil {
		return nil, err
	}
	if reflectSst != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Sst, err = utils.ToInt(reflectSst); err != nil {
			return nil, err
		}
	}

	// Property: upf string
	//encoding gNMI attribute to OAPI
	reflectUpf, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceUpf", params...)
	if err != nil {
		return nil, err
	}
	if reflectUpf != nil {
		attrUpf := reflectUpf.Interface().(string)
		resource.Upf = &attrUpf
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSmallCell converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSmallCell(params ...string) (*types.EnterprisesEnterpriseSiteSmallCell, error) {
	resource := new(types.EnterprisesEnterpriseSiteSmallCell)

	// Property: address string
	//encoding gNMI attribute to OAPI
	reflectAddress, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCellAddress", params...)
	if err != nil {
		return nil, err
	}
	if reflectAddress != nil {
		attrAddress := reflectAddress.Interface().(string)
		resource.Address = &attrAddress
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCellDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCellDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: enable bool
	//encoding gNMI attribute to OAPI
	reflectEnable, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCellEnable", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnable != nil {
		boolEnable := reflectEnable.Interface().(bool)
		resource.Enable = &boolEnable
	}

	// Property: small-cell-id string
	//encoding gNMI attribute to OAPI
	reflectSmallCellId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCellSmallCellId", params...)
	if err != nil {
		return nil, err
	}
	if reflectSmallCellId != nil {
		attrSmallCellId := reflectSmallCellId.Interface().(string)
		resource.SmallCellId = attrSmallCellId
	}

	// Property: tac string
	//encoding gNMI attribute to OAPI
	reflectTac, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSmallCellTac", params...)
	if err != nil {
		return nil, err
	}
	if reflectTac != nil {
		attrTac := reflectTac.Interface().(string)
		resource.Tac = attrTac
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteUpf converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteUpf(params ...string) (*types.EnterprisesEnterpriseSiteUpf, error) {
	resource := new(types.EnterprisesEnterpriseSiteUpf)

	// Property: address string
	//encoding gNMI attribute to OAPI
	reflectAddress, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpfAddress", params...)
	if err != nil {
		return nil, err
	}
	if reflectAddress != nil {
		attrAddress := reflectAddress.Interface().(string)
		resource.Address = attrAddress
	}

	// Property: config-endpoint string
	//encoding gNMI attribute to OAPI
	reflectConfigEndpoint, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpfConfigEndpoint", params...)
	if err != nil {
		return nil, err
	}
	if reflectConfigEndpoint != nil {
		attrConfigEndpoint := reflectConfigEndpoint.Interface().(string)
		resource.ConfigEndpoint = &attrConfigEndpoint
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpfDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpfDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: port int
	//encoding gNMI attribute to OAPI
	reflectPort, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpfPort", params...)
	if err != nil {
		return nil, err
	}
	if reflectPort != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Port, err = utils.ToInt(reflectPort); err != nil {
			return nil, err
		}
	}

	// Property: upf-id string
	//encoding gNMI attribute to OAPI
	reflectUpfId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteUpfUpfId", params...)
	if err != nil {
		return nil, err
	}
	if reflectUpfId != nil {
		attrUpfId := reflectUpfId.Interface().(string)
		resource.UpfId = attrUpfId
	}

	return resource, nil
}

// toEnterprisesEnterpriseTemplateMbr converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseTemplateMbr(params ...string) (*types.EnterprisesEnterpriseTemplateMbr, error) {
	resource := new(types.EnterprisesEnterpriseTemplateMbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateMbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64Ptr(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: downlink-burst-size int32
	//encoding gNMI attribute to OAPI
	reflectDownlinkBurstSize, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateMbrDownlinkBurstSize", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlinkBurstSize != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.DownlinkBurstSize, err = utils.ToInt32Ptr(reflectDownlinkBurstSize); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateMbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64Ptr(reflectUplink); err != nil {
			return nil, err
		}
	}

	// Property: uplink-burst-size int32
	//encoding gNMI attribute to OAPI
	reflectUplinkBurstSize, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseTemplateMbrUplinkBurstSize", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplinkBurstSize != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.UplinkBurstSize, err = utils.ToInt32Ptr(reflectUplinkBurstSize); err != nil {
			return nil, err
		}
	}

	return resource, nil
}

// toEnterprisesEnterpriseApplicationEndpointMbr converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseApplicationEndpointMbr(params ...string) (*types.EnterprisesEnterpriseApplicationEndpointMbr, error) {
	resource := new(types.EnterprisesEnterpriseApplicationEndpointMbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointMbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64Ptr(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseApplicationEndpointMbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64Ptr(reflectUplink); err != nil {
			return nil, err
		}
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDeviceGroupDevice converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteDeviceGroupDevice(params ...string) (*types.EnterprisesEnterpriseSiteDeviceGroupDevice, error) {
	resource := new(types.EnterprisesEnterpriseSiteDeviceGroupDevice)

	// Property: Device-id string
	//encoding gNMI attribute to OAPI
	reflectDeviceId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupDeviceDeviceId", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceId != nil {
		attrDeviceId := reflectDeviceId.Interface().(string)
		resource.DeviceId = attrDeviceId
	}

	// Property: enable bool
	//encoding gNMI attribute to OAPI
	reflectEnable, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupDeviceEnable", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnable != nil {
		boolEnable := reflectEnable.Interface().(bool)
		resource.Enable = &boolEnable
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDeviceGroupMbr converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteDeviceGroupMbr(params ...string) (*types.EnterprisesEnterpriseSiteDeviceGroupMbr, error) {
	resource := new(types.EnterprisesEnterpriseSiteDeviceGroupMbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupMbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteDeviceGroupMbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64(reflectUplink); err != nil {
			return nil, err
		}
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteMonitoringEdgeDevice converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteMonitoringEdgeDevice(params ...string) (*types.EnterprisesEnterpriseSiteMonitoringEdgeDevice, error) {
	resource := new(types.EnterprisesEnterpriseSiteMonitoringEdgeDevice)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteMonitoringEdgeDeviceDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteMonitoringEdgeDeviceDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: edge-Device-id string
	//encoding gNMI attribute to OAPI
	reflectEdgeDeviceId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteMonitoringEdgeDeviceEdgeDeviceId", params...)
	if err != nil {
		return nil, err
	}
	if reflectEdgeDeviceId != nil {
		attrEdgeDeviceId := reflectEdgeDeviceId.Interface().(string)
		resource.EdgeDeviceId = attrEdgeDeviceId
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSliceDeviceGroup converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSliceDeviceGroup(params ...string) (*types.EnterprisesEnterpriseSiteSliceDeviceGroup, error) {
	resource := new(types.EnterprisesEnterpriseSiteSliceDeviceGroup)

	// Property: Device-group string
	//encoding gNMI attribute to OAPI
	reflectDeviceGroup, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceDeviceGroupDeviceGroup", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceGroup != nil {
		attrDeviceGroup := reflectDeviceGroup.Interface().(string)
		resource.DeviceGroup = attrDeviceGroup
	}

	// Property: enable bool
	//encoding gNMI attribute to OAPI
	reflectEnable, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceDeviceGroupEnable", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnable != nil {
		boolEnable := reflectEnable.Interface().(bool)
		resource.Enable = &boolEnable
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSliceFilter converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSliceFilter(params ...string) (*types.EnterprisesEnterpriseSiteSliceFilter, error) {
	resource := new(types.EnterprisesEnterpriseSiteSliceFilter)

	// Property: allow bool
	//encoding gNMI attribute to OAPI
	reflectAllow, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceFilterAllow", params...)
	if err != nil {
		return nil, err
	}
	if reflectAllow != nil {
		boolAllow := reflectAllow.Interface().(bool)
		resource.Allow = &boolAllow
	}

	// Property: application string
	//encoding gNMI attribute to OAPI
	reflectApplication, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceFilterApplication", params...)
	if err != nil {
		return nil, err
	}
	if reflectApplication != nil {
		attrApplication := reflectApplication.Interface().(string)
		resource.Application = attrApplication
	}

	// Property: priority int
	//encoding gNMI attribute to OAPI
	reflectPriority, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceFilterPriority", params...)
	if err != nil {
		return nil, err
	}
	if reflectPriority != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Priority, err = utils.ToIntPtr(reflectPriority); err != nil {
			return nil, err
		}
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSliceMbr converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSliceMbr(params ...string) (*types.EnterprisesEnterpriseSiteSliceMbr, error) {
	resource := new(types.EnterprisesEnterpriseSiteSliceMbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceMbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64Ptr(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: downlink-burst-size int32
	//encoding gNMI attribute to OAPI
	reflectDownlinkBurstSize, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceMbrDownlinkBurstSize", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlinkBurstSize != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.DownlinkBurstSize, err = utils.ToInt32Ptr(reflectDownlinkBurstSize); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceMbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64Ptr(reflectUplink); err != nil {
			return nil, err
		}
	}

	// Property: uplink-burst-size int32
	//encoding gNMI attribute to OAPI
	reflectUplinkBurstSize, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSliceMbrUplinkBurstSize", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplinkBurstSize != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.UplinkBurstSize, err = utils.ToInt32Ptr(reflectUplinkBurstSize); err != nil {
			return nil, err
		}
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSlicePriorityTrafficRule converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSlicePriorityTrafficRule(params ...string) (*types.EnterprisesEnterpriseSiteSlicePriorityTrafficRule, error) {
	resource := new(types.EnterprisesEnterpriseSiteSlicePriorityTrafficRule)

	// Property: application string
	//encoding gNMI attribute to OAPI
	reflectApplication, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleApplication", params...)
	if err != nil {
		return nil, err
	}
	if reflectApplication != nil {
		attrApplication := reflectApplication.Interface().(string)
		resource.Application = attrApplication
	}

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: Device string
	//encoding gNMI attribute to OAPI
	reflectDevice, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleDevice", params...)
	if err != nil {
		return nil, err
	}
	if reflectDevice != nil {
		attrDevice := reflectDevice.Interface().(string)
		resource.Device = attrDevice
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: endpoint string
	//encoding gNMI attribute to OAPI
	reflectEndpoint, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleEndpoint", params...)
	if err != nil {
		return nil, err
	}
	if reflectEndpoint != nil {
		attrEndpoint := reflectEndpoint.Interface().(string)
		resource.Endpoint = attrEndpoint
	}

	// Property: gbr EnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbr
	//Handle object
	attrGbr, err := d.ToEnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Gbr = attrGbr

	// Property: mbr EnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbr
	//Handle object
	attrMbr, err := d.ToEnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Mbr = attrMbr

	// Property: priority-traffic-rule-id string
	//encoding gNMI attribute to OAPI
	reflectPriorityTrafficRuleId, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRulePriorityTrafficRuleId", params...)
	if err != nil {
		return nil, err
	}
	if reflectPriorityTrafficRuleId != nil {
		attrPriorityTrafficRuleId := reflectPriorityTrafficRuleId.Interface().(string)
		resource.PriorityTrafficRuleId = attrPriorityTrafficRuleId
	}

	// Property: traffic-class string
	//encoding gNMI attribute to OAPI
	reflectTrafficClass, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleTrafficClass", params...)
	if err != nil {
		return nil, err
	}
	if reflectTrafficClass != nil {
		attrTrafficClass := reflectTrafficClass.Interface().(string)
		resource.TrafficClass = &attrTrafficClass
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbr converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbr(params ...string) (*types.EnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbr, error) {
	resource := new(types.EnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64Ptr(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleGbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64Ptr(reflectUplink); err != nil {
			return nil, err
		}
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbr converts gNMI to OAPI.
func (d *ModelPluginDevice) ToEnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbr(params ...string) (*types.EnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbr, error) {
	resource := new(types.EnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64Ptr(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.Device, "EnterprisesEnterpriseSiteSlicePriorityTrafficRuleMbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64Ptr(reflectUplink); err != nil {
			return nil, err
		}
	}

	return resource, nil
}
