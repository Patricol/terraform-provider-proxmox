---
layout: page
title: Time
permalink: /data-sources/virtual-environment/time
nav_order: 11
parent: Virtual Environment Data Sources
grand_parent: Data Sources
---

# Data Source: Time

Retrieves the current time for a specific node.

## Example Usage

```
data "proxmox_virtual_environment_time" "first_node_time" {
  node_name = "first-node"
}
```

## Arguments Reference

* `node_name` - (Required) A node name.

## Attributes Reference

* `local_time` - The node's local time.
* `time_zone` - The node's time zone.
* `utc_time` - The node's local time formatted as UTC.
