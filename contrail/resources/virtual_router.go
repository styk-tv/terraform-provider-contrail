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

func SetVirtualRouterFromResource(object *VirtualRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetVirtualRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_router_type"); ok {
		object.SetVirtualRouterType(val.(string))
	}
	if val, ok := d.GetOk("virtual_router_dpdk_enabled"); ok {
		object.SetVirtualRouterDpdkEnabled(val.(bool))
	}
	if val, ok := d.GetOk("virtual_router_ip_address"); ok {
		object.SetVirtualRouterIpAddress(val.(string))
	}
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

func SetRefsVirtualRouterFromResource(object *VirtualRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_machine_refs"); ok {
		log.Printf("Got ref virtual_machine_refs -- will call: object.AddVirtualMachine(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine by Uuid = %v as ref for VirtualMachine on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachine(refObj.(*VirtualMachine))
		}
	}

	return nil
}

func WriteVirtualRouterToResource(object VirtualRouter, d *schema.ResourceData, m interface{}) {

	d.Set("virtual_router_type", object.GetVirtualRouterType())
	d.Set("virtual_router_dpdk_enabled", object.GetVirtualRouterDpdkEnabled())
	d.Set("virtual_router_ip_address", object.GetVirtualRouterIpAddress())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeVirtualRouterAsMap(object *VirtualRouter) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["virtual_router_type"] = object.GetVirtualRouterType()
	omap["virtual_router_dpdk_enabled"] = object.GetVirtualRouterDpdkEnabled()
	omap["virtual_router_ip_address"] = object.GetVirtualRouterIpAddress()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateVirtualRouterFromResource(object *VirtualRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("virtual_router_type") {
		if val, ok := d.GetOk("virtual_router_type"); ok {
			object.SetVirtualRouterType(val.(string))
		}
	}
	if d.HasChange("virtual_router_dpdk_enabled") {
		if val, ok := d.GetOk("virtual_router_dpdk_enabled"); ok {
			object.SetVirtualRouterDpdkEnabled(val.(bool))
		}
	}
	if d.HasChange("virtual_router_ip_address") {
		if val, ok := d.GetOk("virtual_router_ip_address"); ok {
			object.SetVirtualRouterIpAddress(val.(string))
		}
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

func ResourceVirtualRouterCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualRouterCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(VirtualRouter)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceVirtualRouterCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "VirtualRouter", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetVirtualRouterFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceVirtualRouterCreate] Creation of resource VirtualRouter on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceVirtualRouterRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualRouterRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualRouterRefsCreate] Missing 'uuid' field for resource VirtualRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualRouterRefsCreate] Retrieving VirtualRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualRouter := obj.(*VirtualRouter) // Fully set by Contrail backend
	if err := SetRefsVirtualRouterFromResource(objVirtualRouter, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualRouterRefsCreate] Set refs on object VirtualRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualRouter.GetHref())
	if err := client.Update(objVirtualRouter); err != nil {
		return fmt.Errorf("[ResourceVirtualRouterRefsCreate] Update refs for resource VirtualRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualRouter.GetUuid())
	return nil
}

func ResourceVirtualRouterRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualRouterREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualRouterRead] Read resource virtual-router on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualRouter)
	WriteVirtualRouterToResource(*object, d, m)
	return nil
}

func ResourceVirtualRouterRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualRouterRefsREAD")
	return nil
}

func ResourceVirtualRouterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualRouterUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualRouterResourceUpdate] Retrieving VirtualRouter with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualRouter)
	UpdateVirtualRouterFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualRouterUpdate] Update of resource VirtualRouter on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualRouterRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualRouterRefsUpdate")
	return nil
}

func ResourceVirtualRouterDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualRouterDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("virtual-router", d.Id()); err != nil {
		return fmt.Errorf("[ResourceVirtualRouterDelete] Deletion of resource virtual-router on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceVirtualRouterRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualRouterRefsDelete: %v", d.Id())
	return nil
}

func ResourceVirtualRouterSchema() map[string]*schema.Schema {
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
		"virtual_router_type": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"virtual_router_dpdk_enabled": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"virtual_router_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
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

func ResourceVirtualRouterRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_machine_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachine(),
		},
	}
}

func ResourceVirtualRouter() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualRouterCreate,
		Read:   ResourceVirtualRouterRead,
		Update: ResourceVirtualRouterUpdate,
		Delete: ResourceVirtualRouterDelete,
		Schema: ResourceVirtualRouterSchema(),
	}
}

func ResourceVirtualRouterRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualRouterRefsCreate,
		Read:   ResourceVirtualRouterRefsRead,
		Update: ResourceVirtualRouterRefsUpdate,
		Delete: ResourceVirtualRouterRefsDelete,
		Schema: ResourceVirtualRouterRefsSchema(),
	}
}
