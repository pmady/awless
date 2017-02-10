// DO NOT EDIT
// This file was automatically generated with go generate

/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

var AWSTemplatesDefinitions = map[string]string{
	"createvpc":             "create vpc cidr={ vpc.cidr } ",
	"deletevpc":             "delete vpc id={ vpc.id } ",
	"createsubnet":          "create subnet cidr={ subnet.cidr } vpc={ subnet.vpc } ",
	"updatesubnet":          "update subnet id={ subnet.id } ",
	"deletesubnet":          "delete subnet id={ subnet.id } ",
	"createinstance":        "create instance image={ instance.image } type={ instance.type } count={ instance.count } count={ instance.count } subnet={ instance.subnet }  name={ instance.name }",
	"updateinstance":        "update instance id={ instance.id } ",
	"deleteinstance":        "delete instance id={ instance.id } ",
	"startinstance":         "start instance id={ instance.id } ",
	"stopinstance":          "stop instance id={ instance.id } ",
	"checkinstance":         "check instance id={ instance.id } state={ instance.state } timeout={ instance.timeout } ",
	"createsecuritygroup":   "create securitygroup description={ securitygroup.description } name={ securitygroup.name } vpc={ securitygroup.vpc } ",
	"updatesecuritygroup":   "update securitygroup cidr={ securitygroup.cidr } id={ securitygroup.id } protocol={ securitygroup.protocol } ",
	"deletesecuritygroup":   "delete securitygroup id={ securitygroup.id } ",
	"createvolume":          "create volume zone={ volume.zone } size={ volume.size } ",
	"deletevolume":          "delete volume id={ volume.id } ",
	"attachvolume":          "attach volume device={ volume.device } instance={ volume.instance } id={ volume.id } ",
	"createinternetgateway": "create internetgateway ",
	"deleteinternetgateway": "delete internetgateway id={ internetgateway.id } ",
	"attachinternetgateway": "attach internetgateway id={ internetgateway.id } vpc={ internetgateway.vpc } ",
	"detachinternetgateway": "detach internetgateway id={ internetgateway.id } vpc={ internetgateway.vpc } ",
	"createroutetable":      "create routetable vpc={ routetable.vpc } ",
	"deleteroutetable":      "delete routetable id={ routetable.id } ",
	"attachroutetable":      "attach routetable id={ routetable.id } subnet={ routetable.subnet } ",
	"detachroutetable":      "detach routetable association={ routetable.association } ",
	"createroute":           "create route cidr={ route.cidr } gateway={ route.gateway } table={ route.table } ",
	"deleteroute":           "delete route cidr={ route.cidr } table={ route.table } ",
	"createtags":            "create tags resource={ tags.resource } ",
	"createkeypair":         "create keypair name={ keypair.name } ",
	"deletekeypair":         "delete keypair id={ keypair.id } ",
	"createuser":            "create user name={ user.name } ",
	"deleteuser":            "delete user name={ user.name } ",
	"creategroup":           "create group name={ group.name } ",
	"deletegroup":           "delete group name={ group.name } ",
	"attachpolicy":          "attach policy arn={ policy.arn } user={ policy.user } ",
	"detachpolicy":          "detach policy arn={ policy.arn } user={ policy.user } ",
	"createbucket":          "create bucket name={ bucket.name } ",
	"deletebucket":          "delete bucket name={ bucket.name } ",
	"createstorageobject":   "create storageobject file={ storageobject.file } bucket={ storageobject.bucket } ",
	"deletestorageobject":   "delete storageobject bucket={ storageobject.bucket } key={ storageobject.key } ",
}
