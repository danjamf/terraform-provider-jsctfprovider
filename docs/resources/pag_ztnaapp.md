---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jsc_pag_ztnaapp Resource - jsc"
subcategory: ""
description: |-
  
---

# jsc_pag_ztnaapp (Resource)



## Example Usage

```terraform
resource "jsc_pag_ztnaapp" "testztnaapp" {
  name           = "testPAG222"
  routingdnstype = "IPv4"
  routingtype    = "CUSTOM"
  routingid      = "4042"
  hostnames      = ["test234234.com", "test223423423423.com"]
  //securityriskcontrolthreshold = "HIGH"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the ZTNA App. Must be unique

### Optional

- `apptemplateid` (String) App Template ID (if applicable)
- `assignmentallusers` (Boolean) Assign ZTNA App to all users
- `assignmentgroups` (List of String) Group IDs to assign ZTNA App Policy
- `bareips` (List of String) List of bare ips - IPv4 only CIDR notation
- `categoryname` (String) Category Name - supported types are Adult, Advertising, App Counters, App Stores, Audio & Music, Browsers, Business & Industry, Cloud & File Storage, Communication, Content Servers, Custom, Entertainment, Extreme, Finance, Gambling, Games, Generative AI, Illegal, Lifestyle, Medical, Navigation, News & Sport, OS Updates, Productivity, Reference, Shopping, Social, Technology, Travel, Uncategorized, Video & Photo
- `hostnames` (List of String) List of hostnames. Must be unique across all Access Policies and App Templates
- `routingdnstype` (String) Routing IP DNS Resolution Type - IPv4 or IPv6 (default is IPv6)
- `routingid` (String) Routing ID - required when routingtype is CUSTOM. Otherwise must be omitted
- `routingtype` (String) Routing Type - DIRECT or CUSTOM
- `securitydevicemanagementbasedaccessenabled` (Boolean) Enable deviceManagementBasedAccess for ZTNA App Policy
- `securitydevicemanagementbasedaccessnotifications` (Boolean) Enable deviceManagementBasedAccess notifications for ZTNA App Policy
- `securitydohintegrationblocking` (Boolean) Enable DoH blocking for ZTNA App Policy
- `securitydohintegrationnotifications` (Boolean) Enable DoH notifications for ZTNA App Policy
- `securityriskcontrolenabled` (Boolean) Enable device risk security controls for ZTNA App policy
- `securityriskcontrolnotifications` (Boolean) Enable notificatons for device risk security controls
- `securityriskcontrolthreshold` (String) Risk level threshold (when enabled), options of HIGH, MEDIUM, LOW

### Read-Only

- `id` (String) The unique identifier of the ZTNA App datasource set from JSC
