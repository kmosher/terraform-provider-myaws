package myaws

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/builtin/providers/aws"
)

func Provider() terraform.ResourceProvider {
	provider := aws.Provider().(*schema.Provider)
	instance := provider.ResourcesMap["aws_instance"]
	ami := instance.Schema["ami"]
	ami.Required = false
	ami.Optional = true
	ami.Computed = true
	instance.Create = wrapCreate(instance.Create)
	return provider
}

func wrapCreate(create func (*schema.ResourceData, interface{}) error) func (*schema.ResourceData, interface{}) error {
	wrapped := func (d *schema.ResourceData, meta interface{}) error {
		if _, ok := d.GetOk("ami"); !ok {
			d.Set("ami", getAMI())
		}
		return create(d, meta)
	}
	return wrapped
}

func getAMI() string {
	return "ami-39501209"
}
