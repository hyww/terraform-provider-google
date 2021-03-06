// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDNSManagedZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSManagedZoneCreate,
		Read:   resourceDNSManagedZoneRead,
		Update: resourceDNSManagedZoneUpdate,
		Delete: resourceDNSManagedZoneDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDNSManagedZoneImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dns_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The DNS name of this managed zone, for instance "example.com.".`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `User assigned name for this resource.
Must be unique within the project.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A textual description field. Defaults to 'Managed by Terraform'.`,
				Default:     "Managed by Terraform",
			},
			"dnssec_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `DNSSEC configuration`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_key_specs": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `Specifies parameters that will be used for generating initial DnsKeys
for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
you must also provide one for the other.
default_key_specs can only be updated when the state is 'off'.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"algorithm": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512", ""}, false),
										Description:  `String mnemonic specifying the DNSSEC algorithm of this key`,
									},
									"key_length": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: `Length of the keys in bits`,
									},
									"key_type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"keySigning", "zoneSigning", ""}, false),
										Description: `Specifies whether this is a key signing key (KSK) or a zone
signing key (ZSK). Key signing keys have the Secure Entry
Point flag set and, when active, will only be used to sign
resource record sets of type DNSKEY. Zone signing keys do
not have the Secure Entry Point flag set and will be used
to sign all other types of resource record sets.`,
									},
									"kind": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Identifies what kind of resource this is`,
										Default:     "dns#dnsKeySpec",
									},
								},
							},
							AtLeastOneOf: []string{"dnssec_config.0.kind", "dnssec_config.0.non_existence", "dnssec_config.0.state", "dnssec_config.0.default_key_specs"},
						},
						"kind": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `Identifies what kind of resource this is`,
							Default:      "dns#managedZoneDnsSecConfig",
							AtLeastOneOf: []string{"dnssec_config.0.kind", "dnssec_config.0.non_existence", "dnssec_config.0.state", "dnssec_config.0.default_key_specs"},
						},
						"non_existence": {
							Type:         schema.TypeString,
							Computed:     true,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"nsec", "nsec3", ""}, false),
							Description: `Specifies the mechanism used to provide authenticated denial-of-existence responses.
non_existence can only be updated when the state is 'off'.`,
							AtLeastOneOf: []string{"dnssec_config.0.kind", "dnssec_config.0.non_existence", "dnssec_config.0.state", "dnssec_config.0.default_key_specs"},
						},
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"off", "on", "transfer", ""}, false),
							Description:  `Specifies whether DNSSEC is enabled, and what mode it is in`,
							AtLeastOneOf: []string{"dnssec_config.0.kind", "dnssec_config.0.non_existence", "dnssec_config.0.state", "dnssec_config.0.default_key_specs"},
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to this ManagedZone.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"private_visibility_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `For privately visible zones, the set of Virtual Private Cloud
resources that the zone is visible from.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"networks": {
							Type:     schema.TypeSet,
							Required: true,
							Description: `The list of VPC networks that can see this zone. Until the provider updates to use the Terraform 0.12 SDK in a future release, you
may experience issues with this resource while updating. If you've defined a 'networks' block and
add another 'networks' block while keeping the old block, Terraform will see an incorrect diff
and apply an incorrect update to the resource. If you encounter this issue, remove all 'networks'
blocks in an update and then apply another update adding all of them back simultaneously.`,
							Elem: dnsManagedZonePrivateVisibilityConfigNetworksSchema(),
							Set: func(v interface{}) int {
								if v == nil {
									return 0
								}
								raw := v.(map[string]interface{})
								if url, ok := raw["network_url"]; ok {
									return selfLinkNameHash(url)
								}
								var buf bytes.Buffer
								schema.SerializeResourceForHash(&buf, raw, dnsManagedZonePrivateVisibilityConfigNetworksSchema())
								return hashcode.String(buf.String())
							},
						},
					},
				},
			},
			"visibility": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"private", "public", ""}, false),
				DiffSuppressFunc: caseDiffSuppress,
				Description: `The zone's visibility: public zones are exposed to the Internet,
while private zones are visible only to Virtual Private Cloud resources.
Must be one of: 'public', 'private'.`,
				Default: "public",
			},
			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Delegate your managed_zone to these virtual name servers;
defined by the server`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func dnsManagedZonePrivateVisibilityConfigNetworksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_url": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The fully qualified URL of the VPC network to bind to.
This should be formatted like
'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'`,
			},
		},
	}
}

func resourceDNSManagedZoneCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandDNSManagedZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dnsNameProp, err := expandDNSManagedZoneDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns_name"); !isEmptyValue(reflect.ValueOf(dnsNameProp)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	dnssecConfigProp, err := expandDNSManagedZoneDnssecConfig(d.Get("dnssec_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dnssec_config"); !isEmptyValue(reflect.ValueOf(dnssecConfigProp)) && (ok || !reflect.DeepEqual(v, dnssecConfigProp)) {
		obj["dnssecConfig"] = dnssecConfigProp
	}
	nameProp, err := expandDNSManagedZoneName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	labelsProp, err := expandDNSManagedZoneLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	visibilityProp, err := expandDNSManagedZoneVisibility(d.Get("visibility"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("visibility"); !isEmptyValue(reflect.ValueOf(visibilityProp)) && (ok || !reflect.DeepEqual(v, visibilityProp)) {
		obj["visibility"] = visibilityProp
	}
	privateVisibilityConfigProp, err := expandDNSManagedZonePrivateVisibilityConfig(d.Get("private_visibility_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_visibility_config"); !isEmptyValue(reflect.ValueOf(privateVisibilityConfigProp)) && (ok || !reflect.DeepEqual(v, privateVisibilityConfigProp)) {
		obj["privateVisibilityConfig"] = privateVisibilityConfigProp
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ManagedZone: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ManagedZone: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ManagedZone %q: %#v", d.Id(), res)

	return resourceDNSManagedZoneRead(d, meta)
}

func resourceDNSManagedZoneRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DNSManagedZone %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}

	if err := d.Set("description", flattenDNSManagedZoneDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("dns_name", flattenDNSManagedZoneDnsName(res["dnsName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("dnssec_config", flattenDNSManagedZoneDnssecConfig(res["dnssecConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("name", flattenDNSManagedZoneName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("name_servers", flattenDNSManagedZoneNameServers(res["nameServers"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("labels", flattenDNSManagedZoneLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("visibility", flattenDNSManagedZoneVisibility(res["visibility"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("private_visibility_config", flattenDNSManagedZonePrivateVisibilityConfig(res["privateVisibilityConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}

	return nil
}

func resourceDNSManagedZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandDNSManagedZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dnssecConfigProp, err := expandDNSManagedZoneDnssecConfig(d.Get("dnssec_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dnssec_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dnssecConfigProp)) {
		obj["dnssecConfig"] = dnssecConfigProp
	}
	labelsProp, err := expandDNSManagedZoneLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	privateVisibilityConfigProp, err := expandDNSManagedZonePrivateVisibilityConfig(d.Get("private_visibility_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("private_visibility_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, privateVisibilityConfigProp)) {
		obj["privateVisibilityConfig"] = privateVisibilityConfigProp
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ManagedZone %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ManagedZone %q: %s", d.Id(), err)
	}

	return resourceDNSManagedZoneRead(d, meta)
}

func resourceDNSManagedZoneDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ManagedZone %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ManagedZone")
	}

	log.Printf("[DEBUG] Finished deleting ManagedZone %q: %#v", d.Id(), res)
	return nil
}

func resourceDNSManagedZoneImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/managedZones/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDNSManagedZoneDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kind"] =
		flattenDNSManagedZoneDnssecConfigKind(original["kind"], d, config)
	transformed["non_existence"] =
		flattenDNSManagedZoneDnssecConfigNonExistence(original["nonExistence"], d, config)
	transformed["state"] =
		flattenDNSManagedZoneDnssecConfigState(original["state"], d, config)
	transformed["default_key_specs"] =
		flattenDNSManagedZoneDnssecConfigDefaultKeySpecs(original["defaultKeySpecs"], d, config)
	return []interface{}{transformed}
}
func flattenDNSManagedZoneDnssecConfigKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigNonExistence(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"algorithm":  flattenDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(original["algorithm"], d, config),
			"key_length": flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(original["keyLength"], d, config),
			"key_type":   flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(original["keyType"], d, config),
			"kind":       flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKind(original["kind"], d, config),
		})
	}
	return transformed
}
func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneDnssecConfigDefaultKeySpecsKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneNameServers(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDNSManagedZoneVisibility(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil || isEmptyValue(reflect.ValueOf(v)) {
		return "public"
	}

	return v
}

func flattenDNSManagedZonePrivateVisibilityConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["networks"] =
		flattenDNSManagedZonePrivateVisibilityConfigNetworks(original["networks"], d, config)
	return []interface{}{transformed}
}
func flattenDNSManagedZonePrivateVisibilityConfigNetworks(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		if v == nil {
			return 0
		}
		raw := v.(map[string]interface{})
		if url, ok := raw["network_url"]; ok {
			return selfLinkNameHash(url)
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsManagedZonePrivateVisibilityConfigNetworksSchema())
		return hashcode.String(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"network_url": flattenDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(original["networkUrl"], d, config),
		})
	}
	return transformed
}
func flattenDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDNSManagedZoneDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKind, err := expandDNSManagedZoneDnssecConfigKind(original["kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
		transformed["kind"] = transformedKind
	}

	transformedNonExistence, err := expandDNSManagedZoneDnssecConfigNonExistence(original["non_existence"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNonExistence); val.IsValid() && !isEmptyValue(val) {
		transformed["nonExistence"] = transformedNonExistence
	}

	transformedState, err := expandDNSManagedZoneDnssecConfigState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedDefaultKeySpecs, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecs(original["default_key_specs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultKeySpecs); val.IsValid() && !isEmptyValue(val) {
		transformed["defaultKeySpecs"] = transformedDefaultKeySpecs
	}

	return transformed, nil
}

func expandDNSManagedZoneDnssecConfigKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigNonExistence(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAlgorithm, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(original["algorithm"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !isEmptyValue(val) {
			transformed["algorithm"] = transformedAlgorithm
		}

		transformedKeyLength, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(original["key_length"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyLength); val.IsValid() && !isEmptyValue(val) {
			transformed["keyLength"] = transformedKeyLength
		}

		transformedKeyType, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(original["key_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyType); val.IsValid() && !isEmptyValue(val) {
			transformed["keyType"] = transformedKeyType
		}

		transformedKind, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKind(original["kind"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
			transformed["kind"] = transformedKind
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDNSManagedZoneVisibility(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZonePrivateVisibilityConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworks, err := expandDNSManagedZonePrivateVisibilityConfigNetworks(original["networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworks); val.IsValid() && !isEmptyValue(val) {
		transformed["networks"] = transformedNetworks
	}

	return transformed, nil
}

func expandDNSManagedZonePrivateVisibilityConfigNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
