//
// Automatically generated. DO NOT EDIT.
// (Object)
//

package resources

import (
	"fmt"
	"log"
	"strings"

	"github.com/Juniper/contrail-go-api"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
)

var _ = spew.Dump // Avoid import errors if not used

func SetBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetBgpRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("id_perms"); ok {
		member := new(IdPermsType)
		SetIdPermsTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetIdPerms(member)
	}
	if val, ok := d.GetOk("perms2"); ok {
		member := new(PermType2)
		SetPermType2FromMap(member, d, m, (val.([]interface{}))[0])
		object.SetPerms2(member)
	}
	if val, ok := d.GetOk("annotations"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetAnnotations(member)
	}
	if val, ok := d.GetOk("display_name"); ok {
		object.SetDisplayName(val.(string))
	}

}

func SetRefsBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsBgpRouterFromResource] key = %v, prefix = %v", key, prefix)

	return nil
}

func WriteBgpRouterToResource(object BgpRouter, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeBgpRouterAsMap(object *BgpRouter) map[string]interface{} {
	omap := make(map[string]interface{})

	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("id_perms") {
		if val, ok := d.GetOk("id_perms"); ok {
			member := new(IdPermsType)
			SetIdPermsTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetIdPerms(member)
		}
	}
	if d.HasChange("perms2") {
		if val, ok := d.GetOk("perms2"); ok {
			member := new(PermType2)
			SetPermType2FromMap(member, d, m, (val.([]interface{}))[0])
			object.SetPerms2(member)
		}
	}
	if d.HasChange("annotations") {
		if val, ok := d.GetOk("annotations"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetAnnotations(member)
		}
	}
	if d.HasChange("display_name") {
		if val, ok := d.GetOk("display_name"); ok {
			object.SetDisplayName(val.(string))
		}
	}

}

func ResourceBgpRouterCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpRouterCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(BgpRouter)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceBgpRouterCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "BgpRouter", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetBgpRouterFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceBgpRouterCreate] Creation of resource BgpRouter on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceBgpRouterRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpRouterRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Missing 'uuid' field for resource BgpRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgp-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Retrieving BgpRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpRouter := obj.(*BgpRouter) // Fully set by Contrail backend
	if err := SetRefsBgpRouterFromResource(objBgpRouter, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Set refs on object BgpRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpRouter.GetHref())
	if err := client.Update(objBgpRouter); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Update refs for resource BgpRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpRouter.GetUuid())
	return nil
}

func ResourceBgpRouterRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgp-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRead] Read resource bgp-router on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BgpRouter)
	WriteBgpRouterToResource(*object, d, m)
	return nil
}

func ResourceBgpRouterRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterRefsREAD")
	return nil
}

func ResourceBgpRouterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgp-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterResourceUpdate] Retrieving BgpRouter with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BgpRouter)
	UpdateBgpRouterFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpRouterUpdate] Update of resource BgpRouter on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpRouterRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterRefsUpdate")
	return nil
}

func ResourceBgpRouterDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("bgp-router", d.Id()); err != nil {
		return fmt.Errorf("[ResourceBgpRouterDelete] Deletion of resource bgp-router on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceBgpRouterRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterRefsDelete: %v", d.Id())
	return nil
}

func ResourceBgpRouterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"parent_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		// Properties
		"id_perms": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIdPermsType(),
		},
		"perms2": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePermType2(),
		},
		"annotations": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceKeyValuePairs(),
		},
		"display_name": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceBgpRouter() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpRouterCreate,
		Read:   ResourceBgpRouterRead,
		Update: ResourceBgpRouterUpdate,
		Delete: ResourceBgpRouterDelete,
		Schema: ResourceBgpRouterSchema(),
	}
}
