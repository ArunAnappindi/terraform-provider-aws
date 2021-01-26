package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsCloudFrontCachePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsCloudFrontCachePolicyCreate,
		Read:   resourceAwsCloudFrontCachePolicyRead,
		Update: resourceAwsCloudFrontCachePolicyUpdate,
		Delete: resourceAwsCloudFrontCachePolicyDelete,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_ttl": {
				Type: schema.TypeInt,
				Optional: true,
				Deafult: 86400,
			},
			"max_ttl": {
				Type: schema.TypeInt,
				Optional: true,
				Default: 31536000,
			},
			"min_ttl": {
				Type: schema.TypeInt,
				Optional: true,
				Default: 0,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters_in_cache_key_and_forward_to_origin": {
				Type: schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: map[string]schema.Schema {
					"cookies_config": {
					}
				}
			}

		},
	}
}

func resourceAwsCloudFrontCachePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).cloudfrontconn

	cloudfront.CachePolicy
	request := &cloudfront.CreateCachePolicyInput{
		CachePolicyConfig: &cloudfront.CachePolicyConfig{
			Comment:    nil,
			DefaultTTL: nil,
			MaxTTL:     nil,
			MinTTL:     nil,
			Name:       nil,
			ParametersInCacheKeyAndForwardedToOrigin: &cloudfront.ParametersInCacheKeyAndForwardedToOrigin{
				CookiesConfig: &cloudfront.CachePolicyCookiesConfig{
					CookieBehavior: nil,
					Cookies: &cloudfront.CookieNames{
						Items:    nil,
						Quantity: nil,
					},
				},
				EnableAcceptEncodingBrotli: nil,
				EnableAcceptEncodingGzip:   nil,
				HeadersConfig: &cloudfront.CachePolicyHeadersConfig{
					HeaderBehavior: nil,
					Headers: &cloudfront.Headers{
						Items:    nil,
						Quantity: nil,
					},
				},
				QueryStringsConfig: &cloudfront.CachePolicyQueryStringsConfig{
					QueryStringBehavior: nil,
					QueryStrings: &cloudfront.QueryStringNames{
						Items:    nil,
						Quantity: nil,
					},
				},
			},
		},
	}



	return resourceAwsCloudFrontCachePolicyRead(d, meta)
}

func resourceAwsCloudFrontCachePolicyRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).cloudfrontconn
	request := &cloudfront.GetCachePolicyInput{
		Id: aws.String(d.Id()),
	}
	conn.GetCachePolicy

	return nil
}

func resourceAwsCloudFrontCachePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).cloudfrontconn

	return resourceAwsCloudFrontCachePolicyRead(d, meta)
}

func resourceAwsCloudFrontCachePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).cloudfrontconn

	request := &cloudfront.DeleteCachePolicyInput{
		Id:      aws.String(d.Id()),
		IfMatch: aws.String(d.Get("etag").(string)),
	}

	_, err := conn.DeleteCachePolicy(request)
	if err != nil {
		if isAWSErr(err, cloudfront.ErrCodeNoSuchCachePolicy, "") {
			return nil
		}
		return err
	}

	return nil
}
