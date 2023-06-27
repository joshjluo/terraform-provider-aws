package opensearchserverless

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

// @SDKDataSource("aws_opensearchserverless_access_policy")
func DataSourceAccessPolicy() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceAccessPolicyRead,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(3, 32),
					validation.StringMatch(regexp.MustCompile(`^[a-z][a-z0-9-]+$`), `must start with any lower case letter and can include any lower case letter, number, or "-"`),
				),
			},
			"policy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: enum.Validate[types.AccessPolicyType](),
			},
		},
	}
}

func dataSourceAccessPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).OpenSearchServerlessClient(ctx)

	accessPolicyName := d.Get("name").(string)
	accessPolicyType := d.Get("type").(string)
	accessPolicy, err := FindAccessPolicyByNameAndType(ctx, conn, accessPolicyName, accessPolicyType)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading AccessPolicy with name (%s) and type (%s): %s", accessPolicyName, accessPolicyType, err)
	}

	policyBytes, err := accessPolicy.Policy.MarshalSmithyDocument()
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading JSON policy document for AccessPolicy with name %s and type %s: %s", accessPolicyName, accessPolicyType, err)
	}

	d.SetId(aws.ToString(accessPolicy.Name))
	d.Set("description", accessPolicy.Description)
	d.Set("name", accessPolicy.Name)
	d.Set("policy", string(policyBytes))
	d.Set("policy_version", accessPolicy.PolicyVersion)
	d.Set("type", accessPolicy.Type)

	return diags
}
