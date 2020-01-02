/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/danitso/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

const (
	dvResourceVirtualEnvironmentVMACPI                              = true
	dvResourceVirtualEnvironmentVMAgentEnabled                      = false
	dvResourceVirtualEnvironmentVMAgentTrim                         = false
	dvResourceVirtualEnvironmentVMAgentType                         = "virtio"
	dvResourceVirtualEnvironmentVMBIOS                              = "seabios"
	dvResourceVirtualEnvironmentVMCDROMEnabled                      = false
	dvResourceVirtualEnvironmentVMCDROMFileID                       = ""
	dvResourceVirtualEnvironmentVMCPUArchitecture                   = "x86_64"
	dvResourceVirtualEnvironmentVMCPUCores                          = 1
	dvResourceVirtualEnvironmentVMCPUHotplugged                     = 0
	dvResourceVirtualEnvironmentVMCPUSockets                        = 1
	dvResourceVirtualEnvironmentVMCPUType                           = "qemu64"
	dvResourceVirtualEnvironmentVMCPUUnits                          = 1024
	dvResourceVirtualEnvironmentVMDescription                       = ""
	dvResourceVirtualEnvironmentVMDiskDatastoreID                   = "local-lvm"
	dvResourceVirtualEnvironmentVMDiskFileFormat                    = "qcow2"
	dvResourceVirtualEnvironmentVMDiskFileID                        = ""
	dvResourceVirtualEnvironmentVMDiskSize                          = 8
	dvResourceVirtualEnvironmentVMDiskSpeedRead                     = 0
	dvResourceVirtualEnvironmentVMDiskSpeedReadBurstable            = 0
	dvResourceVirtualEnvironmentVMDiskSpeedWrite                    = 0
	dvResourceVirtualEnvironmentVMDiskSpeedWriteBurstable           = 0
	dvResourceVirtualEnvironmentVMInitializationDNSDomain           = ""
	dvResourceVirtualEnvironmentVMInitializationDNSServer           = ""
	dvResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address = ""
	dvResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway = ""
	dvResourceVirtualEnvironmentVMInitializationIPConfigIPv6Address = ""
	dvResourceVirtualEnvironmentVMInitializationIPConfigIPv6Gateway = ""
	dvResourceVirtualEnvironmentVMInitializationUserAccountPassword = ""
	dvResourceVirtualEnvironmentVMInitializationUserDataFileID      = ""
	dvResourceVirtualEnvironmentVMKeyboardLayout                    = "en-us"
	dvResourceVirtualEnvironmentVMMemoryDedicated                   = 512
	dvResourceVirtualEnvironmentVMMemoryFloating                    = 0
	dvResourceVirtualEnvironmentVMMemoryShared                      = 0
	dvResourceVirtualEnvironmentVMName                              = ""
	dvResourceVirtualEnvironmentVMNetworkDeviceBridge               = "vmbr0"
	dvResourceVirtualEnvironmentVMNetworkDeviceEnabled              = true
	dvResourceVirtualEnvironmentVMNetworkDeviceMACAddress           = ""
	dvResourceVirtualEnvironmentVMNetworkDeviceModel                = "virtio"
	dvResourceVirtualEnvironmentVMNetworkDeviceRateLimit            = 0
	dvResourceVirtualEnvironmentVMNetworkDeviceVLANID               = 0
	dvResourceVirtualEnvironmentVMOperatingSystemType               = "other"
	dvResourceVirtualEnvironmentVMPoolID                            = ""
	dvResourceVirtualEnvironmentVMStarted                           = true
	dvResourceVirtualEnvironmentVMTabletDevice                      = true
	dvResourceVirtualEnvironmentVMVGAEnabled                        = true
	dvResourceVirtualEnvironmentVMVGAMemory                         = 16
	dvResourceVirtualEnvironmentVMVGAType                           = "std"
	dvResourceVirtualEnvironmentVMVMID                              = -1

	mkResourceVirtualEnvironmentVMACPI                              = "acpi"
	mkResourceVirtualEnvironmentVMAgent                             = "agent"
	mkResourceVirtualEnvironmentVMAgentEnabled                      = "enabled"
	mkResourceVirtualEnvironmentVMAgentTrim                         = "trim"
	mkResourceVirtualEnvironmentVMAgentType                         = "type"
	mkResourceVirtualEnvironmentVMBIOS                              = "bios"
	mkResourceVirtualEnvironmentVMCDROM                             = "cdrom"
	mkResourceVirtualEnvironmentVMCDROMEnabled                      = "enabled"
	mkResourceVirtualEnvironmentVMCDROMFileID                       = "file_id"
	mkResourceVirtualEnvironmentVMCPU                               = "cpu"
	mkResourceVirtualEnvironmentVMCPUArchitecture                   = "architecture"
	mkResourceVirtualEnvironmentVMCPUCores                          = "cores"
	mkResourceVirtualEnvironmentVMCPUFlags                          = "flags"
	mkResourceVirtualEnvironmentVMCPUHotplugged                     = "hotplugged"
	mkResourceVirtualEnvironmentVMCPUSockets                        = "sockets"
	mkResourceVirtualEnvironmentVMCPUType                           = "type"
	mkResourceVirtualEnvironmentVMCPUUnits                          = "units"
	mkResourceVirtualEnvironmentVMDescription                       = "description"
	mkResourceVirtualEnvironmentVMDisk                              = "disk"
	mkResourceVirtualEnvironmentVMDiskDatastoreID                   = "datastore_id"
	mkResourceVirtualEnvironmentVMDiskFileFormat                    = "file_format"
	mkResourceVirtualEnvironmentVMDiskFileID                        = "file_id"
	mkResourceVirtualEnvironmentVMDiskSize                          = "size"
	mkResourceVirtualEnvironmentVMDiskSpeed                         = "speed"
	mkResourceVirtualEnvironmentVMDiskSpeedRead                     = "read"
	mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable            = "read_burstable"
	mkResourceVirtualEnvironmentVMDiskSpeedWrite                    = "write"
	mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable           = "write_burstable"
	mkResourceVirtualEnvironmentVMInitialization                    = "initialization"
	mkResourceVirtualEnvironmentVMInitializationDNS                 = "dns"
	mkResourceVirtualEnvironmentVMInitializationDNSDomain           = "domain"
	mkResourceVirtualEnvironmentVMInitializationDNSServer           = "server"
	mkResourceVirtualEnvironmentVMInitializationIPConfig            = "ip_config"
	mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4        = "ipv4"
	mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address = "address"
	mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway = "gateway"
	mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6        = "ipv6"
	mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6Address = "address"
	mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6Gateway = "gateway"
	mkResourceVirtualEnvironmentVMInitializationUserAccount         = "user_account"
	mkResourceVirtualEnvironmentVMInitializationUserAccountKeys     = "keys"
	mkResourceVirtualEnvironmentVMInitializationUserAccountPassword = "password"
	mkResourceVirtualEnvironmentVMInitializationUserAccountUsername = "username"
	mkResourceVirtualEnvironmentVMInitializationUserDataFileID      = "user_data_file_id"
	mkResourceVirtualEnvironmentVMIPv4Addresses                     = "ipv4_addresses"
	mkResourceVirtualEnvironmentVMIPv6Addresses                     = "ipv6_addresses"
	mkResourceVirtualEnvironmentVMKeyboardLayout                    = "keyboard_layout"
	mkResourceVirtualEnvironmentVMMACAddresses                      = "mac_addresses"
	mkResourceVirtualEnvironmentVMMemory                            = "memory"
	mkResourceVirtualEnvironmentVMMemoryDedicated                   = "dedicated"
	mkResourceVirtualEnvironmentVMMemoryFloating                    = "floating"
	mkResourceVirtualEnvironmentVMMemoryShared                      = "shared"
	mkResourceVirtualEnvironmentVMName                              = "name"
	mkResourceVirtualEnvironmentVMNetworkDevice                     = "network_device"
	mkResourceVirtualEnvironmentVMNetworkDeviceBridge               = "bridge"
	mkResourceVirtualEnvironmentVMNetworkDeviceEnabled              = "enabled"
	mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress           = "mac_address"
	mkResourceVirtualEnvironmentVMNetworkDeviceModel                = "model"
	mkResourceVirtualEnvironmentVMNetworkDeviceRateLimit            = "rate_limit"
	mkResourceVirtualEnvironmentVMNetworkDeviceVLANID               = "vlan_id"
	mkResourceVirtualEnvironmentVMNetworkInterfaceNames             = "network_interface_names"
	mkResourceVirtualEnvironmentVMNodeName                          = "node_name"
	mkResourceVirtualEnvironmentVMOperatingSystem                   = "operating_system"
	mkResourceVirtualEnvironmentVMOperatingSystemType               = "type"
	mkResourceVirtualEnvironmentVMPoolID                            = "pool_id"
	mkResourceVirtualEnvironmentVMStarted                           = "started"
	mkResourceVirtualEnvironmentVMTabletDevice                      = "tablet_device"
	mkResourceVirtualEnvironmentVMVGA                               = "vga"
	mkResourceVirtualEnvironmentVMVGAEnabled                        = "enabled"
	mkResourceVirtualEnvironmentVMVGAMemory                         = "memory"
	mkResourceVirtualEnvironmentVMVGAType                           = "type"
	mkResourceVirtualEnvironmentVMVMID                              = "vm_id"
)

func resourceVirtualEnvironmentVM() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkResourceVirtualEnvironmentVMACPI: {
				Type:        schema.TypeBool,
				Description: "Whether to enable ACPI",
				Optional:    true,
				Default:     dvResourceVirtualEnvironmentVMACPI,
			},
			mkResourceVirtualEnvironmentVMAgent: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The QEMU agent configuration",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMAgentEnabled] = dvResourceVirtualEnvironmentVMAgentEnabled
					defaultMap[mkResourceVirtualEnvironmentVMAgentTrim] = dvResourceVirtualEnvironmentVMAgentTrim
					defaultMap[mkResourceVirtualEnvironmentVMAgentType] = dvResourceVirtualEnvironmentVMAgentType

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMAgentEnabled: {
							Type:        schema.TypeBool,
							Description: "Whether to enable the QEMU agent",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMAgentEnabled,
						},
						mkResourceVirtualEnvironmentVMAgentTrim: {
							Type:        schema.TypeBool,
							Description: "Whether to enable the FSTRIM feature in the QEMU agent",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMAgentTrim,
						},
						mkResourceVirtualEnvironmentVMAgentType: {
							Type:         schema.TypeString,
							Description:  "The QEMU agent interface type",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMAgentType,
							ValidateFunc: getQEMUAgentTypeValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMBIOS: {
				Type:         schema.TypeString,
				Description:  "The BIOS implementation",
				Optional:     true,
				Default:      dvResourceVirtualEnvironmentVMBIOS,
				ValidateFunc: getBIOSValidator(),
			},
			mkResourceVirtualEnvironmentVMCDROM: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The CDROM drive",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMCDROMEnabled] = dvResourceVirtualEnvironmentVMCDROMEnabled
					defaultMap[mkResourceVirtualEnvironmentVMCDROMFileID] = dvResourceVirtualEnvironmentVMCDROMFileID

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMCDROMEnabled: {
							Type:        schema.TypeBool,
							Description: "Whether to enable the CDROM drive",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMCDROMEnabled,
						},
						mkResourceVirtualEnvironmentVMCDROMFileID: {
							Type:         schema.TypeString,
							Description:  "The file id",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCDROMFileID,
							ValidateFunc: getFileIDValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMCPU: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The CPU allocation",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMCPUArchitecture] = dvResourceVirtualEnvironmentVMCPUArchitecture
					defaultMap[mkResourceVirtualEnvironmentVMCPUCores] = dvResourceVirtualEnvironmentVMCPUCores
					defaultMap[mkResourceVirtualEnvironmentVMCPUFlags] = []interface{}{}
					defaultMap[mkResourceVirtualEnvironmentVMCPUHotplugged] = dvResourceVirtualEnvironmentVMCPUHotplugged
					defaultMap[mkResourceVirtualEnvironmentVMCPUSockets] = dvResourceVirtualEnvironmentVMCPUSockets
					defaultMap[mkResourceVirtualEnvironmentVMCPUType] = dvResourceVirtualEnvironmentVMCPUType
					defaultMap[mkResourceVirtualEnvironmentVMCPUUnits] = dvResourceVirtualEnvironmentVMCPUUnits

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMCPUArchitecture: {
							Type:         schema.TypeString,
							Description:  "The CPU architecture",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCPUArchitecture,
							ValidateFunc: resourceVirtualEnvironmentVMGetCPUArchitectureValidator(),
						},
						mkResourceVirtualEnvironmentVMCPUCores: {
							Type:         schema.TypeInt,
							Description:  "The number of CPU cores",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCPUCores,
							ValidateFunc: validation.IntBetween(1, 2304),
						},
						mkResourceVirtualEnvironmentVMCPUFlags: {
							Type:        schema.TypeList,
							Description: "The CPU flags",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								return []interface{}{}, nil
							},
							Elem: &schema.Schema{Type: schema.TypeString},
						},
						mkResourceVirtualEnvironmentVMCPUHotplugged: {
							Type:         schema.TypeInt,
							Description:  "The number of hotplugged vCPUs",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCPUHotplugged,
							ValidateFunc: validation.IntBetween(0, 2304),
						},
						mkResourceVirtualEnvironmentVMCPUSockets: {
							Type:         schema.TypeInt,
							Description:  "The number of CPU sockets",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCPUSockets,
							ValidateFunc: validation.IntBetween(1, 16),
						},
						mkResourceVirtualEnvironmentVMCPUType: {
							Type:         schema.TypeString,
							Description:  "The emulated CPU type",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCPUType,
							ValidateFunc: getCPUTypeValidator(),
						},
						mkResourceVirtualEnvironmentVMCPUUnits: {
							Type:         schema.TypeInt,
							Description:  "The CPU units",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMCPUUnits,
							ValidateFunc: validation.IntBetween(2, 262144),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMDescription: {
				Type:        schema.TypeString,
				Description: "The description",
				Optional:    true,
				Default:     dvResourceVirtualEnvironmentVMDescription,
			},
			mkResourceVirtualEnvironmentVMDisk: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The disk devices",
				Optional:    true,
				ForceNew:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMDiskDatastoreID] = dvResourceVirtualEnvironmentVMDiskDatastoreID
					defaultMap[mkResourceVirtualEnvironmentVMDiskFileFormat] = dvResourceVirtualEnvironmentVMDiskFileFormat
					defaultMap[mkResourceVirtualEnvironmentVMDiskFileID] = dvResourceVirtualEnvironmentVMDiskFileID
					defaultMap[mkResourceVirtualEnvironmentVMDiskSize] = dvResourceVirtualEnvironmentVMDiskSize

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMDiskDatastoreID: {
							Type:        schema.TypeString,
							Description: "The datastore id",
							Optional:    true,
							ForceNew:    true,
							Default:     dvResourceVirtualEnvironmentVMDiskDatastoreID,
						},
						mkResourceVirtualEnvironmentVMDiskFileFormat: {
							Type:         schema.TypeString,
							Description:  "The file format",
							Optional:     true,
							ForceNew:     true,
							Default:      dvResourceVirtualEnvironmentVMDiskFileFormat,
							ValidateFunc: getFileFormatValidator(),
						},
						mkResourceVirtualEnvironmentVMDiskFileID: {
							Type:         schema.TypeString,
							Description:  "The file id for a disk image",
							Optional:     true,
							ForceNew:     true,
							Default:      dvResourceVirtualEnvironmentVMDiskFileID,
							ValidateFunc: getFileIDValidator(),
						},
						mkResourceVirtualEnvironmentVMDiskSize: {
							Type:         schema.TypeInt,
							Description:  "The disk size in gigabytes",
							Optional:     true,
							ForceNew:     true,
							Default:      dvResourceVirtualEnvironmentVMDiskSize,
							ValidateFunc: validation.IntBetween(1, 8192),
						},
						mkResourceVirtualEnvironmentVMDiskSpeed: {
							Type:        schema.TypeList,
							Description: "The speed limits",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								defaultList := make([]interface{}, 1)
								defaultMap := map[string]interface{}{}

								defaultMap[mkResourceVirtualEnvironmentVMDiskSpeedRead] = dvResourceVirtualEnvironmentVMDiskSpeedRead
								defaultMap[mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable] = dvResourceVirtualEnvironmentVMDiskSpeedReadBurstable
								defaultMap[mkResourceVirtualEnvironmentVMDiskSpeedWrite] = dvResourceVirtualEnvironmentVMDiskSpeedWrite
								defaultMap[mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable] = dvResourceVirtualEnvironmentVMDiskSpeedWriteBurstable

								defaultList[0] = defaultMap

								return defaultList, nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMDiskSpeedRead: {
										Type:        schema.TypeInt,
										Description: "The maximum read speed in megabytes per second",
										Optional:    true,
										Default:     dvResourceVirtualEnvironmentVMDiskSpeedRead,
									},
									mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable: {
										Type:        schema.TypeInt,
										Description: "The maximum burstable read speed in megabytes per second",
										Optional:    true,
										Default:     dvResourceVirtualEnvironmentVMDiskSpeedReadBurstable,
									},
									mkResourceVirtualEnvironmentVMDiskSpeedWrite: {
										Type:        schema.TypeInt,
										Description: "The maximum write speed in megabytes per second",
										Optional:    true,
										Default:     dvResourceVirtualEnvironmentVMDiskSpeedWrite,
									},
									mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable: {
										Type:        schema.TypeInt,
										Description: "The maximum burstable write speed in megabytes per second",
										Optional:    true,
										Default:     dvResourceVirtualEnvironmentVMDiskSpeedWriteBurstable,
									},
								},
							},
							MaxItems: 1,
							MinItems: 0,
						},
					},
				},
				MaxItems: 14,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMInitialization: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The cloud-init configuration",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{}, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMInitializationDNS: {
							Type:        schema.TypeList,
							Description: "The DNS configuration",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								return []interface{}{}, nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMInitializationDNSDomain: {
										Type:        schema.TypeString,
										Description: "The DNS search domain",
										Optional:    true,
										Default:     dvResourceVirtualEnvironmentVMInitializationDNSDomain,
									},
									mkResourceVirtualEnvironmentVMInitializationDNSServer: {
										Type:        schema.TypeString,
										Description: "The DNS server",
										Optional:    true,
										Default:     dvResourceVirtualEnvironmentVMInitializationDNSServer,
									},
								},
							},
							MaxItems: 1,
							MinItems: 0,
						},
						mkResourceVirtualEnvironmentVMInitializationIPConfig: {
							Type:        schema.TypeList,
							Description: "The IP configuration",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								return []interface{}{}, nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4: {
										Type:        schema.TypeList,
										Description: "The IPv4 configuration",
										Optional:    true,
										DefaultFunc: func() (interface{}, error) {
											return []interface{}{}, nil
										},
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address: {
													Type:        schema.TypeString,
													Description: "The IPv4 address",
													Optional:    true,
													Default:     dvResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address,
												},
												mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway: {
													Type:        schema.TypeString,
													Description: "The IPv4 gateway",
													Optional:    true,
													Default:     dvResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway,
												},
											},
										},
										MaxItems: 1,
										MinItems: 0,
									},
									mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6: {
										Type:        schema.TypeList,
										Description: "The IPv6 configuration",
										Optional:    true,
										DefaultFunc: func() (interface{}, error) {
											return []interface{}{}, nil
										},
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6Address: {
													Type:        schema.TypeString,
													Description: "The IPv6 address",
													Optional:    true,
													Default:     dvResourceVirtualEnvironmentVMInitializationIPConfigIPv6Address,
												},
												mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6Gateway: {
													Type:        schema.TypeString,
													Description: "The IPv6 gateway",
													Optional:    true,
													Default:     dvResourceVirtualEnvironmentVMInitializationIPConfigIPv6Gateway,
												},
											},
										},
										MaxItems: 1,
										MinItems: 0,
									},
								},
							},
							MaxItems: 8,
							MinItems: 0,
						},
						mkResourceVirtualEnvironmentVMInitializationUserAccount: {
							Type:        schema.TypeList,
							Description: "The user account configuration",
							Optional:    true,
							ForceNew:    true,
							DefaultFunc: func() (interface{}, error) {
								return []interface{}{}, nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMInitializationUserAccountKeys: {
										Type:        schema.TypeList,
										Description: "The SSH keys",
										Optional:    true,
										ForceNew:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									mkResourceVirtualEnvironmentVMInitializationUserAccountPassword: {
										Type:        schema.TypeString,
										Description: "The SSH password",
										Optional:    true,
										ForceNew:    true,
										Sensitive:   true,
										Default:     dvResourceVirtualEnvironmentVMInitializationUserAccountPassword,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											return len(old) > 0 && strings.ReplaceAll(old, "*", "") == ""
										},
									},
									mkResourceVirtualEnvironmentVMInitializationUserAccountUsername: {
										Type:        schema.TypeString,
										Description: "The SSH username",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							MaxItems: 1,
							MinItems: 0,
						},
						mkResourceVirtualEnvironmentVMInitializationUserDataFileID: {
							Type:         schema.TypeString,
							Description:  "The ID of a file containing custom user data",
							Optional:     true,
							ForceNew:     true,
							Default:      dvResourceVirtualEnvironmentVMInitializationUserDataFileID,
							ValidateFunc: getFileIDValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMIPv4Addresses: {
				Type:        schema.TypeList,
				Description: "The IPv4 addresses published by the QEMU agent",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			mkResourceVirtualEnvironmentVMIPv6Addresses: {
				Type:        schema.TypeList,
				Description: "The IPv6 addresses published by the QEMU agent",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			mkResourceVirtualEnvironmentVMKeyboardLayout: {
				Type:         schema.TypeString,
				Description:  "The keyboard layout",
				Optional:     true,
				Default:      dvResourceVirtualEnvironmentVMKeyboardLayout,
				ValidateFunc: getKeyboardLayoutValidator(),
			},
			mkResourceVirtualEnvironmentVMMACAddresses: {
				Type:        schema.TypeList,
				Description: "The MAC addresses for the network interfaces",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkResourceVirtualEnvironmentVMMemory: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The memory allocation",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMMemoryDedicated] = dvResourceVirtualEnvironmentVMMemoryDedicated
					defaultMap[mkResourceVirtualEnvironmentVMMemoryFloating] = dvResourceVirtualEnvironmentVMMemoryFloating
					defaultMap[mkResourceVirtualEnvironmentVMMemoryShared] = dvResourceVirtualEnvironmentVMMemoryShared

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMMemoryDedicated: {
							Type:         schema.TypeInt,
							Description:  "The dedicated memory in megabytes",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMMemoryDedicated,
							ValidateFunc: validation.IntBetween(64, 268435456),
						},
						mkResourceVirtualEnvironmentVMMemoryFloating: {
							Type:         schema.TypeInt,
							Description:  "The floating memory in megabytes (balloon)",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMMemoryFloating,
							ValidateFunc: validation.IntBetween(0, 268435456),
						},
						mkResourceVirtualEnvironmentVMMemoryShared: {
							Type:         schema.TypeInt,
							Description:  "The shared memory in megabytes",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMMemoryShared,
							ValidateFunc: validation.IntBetween(0, 268435456),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMName: {
				Type:        schema.TypeString,
				Description: "The name",
				Optional:    true,
				Default:     dvResourceVirtualEnvironmentVMName,
			},
			mkResourceVirtualEnvironmentVMNetworkDevice: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The network devices",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return make([]interface{}, 1), nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMNetworkDeviceBridge: {
							Type:        schema.TypeString,
							Description: "The bridge",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceBridge,
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceEnabled: {
							Type:        schema.TypeBool,
							Description: "Whether to enable the network device",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceEnabled,
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress: {
							Type:        schema.TypeString,
							Description: "The MAC address",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceMACAddress,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
							ValidateFunc: getMACAddressValidator(),
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceModel: {
							Type:         schema.TypeString,
							Description:  "The model",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMNetworkDeviceModel,
							ValidateFunc: getNetworkDeviceModelValidator(),
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceRateLimit: {
							Type:        schema.TypeFloat,
							Description: "The rate limit in megabytes per second",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceRateLimit,
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceVLANID: {
							Type:        schema.TypeInt,
							Description: "The VLAN identifier",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceVLANID,
						},
					},
				},
				MaxItems: 8,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMNetworkInterfaceNames: {
				Type:        schema.TypeList,
				Description: "The network interface names published by the QEMU agent",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkResourceVirtualEnvironmentVMNodeName: &schema.Schema{
				Type:        schema.TypeString,
				Description: "The node name",
				Required:    true,
				ForceNew:    true,
			},
			mkResourceVirtualEnvironmentVMOperatingSystem: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The operating system configuration",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMOperatingSystemType] = dvResourceVirtualEnvironmentVMOperatingSystemType

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMOperatingSystemType: {
							Type:         schema.TypeString,
							Description:  "The type",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMOperatingSystemType,
							ValidateFunc: resourceVirtualEnvironmentVMGetOperatingSystemTypeValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMPoolID: {
				Type:        schema.TypeString,
				Description: "The ID of the pool to assign the virtual machine to",
				Optional:    true,
				ForceNew:    true,
				Default:     dvResourceVirtualEnvironmentVMPoolID,
			},
			mkResourceVirtualEnvironmentVMStarted: {
				Type:        schema.TypeBool,
				Description: "Whether to start the virtual machine",
				Optional:    true,
				Default:     dvResourceVirtualEnvironmentVMStarted,
			},
			mkResourceVirtualEnvironmentVMTabletDevice: {
				Type:        schema.TypeBool,
				Description: "Whether to enable the USB tablet device",
				Optional:    true,
				Default:     dvResourceVirtualEnvironmentVMTabletDevice,
			},
			mkResourceVirtualEnvironmentVMVGA: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The VGA configuration",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := map[string]interface{}{}

					defaultMap[mkResourceVirtualEnvironmentVMVGAEnabled] = dvResourceVirtualEnvironmentVMVGAEnabled
					defaultMap[mkResourceVirtualEnvironmentVMVGAMemory] = dvResourceVirtualEnvironmentVMVGAMemory
					defaultMap[mkResourceVirtualEnvironmentVMVGAType] = dvResourceVirtualEnvironmentVMVGAType

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMVGAEnabled: {
							Type:        schema.TypeBool,
							Description: "Whether to enable the VGA device",
							Optional:    true,
							Default:     dvResourceVirtualEnvironmentVMVGAEnabled,
						},
						mkResourceVirtualEnvironmentVMVGAMemory: {
							Type:         schema.TypeInt,
							Description:  "The VGA memory in megabytes (4-512 MB)",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMVGAMemory,
							ValidateFunc: getVGAMemoryValidator(),
						},
						mkResourceVirtualEnvironmentVMVGAType: {
							Type:         schema.TypeString,
							Description:  "The VGA type",
							Optional:     true,
							Default:      dvResourceVirtualEnvironmentVMVGAType,
							ValidateFunc: getVGATypeValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMVMID: {
				Type:         schema.TypeInt,
				Description:  "The VM identifier",
				Optional:     true,
				ForceNew:     true,
				Default:      dvResourceVirtualEnvironmentVMVMID,
				ValidateFunc: getVMIDValidator(),
			},
		},
		Create: resourceVirtualEnvironmentVMCreate,
		Read:   resourceVirtualEnvironmentVMRead,
		Update: resourceVirtualEnvironmentVMUpdate,
		Delete: resourceVirtualEnvironmentVMDelete,
	}
}

func resourceVirtualEnvironmentVMCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	resource := resourceVirtualEnvironmentVM()

	acpi := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentVMACPI).(bool))

	agentBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMAgent}, 0, true)

	if err != nil {
		return err
	}

	agentEnabled := proxmox.CustomBool(agentBlock[mkResourceVirtualEnvironmentVMAgentEnabled].(bool))
	agentTrim := proxmox.CustomBool(agentBlock[mkResourceVirtualEnvironmentVMAgentTrim].(bool))
	agentType := agentBlock[mkResourceVirtualEnvironmentVMAgentType].(string)

	bios := d.Get(mkResourceVirtualEnvironmentVMBIOS).(string)

	cdromBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMCDROM}, 0, true)

	if err != nil {
		return err
	}

	cdromEnabled := cdromBlock[mkResourceVirtualEnvironmentVMCDROMEnabled].(bool)
	cdromFileID := cdromBlock[mkResourceVirtualEnvironmentVMCDROMFileID].(string)

	if cdromFileID == "" {
		cdromFileID = "cdrom"
	}

	cpuBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMCPU}, 0, true)

	if err != nil {
		return err
	}

	cpuArchitecture := cpuBlock[mkResourceVirtualEnvironmentVMCPUArchitecture].(string)
	cpuCores := cpuBlock[mkResourceVirtualEnvironmentVMCPUCores].(int)
	cpuFlags := cpuBlock[mkResourceVirtualEnvironmentVMCPUFlags].([]interface{})
	cpuHotplugged := cpuBlock[mkResourceVirtualEnvironmentVMCPUHotplugged].(int)
	cpuSockets := cpuBlock[mkResourceVirtualEnvironmentVMCPUSockets].(int)
	cpuType := cpuBlock[mkResourceVirtualEnvironmentVMCPUType].(string)
	cpuUnits := cpuBlock[mkResourceVirtualEnvironmentVMCPUUnits].(int)

	description := d.Get(mkResourceVirtualEnvironmentVMDescription).(string)
	diskDeviceObjects, err := resourceVirtualEnvironmentVMGetDiskDeviceObjects(d, m)

	if err != nil {
		return err
	}

	initializationConfig, err := resourceVirtualEnvironmentVMGetCloudInitConfig(d, m)

	if err != nil {
		return err
	}

	if initializationConfig != nil {
		cdromEnabled = true
		cdromFileID = "local-lvm:cloudinit"
	}

	keyboardLayout := d.Get(mkResourceVirtualEnvironmentVMKeyboardLayout).(string)
	memoryBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMMemory}, 0, true)

	if err != nil {
		return err
	}

	memoryDedicated := memoryBlock[mkResourceVirtualEnvironmentVMMemoryDedicated].(int)
	memoryFloating := memoryBlock[mkResourceVirtualEnvironmentVMMemoryFloating].(int)
	memoryShared := memoryBlock[mkResourceVirtualEnvironmentVMMemoryShared].(int)

	name := d.Get(mkResourceVirtualEnvironmentVMName).(string)

	networkDeviceObjects, err := resourceVirtualEnvironmentVMGetNetworkDeviceObjects(d, m)

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)

	operatingSystem, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMOperatingSystem}, 0, true)

	if err != nil {
		return err
	}

	operatingSystemType := operatingSystem[mkResourceVirtualEnvironmentVMOperatingSystemType].(string)

	poolID := d.Get(mkResourceVirtualEnvironmentVMPoolID).(string)
	started := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentVMStarted).(bool))
	tabletDevice := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentVMTabletDevice).(bool))

	vgaDevice, err := resourceVirtualEnvironmentVMGetVGADeviceObject(d, m)

	if err != nil {
		return err
	}

	vmID := d.Get(mkResourceVirtualEnvironmentVMVMID).(int)

	if vmID == -1 {
		vmIDNew, err := veClient.GetVMID()

		if err != nil {
			return err
		}

		vmID = *vmIDNew
	}

	var memorySharedObject *proxmox.CustomSharedMemory

	bootDisk := "scsi0"
	bootOrder := "c"

	if cdromEnabled {
		bootOrder = "cd"
	}

	cpuFlagsConverted := make([]string, len(cpuFlags))

	for fi, flag := range cpuFlags {
		cpuFlagsConverted[fi] = flag.(string)
	}

	ideDevice2Media := "cdrom"
	ideDevices := proxmox.CustomStorageDevices{
		proxmox.CustomStorageDevice{
			Enabled: false,
		},
		proxmox.CustomStorageDevice{
			Enabled: false,
		},
		proxmox.CustomStorageDevice{
			Enabled:    cdromEnabled,
			FileVolume: cdromFileID,
			Media:      &ideDevice2Media,
		},
	}

	if memoryShared > 0 {
		memorySharedName := fmt.Sprintf("vm-%d-ivshmem", vmID)
		memorySharedObject = &proxmox.CustomSharedMemory{
			Name: &memorySharedName,
			Size: memoryShared,
		}
	}

	scsiHardware := "virtio-scsi-pci"

	body := &proxmox.VirtualEnvironmentVMCreateRequestBody{
		ACPI: &acpi,
		Agent: &proxmox.CustomAgent{
			Enabled:         &agentEnabled,
			TrimClonedDisks: &agentTrim,
			Type:            &agentType,
		},
		BIOS:            &bios,
		BootDisk:        &bootDisk,
		BootOrder:       &bootOrder,
		CloudInitConfig: initializationConfig,
		CPUArchitecture: &cpuArchitecture,
		CPUCores:        &cpuCores,
		CPUEmulation: &proxmox.CustomCPUEmulation{
			Flags: &cpuFlagsConverted,
			Type:  cpuType,
		},
		CPUSockets:          &cpuSockets,
		CPUUnits:            &cpuUnits,
		DedicatedMemory:     &memoryDedicated,
		FloatingMemory:      &memoryFloating,
		IDEDevices:          ideDevices,
		KeyboardLayout:      &keyboardLayout,
		NetworkDevices:      networkDeviceObjects,
		OSType:              &operatingSystemType,
		PoolID:              &poolID,
		SCSIDevices:         diskDeviceObjects,
		SCSIHardware:        &scsiHardware,
		SerialDevices:       []string{"socket"},
		SharedMemory:        memorySharedObject,
		StartOnBoot:         &started,
		TabletDeviceEnabled: &tabletDevice,
		VGADevice:           vgaDevice,
		VMID:                &vmID,
	}

	if cpuHotplugged > 0 {
		body.VirtualCPUCount = &cpuHotplugged
	}

	if description != "" {
		body.Description = &description
	}

	if name != "" {
		body.Name = &name
	}

	err = veClient.CreateVM(nodeName, body)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(vmID))

	return resourceVirtualEnvironmentVMCreateImportedDisks(d, m)
}

func resourceVirtualEnvironmentVMCreateImportedDisks(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	commands := []string{}

	// Determine the ID of the next disk.
	disk := d.Get(mkResourceVirtualEnvironmentVMDisk).([]interface{})
	diskCount := 0

	for _, d := range disk {
		block := d.(map[string]interface{})
		fileID, _ := block[mkResourceVirtualEnvironmentVMDiskFileID].(string)

		if fileID == "" {
			diskCount++
		}
	}

	// Retrieve some information about the disk schema.
	resourceSchema := resourceVirtualEnvironmentVM().Schema
	diskSchemaElem := resourceSchema[mkResourceVirtualEnvironmentVMDisk].Elem
	diskSchemaResource := diskSchemaElem.(*schema.Resource)
	diskSpeedResource := diskSchemaResource.Schema[mkResourceVirtualEnvironmentVMDiskSpeed]

	// Generate the commands required to import the specified disks.
	importedDiskCount := 0

	for i, d := range disk {
		block := d.(map[string]interface{})

		fileID, _ := block[mkResourceVirtualEnvironmentVMDiskFileID].(string)

		if fileID == "" {
			continue
		}

		datastoreID, _ := block[mkResourceVirtualEnvironmentVMDiskDatastoreID].(string)
		fileFormat, _ := block[mkResourceVirtualEnvironmentVMDiskFileFormat].(string)
		size, _ := block[mkResourceVirtualEnvironmentVMDiskSize].(int)
		speed := block[mkResourceVirtualEnvironmentVMDiskSpeed].([]interface{})

		if len(speed) == 0 {
			diskSpeedDefault, err := diskSpeedResource.DefaultValue()

			if err != nil {
				return err
			}

			speed = diskSpeedDefault.([]interface{})
		}

		speedBlock := speed[0].(map[string]interface{})
		speedLimitRead := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedRead].(int)
		speedLimitReadBurstable := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable].(int)
		speedLimitWrite := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedWrite].(int)
		speedLimitWriteBurstable := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable].(int)

		diskOptions := ""

		if speedLimitRead > 0 {
			diskOptions += fmt.Sprintf(",mbps_rd=%d", speedLimitRead)
		}

		if speedLimitReadBurstable > 0 {
			diskOptions += fmt.Sprintf(",mbps_rd_max=%d", speedLimitReadBurstable)
		}

		if speedLimitWrite > 0 {
			diskOptions += fmt.Sprintf(",mbps_wr=%d", speedLimitWrite)
		}

		if speedLimitWriteBurstable > 0 {
			diskOptions += fmt.Sprintf(",mbps_wr_max=%d", speedLimitWriteBurstable)
		}

		fileIDParts := strings.Split(fileID, ":")
		filePath := ""

		if strings.HasPrefix(fileIDParts[1], "iso/") {
			filePath = fmt.Sprintf("/template/%s", fileIDParts[1])
		} else {
			filePath = fmt.Sprintf("/%s", fileIDParts[1])
		}

		filePathTmp := fmt.Sprintf("/tmp/vm-%d-disk-%d.%s", vmID, diskCount+importedDiskCount, fileFormat)

		commands = append(
			commands,
			`set -e`,
			fmt.Sprintf(`cp "$(grep -Pzo ': %s\s+path\s+[^\s]+' /etc/pve/storage.cfg | grep -Pzo '/[^\s]*' | tr -d '\000')%s" %s`, fileIDParts[0], filePath, filePathTmp),
			fmt.Sprintf(`qemu-img resize %s %dG`, filePathTmp, size),
			fmt.Sprintf(`qm importdisk %d %s %s -format qcow2`, vmID, filePathTmp, datastoreID),
			fmt.Sprintf(`qm set %d -scsi%d %s:vm-%d-disk-%d%s`, vmID, i, datastoreID, vmID, diskCount+importedDiskCount, diskOptions),
			fmt.Sprintf(`rm -f %s`, filePathTmp),
		)

		importedDiskCount++
	}

	// Execute the commands on the node and wait for the result.
	// This is a highly experimental approach to disk imports and is not recommended by Proxmox.
	if len(commands) > 0 {
		err = veClient.ExecuteNodeCommands(nodeName, commands)

		if err != nil {
			return err
		}
	}

	return resourceVirtualEnvironmentVMCreateStart(d, m)
}

func resourceVirtualEnvironmentVMCreateStart(d *schema.ResourceData, m interface{}) error {
	started := d.Get(mkResourceVirtualEnvironmentVMStarted).(bool)

	if !started {
		return resourceVirtualEnvironmentVMRead(d, m)
	}

	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	// Start the virtual machine and wait for it to reach a running state before continuing.
	err = veClient.StartVM(nodeName, vmID)

	if err != nil {
		return err
	}

	err = veClient.WaitForVMState(nodeName, vmID, "running", 120, 5)

	if err != nil {
		return err
	}

	return resourceVirtualEnvironmentVMRead(d, m)
}

func resourceVirtualEnvironmentVMGetCloudInitConfig(d *schema.ResourceData, m interface{}) (*proxmox.CustomCloudInitConfig, error) {
	var initializationConfig *proxmox.CustomCloudInitConfig

	initialization := d.Get(mkResourceVirtualEnvironmentVMInitialization).([]interface{})

	if len(initialization) > 0 {
		initializationBlock := initialization[0].(map[string]interface{})
		initializationConfig = &proxmox.CustomCloudInitConfig{}
		initializationDNS := initializationBlock[mkResourceVirtualEnvironmentVMInitializationDNS].([]interface{})

		if len(initializationDNS) > 0 {
			initializationDNSBlock := initializationDNS[0].(map[string]interface{})
			domain := initializationDNSBlock[mkResourceVirtualEnvironmentVMInitializationDNSDomain].(string)

			if domain != "" {
				initializationConfig.SearchDomain = &domain
			}

			server := initializationDNSBlock[mkResourceVirtualEnvironmentVMInitializationDNSServer].(string)

			if server != "" {
				initializationConfig.Nameserver = &server
			}
		}

		initializationIPConfig := initializationBlock[mkResourceVirtualEnvironmentVMInitializationIPConfig].([]interface{})
		initializationConfig.IPConfig = make([]proxmox.CustomCloudInitIPConfig, len(initializationIPConfig))

		for i, c := range initializationIPConfig {
			configBlock := c.(map[string]interface{})
			ipv4 := configBlock[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4].([]interface{})

			if len(ipv4) > 0 {
				ipv4Block := ipv4[0].(map[string]interface{})
				ipv4Address := ipv4Block[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address].(string)

				if ipv4Address != "" {
					initializationConfig.IPConfig[i].IPv4 = &ipv4Address
				}

				ipv4Gateway := ipv4Block[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway].(string)

				if ipv4Gateway != "" {
					initializationConfig.IPConfig[i].GatewayIPv4 = &ipv4Gateway
				}
			}

			ipv6 := configBlock[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6].([]interface{})

			if len(ipv6) > 0 {
				ipv6Block := ipv6[0].(map[string]interface{})
				ipv6Address := ipv6Block[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6Address].(string)

				if ipv6Address != "" {
					initializationConfig.IPConfig[i].IPv6 = &ipv6Address
				}

				ipv6Gateway := ipv6Block[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6Gateway].(string)

				if ipv6Gateway != "" {
					initializationConfig.IPConfig[i].GatewayIPv6 = &ipv6Gateway
				}
			}
		}

		initializationUserAccount := initializationBlock[mkResourceVirtualEnvironmentVMInitializationUserAccount].([]interface{})

		if len(initializationUserAccount) > 0 {
			initializationUserAccountBlock := initializationUserAccount[0].(map[string]interface{})
			keys := initializationUserAccountBlock[mkResourceVirtualEnvironmentVMInitializationUserAccountKeys].([]interface{})

			if len(keys) > 0 {
				sshKeys := make(proxmox.CustomCloudInitSSHKeys, len(keys))

				for i, k := range keys {
					sshKeys[i] = k.(string)
				}

				initializationConfig.SSHKeys = &sshKeys
			}

			password := initializationUserAccountBlock[mkResourceVirtualEnvironmentVMInitializationUserAccountPassword].(string)

			if password != "" {
				initializationConfig.Password = &password
			}

			username := initializationUserAccountBlock[mkResourceVirtualEnvironmentVMInitializationUserAccountUsername].(string)

			initializationConfig.Username = &username
		}

		initializationUserDataFileID := initializationBlock[mkResourceVirtualEnvironmentVMInitializationUserDataFileID].(string)

		if initializationUserDataFileID != "" {
			initializationConfig.Files = &proxmox.CustomCloudInitFiles{
				UserVolume: &initializationUserDataFileID,
			}
		}
	}

	return initializationConfig, nil
}

func resourceVirtualEnvironmentVMGetCPUArchitectureValidator() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		"aarch64",
		"x86_64",
	}, false)
}

func resourceVirtualEnvironmentVMGetDiskDeviceObjects(d *schema.ResourceData, m interface{}) (proxmox.CustomStorageDevices, error) {
	diskDevice := d.Get(mkResourceVirtualEnvironmentVMDisk).([]interface{})
	diskDeviceObjects := make(proxmox.CustomStorageDevices, len(diskDevice))
	resource := resourceVirtualEnvironmentVM()

	for i, diskEntry := range diskDevice {
		diskDevice := proxmox.CustomStorageDevice{
			Enabled: true,
		}

		block := diskEntry.(map[string]interface{})
		datastoreID, _ := block[mkResourceVirtualEnvironmentVMDiskDatastoreID].(string)
		fileID, _ := block[mkResourceVirtualEnvironmentVMDiskFileID].(string)
		size, _ := block[mkResourceVirtualEnvironmentVMDiskSize].(int)

		speedBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMDisk, mkResourceVirtualEnvironmentVMDiskSpeed}, 0, false)

		if err != nil {
			return diskDeviceObjects, err
		}

		if fileID != "" {
			diskDevice.Enabled = false
		} else {
			diskDevice.FileVolume = fmt.Sprintf("%s:%d", datastoreID, size)
		}

		if len(speedBlock) > 0 {
			speedLimitRead := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedRead].(int)
			speedLimitReadBurstable := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable].(int)
			speedLimitWrite := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedWrite].(int)
			speedLimitWriteBurstable := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable].(int)

			if speedLimitRead > 0 {
				diskDevice.MaxReadSpeedMbps = &speedLimitRead
			}

			if speedLimitReadBurstable > 0 {
				diskDevice.BurstableReadSpeedMbps = &speedLimitReadBurstable
			}

			if speedLimitWrite > 0 {
				diskDevice.MaxWriteSpeedMbps = &speedLimitWrite
			}

			if speedLimitWriteBurstable > 0 {
				diskDevice.BurstableWriteSpeedMbps = &speedLimitWriteBurstable
			}
		}

		diskDeviceObjects[i] = diskDevice
	}

	return diskDeviceObjects, nil
}

func resourceVirtualEnvironmentVMGetNetworkDeviceObjects(d *schema.ResourceData, m interface{}) (proxmox.CustomNetworkDevices, error) {
	networkDevice := d.Get(mkResourceVirtualEnvironmentVMNetworkDevice).([]interface{})
	networkDeviceObjects := make(proxmox.CustomNetworkDevices, len(networkDevice))

	for i, networkDeviceEntry := range networkDevice {
		block := networkDeviceEntry.(map[string]interface{})

		bridge, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceBridge].(string)
		enabled, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceEnabled].(bool)
		macAddress, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress].(string)
		model, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceModel].(string)
		rateLimit, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceRateLimit].(float64)
		vlanID, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceVLANID].(int)

		device := proxmox.CustomNetworkDevice{
			Enabled: enabled,
			Model:   model,
		}

		if bridge != "" {
			device.Bridge = &bridge
		}

		if macAddress != "" {
			device.MACAddress = &macAddress
		}

		if rateLimit != 0 {
			device.RateLimit = &rateLimit
		}

		if vlanID != 0 {
			device.Tag = &vlanID
		}

		networkDeviceObjects[i] = device
	}

	return networkDeviceObjects, nil
}

func resourceVirtualEnvironmentVMGetOperatingSystemTypeValidator() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		"l24",
		"l26",
		"other",
		"solaris",
		"w2k",
		"w2k3",
		"w2k8",
		"win7",
		"win8",
		"win10",
		"wvista",
		"wxp",
	}, false)
}

func resourceVirtualEnvironmentVMGetVGADeviceObject(d *schema.ResourceData, m interface{}) (*proxmox.CustomVGADevice, error) {
	resource := resourceVirtualEnvironmentVM()

	vgaBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMVGA}, 0, true)

	if err != nil {
		return nil, err
	}

	vgaEnabled := proxmox.CustomBool(vgaBlock[mkResourceVirtualEnvironmentVMAgentEnabled].(bool))
	vgaMemory := vgaBlock[mkResourceVirtualEnvironmentVMVGAMemory].(int)
	vgaType := vgaBlock[mkResourceVirtualEnvironmentVMVGAType].(string)

	vgaDevice := &proxmox.CustomVGADevice{}

	if vgaEnabled {
		if vgaMemory > 0 {
			vgaDevice.Memory = &vgaMemory
		}

		vgaDevice.Type = &vgaType
	} else {
		vgaType = "none"

		vgaDevice = &proxmox.CustomVGADevice{
			Type: &vgaType,
		}
	}

	return vgaDevice, nil
}

func resourceVirtualEnvironmentVMRead(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	// Retrieve the entire configuration in order to compare it to the state.
	vmConfig, err := veClient.GetVM(nodeName, vmID)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP 404") ||
			(strings.Contains(err.Error(), "HTTP 500") && strings.Contains(err.Error(), "does not exist")) {
			d.SetId("")

			return nil
		}

		return err
	}

	// Compare some primitive arguments to the values stored in the state.
	if vmConfig.ACPI != nil {
		d.Set(mkResourceVirtualEnvironmentVMACPI, bool(*vmConfig.ACPI))
	} else {
		// Default value of "acpi" is "1" according to the API documentation.
		d.Set(mkResourceVirtualEnvironmentVMACPI, true)
	}

	if vmConfig.BIOS != nil {
		d.Set(mkResourceVirtualEnvironmentVMBIOS, *vmConfig.BIOS)
	} else {
		// Default value of "bios" is "seabios" according to the API documentation.
		d.Set(mkResourceVirtualEnvironmentVMBIOS, "seabios")
	}

	if vmConfig.Description != nil {
		d.Set(mkResourceVirtualEnvironmentVMDescription, *vmConfig.Description)
	} else {
		d.Set(mkResourceVirtualEnvironmentVMDescription, "")
	}

	if vmConfig.KeyboardLayout != nil {
		d.Set(mkResourceVirtualEnvironmentVMKeyboardLayout, *vmConfig.KeyboardLayout)
	} else {
		d.Set(mkResourceVirtualEnvironmentVMKeyboardLayout, "")
	}

	if vmConfig.TabletDeviceEnabled != nil {
		d.Set(mkResourceVirtualEnvironmentVMTabletDevice, bool(*vmConfig.TabletDeviceEnabled))
	} else {
		// Default value of "tablet" is "1" according to the API documentation.
		d.Set(mkResourceVirtualEnvironmentVMTabletDevice, true)
	}

	// Compare the agent configuration to the one stored in the state.
	if vmConfig.Agent != nil {
		agent := map[string]interface{}{}

		if vmConfig.Agent.Enabled != nil {
			agent[mkResourceVirtualEnvironmentVMAgentEnabled] = bool(*vmConfig.Agent.Enabled)
		} else {
			agent[mkResourceVirtualEnvironmentVMAgentEnabled] = false
		}

		if vmConfig.Agent.TrimClonedDisks != nil {
			agent[mkResourceVirtualEnvironmentVMAgentTrim] = bool(*vmConfig.Agent.TrimClonedDisks)
		} else {
			agent[mkResourceVirtualEnvironmentVMAgentTrim] = false
		}

		if vmConfig.Agent.Type != nil {
			agent[mkResourceVirtualEnvironmentVMAgentType] = *vmConfig.Agent.Type
		} else {
			agent[mkResourceVirtualEnvironmentVMAgentType] = ""
		}

		currentAgent := d.Get(mkResourceVirtualEnvironmentVMAgent).([]interface{})

		if len(currentAgent) > 0 ||
			agent[mkResourceVirtualEnvironmentVMAgentEnabled] != dvResourceVirtualEnvironmentVMAgentEnabled ||
			agent[mkResourceVirtualEnvironmentVMAgentTrim] != dvResourceVirtualEnvironmentVMAgentTrim ||
			agent[mkResourceVirtualEnvironmentVMAgentType] != dvResourceVirtualEnvironmentVMAgentType {
			d.Set(mkResourceVirtualEnvironmentVMAgent, []interface{}{agent})
		}
	} else {
		d.Set(mkResourceVirtualEnvironmentVMAgent, []interface{}{})
	}

	// Compare the IDE devices to the CDROM and cloud-init configurations stored in the state.
	if vmConfig.IDEDevice2 != nil {
		if *vmConfig.IDEDevice2.Media == "cdrom" {
			if strings.Contains(vmConfig.IDEDevice2.FileVolume, fmt.Sprintf("vm-%d-cloudinit", vmID)) {
				d.Set(mkResourceVirtualEnvironmentVMCDROM, []interface{}{})
			} else {
				d.Set(mkResourceVirtualEnvironmentVMInitialization, []interface{}{})

				cdrom := make([]interface{}, 1)
				cdromBlock := map[string]interface{}{}

				cdromBlock[mkResourceVirtualEnvironmentVMCDROMEnabled] = true
				cdromBlock[mkResourceVirtualEnvironmentVMCDROMFileID] = vmConfig.IDEDevice2.FileVolume

				cdrom[0] = cdromBlock

				d.Set(mkResourceVirtualEnvironmentVMCDROM, cdrom)
			}
		} else {
			d.Set(mkResourceVirtualEnvironmentVMCDROM, []interface{}{})
			d.Set(mkResourceVirtualEnvironmentVMInitialization, []interface{}{})
		}
	} else {
		d.Set(mkResourceVirtualEnvironmentVMCDROM, []interface{}{})
		d.Set(mkResourceVirtualEnvironmentVMInitialization, []interface{}{})
	}

	// Compare the CPU configuration to the one stored in the state.
	cpu := map[string]interface{}{}

	if vmConfig.CPUArchitecture != nil {
		cpu[mkResourceVirtualEnvironmentVMCPUArchitecture] = *vmConfig.CPUArchitecture
	} else {
		// Default value of "arch" is "" according to the API documentation.
		cpu[mkResourceVirtualEnvironmentVMCPUArchitecture] = ""
	}

	if vmConfig.CPUCores != nil {
		cpu[mkResourceVirtualEnvironmentVMCPUCores] = *vmConfig.CPUCores
	} else {
		// Default value of "cores" is "1" according to the API documentation.
		cpu[mkResourceVirtualEnvironmentVMCPUCores] = 1
	}

	if vmConfig.VirtualCPUCount != nil {
		cpu[mkResourceVirtualEnvironmentVMCPUHotplugged] = *vmConfig.VirtualCPUCount
	} else {
		// Default value of "vcpus" is "1" according to the API documentation.
		cpu[mkResourceVirtualEnvironmentVMCPUHotplugged] = 0
	}

	if vmConfig.CPUSockets != nil {
		cpu[mkResourceVirtualEnvironmentVMCPUSockets] = *vmConfig.CPUSockets
	} else {
		// Default value of "sockets" is "1" according to the API documentation.
		cpu[mkResourceVirtualEnvironmentVMCPUSockets] = 1
	}

	if vmConfig.CPUEmulation != nil {
		if vmConfig.CPUEmulation.Flags != nil {
			convertedFlags := make([]interface{}, len(*vmConfig.CPUEmulation.Flags))

			for fi, fv := range *vmConfig.CPUEmulation.Flags {
				convertedFlags[fi] = fv
			}

			cpu[mkResourceVirtualEnvironmentVMCPUFlags] = convertedFlags
		} else {
			cpu[mkResourceVirtualEnvironmentVMCPUFlags] = []interface{}{}
		}

		cpu[mkResourceVirtualEnvironmentVMCPUType] = vmConfig.CPUEmulation.Type
	} else {
		cpu[mkResourceVirtualEnvironmentVMCPUFlags] = []interface{}{}
		// Default value of "cputype" is "qemu64" according to the QEMU documentation.
		cpu[mkResourceVirtualEnvironmentVMCPUType] = "qemu64"
	}

	if vmConfig.CPUUnits != nil {
		cpu[mkResourceVirtualEnvironmentVMCPUUnits] = *vmConfig.CPUUnits
	} else {
		// Default value of "cpuunits" is "1024" according to the API documentation.
		cpu[mkResourceVirtualEnvironmentVMCPUUnits] = 1024
	}

	currentCPU := d.Get(mkResourceVirtualEnvironmentVMCPU).([]interface{})

	if len(currentCPU) > 0 ||
		cpu[mkResourceVirtualEnvironmentVMCPUArchitecture] != dvResourceVirtualEnvironmentVMCPUArchitecture ||
		cpu[mkResourceVirtualEnvironmentVMCPUCores] != dvResourceVirtualEnvironmentVMCPUCores ||
		len(cpu[mkResourceVirtualEnvironmentVMCPUFlags].([]interface{})) > 0 ||
		cpu[mkResourceVirtualEnvironmentVMCPUHotplugged] != dvResourceVirtualEnvironmentVMCPUHotplugged ||
		cpu[mkResourceVirtualEnvironmentVMCPUSockets] != dvResourceVirtualEnvironmentVMCPUSockets ||
		cpu[mkResourceVirtualEnvironmentVMCPUType] != dvResourceVirtualEnvironmentVMCPUType ||
		cpu[mkResourceVirtualEnvironmentVMCPUUnits] != dvResourceVirtualEnvironmentVMCPUUnits {
		d.Set(mkResourceVirtualEnvironmentVMCPU, []interface{}{cpu})
	}

	// Compare the disks to those stored in the state.
	currentDiskList := d.Get(mkResourceVirtualEnvironmentVMDisk).([]interface{})

	diskList := []interface{}{}
	diskObjects := []*proxmox.CustomStorageDevice{
		vmConfig.SCSIDevice0,
		vmConfig.SCSIDevice1,
		vmConfig.SCSIDevice2,
		vmConfig.SCSIDevice3,
		vmConfig.SCSIDevice4,
		vmConfig.SCSIDevice5,
		vmConfig.SCSIDevice6,
		vmConfig.SCSIDevice7,
		vmConfig.SCSIDevice8,
		vmConfig.SCSIDevice9,
		vmConfig.SCSIDevice10,
		vmConfig.SCSIDevice11,
		vmConfig.SCSIDevice12,
		vmConfig.SCSIDevice13,
	}

	for di, dd := range diskObjects {
		disk := map[string]interface{}{}

		if dd == nil {
			continue
		}

		fileIDParts := strings.Split(dd.FileVolume, ":")

		disk[mkResourceVirtualEnvironmentVMDiskDatastoreID] = fileIDParts[0]

		if len(currentDiskList) > di {
			currentDisk := currentDiskList[di].(map[string]interface{})

			disk[mkResourceVirtualEnvironmentVMDiskFileFormat] = currentDisk[mkResourceVirtualEnvironmentVMDiskFileFormat]
			disk[mkResourceVirtualEnvironmentVMDiskFileID] = currentDisk[mkResourceVirtualEnvironmentVMDiskFileID]
		}

		diskSize := 0

		if dd.Size != nil {
			if strings.HasSuffix(*dd.Size, "T") {
				diskSize, err = strconv.Atoi(strings.TrimSuffix(*dd.Size, "T"))

				if err != nil {
					return err
				}

				diskSize = int(math.Ceil(float64(diskSize) * 1024))
			} else if strings.HasSuffix(*dd.Size, "G") {
				diskSize, err = strconv.Atoi(strings.TrimSuffix(*dd.Size, "G"))

				if err != nil {
					return err
				}
			} else if strings.HasSuffix(*dd.Size, "M") {
				diskSize, err = strconv.Atoi(strings.TrimSuffix(*dd.Size, "M"))

				if err != nil {
					return err
				}

				diskSize = int(math.Ceil(float64(diskSize) / 1024))
			} else {
				return fmt.Errorf("Cannot parse storage size \"%s\"", *dd.Size)
			}
		}

		disk[mkResourceVirtualEnvironmentVMDiskSize] = diskSize

		if dd.BurstableReadSpeedMbps != nil ||
			dd.BurstableWriteSpeedMbps != nil ||
			dd.MaxReadSpeedMbps != nil ||
			dd.MaxWriteSpeedMbps != nil {
			speed := map[string]interface{}{}

			if dd.MaxReadSpeedMbps != nil {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedRead] = *dd.MaxReadSpeedMbps
			} else {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedRead] = 0
			}

			if dd.BurstableReadSpeedMbps != nil {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable] = *dd.BurstableReadSpeedMbps
			} else {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedReadBurstable] = 0
			}

			if dd.MaxWriteSpeedMbps != nil {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedWrite] = *dd.MaxWriteSpeedMbps
			} else {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedWrite] = 0
			}

			if dd.BurstableWriteSpeedMbps != nil {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable] = *dd.BurstableWriteSpeedMbps
			} else {
				speed[mkResourceVirtualEnvironmentVMDiskSpeedWriteBurstable] = 0
			}

			disk[mkResourceVirtualEnvironmentVMDiskSpeed] = []interface{}{speed}
		} else {
			disk[mkResourceVirtualEnvironmentVMDiskSpeed] = []interface{}{}
		}

		diskList = append(diskList, disk)
	}

	if len(currentDiskList) > 0 || len(diskList) > 0 {
		d.Set(mkResourceVirtualEnvironmentVMDisk, diskList)
	}

	// Compare the cloud-init configuration to the one stored in the state.
	initialization := map[string]interface{}{}

	if vmConfig.CloudInitDNSDomain != nil || vmConfig.CloudInitDNSServer != nil {
		initializationDNS := map[string]interface{}{}

		if vmConfig.CloudInitDNSDomain != nil {
			initializationDNS[mkResourceVirtualEnvironmentVMInitializationDNSDomain] = *vmConfig.CloudInitDNSDomain
		} else {
			initializationDNS[mkResourceVirtualEnvironmentVMInitializationDNSDomain] = ""
		}

		if vmConfig.CloudInitDNSServer != nil {
			initializationDNS[mkResourceVirtualEnvironmentVMInitializationDNSServer] = *vmConfig.CloudInitDNSServer
		} else {
			initializationDNS[mkResourceVirtualEnvironmentVMInitializationDNSServer] = ""
		}

		initialization[mkResourceVirtualEnvironmentVMInitializationDNS] = []interface{}{initializationDNS}
	}

	ipConfigLast := -1
	ipConfigObjects := []*proxmox.CustomCloudInitIPConfig{
		vmConfig.IPConfig0,
		vmConfig.IPConfig1,
		vmConfig.IPConfig2,
		vmConfig.IPConfig3,
		vmConfig.IPConfig4,
		vmConfig.IPConfig5,
		vmConfig.IPConfig6,
		vmConfig.IPConfig7,
	}
	ipConfigList := make([]interface{}, len(ipConfigObjects))

	for ipConfigIndex, ipConfig := range ipConfigObjects {
		ipConfigItem := map[string]interface{}{}

		if ipConfig != nil {
			ipConfigLast = ipConfigIndex

			if ipConfig.GatewayIPv4 != nil || ipConfig.IPv4 != nil {
				ipv4 := map[string]interface{}{}

				if ipConfig.IPv4 != nil {
					ipv4[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address] = *ipConfig.IPv4
				} else {
					ipv4[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address] = ""
				}

				if ipConfig.GatewayIPv4 != nil {
					ipv4[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway] = *ipConfig.GatewayIPv4
				} else {
					ipv4[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway] = ""
				}

				ipConfigItem[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4] = []interface{}{ipv4}
			} else {
				ipConfigItem[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4] = []interface{}{}
			}

			if ipConfig.GatewayIPv6 != nil || ipConfig.IPv6 != nil {
				ipv6 := map[string]interface{}{}

				if ipConfig.IPv4 != nil {
					ipv6[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address] = *ipConfig.IPv6
				} else {
					ipv6[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Address] = ""
				}

				if ipConfig.GatewayIPv4 != nil {
					ipv6[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway] = *ipConfig.GatewayIPv6
				} else {
					ipv6[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4Gateway] = ""
				}

				ipConfigItem[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6] = []interface{}{ipv6}
			} else {
				ipConfigItem[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6] = []interface{}{}
			}
		} else {
			ipConfigItem[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv4] = []interface{}{}
			ipConfigItem[mkResourceVirtualEnvironmentVMInitializationIPConfigIPv6] = []interface{}{}
		}

		ipConfigList[ipConfigIndex] = ipConfigItem
	}

	initialization[mkResourceVirtualEnvironmentVMInitializationIPConfig] = ipConfigList[:ipConfigLast+1]

	if vmConfig.CloudInitPassword != nil || vmConfig.CloudInitSSHKeys != nil || vmConfig.CloudInitUsername != nil {
		initializationUserAccount := map[string]interface{}{}

		if vmConfig.CloudInitSSHKeys != nil {
			initializationUserAccount[mkResourceVirtualEnvironmentVMInitializationUserAccountKeys] = []string(*vmConfig.CloudInitSSHKeys)
		} else {
			initializationUserAccount[mkResourceVirtualEnvironmentVMInitializationUserAccountKeys] = []string{}
		}

		if vmConfig.CloudInitPassword != nil {
			initializationUserAccount[mkResourceVirtualEnvironmentVMInitializationUserAccountPassword] = *vmConfig.CloudInitPassword
		} else {
			initializationUserAccount[mkResourceVirtualEnvironmentVMInitializationUserAccountPassword] = ""
		}

		if vmConfig.CloudInitUsername != nil {
			initializationUserAccount[mkResourceVirtualEnvironmentVMInitializationUserAccountUsername] = *vmConfig.CloudInitUsername
		} else {
			initializationUserAccount[mkResourceVirtualEnvironmentVMInitializationUserAccountUsername] = ""
		}

		initialization[mkResourceVirtualEnvironmentVMInitializationUserAccount] = []interface{}{initializationUserAccount}
	}

	if vmConfig.CloudInitFiles != nil {
		if vmConfig.CloudInitFiles.UserVolume != nil {
			initialization[mkResourceVirtualEnvironmentVMInitializationUserDataFileID] = *vmConfig.CloudInitFiles.UserVolume
		} else {
			initialization[mkResourceVirtualEnvironmentVMInitializationUserDataFileID] = ""
		}
	} else {
		initialization[mkResourceVirtualEnvironmentVMInitializationUserDataFileID] = ""
	}

	if len(initialization) > 0 {
		d.Set(mkResourceVirtualEnvironmentVMInitialization, []interface{}{initialization})
	} else {
		d.Set(mkResourceVirtualEnvironmentVMInitialization, []interface{}{})
	}

	// Compare the memory configuration to the one stored in the state.
	memory := map[string]interface{}{}

	if vmConfig.DedicatedMemory != nil {
		memory[mkResourceVirtualEnvironmentVMMemoryDedicated] = *vmConfig.DedicatedMemory
	} else {
		memory[mkResourceVirtualEnvironmentVMMemoryDedicated] = 0
	}

	if vmConfig.FloatingMemory != nil {
		memory[mkResourceVirtualEnvironmentVMMemoryFloating] = *vmConfig.FloatingMemory
	} else {
		memory[mkResourceVirtualEnvironmentVMMemoryFloating] = 0
	}

	if vmConfig.SharedMemory != nil {
		memory[mkResourceVirtualEnvironmentVMMemoryShared] = vmConfig.SharedMemory.Size
	} else {
		memory[mkResourceVirtualEnvironmentVMMemoryShared] = 0
	}

	currentMemory := d.Get(mkResourceVirtualEnvironmentVMMemory).([]interface{})

	if len(currentMemory) > 0 ||
		memory[mkResourceVirtualEnvironmentVMMemoryDedicated] != dvResourceVirtualEnvironmentVMMemoryDedicated ||
		memory[mkResourceVirtualEnvironmentVMMemoryFloating] != dvResourceVirtualEnvironmentVMMemoryFloating ||
		memory[mkResourceVirtualEnvironmentVMMemoryShared] != dvResourceVirtualEnvironmentVMMemoryShared {
		d.Set(mkResourceVirtualEnvironmentVMMemory, []interface{}{memory})
	}

	// Compare the name to the value stored in the state.
	if vmConfig.Name != nil {
		d.Set(mkResourceVirtualEnvironmentVMName, *vmConfig.Name)
	} else {
		d.Set(mkResourceVirtualEnvironmentVMName, "")
	}

	// Compare the network devices to those stored in the state.
	currentNetworkDeviceList := d.Get(mkResourceVirtualEnvironmentVMNetworkDevice).([]interface{})

	macAddresses := make([]interface{}, 8)
	networkDeviceLast := -1
	networkDeviceList := make([]interface{}, 8)
	networkDeviceObjects := []*proxmox.CustomNetworkDevice{
		vmConfig.NetworkDevice0,
		vmConfig.NetworkDevice1,
		vmConfig.NetworkDevice2,
		vmConfig.NetworkDevice3,
		vmConfig.NetworkDevice4,
		vmConfig.NetworkDevice5,
		vmConfig.NetworkDevice6,
		vmConfig.NetworkDevice7,
	}

	for ni, nd := range networkDeviceObjects {
		networkDevice := map[string]interface{}{}

		if nd != nil {
			networkDeviceLast = ni

			if nd.Bridge != nil {
				networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceBridge] = *nd.Bridge
			} else {
				networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceBridge] = ""
			}

			networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceEnabled] = nd.Enabled

			if nd.MACAddress != nil {
				macAddresses[ni] = *nd.MACAddress
			} else {
				macAddresses[ni] = ""
			}

			networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress] = macAddresses[ni]
			networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceModel] = nd.Model

			if nd.RateLimit != nil {
				networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceRateLimit] = *nd.RateLimit
			} else {
				networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceRateLimit] = 0
			}

			if nd.Tag != nil {
				networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceVLANID] = nd.Tag
			} else {
				networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceVLANID] = 0
			}
		} else {
			macAddresses[ni] = ""
			networkDevice[mkResourceVirtualEnvironmentVMNetworkDeviceEnabled] = false
		}

		networkDeviceList[ni] = networkDevice
	}

	d.Set(mkResourceVirtualEnvironmentVMMACAddresses, macAddresses[0:len(currentNetworkDeviceList)])

	if len(currentNetworkDeviceList) > 0 || networkDeviceLast > -1 {
		d.Set(mkResourceVirtualEnvironmentVMNetworkDevice, networkDeviceList[:networkDeviceLast+1])
	}

	// Compare the operating system configuration to the one stored in the state.
	operatingSystem := map[string]interface{}{}

	if vmConfig.OSType != nil {
		operatingSystem[mkResourceVirtualEnvironmentVMOperatingSystemType] = *vmConfig.OSType
	} else {
		operatingSystem[mkResourceVirtualEnvironmentVMOperatingSystemType] = ""
	}

	currentOperatingSystem := d.Get(mkResourceVirtualEnvironmentVMOperatingSystem).([]interface{})

	if len(currentOperatingSystem) > 0 ||
		operatingSystem[mkResourceVirtualEnvironmentVMOperatingSystemType] != dvResourceVirtualEnvironmentVMOperatingSystemType {
		d.Set(mkResourceVirtualEnvironmentVMOperatingSystem, []interface{}{operatingSystem})
	} else {
		d.Set(mkResourceVirtualEnvironmentVMOperatingSystem, []interface{}{})
	}

	// Compare the pool ID to the values stored in the state.
	if vmConfig.PoolID != nil {
		d.Set(mkResourceVirtualEnvironmentVMPoolID, *vmConfig.PoolID)
	}

	// Compare the VGA configuration to the one stored in the state.
	vga := map[string]interface{}{}

	if vmConfig.VGADevice != nil {
		vgaEnabled := true

		if vmConfig.VGADevice.Type != nil {
			vgaEnabled = *vmConfig.VGADevice.Type != "none"
		}

		vga[mkResourceVirtualEnvironmentVMVGAEnabled] = vgaEnabled

		if vmConfig.VGADevice.Memory != nil {
			vga[mkResourceVirtualEnvironmentVMVGAMemory] = *vmConfig.VGADevice.Memory
		} else {
			vga[mkResourceVirtualEnvironmentVMVGAMemory] = 0
		}

		if vgaEnabled {
			if vmConfig.VGADevice.Type != nil {
				vga[mkResourceVirtualEnvironmentVMVGAType] = *vmConfig.VGADevice.Type
			} else {
				vga[mkResourceVirtualEnvironmentVMVGAType] = ""
			}
		}
	} else {
		vga[mkResourceVirtualEnvironmentVMVGAEnabled] = true
		vga[mkResourceVirtualEnvironmentVMVGAMemory] = 0
		vga[mkResourceVirtualEnvironmentVMVGAType] = ""
	}

	currentVGA := d.Get(mkResourceVirtualEnvironmentVMVGA).([]interface{})

	if len(currentVGA) > 0 ||
		vga[mkResourceVirtualEnvironmentVMVGAEnabled] != dvResourceVirtualEnvironmentVMVGAEnabled ||
		vga[mkResourceVirtualEnvironmentVMVGAMemory] != dvResourceVirtualEnvironmentVMVGAMemory ||
		vga[mkResourceVirtualEnvironmentVMVGAType] != dvResourceVirtualEnvironmentVMVGAType {
		d.Set(mkResourceVirtualEnvironmentVMVGA, []interface{}{vga})
	} else {
		d.Set(mkResourceVirtualEnvironmentVMVGA, []interface{}{})
	}

	// Determine the state of the virtual machine in order to update the "started" argument.
	status, err := veClient.GetVMStatus(nodeName, vmID)

	if err != nil {
		return err
	}

	d.Set(mkResourceVirtualEnvironmentVMStarted, status.Status == "running")

	// Populate the attributes that rely on the QEMU agent.
	ipv4Addresses := []interface{}{}
	ipv6Addresses := []interface{}{}
	networkInterfaceNames := []interface{}{}

	if vmConfig.Agent != nil && vmConfig.Agent.Enabled != nil && *vmConfig.Agent.Enabled {
		networkInterfaces, err := veClient.WaitForNetworkInterfacesFromVMAgent(nodeName, vmID, 1800, 5)

		if err == nil && networkInterfaces.Result != nil {
			ipv4Addresses = make([]interface{}, len(*networkInterfaces.Result))
			ipv6Addresses = make([]interface{}, len(*networkInterfaces.Result))
			macAddresses = make([]interface{}, len(*networkInterfaces.Result))
			networkInterfaceNames = make([]interface{}, len(*networkInterfaces.Result))

			for ri, rv := range *networkInterfaces.Result {
				rvIPv4Addresses := []interface{}{}
				rvIPv6Addresses := []interface{}{}

				for _, ip := range *rv.IPAddresses {
					switch ip.Type {
					case "ipv4":
						rvIPv4Addresses = append(rvIPv4Addresses, ip.Address)
					case "ipv6":
						rvIPv6Addresses = append(rvIPv6Addresses, ip.Address)
					}
				}

				ipv4Addresses[ri] = rvIPv4Addresses
				ipv6Addresses[ri] = rvIPv6Addresses
				macAddresses[ri] = strings.ToUpper(rv.MACAddress)
				networkInterfaceNames[ri] = rv.Name
			}
		}
	}

	d.Set(mkResourceVirtualEnvironmentVMIPv4Addresses, ipv4Addresses)
	d.Set(mkResourceVirtualEnvironmentVMIPv6Addresses, ipv6Addresses)
	d.Set(mkResourceVirtualEnvironmentVMMACAddresses, macAddresses)
	d.Set(mkResourceVirtualEnvironmentVMNetworkInterfaceNames, networkInterfaceNames)

	return nil
}

func resourceVirtualEnvironmentVMUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	rebootRequired := false

	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	body := &proxmox.VirtualEnvironmentVMUpdateRequestBody{
		IDEDevices: proxmox.CustomStorageDevices{
			proxmox.CustomStorageDevice{
				Enabled: false,
			},
			proxmox.CustomStorageDevice{
				Enabled: false,
			},
			proxmox.CustomStorageDevice{
				Enabled: false,
			},
		},
	}

	resource := resourceVirtualEnvironmentVM()

	// Retrieve the entire configuration as we need to process certain values.
	vmConfig, err := veClient.GetVM(nodeName, vmID)

	if err != nil {
		return err
	}

	// Prepare the new primitive configuration values.
	acpi := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentVMACPI).(bool))
	bios := d.Get(mkResourceVirtualEnvironmentVMBIOS).(string)
	delete := []string{}
	description := d.Get(mkResourceVirtualEnvironmentVMDescription).(string)
	keyboardLayout := d.Get(mkResourceVirtualEnvironmentVMKeyboardLayout).(string)
	name := d.Get(mkResourceVirtualEnvironmentVMName).(string)
	tabletDevice := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentVMTabletDevice).(bool))

	body.ACPI = &acpi
	body.BIOS = &bios

	if description != "" {
		body.Description = &description
	}

	body.KeyboardLayout = &keyboardLayout

	if name != "" {
		body.Name = &name
	}

	body.TabletDeviceEnabled = &tabletDevice

	if d.HasChange(mkResourceVirtualEnvironmentVMACPI) ||
		d.HasChange(mkResourceVirtualEnvironmentVMBIOS) ||
		d.HasChange(mkResourceVirtualEnvironmentVMKeyboardLayout) ||
		d.HasChange(mkResourceVirtualEnvironmentVMOperatingSystemType) ||
		d.HasChange(mkResourceVirtualEnvironmentVMTabletDevice) {
		rebootRequired = true
	}

	// Prepare the new agent configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMAgent) {
		agentBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMAgent}, 0, true)

		if err != nil {
			return err
		}

		agentEnabled := proxmox.CustomBool(agentBlock[mkResourceVirtualEnvironmentVMAgentEnabled].(bool))
		agentTrim := proxmox.CustomBool(agentBlock[mkResourceVirtualEnvironmentVMAgentTrim].(bool))
		agentType := agentBlock[mkResourceVirtualEnvironmentVMAgentType].(string)

		body.Agent = &proxmox.CustomAgent{
			Enabled:         &agentEnabled,
			TrimClonedDisks: &agentTrim,
			Type:            &agentType,
		}

		rebootRequired = true
	}

	// Prepare the new CDROM configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMCDROM) {
		cdromBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMCDROM}, 0, true)

		if err != nil {
			return err
		}

		cdromEnabled := cdromBlock[mkResourceVirtualEnvironmentVMCDROMEnabled].(bool)
		cdromFileID := cdromBlock[mkResourceVirtualEnvironmentVMCDROMFileID].(string)

		if cdromFileID == "" {
			cdromFileID = "cdrom"
		}

		cdromMedia := "cdrom"

		body.IDEDevices[2] = proxmox.CustomStorageDevice{
			Enabled:    cdromEnabled,
			FileVolume: cdromFileID,
			Media:      &cdromMedia,
		}
	}

	// Prepare the new CPU configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMCPU) {
		cpuBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMCPU}, 0, true)

		if err != nil {
			return err
		}

		cpuArchitecture := cpuBlock[mkResourceVirtualEnvironmentVMCPUArchitecture].(string)
		cpuCores := cpuBlock[mkResourceVirtualEnvironmentVMCPUCores].(int)
		cpuFlags := cpuBlock[mkResourceVirtualEnvironmentVMCPUFlags].([]interface{})
		cpuHotplugged := cpuBlock[mkResourceVirtualEnvironmentVMCPUHotplugged].(int)
		cpuSockets := cpuBlock[mkResourceVirtualEnvironmentVMCPUSockets].(int)
		cpuType := cpuBlock[mkResourceVirtualEnvironmentVMCPUType].(string)
		cpuUnits := cpuBlock[mkResourceVirtualEnvironmentVMCPUUnits].(int)

		body.CPUArchitecture = &cpuArchitecture
		body.CPUCores = &cpuCores
		body.CPUSockets = &cpuSockets
		body.CPUUnits = &cpuUnits

		if cpuHotplugged > 0 {
			body.VirtualCPUCount = &cpuHotplugged
		} else {
			delete = append(delete, "vcpus")
		}

		cpuFlagsConverted := make([]string, len(cpuFlags))

		for fi, flag := range cpuFlags {
			cpuFlagsConverted[fi] = flag.(string)
		}

		body.CPUEmulation = &proxmox.CustomCPUEmulation{
			Flags: &cpuFlagsConverted,
			Type:  cpuType,
		}

		rebootRequired = true
	}

	// Prepare the new disk device configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMDisk) {
		diskDeviceObjects, err := resourceVirtualEnvironmentVMGetDiskDeviceObjects(d, m)

		if err != nil {
			return err
		}

		scsiDevices := []*proxmox.CustomStorageDevice{
			vmConfig.SCSIDevice0,
			vmConfig.SCSIDevice1,
			vmConfig.SCSIDevice2,
			vmConfig.SCSIDevice3,
			vmConfig.SCSIDevice4,
			vmConfig.SCSIDevice5,
			vmConfig.SCSIDevice6,
			vmConfig.SCSIDevice7,
			vmConfig.SCSIDevice8,
			vmConfig.SCSIDevice9,
			vmConfig.SCSIDevice10,
			vmConfig.SCSIDevice11,
			vmConfig.SCSIDevice12,
			vmConfig.SCSIDevice13,
		}

		body.SCSIDevices = make(proxmox.CustomStorageDevices, len(diskDeviceObjects))

		for di, do := range diskDeviceObjects {
			if scsiDevices[di] == nil {
				return fmt.Errorf("Missing SCSI device %d (scsi%d)", di, di)
			}

			body.SCSIDevices[di] = *scsiDevices[di]
			body.SCSIDevices[di].BurstableReadSpeedMbps = do.BurstableReadSpeedMbps
			body.SCSIDevices[di].BurstableWriteSpeedMbps = do.BurstableWriteSpeedMbps
			body.SCSIDevices[di].MaxReadSpeedMbps = do.MaxReadSpeedMbps
			body.SCSIDevices[di].MaxWriteSpeedMbps = do.MaxWriteSpeedMbps
		}

		rebootRequired = true
	}

	// Prepare the new cloud-init configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMInitialization) {
		initializationConfig, err := resourceVirtualEnvironmentVMGetCloudInitConfig(d, m)

		if err != nil {
			return err
		}

		body.CloudInitConfig = initializationConfig

		if body.CloudInitConfig != nil {
			cdromMedia := "cdrom"

			body.IDEDevices[2] = proxmox.CustomStorageDevice{
				Enabled:    true,
				FileVolume: "local-lvm:cloudinit",
				Media:      &cdromMedia,
			}

			if vmConfig.IDEDevice2 != nil {
				if strings.Contains(vmConfig.IDEDevice2.FileVolume, fmt.Sprintf("vm-%d-cloudinit", vmID)) {
					body.IDEDevices[2].Enabled = false
				}
			}
		}

		rebootRequired = true
	}

	// Prepare the new memory configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMMemory) {
		memoryBlock, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMMemory}, 0, true)

		if err != nil {
			return err
		}

		memoryDedicated := memoryBlock[mkResourceVirtualEnvironmentVMMemoryDedicated].(int)
		memoryFloating := memoryBlock[mkResourceVirtualEnvironmentVMMemoryFloating].(int)
		memoryShared := memoryBlock[mkResourceVirtualEnvironmentVMMemoryShared].(int)

		body.DedicatedMemory = &memoryDedicated
		body.FloatingMemory = &memoryFloating

		if memoryShared > 0 {
			memorySharedName := fmt.Sprintf("vm-%d-ivshmem", vmID)

			body.SharedMemory = &proxmox.CustomSharedMemory{
				Name: &memorySharedName,
				Size: memoryShared,
			}
		}

		rebootRequired = true
	}

	// Prepare the new network device configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMNetworkDevice) {
		body.NetworkDevices, err = resourceVirtualEnvironmentVMGetNetworkDeviceObjects(d, m)

		if err != nil {
			return err
		}

		rebootRequired = true
	}

	// Prepare the new operating system configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMOperatingSystem) {
		operatingSystem, err := getSchemaBlock(resource, d, m, []string{mkResourceVirtualEnvironmentVMOperatingSystem}, 0, true)

		if err != nil {
			return err
		}

		operatingSystemType := operatingSystem[mkResourceVirtualEnvironmentVMOperatingSystemType].(string)

		body.OSType = &operatingSystemType

		rebootRequired = true
	}

	// Prepare the new VGA configuration.
	if d.HasChange(mkResourceVirtualEnvironmentVMVGA) {
		body.VGADevice, err = resourceVirtualEnvironmentVMGetVGADeviceObject(d, m)

		if err != nil {
			return err
		}

		rebootRequired = true
	}

	// Update the configuration now that everything has been prepared.
	body.Delete = delete

	err = veClient.UpdateVM(nodeName, vmID, body)

	if err != nil {
		return err
	}

	// Determine if the state of the virtual machine needs to be changed.
	if d.HasChange(mkResourceVirtualEnvironmentVMStarted) {
		started := d.Get(mkResourceVirtualEnvironmentVMStarted).(bool)

		if started {
			err = veClient.StartVM(nodeName, vmID)

			if err != nil {
				return err
			}

			err = veClient.WaitForVMState(nodeName, vmID, "running", 120, 5)

			if err != nil {
				return err
			}
		} else {
			forceStop := proxmox.CustomBool(true)
			shutdownTimeout := 300

			err = veClient.ShutdownVM(nodeName, vmID, &proxmox.VirtualEnvironmentVMShutdownRequestBody{
				ForceStop: &forceStop,
				Timeout:   &shutdownTimeout,
			})

			if err != nil {
				return err
			}

			err = veClient.WaitForVMState(nodeName, vmID, "stopped", 30, 5)

			if err != nil {
				return err
			}

			rebootRequired = false
		}
	}

	// As a final step in the update procedure, we might need to reboot the virtual machine.
	if rebootRequired {
		rebootTimeout := 300

		err = veClient.RebootVM(nodeName, vmID, &proxmox.VirtualEnvironmentVMRebootRequestBody{
			Timeout: &rebootTimeout,
		})

		if err != nil {
			return err
		}

		// Wait for the agent to unpublish the network interfaces, if it's enabled.
		if vmConfig.Agent != nil && vmConfig.Agent.Enabled != nil && *vmConfig.Agent.Enabled {
			err = veClient.WaitForNoNetworkInterfacesFromVMAgent(nodeName, vmID, 300, 5)

			if err != nil {
				return err
			}
		}
	}

	return resourceVirtualEnvironmentVMRead(d, m)
}

func resourceVirtualEnvironmentVMDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	// Shut down the virtual machine before deleting it.
	status, err := veClient.GetVMStatus(nodeName, vmID)

	if err != nil {
		return err
	}

	if status.Status != "stopped" {
		forceStop := proxmox.CustomBool(true)
		shutdownTimeout := 300

		err = veClient.ShutdownVM(nodeName, vmID, &proxmox.VirtualEnvironmentVMShutdownRequestBody{
			ForceStop: &forceStop,
			Timeout:   &shutdownTimeout,
		})

		if err != nil {
			return err
		}

		err = veClient.WaitForVMState(nodeName, vmID, "stopped", 30, 5)

		if err != nil {
			return err
		}
	}

	err = veClient.DeleteVM(nodeName, vmID)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP 404") ||
			(strings.Contains(err.Error(), "HTTP 500") && strings.Contains(err.Error(), "does not exist")) {
			d.SetId("")

			return nil
		}

		return err
	}

	// Wait for the state to become unavailable as that clearly indicates the destruction of the VM.
	err = veClient.WaitForVMState(nodeName, vmID, "", 60, 2)

	if err == nil {
		return fmt.Errorf("Failed to delete VM \"%d\"", vmID)
	}

	d.SetId("")

	return nil
}
