---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jsc_pag_ztnaapp Data Source - jsc"
subcategory: ""
description: |-
---

# jsc_pag_ztnaapp (Data Source)

## Example Usage

```terraform
data "jsc_pag_ztnaapp" "ztnaapptest001" {
  name = "Login Redirect"
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `name` (String) The name of the ZTNA App

### Read-Only

- `apptemplateid` (String) App Template ID (if applicable)
- `bareips` (List of String) List of bare ips
- `categoryname` (String) Category Name
- `hostnames` (List of String) List of hostnames
- `id` (String) The unique identifier of the ZTNA App datasource set from JSC
- `routingdnstype` (String) Routing IP DNS Resolution Type
- `routingid` (String) Routing ID
- `routingtype` (String) Routing Type