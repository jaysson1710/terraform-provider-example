// resource_server.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "log"
	// "net/http"
	//"fmt"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"env": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type Article struct {
	Nombre string
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	env := d.Get("env").(string)
	resource := d.Get("resource").(string)

	// a := Article{
	// 	Nombre: "RG-"+env+"-" + resource,
	// }

	d.SetId("1")

	d.Set("name", "RG-"+env+"-" + resource)

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
