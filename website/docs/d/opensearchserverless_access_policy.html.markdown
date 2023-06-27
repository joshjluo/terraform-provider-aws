---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_access_policy"
description: |-
  Get information on an OpenSearch Serverless Access Policy.
---

# Data Source: aws_opensearchserverless_access_policy

Use this data source to get information about an AWS OpenSearch Serverless Access Policy.

## Example Usage

```terraform
data "aws_opensearchserverless_access_policy" "example" {
  name = "example-access-policy"
  type = "data"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the policy
* `type` - (Required) Type of access policy. Currently, the only supported value is `data`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Description of the access policy.
* `policy` - The JSON policy document without any whitespaces.
* `policy_version` - Version of the policy.