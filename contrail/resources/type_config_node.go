//
// Automatically generated. DO NOT EDIT.
//

package resources

import (
	"encoding/json"
	"math/big"

	"github.com/Juniper/contrail-go-api"
)

const (
	config_node_config_node_ip_address int = iota
	config_node_id_perms
	config_node_perms2
	config_node_annotations
	config_node_display_name
)

type ConfigNode struct {
	contrail.ObjectBase
	config_node_ip_address string
	id_perms               IdPermsType
	perms2                 PermType2
	annotations            KeyValuePairs
	display_name           string
	valid                  big.Int
	modified               big.Int
	baseMap                map[string]contrail.ReferenceList
}

func (obj *ConfigNode) GetType() string {
	return "config-node"
}

func (obj *ConfigNode) GetDefaultParent() []string {
	name := []string{"default-global-system-config"}
	return name
}

func (obj *ConfigNode) GetDefaultParentType() string {
	return "global-system-config"
}

func (obj *ConfigNode) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *ConfigNode) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *ConfigNode) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *ConfigNode) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *ConfigNode) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *ConfigNode) GetConfigNodeIpAddress() string {
	return obj.config_node_ip_address
}

func (obj *ConfigNode) SetConfigNodeIpAddress(value string) {
	obj.config_node_ip_address = value
	obj.modified.SetBit(&obj.modified, config_node_config_node_ip_address, 1)
}

func (obj *ConfigNode) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *ConfigNode) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, config_node_id_perms, 1)
}

func (obj *ConfigNode) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *ConfigNode) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, config_node_perms2, 1)
}

func (obj *ConfigNode) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *ConfigNode) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, config_node_annotations, 1)
}

func (obj *ConfigNode) GetDisplayName() string {
	return obj.display_name
}

func (obj *ConfigNode) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, config_node_display_name, 1)
}

func (obj *ConfigNode) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(config_node_config_node_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.config_node_ip_address)
		if err != nil {
			return nil, err
		}
		msg["config_node_ip_address"] = &value
	}

	if obj.modified.Bit(config_node_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(config_node_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(config_node_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(config_node_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	return json.Marshal(msg)
}

func (obj *ConfigNode) UnmarshalJSON(body []byte) error {
	var m map[string]json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	err = obj.UnmarshalCommon(m)
	if err != nil {
		return err
	}
	for key, value := range m {
		switch key {
		case "config_node_ip_address":
			err = json.Unmarshal(value, &obj.config_node_ip_address)
			if err == nil {
				obj.valid.SetBit(&obj.valid, config_node_config_node_ip_address, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, config_node_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, config_node_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, config_node_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, config_node_display_name, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ConfigNode) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(config_node_config_node_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.config_node_ip_address)
		if err != nil {
			return nil, err
		}
		msg["config_node_ip_address"] = &value
	}

	if obj.modified.Bit(config_node_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(config_node_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(config_node_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(config_node_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	return json.Marshal(msg)
}

func (obj *ConfigNode) UpdateReferences() error {

	return nil
}

func ConfigNodeByName(c contrail.ApiClient, fqn string) (*ConfigNode, error) {
	obj, err := c.FindByName("config-node", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*ConfigNode), nil
}

func ConfigNodeByUuid(c contrail.ApiClient, uuid string) (*ConfigNode, error) {
	obj, err := c.FindByUuid("config-node", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*ConfigNode), nil
}
