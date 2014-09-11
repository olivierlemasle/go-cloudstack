//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type AddTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *AddTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervnetworklabel"]; found {
		u.Set("hypervnetworklabel", v.(string))
	}
	if v, found := p.p["isolationmethod"]; found {
		u.Set("isolationmethod", v.(string))
	}
	if v, found := p.p["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vmwarenetworklabel"]; found {
		u.Set("vmwarenetworklabel", v.(string))
	}
	if v, found := p.p["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (p *AddTrafficTypeParams) SetHypervnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervnetworklabel"] = v
	return
}

func (p *AddTrafficTypeParams) SetIsolationmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolationmethod"] = v
	return
}

func (p *AddTrafficTypeParams) SetKvmnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kvmnetworklabel"] = v
	return
}

func (p *AddTrafficTypeParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddTrafficTypeParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

func (p *AddTrafficTypeParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

func (p *AddTrafficTypeParams) SetVmwarenetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmwarenetworklabel"] = v
	return
}

func (p *AddTrafficTypeParams) SetXennetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["xennetworklabel"] = v
	return
}

// You should always use this function to get a new AddTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficTypeParams(physicalnetworkid string, traffictype string) *AddTrafficTypeParams {
	p := &AddTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["traffictype"] = traffictype
	return p
}

// Adds traffic type to a physical network
func (s *UsageService) AddTrafficType(p *AddTrafficTypeParams) (*AddTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("addTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficTypeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r AddTrafficTypeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddTrafficTypeResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Kvmnetworklabel    string `json:"kvmnetworklabel,omitempty"`
	Id                 string `json:"id,omitempty"`
	Hypervnetworklabel string `json:"hypervnetworklabel,omitempty"`
	Traffictype        string `json:"traffictype,omitempty"`
	Xennetworklabel    string `json:"xennetworklabel,omitempty"`
	Physicalnetworkid  string `json:"physicalnetworkid,omitempty"`
	Vmwarenetworklabel string `json:"vmwarenetworklabel,omitempty"`
}

type DeleteTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *DeleteTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteTrafficTypeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficTypeParams(id string) *DeleteTrafficTypeParams {
	p := &DeleteTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes traffic type of a physical network
func (s *UsageService) DeleteTrafficType(p *DeleteTrafficTypeParams) (*DeleteTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("deleteTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficTypeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r DeleteTrafficTypeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteTrafficTypeResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListTrafficTypesParams struct {
	p map[string]interface{}
}

func (p *ListTrafficTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListTrafficTypesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListTrafficTypesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListTrafficTypesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListTrafficTypesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

// You should always use this function to get a new ListTrafficTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypesParams(physicalnetworkid string) *ListTrafficTypesParams {
	p := &ListTrafficTypesParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UsageService) GetTrafficTypeID(keyword string, physicalnetworkid string) (string, error) {
	p := &ListTrafficTypesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["physicalnetworkid"] = physicalnetworkid

	l, err := s.ListTrafficTypes(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.TrafficTypes[0].Id, nil
}

// Lists traffic types of a given physical network.
func (s *UsageService) ListTrafficTypes(p *ListTrafficTypesParams) (*ListTrafficTypesResponse, error) {
	resp, err := s.cs.newRequest("listTrafficTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListTrafficTypesResponse struct {
	Count        int            `json:"count"`
	TrafficTypes []*TrafficType `json:"traffictype"`
}

type TrafficType struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
	Id                           string   `json:"id,omitempty"`
	Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Servicelist                  []string `json:"servicelist,omitempty"`
	State                        string   `json:"state,omitempty"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
}

type UpdateTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *UpdateTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervnetworklabel"]; found {
		u.Set("hypervnetworklabel", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := p.p["vmwarenetworklabel"]; found {
		u.Set("vmwarenetworklabel", v.(string))
	}
	if v, found := p.p["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (p *UpdateTrafficTypeParams) SetHypervnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervnetworklabel"] = v
	return
}

func (p *UpdateTrafficTypeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateTrafficTypeParams) SetKvmnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kvmnetworklabel"] = v
	return
}

func (p *UpdateTrafficTypeParams) SetVmwarenetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmwarenetworklabel"] = v
	return
}

func (p *UpdateTrafficTypeParams) SetXennetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["xennetworklabel"] = v
	return
}

// You should always use this function to get a new UpdateTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewUpdateTrafficTypeParams(id string) *UpdateTrafficTypeParams {
	p := &UpdateTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates traffic type of a physical network
func (s *UsageService) UpdateTrafficType(p *UpdateTrafficTypeParams) (*UpdateTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("updateTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTrafficTypeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r UpdateTrafficTypeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateTrafficTypeResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Physicalnetworkid  string `json:"physicalnetworkid,omitempty"`
	Vmwarenetworklabel string `json:"vmwarenetworklabel,omitempty"`
	Kvmnetworklabel    string `json:"kvmnetworklabel,omitempty"`
	Id                 string `json:"id,omitempty"`
	Hypervnetworklabel string `json:"hypervnetworklabel,omitempty"`
	Xennetworklabel    string `json:"xennetworklabel,omitempty"`
	Traffictype        string `json:"traffictype,omitempty"`
}

type ListTrafficTypeImplementorsParams struct {
	p map[string]interface{}
}

func (p *ListTrafficTypeImplementorsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *ListTrafficTypeImplementorsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListTrafficTypeImplementorsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListTrafficTypeImplementorsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListTrafficTypeImplementorsParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

// You should always use this function to get a new ListTrafficTypeImplementorsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypeImplementorsParams() *ListTrafficTypeImplementorsParams {
	p := &ListTrafficTypeImplementorsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists implementors of implementor of a network traffic type or implementors of all network traffic types
func (s *UsageService) ListTrafficTypeImplementors(p *ListTrafficTypeImplementorsParams) (*ListTrafficTypeImplementorsResponse, error) {
	resp, err := s.cs.newRequest("listTrafficTypeImplementors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficTypeImplementorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListTrafficTypeImplementorsResponse struct {
	Count                   int                       `json:"count"`
	TrafficTypeImplementors []*TrafficTypeImplementor `json:"traffictypeimplementor"`
}

type TrafficTypeImplementor struct {
	Traffictype            string `json:"traffictype,omitempty"`
	Traffictypeimplementor string `json:"traffictypeimplementor,omitempty"`
}

type GenerateUsageRecordsParams struct {
	p map[string]interface{}
}

func (p *GenerateUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *GenerateUsageRecordsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *GenerateUsageRecordsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
	return
}

func (p *GenerateUsageRecordsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
	return
}

// You should always use this function to get a new GenerateUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewGenerateUsageRecordsParams(enddate string, startdate string) *GenerateUsageRecordsParams {
	p := &GenerateUsageRecordsParams{}
	p.p = make(map[string]interface{})
	p.p["enddate"] = enddate
	p.p["startdate"] = startdate
	return p
}

// Generates usage records. This will generate records only if there any records to be generated, i.e if the scheduled usage job was not run or failed
func (s *UsageService) GenerateUsageRecords(p *GenerateUsageRecordsParams) (*GenerateUsageRecordsResponse, error) {
	resp, err := s.cs.newRequest("generateUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GenerateUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type GenerateUsageRecordsResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListUsageRecordsParams struct {
	p map[string]interface{}
}

func (p *ListUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["type"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("type", vv)
	}
	return u
}

func (p *ListUsageRecordsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListUsageRecordsParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
	return
}

func (p *ListUsageRecordsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListUsageRecordsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
	return
}

func (p *ListUsageRecordsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUsageRecordsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUsageRecordsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUsageRecordsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListUsageRecordsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
	return
}

func (p *ListUsageRecordsParams) SetType(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usageType"] = v
	return
}

// You should always use this function to get a new ListUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageRecordsParams(enddate string, startdate string) *ListUsageRecordsParams {
	p := &ListUsageRecordsParams{}
	p.p = make(map[string]interface{})
	p.p["enddate"] = enddate
	p.p["startdate"] = startdate
	return p
}

// Lists usage records for accounts
func (s *UsageService) ListUsageRecords(p *ListUsageRecordsParams) (*ListUsageRecordsResponse, error) {
	resp, err := s.cs.newRequest("listUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUsageRecordsResponse struct {
	Count        int            `json:"count"`
	UsageRecords []*UsageRecord `json:"usagerecord"`
}

type UsageRecord struct {
	Projectid        string `json:"projectid,omitempty"`
	Isdefault        bool   `json:"isdefault,omitempty"`
	Enddate          string `json:"enddate,omitempty"`
	Rawusage         string `json:"rawusage,omitempty"`
	Offeringid       string `json:"offeringid,omitempty"`
	Name             string `json:"name,omitempty"`
	Issourcenat      bool   `json:"issourcenat,omitempty"`
	Domainid         string `json:"domainid,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Usage            string `json:"usage,omitempty"`
	Issystem         bool   `json:"issystem,omitempty"`
	Type             string `json:"type,omitempty"`
	Description      string `json:"description,omitempty"`
	Usageid          string `json:"usageid,omitempty"`
	Networkid        string `json:"networkid,omitempty"`
	Project          string `json:"project,omitempty"`
	Startdate        string `json:"startdate,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Account          string `json:"account,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
	Size             int    `json:"size,omitempty"`
	Virtualsize      int    `json:"virtualsize,omitempty"`
	Usagetype        int    `json:"usagetype,omitempty"`
	Templateid       string `json:"templateid,omitempty"`
}

type ListUsageTypesParams struct {
	p map[string]interface{}
}

func (p *ListUsageTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListUsageTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageTypesParams() *ListUsageTypesParams {
	p := &ListUsageTypesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List Usage Types
func (s *UsageService) ListUsageTypes(p *ListUsageTypesParams) (*ListUsageTypesResponse, error) {
	resp, err := s.cs.newRequest("listUsageTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUsageTypesResponse struct {
	Count      int          `json:"count"`
	UsageTypes []*UsageType `json:"usagetype"`
}

type UsageType struct {
	Usagetypeid int    `json:"usagetypeid,omitempty"`
	Description string `json:"description,omitempty"`
}

type AddTrafficMonitorParams struct {
	p map[string]interface{}
}

func (p *AddTrafficMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["excludezones"]; found {
		u.Set("excludezones", v.(string))
	}
	if v, found := p.p["includezones"]; found {
		u.Set("includezones", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddTrafficMonitorParams) SetExcludezones(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["excludezones"] = v
	return
}

func (p *AddTrafficMonitorParams) SetIncludezones(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["includezones"] = v
	return
}

func (p *AddTrafficMonitorParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddTrafficMonitorParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddTrafficMonitorParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficMonitorParams(url string, zoneid string) *AddTrafficMonitorParams {
	p := &AddTrafficMonitorParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Adds Traffic Monitor Host for Direct Network Usage
func (s *UsageService) AddTrafficMonitor(p *AddTrafficMonitorParams) (*AddTrafficMonitorResponse, error) {
	resp, err := s.cs.newRequest("addTrafficMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddTrafficMonitorResponse struct {
	Timeout    string `json:"timeout,omitempty"`
	Id         string `json:"id,omitempty"`
	Zoneid     string `json:"zoneid,omitempty"`
	Numretries string `json:"numretries,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
}

type DeleteTrafficMonitorParams struct {
	p map[string]interface{}
}

func (p *DeleteTrafficMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteTrafficMonitorParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteTrafficMonitorParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficMonitorParams(id string) *DeleteTrafficMonitorParams {
	p := &DeleteTrafficMonitorParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an traffic monitor host.
func (s *UsageService) DeleteTrafficMonitor(p *DeleteTrafficMonitorParams) (*DeleteTrafficMonitorResponse, error) {
	resp, err := s.cs.newRequest("deleteTrafficMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteTrafficMonitorResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListTrafficMonitorsParams struct {
	p map[string]interface{}
}

func (p *ListTrafficMonitorsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTrafficMonitorsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListTrafficMonitorsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListTrafficMonitorsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListTrafficMonitorsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListTrafficMonitorsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficMonitorsParams(zoneid string) *ListTrafficMonitorsParams {
	p := &ListTrafficMonitorsParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UsageService) GetTrafficMonitorID(keyword string, zoneid string) (string, error) {
	p := &ListTrafficMonitorsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["zoneid"] = zoneid

	l, err := s.ListTrafficMonitors(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.TrafficMonitors[0].Id, nil
}

// List traffic monitor Hosts.
func (s *UsageService) ListTrafficMonitors(p *ListTrafficMonitorsParams) (*ListTrafficMonitorsResponse, error) {
	resp, err := s.cs.newRequest("listTrafficMonitors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficMonitorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListTrafficMonitorsResponse struct {
	Count           int               `json:"count"`
	TrafficMonitors []*TrafficMonitor `json:"trafficmonitor"`
}

type TrafficMonitor struct {
	Timeout    string `json:"timeout,omitempty"`
	Zoneid     string `json:"zoneid,omitempty"`
	Numretries string `json:"numretries,omitempty"`
	Id         string `json:"id,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
}