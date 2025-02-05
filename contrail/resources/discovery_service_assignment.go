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

func SetDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetDiscoveryServiceAssignmentFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsDiscoveryServiceAssignmentFromResource] key = %v, prefix = %v", key, prefix)

	return nil
}

func WriteDiscoveryServiceAssignmentToResource(object DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeDiscoveryServiceAssignmentAsMap(object *DiscoveryServiceAssignment) map[string]interface{} {
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

func UpdateDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceDiscoveryServiceAssignmentCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDiscoveryServiceAssignmentCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(DiscoveryServiceAssignment)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceDiscoveryServiceAssignmentCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "DiscoveryServiceAssignment", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetDiscoveryServiceAssignmentFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentCreate] Creation of resource DiscoveryServiceAssignment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDiscoveryServiceAssignmentRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Missing 'uuid' field for resource DiscoveryServiceAssignment")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("discovery-service-assignment", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Retrieving DiscoveryServiceAssignment with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDiscoveryServiceAssignment := obj.(*DiscoveryServiceAssignment) // Fully set by Contrail backend
	if err := SetRefsDiscoveryServiceAssignmentFromResource(objDiscoveryServiceAssignment, d, m); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Set refs on object DiscoveryServiceAssignment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDiscoveryServiceAssignment.GetHref())
	if err := client.Update(objDiscoveryServiceAssignment); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Update refs for resource DiscoveryServiceAssignment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDiscoveryServiceAssignment.GetUuid())
	return nil
}

func ResourceDiscoveryServiceAssignmentRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("discovery-service-assignment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRead] Read resource discovery-service-assignment on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*DiscoveryServiceAssignment)
	WriteDiscoveryServiceAssignmentToResource(*object, d, m)
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentRefsREAD")
	return nil
}

func ResourceDiscoveryServiceAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("discovery-service-assignment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentResourceUpdate] Retrieving DiscoveryServiceAssignment with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*DiscoveryServiceAssignment)
	UpdateDiscoveryServiceAssignmentFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentUpdate] Update of resource DiscoveryServiceAssignment on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentRefsUpdate")
	return nil
}

func ResourceDiscoveryServiceAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("discovery-service-assignment", d.Id()); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentDelete] Deletion of resource discovery-service-assignment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentRefsDelete: %v", d.Id())
	return nil
}

func ResourceDiscoveryServiceAssignmentSchema() map[string]*schema.Schema {
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

func ResourceDiscoveryServiceAssignment() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDiscoveryServiceAssignmentCreate,
		Read:   ResourceDiscoveryServiceAssignmentRead,
		Update: ResourceDiscoveryServiceAssignmentUpdate,
		Delete: ResourceDiscoveryServiceAssignmentDelete,
		Schema: ResourceDiscoveryServiceAssignmentSchema(),
	}
}
