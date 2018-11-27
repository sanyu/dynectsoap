package dynectsoap

import (
	"github.com/hooklift/gowsdl/soap"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	PollingInterval = 2 * time.Second
)

type dynect struct {
	client *soap.Client
}

func NewDynect(client *soap.Client) *dynect {
	return &dynect{
		client: client,
	}
}

func (service *dynect) GetJobRetry(request *GetJobRequestType) (*GetJobResponseType, error) {
	var err error
	response := new(GetJobResponseType)

	for {
		select {
		case <-time.After(PollingInterval):
			err = service.client.Call("https://api2.dynect.net/SOAP/", request, response)
			if err != nil {
				return nil, err
			}

			realResponse := response.Data.(JobResponseInterface)
			status := realResponse.status()

			switch status {
			case "incomplete":
				//Loop around again
				log.Infof("Tryingt to retrieve Job %v again, status is incomplete", response.Job_id)
				continue
			default:
				return response, nil
			}
		}
	}

	return response, nil
}

func (service *dynect) GetAllRecords(request *GetAllRecordsRequestType) (*GetAllRecordsResponseType, error) {
	response := new(GetAllRecordsResponseType)

	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetJob(request *GetJobRequestType) (*GetJobResponseType, error) {
	response := new(GetJobResponseType)

	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionLogin(request *SessionLoginRequestType) (*SessionLoginResponseType, error) {
	response := new(SessionLoginResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionLogout(request *SessionLogoutRequestType) (*SessionLogoutResponseType, error) {
	response := new(SessionLogoutResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionIsAlive(request *SessionIsAliveRequestType) (*SessionIsAliveResponseType, error) {
	response := new(SessionIsAliveResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionKeepAlive(request *SessionKeepAliveRequestType) (*SessionKeepAliveResponseType, error) {
	response := new(SessionKeepAliveResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ScopeIn(request *ScopeInRequestType) (*ScopeInResponseType, error) {
	response := new(ScopeInResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ScopeAs(request *ScopeAsRequestType) (*ScopeAsResponseType, error) {
	response := new(ScopeAsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) Unscope(request *UnscopeRequestType) (*UnscopeResponseType, error) {
	response := new(UnscopeResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetQueryStats(request *GetQueryStatsRequestType) (*GetQueryStatsResponseType, error) {
	response := new(GetQueryStatsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateGeo(request *CreateGeoRequestType) (*CreateGeoResponseType, error) {
	response := new(CreateGeoResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateGeo(request *UpdateGeoRequestType) (*UpdateGeoResponseType, error) {
	response := new(UpdateGeoResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetGeos(request *GetGeosRequestType) (*GetGeosResponseType, error) {
	response := new(GetGeosResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneGeo(request *GetOneGeoRequestType) (*GetOneGeoResponseType, error) {
	response := new(GetOneGeoResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneGeo(request *DeleteOneGeoRequestType) (*DeleteOneGeoResponseType, error) {
	response := new(DeleteOneGeoResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateGeo(request *ActivateGeoRequestType) (*ActivateGeoResponseType, error) {
	response := new(ActivateGeoResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateGeo(request *DeactivateGeoRequestType) (*DeactivateGeoResponseType, error) {
	response := new(DeactivateGeoResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateGeoRegionGroup(request *CreateGeoRegionGroupRequestType) (*CreateGeoRegionGroupResponseType, error) {
	response := new(CreateGeoRegionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateGeoRegionGroup(request *UpdateGeoRegionGroupRequestType) (*UpdateGeoRegionGroupResponseType, error) {
	response := new(UpdateGeoRegionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneGeoRegionGroup(request *DeleteOneGeoRegionGroupRequestType) (*DeleteOneGeoRegionGroupResponseType, error) {
	response := new(DeleteOneGeoRegionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetGeoRegionGroups(request *GetGeoRegionGroupsRequestType) (*GetGeoRegionGroupsResponseType, error) {
	response := new(GetGeoRegionGroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneGeoRegionGroup(request *GetOneGeoRegionGroupRequestType) (*GetOneGeoRegionGroupResponseType, error) {
	response := new(GetOneGeoRegionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateGeoNode(request *CreateGeoNodeRequestType) (*CreateGeoNodeResponseType, error) {
	response := new(CreateGeoNodeResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneGeoNode(request *DeleteOneGeoNodeRequestType) (*DeleteOneGeoNodeResponseType, error) {
	response := new(DeleteOneGeoNodeResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetGeoNodes(request *GetGeoNodesRequestType) (*GetGeoNodesResponseType, error) {
	response := new(GetGeoNodesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSF(request *CreateDSFRequestType) (*CreateDSFResponseType, error) {
	response := new(CreateDSFResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSF(request *UpdateDSFRequestType) (*UpdateDSFResponseType, error) {
	response := new(UpdateDSFResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFs(request *GetDSFsRequestType) (*GetDSFsResponseType, error) {
	response := new(GetDSFsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFNotifiers(request *GetDSFNotifiersRequestType) (*GetDSFNotifiersResponseType, error) {
	response := new(GetDSFNotifiersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSF(request *DeleteOneDSFRequestType) (*DeleteOneDSFResponseType, error) {
	response := new(DeleteOneDSFResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSF(request *GetOneDSFRequestType) (*GetOneDSFResponseType, error) {
	response := new(GetOneDSFResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RevertDSF(request *RevertDSFRequestType) (*RevertDSFResponseType, error) {
	response := new(RevertDSFResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) PublishDSF(request *PublishDSFRequestType) (*PublishDSFResponseType, error) {
	response := new(PublishDSFResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddDSFNotifier(request *AddDSFNotifierRequestType) (*AddDSFNotifierResponseType, error) {
	response := new(AddDSFNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveDSFNotifier(request *RemoveDSFNotifierRequestType) (*RemoveDSFNotifierResponseType, error) {
	response := new(RemoveDSFNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSFRuleset(request *CreateDSFRulesetRequestType) (*CreateDSFRulesetResponseType, error) {
	response := new(CreateDSFRulesetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFRuleset(request *UpdateDSFRulesetRequestType) (*UpdateDSFRulesetResponseType, error) {
	response := new(UpdateDSFRulesetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFRulesets(request *GetDSFRulesetsRequestType) (*GetDSFRulesetsResponseType, error) {
	response := new(GetDSFRulesetsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSFRuleset(request *GetOneDSFRulesetRequestType) (*GetOneDSFRulesetResponseType, error) {
	response := new(GetOneDSFRulesetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFRuleset(request *DeleteOneDSFRulesetRequestType) (*DeleteOneDSFRulesetResponseType, error) {
	response := new(DeleteOneDSFRulesetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSFResponsePool(request *CreateDSFResponsePoolRequestType) (*CreateDSFResponsePoolResponseType, error) {
	response := new(CreateDSFResponsePoolResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFResponsePool(request *UpdateDSFResponsePoolRequestType) (*UpdateDSFResponsePoolResponseType, error) {
	response := new(UpdateDSFResponsePoolResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFResponsePools(request *GetDSFResponsePoolsRequestType) (*GetDSFResponsePoolsResponseType, error) {
	response := new(GetDSFResponsePoolsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSFResponsePool(request *GetOneDSFResponsePoolRequestType) (*GetOneDSFResponsePoolResponseType, error) {
	response := new(GetOneDSFResponsePoolResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFResponsePool(request *DeleteOneDSFResponsePoolRequestType) (*DeleteOneDSFResponsePoolResponseType, error) {
	response := new(DeleteOneDSFResponsePoolResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSFRecordSetFailoverChain(request *CreateDSFRecordSetFailoverChainRequestType) (*CreateDSFRecordSetFailoverChainResponseType, error) {
	response := new(CreateDSFRecordSetFailoverChainResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFRecordSetFailoverChain(request *UpdateDSFRecordSetFailoverChainRequestType) (*UpdateDSFRecordSetFailoverChainResponseType, error) {
	response := new(UpdateDSFRecordSetFailoverChainResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFRecordSetFailoverChains(request *GetDSFRecordSetFailoverChainsRequestType) (*GetDSFRecordSetFailoverChainsResponseType, error) {
	response := new(GetDSFRecordSetFailoverChainsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSFRecordSetFailoverChain(request *GetOneDSFRecordSetFailoverChainRequestType) (*GetOneDSFRecordSetFailoverChainResponseType, error) {
	response := new(GetOneDSFRecordSetFailoverChainResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFRecordSetFailoverChain(request *DeleteOneDSFRecordSetFailoverChainRequestType) (*DeleteOneDSFRecordSetFailoverChainResponseType, error) {
	response := new(DeleteOneDSFRecordSetFailoverChainResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSFRecordSet(request *CreateDSFRecordSetRequestType) (*CreateDSFRecordSetResponseType, error) {
	response := new(CreateDSFRecordSetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFRecordSet(request *UpdateDSFRecordSetRequestType) (*UpdateDSFRecordSetResponseType, error) {
	response := new(UpdateDSFRecordSetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSFRecordSet(request *GetOneDSFRecordSetRequestType) (*GetOneDSFRecordSetResponseType, error) {
	response := new(GetOneDSFRecordSetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFRecordSets(request *GetDSFRecordSetsRequestType) (*GetDSFRecordSetsResponseType, error) {
	response := new(GetDSFRecordSetsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFRecordSet(request *DeleteOneDSFRecordSetRequestType) (*DeleteOneDSFRecordSetResponseType, error) {
	response := new(DeleteOneDSFRecordSetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSFRecord(request *CreateDSFRecordRequestType) (*CreateDSFRecordResponseType, error) {
	response := new(CreateDSFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFRecord(request *UpdateDSFRecordRequestType) (*UpdateDSFRecordResponseType, error) {
	response := new(UpdateDSFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSFRecord(request *GetOneDSFRecordRequestType) (*GetOneDSFRecordResponseType, error) {
	response := new(GetOneDSFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFRecords(request *GetDSFRecordsRequestType) (*GetDSFRecordsResponseType, error) {
	response := new(GetDSFRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFRecord(request *DeleteOneDSFRecordRequestType) (*DeleteOneDSFRecordResponseType, error) {
	response := new(DeleteOneDSFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddDSFNode(request *AddDSFNodeRequestType) (*AddDSFNodeResponseType, error) {
	response := new(AddDSFNodeResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFNodes(request *UpdateDSFNodesRequestType) (*UpdateDSFNodesResponseType, error) {
	response := new(UpdateDSFNodesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFNodes(request *GetDSFNodesRequestType) (*GetDSFNodesResponseType, error) {
	response := new(GetDSFNodesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFNode(request *DeleteOneDSFNodeRequestType) (*DeleteOneDSFNodeResponseType, error) {
	response := new(DeleteOneDSFNodeResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSFMonitor(request *CreateDSFMonitorRequestType) (*CreateDSFMonitorResponseType, error) {
	response := new(CreateDSFMonitorResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSFMonitor(request *UpdateDSFMonitorRequestType) (*UpdateDSFMonitorResponseType, error) {
	response := new(UpdateDSFMonitorResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSFMonitor(request *GetOneDSFMonitorRequestType) (*GetOneDSFMonitorResponseType, error) {
	response := new(GetOneDSFMonitorResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFMonitors(request *GetDSFMonitorsRequestType) (*GetDSFMonitorsResponseType, error) {
	response := new(GetDSFMonitorsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSFMonitor(request *DeleteOneDSFMonitorRequestType) (*DeleteOneDSFMonitorResponseType, error) {
	response := new(DeleteOneDSFMonitorResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddDSFMonitorNotifier(request *AddDSFMonitorNotifierRequestType) (*AddDSFMonitorNotifierResponseType, error) {
	response := new(AddDSFMonitorNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSFMonitorSites(request *GetDSFMonitorSitesRequestType) (*GetDSFMonitorSitesResponseType, error) {
	response := new(GetDSFMonitorSitesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateNotifier(request *CreateNotifierRequestType) (*CreateNotifierResponseType, error) {
	response := new(CreateNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateNotifier(request *UpdateNotifierRequestType) (*UpdateNotifierResponseType, error) {
	response := new(UpdateNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneNotifier(request *GetOneNotifierRequestType) (*GetOneNotifierResponseType, error) {
	response := new(GetOneNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetNotifiers(request *GetNotifiersRequestType) (*GetNotifiersResponseType, error) {
	response := new(GetNotifiersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneNotifier(request *DeleteOneNotifierRequestType) (*DeleteOneNotifierResponseType, error) {
	response := new(DeleteOneNotifierResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateConfigLimit(request *CreateConfigLimitRequestType) (*CreateConfigLimitResponseType, error) {
	response := new(CreateConfigLimitResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneConfigLimit(request *GetOneConfigLimitRequestType) (*GetOneConfigLimitResponseType, error) {
	response := new(GetOneConfigLimitResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetConfigLimits(request *GetConfigLimitsRequestType) (*GetConfigLimitsResponseType, error) {
	response := new(GetConfigLimitsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateConfigLimit(request *UpdateConfigLimitRequestType) (*UpdateConfigLimitResponseType, error) {
	response := new(UpdateConfigLimitResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneConfigLimit(request *DeleteOneConfigLimitRequestType) (*DeleteOneConfigLimitResponseType, error) {
	response := new(DeleteOneConfigLimitResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreatePermissionGroup(request *CreatePermissionGroupRequestType) (*CreatePermissionGroupResponseType, error) {
	response := new(CreatePermissionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOnePermissionGroup(request *GetOnePermissionGroupRequestType) (*GetOnePermissionGroupResponseType, error) {
	response := new(GetOnePermissionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetPermissionGroups(request *GetPermissionGroupsRequestType) (*GetPermissionGroupsResponseType, error) {
	response := new(GetPermissionGroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOnePermissionGroup(request *DeleteOnePermissionGroupRequestType) (*DeleteOnePermissionGroupResponseType, error) {
	response := new(DeleteOnePermissionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdatePermissionGroup(request *UpdatePermissionGroupRequestType) (*UpdatePermissionGroupResponseType, error) {
	response := new(UpdatePermissionGroupResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCustomerPermissions(request *GetCustomerPermissionsRequestType) (*GetCustomerPermissionsResponseType, error) {
	response := new(GetCustomerPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetUserPermissions(request *GetUserPermissionsRequestType) (*GetUserPermissionsResponseType, error) {
	response := new(GetUserPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CheckPermissions(request *CheckPermissionsRequestType) (*CheckPermissionsResponseType, error) {
	response := new(CheckPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddPermissionGroupUsers(request *AddPermissionGroupUsersRequestType) (*AddPermissionGroupUsersResponseType, error) {
	response := new(AddPermissionGroupUsersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetPermissionGroupUsers(request *SetPermissionGroupUsersRequestType) (*SetPermissionGroupUsersResponseType, error) {
	response := new(SetPermissionGroupUsersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemovePermissionGroupUsers(request *RemovePermissionGroupUsersRequestType) (*RemovePermissionGroupUsersResponseType, error) {
	response := new(RemovePermissionGroupUsersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddPermissionGroupSubgroups(request *AddPermissionGroupSubgroupsRequestType) (*AddPermissionGroupSubgroupsResponseType, error) {
	response := new(AddPermissionGroupSubgroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetPermissionGroupSubgroups(request *SetPermissionGroupSubgroupsRequestType) (*SetPermissionGroupSubgroupsResponseType, error) {
	response := new(SetPermissionGroupSubgroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemovePermissionGroupSubgroups(request *RemovePermissionGroupSubgroupsRequestType) (*RemovePermissionGroupSubgroupsResponseType, error) {
	response := new(RemovePermissionGroupSubgroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddPermissionGroupPermissions(request *AddPermissionGroupPermissionsRequestType) (*AddPermissionGroupPermissionsResponseType, error) {
	response := new(AddPermissionGroupPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetPermissionGroupPermissions(request *SetPermissionGroupPermissionsRequestType) (*SetPermissionGroupPermissionsResponseType, error) {
	response := new(SetPermissionGroupPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemovePermissionGroupPermissions(request *RemovePermissionGroupPermissionsRequestType) (*RemovePermissionGroupPermissionsResponseType, error) {
	response := new(RemovePermissionGroupPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddPermissionGroupZones(request *AddPermissionGroupZonesRequestType) (*AddPermissionGroupZonesResponseType, error) {
	response := new(AddPermissionGroupZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetPermissionGroupZones(request *SetPermissionGroupZonesRequestType) (*SetPermissionGroupZonesResponseType, error) {
	response := new(SetPermissionGroupZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemovePermissionGroupZones(request *RemovePermissionGroupZonesRequestType) (*RemovePermissionGroupZonesResponseType, error) {
	response := new(RemovePermissionGroupZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddUserGroups(request *AddUserGroupsRequestType) (*AddUserGroupsResponseType, error) {
	response := new(AddUserGroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetUserGroups(request *SetUserGroupsRequestType) (*SetUserGroupsResponseType, error) {
	response := new(SetUserGroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveUserGroups(request *RemoveUserGroupsRequestType) (*RemoveUserGroupsResponseType, error) {
	response := new(RemoveUserGroupsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddUserZones(request *AddUserZonesRequestType) (*AddUserZonesResponseType, error) {
	response := new(AddUserZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetUserZones(request *SetUserZonesRequestType) (*SetUserZonesResponseType, error) {
	response := new(SetUserZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveUserZones(request *RemoveUserZonesRequestType) (*RemoveUserZonesResponseType, error) {
	response := new(RemoveUserZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddUserPermissions(request *AddUserPermissionsRequestType) (*AddUserPermissionsResponseType, error) {
	response := new(AddUserPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetUserPermissions(request *SetUserPermissionsRequestType) (*SetUserPermissionsResponseType, error) {
	response := new(SetUserPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveUserPermissions(request *RemoveUserPermissionsRequestType) (*RemoveUserPermissionsResponseType, error) {
	response := new(RemoveUserPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddUserForbids(request *AddUserForbidsRequestType) (*AddUserForbidsResponseType, error) {
	response := new(AddUserForbidsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetUserForbids(request *SetUserForbidsRequestType) (*SetUserForbidsResponseType, error) {
	response := new(SetUserForbidsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveUserForbids(request *RemoveUserForbidsRequestType) (*RemoveUserForbidsResponseType, error) {
	response := new(RemoveUserForbidsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddCustomerPermissions(request *AddCustomerPermissionsRequestType) (*AddCustomerPermissionsResponseType, error) {
	response := new(AddCustomerPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetCustomerPermissions(request *SetCustomerPermissionsRequestType) (*SetCustomerPermissionsResponseType, error) {
	response := new(SetCustomerPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveCustomerPermissions(request *RemoveCustomerPermissionsRequestType) (*RemoveCustomerPermissionsResponseType, error) {
	response := new(RemoveCustomerPermissionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddCustomerForbids(request *AddCustomerForbidsRequestType) (*AddCustomerForbidsResponseType, error) {
	response := new(AddCustomerForbidsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetCustomerForbids(request *SetCustomerForbidsRequestType) (*SetCustomerForbidsResponseType, error) {
	response := new(SetCustomerForbidsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RemoveCustomerForbids(request *RemoveCustomerForbidsRequestType) (*RemoveCustomerForbidsResponseType, error) {
	response := new(RemoveCustomerForbidsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetHostStatsFlags(request *GetHostStatsFlagsRequestType) (*GetHostStatsFlagsResponseType, error) {
	response := new(GetHostStatsFlagsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetHostStatsFlags(request *SetHostStatsFlagsRequestType) (*SetHostStatsFlagsResponseType, error) {
	response := new(SetHostStatsFlagsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateTSIGKey(request *CreateTSIGKeyRequestType) (*CreateTSIGKeyResponseType, error) {
	response := new(CreateTSIGKeyResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneTSIGKey(request *GetOneTSIGKeyRequestType) (*GetOneTSIGKeyResponseType, error) {
	response := new(GetOneTSIGKeyResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetTSIGKeys(request *GetTSIGKeysRequestType) (*GetTSIGKeysResponseType, error) {
	response := new(GetTSIGKeysResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateTSIGKey(request *UpdateTSIGKeyRequestType) (*UpdateTSIGKeyResponseType, error) {
	response := new(UpdateTSIGKeyResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneTSIGKey(request *DeleteOneTSIGKeyRequestType) (*DeleteOneTSIGKeyResponseType, error) {
	response := new(DeleteOneTSIGKeyResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateZone(request *CreateZoneRequestType) (*CreateZoneResponseType, error) {
	response := new(CreateZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneZone(request *GetOneZoneRequestType) (*GetOneZoneResponseType, error) {
	response := new(GetOneZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetZones(request *GetZonesRequestType) (*GetZonesResponseType, error) {
	response := new(GetZonesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneZone(request *DeleteOneZoneRequestType) (*DeleteOneZoneResponseType, error) {
	response := new(DeleteOneZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateSecondaryZone(request *CreateSecondaryZoneRequestType) (*CreateSecondaryZoneResponseType, error) {
	response := new(CreateSecondaryZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateSecondary(request *UpdateSecondaryRequestType) (*UpdateSecondaryResponseType, error) {
	response := new(UpdateSecondaryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateSecondary(request *ActivateSecondaryRequestType) (*ActivateSecondaryResponseType, error) {
	response := new(ActivateSecondaryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateSecondary(request *DeactivateSecondaryRequestType) (*DeactivateSecondaryResponseType, error) {
	response := new(DeactivateSecondaryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RetransferSecondary(request *RetransferSecondaryRequestType) (*RetransferSecondaryResponseType, error) {
	response := new(RetransferSecondaryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneSecondary(request *GetOneSecondaryRequestType) (*GetOneSecondaryResponseType, error) {
	response := new(GetOneSecondaryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetSecondaries(request *GetSecondariesRequestType) (*GetSecondariesResponseType, error) {
	response := new(GetSecondariesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetZoneApex(request *GetZoneApexRequestType) (*GetZoneApexResponseType, error) {
	response := new(GetZoneApexResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateARecord(request *CreateARecordRequestType) (*CreateARecordResponseType, error) {
	response := new(CreateARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneARecord(request *GetOneARecordRequestType) (*GetOneARecordResponseType, error) {
	response := new(GetOneARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetARecords(request *GetARecordsRequestType) (*GetARecordsResponseType, error) {
	response := new(GetARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateARecord(request *UpdateARecordRequestType) (*UpdateARecordResponseType, error) {
	response := new(UpdateARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteARecords(request *DeleteARecordsRequestType) (*DeleteARecordsResponseType, error) {
	response := new(DeleteARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneARecord(request *DeleteOneARecordRequestType) (*DeleteOneARecordResponseType, error) {
	response := new(DeleteOneARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateAAAARecord(request *CreateAAAARecordRequestType) (*CreateAAAARecordResponseType, error) {
	response := new(CreateAAAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneAAAARecord(request *GetOneAAAARecordRequestType) (*GetOneAAAARecordResponseType, error) {
	response := new(GetOneAAAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetAAAARecords(request *GetAAAARecordsRequestType) (*GetAAAARecordsResponseType, error) {
	response := new(GetAAAARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateAAAARecord(request *UpdateAAAARecordRequestType) (*UpdateAAAARecordResponseType, error) {
	response := new(UpdateAAAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteAAAARecords(request *DeleteAAAARecordsRequestType) (*DeleteAAAARecordsResponseType, error) {
	response := new(DeleteAAAARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneAAAARecord(request *DeleteOneAAAARecordRequestType) (*DeleteOneAAAARecordResponseType, error) {
	response := new(DeleteOneAAAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateALIASRecord(request *CreateALIASRecordRequestType) (*CreateALIASRecordResponseType, error) {
	response := new(CreateALIASRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneALIASRecord(request *GetOneALIASRecordRequestType) (*GetOneALIASRecordResponseType, error) {
	response := new(GetOneALIASRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetALIASRecords(request *GetALIASRecordsRequestType) (*GetALIASRecordsResponseType, error) {
	response := new(GetALIASRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateALIASRecord(request *UpdateALIASRecordRequestType) (*UpdateALIASRecordResponseType, error) {
	response := new(UpdateALIASRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteALIASRecords(request *DeleteALIASRecordsRequestType) (*DeleteALIASRecordsResponseType, error) {
	response := new(DeleteALIASRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneALIASRecord(request *DeleteOneALIASRecordRequestType) (*DeleteOneALIASRecordResponseType, error) {
	response := new(DeleteOneALIASRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCAARecord(request *CreateCAARecordRequestType) (*CreateCAARecordResponseType, error) {
	response := new(CreateCAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCAARecord(request *GetOneCAARecordRequestType) (*GetOneCAARecordResponseType, error) {
	response := new(GetOneCAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCAARecords(request *GetCAARecordsRequestType) (*GetCAARecordsResponseType, error) {
	response := new(GetCAARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCAARecord(request *UpdateCAARecordRequestType) (*UpdateCAARecordResponseType, error) {
	response := new(UpdateCAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCAARecords(request *DeleteCAARecordsRequestType) (*DeleteCAARecordsResponseType, error) {
	response := new(DeleteCAARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCAARecord(request *DeleteOneCAARecordRequestType) (*DeleteOneCAARecordResponseType, error) {
	response := new(DeleteOneCAARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCDNSKEYRecord(request *CreateCDNSKEYRecordRequestType) (*CreateCDNSKEYRecordResponseType, error) {
	response := new(CreateCDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCDNSKEYRecord(request *GetOneCDNSKEYRecordRequestType) (*GetOneCDNSKEYRecordResponseType, error) {
	response := new(GetOneCDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCDNSKEYRecords(request *GetCDNSKEYRecordsRequestType) (*GetCDNSKEYRecordsResponseType, error) {
	response := new(GetCDNSKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCDNSKEYRecord(request *UpdateCDNSKEYRecordRequestType) (*UpdateCDNSKEYRecordResponseType, error) {
	response := new(UpdateCDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCDNSKEYRecords(request *DeleteCDNSKEYRecordsRequestType) (*DeleteCDNSKEYRecordsResponseType, error) {
	response := new(DeleteCDNSKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCDNSKEYRecord(request *DeleteOneCDNSKEYRecordRequestType) (*DeleteOneCDNSKEYRecordResponseType, error) {
	response := new(DeleteOneCDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCDSRecord(request *CreateCDSRecordRequestType) (*CreateCDSRecordResponseType, error) {
	response := new(CreateCDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCDSRecord(request *GetOneCDSRecordRequestType) (*GetOneCDSRecordResponseType, error) {
	response := new(GetOneCDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCDSRecords(request *GetCDSRecordsRequestType) (*GetCDSRecordsResponseType, error) {
	response := new(GetCDSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCDSRecord(request *UpdateCDSRecordRequestType) (*UpdateCDSRecordResponseType, error) {
	response := new(UpdateCDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCDSRecords(request *DeleteCDSRecordsRequestType) (*DeleteCDSRecordsResponseType, error) {
	response := new(DeleteCDSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCDSRecord(request *DeleteOneCDSRecordRequestType) (*DeleteOneCDSRecordResponseType, error) {
	response := new(DeleteOneCDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCERTRecord(request *CreateCERTRecordRequestType) (*CreateCERTRecordResponseType, error) {
	response := new(CreateCERTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCERTRecord(request *GetOneCERTRecordRequestType) (*GetOneCERTRecordResponseType, error) {
	response := new(GetOneCERTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCERTRecords(request *GetCERTRecordsRequestType) (*GetCERTRecordsResponseType, error) {
	response := new(GetCERTRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCERTRecord(request *UpdateCERTRecordRequestType) (*UpdateCERTRecordResponseType, error) {
	response := new(UpdateCERTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCERTRecords(request *DeleteCERTRecordsRequestType) (*DeleteCERTRecordsResponseType, error) {
	response := new(DeleteCERTRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCERTRecord(request *DeleteOneCERTRecordRequestType) (*DeleteOneCERTRecordResponseType, error) {
	response := new(DeleteOneCERTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCNAMERecord(request *CreateCNAMERecordRequestType) (*CreateCNAMERecordResponseType, error) {
	response := new(CreateCNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCNAMERecord(request *GetOneCNAMERecordRequestType) (*GetOneCNAMERecordResponseType, error) {
	response := new(GetOneCNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCNAMERecords(request *GetCNAMERecordsRequestType) (*GetCNAMERecordsResponseType, error) {
	response := new(GetCNAMERecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCNAMERecord(request *UpdateCNAMERecordRequestType) (*UpdateCNAMERecordResponseType, error) {
	response := new(UpdateCNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCNAMERecords(request *DeleteCNAMERecordsRequestType) (*DeleteCNAMERecordsResponseType, error) {
	response := new(DeleteCNAMERecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCNAMERecord(request *DeleteOneCNAMERecordRequestType) (*DeleteOneCNAMERecordResponseType, error) {
	response := new(DeleteOneCNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCSYNCRecord(request *CreateCSYNCRecordRequestType) (*CreateCSYNCRecordResponseType, error) {
	response := new(CreateCSYNCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCSYNCRecord(request *GetOneCSYNCRecordRequestType) (*GetOneCSYNCRecordResponseType, error) {
	response := new(GetOneCSYNCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCSYNCRecords(request *GetCSYNCRecordsRequestType) (*GetCSYNCRecordsResponseType, error) {
	response := new(GetCSYNCRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCSYNCRecord(request *UpdateCSYNCRecordRequestType) (*UpdateCSYNCRecordResponseType, error) {
	response := new(UpdateCSYNCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCSYNCRecords(request *DeleteCSYNCRecordsRequestType) (*DeleteCSYNCRecordsResponseType, error) {
	response := new(DeleteCSYNCRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCSYNCRecord(request *DeleteOneCSYNCRecordRequestType) (*DeleteOneCSYNCRecordResponseType, error) {
	response := new(DeleteOneCSYNCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDHCIDRecord(request *CreateDHCIDRecordRequestType) (*CreateDHCIDRecordResponseType, error) {
	response := new(CreateDHCIDRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDHCIDRecord(request *GetOneDHCIDRecordRequestType) (*GetOneDHCIDRecordResponseType, error) {
	response := new(GetOneDHCIDRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDHCIDRecords(request *GetDHCIDRecordsRequestType) (*GetDHCIDRecordsResponseType, error) {
	response := new(GetDHCIDRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDHCIDRecord(request *UpdateDHCIDRecordRequestType) (*UpdateDHCIDRecordResponseType, error) {
	response := new(UpdateDHCIDRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteDHCIDRecords(request *DeleteDHCIDRecordsRequestType) (*DeleteDHCIDRecordsResponseType, error) {
	response := new(DeleteDHCIDRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDHCIDRecord(request *DeleteOneDHCIDRecordRequestType) (*DeleteOneDHCIDRecordResponseType, error) {
	response := new(DeleteOneDHCIDRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDNAMERecord(request *CreateDNAMERecordRequestType) (*CreateDNAMERecordResponseType, error) {
	response := new(CreateDNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDNAMERecord(request *GetOneDNAMERecordRequestType) (*GetOneDNAMERecordResponseType, error) {
	response := new(GetOneDNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDNAMERecords(request *GetDNAMERecordsRequestType) (*GetDNAMERecordsResponseType, error) {
	response := new(GetDNAMERecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDNAMERecord(request *UpdateDNAMERecordRequestType) (*UpdateDNAMERecordResponseType, error) {
	response := new(UpdateDNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteDNAMERecords(request *DeleteDNAMERecordsRequestType) (*DeleteDNAMERecordsResponseType, error) {
	response := new(DeleteDNAMERecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDNAMERecord(request *DeleteOneDNAMERecordRequestType) (*DeleteOneDNAMERecordResponseType, error) {
	response := new(DeleteOneDNAMERecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDNSKEYRecord(request *CreateDNSKEYRecordRequestType) (*CreateDNSKEYRecordResponseType, error) {
	response := new(CreateDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDNSKEYRecord(request *GetOneDNSKEYRecordRequestType) (*GetOneDNSKEYRecordResponseType, error) {
	response := new(GetOneDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDNSKEYRecords(request *GetDNSKEYRecordsRequestType) (*GetDNSKEYRecordsResponseType, error) {
	response := new(GetDNSKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDNSKEYRecord(request *UpdateDNSKEYRecordRequestType) (*UpdateDNSKEYRecordResponseType, error) {
	response := new(UpdateDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteDNSKEYRecords(request *DeleteDNSKEYRecordsRequestType) (*DeleteDNSKEYRecordsResponseType, error) {
	response := new(DeleteDNSKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDNSKEYRecord(request *DeleteOneDNSKEYRecordRequestType) (*DeleteOneDNSKEYRecordResponseType, error) {
	response := new(DeleteOneDNSKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDSRecord(request *CreateDSRecordRequestType) (*CreateDSRecordResponseType, error) {
	response := new(CreateDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDSRecord(request *GetOneDSRecordRequestType) (*GetOneDSRecordResponseType, error) {
	response := new(GetOneDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDSRecords(request *GetDSRecordsRequestType) (*GetDSRecordsResponseType, error) {
	response := new(GetDSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDSRecord(request *UpdateDSRecordRequestType) (*UpdateDSRecordResponseType, error) {
	response := new(UpdateDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteDSRecords(request *DeleteDSRecordsRequestType) (*DeleteDSRecordsResponseType, error) {
	response := new(DeleteDSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDSRecord(request *DeleteOneDSRecordRequestType) (*DeleteOneDSRecordResponseType, error) {
	response := new(DeleteOneDSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateIPSECKEYRecord(request *CreateIPSECKEYRecordRequestType) (*CreateIPSECKEYRecordResponseType, error) {
	response := new(CreateIPSECKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneIPSECKEYRecord(request *GetOneIPSECKEYRecordRequestType) (*GetOneIPSECKEYRecordResponseType, error) {
	response := new(GetOneIPSECKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetIPSECKEYRecords(request *GetIPSECKEYRecordsRequestType) (*GetIPSECKEYRecordsResponseType, error) {
	response := new(GetIPSECKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateIPSECKEYRecord(request *UpdateIPSECKEYRecordRequestType) (*UpdateIPSECKEYRecordResponseType, error) {
	response := new(UpdateIPSECKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteIPSECKEYRecords(request *DeleteIPSECKEYRecordsRequestType) (*DeleteIPSECKEYRecordsResponseType, error) {
	response := new(DeleteIPSECKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneIPSECKEYRecord(request *DeleteOneIPSECKEYRecordRequestType) (*DeleteOneIPSECKEYRecordResponseType, error) {
	response := new(DeleteOneIPSECKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateKEYRecord(request *CreateKEYRecordRequestType) (*CreateKEYRecordResponseType, error) {
	response := new(CreateKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneKEYRecord(request *GetOneKEYRecordRequestType) (*GetOneKEYRecordResponseType, error) {
	response := new(GetOneKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetKEYRecords(request *GetKEYRecordsRequestType) (*GetKEYRecordsResponseType, error) {
	response := new(GetKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateKEYRecord(request *UpdateKEYRecordRequestType) (*UpdateKEYRecordResponseType, error) {
	response := new(UpdateKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteKEYRecords(request *DeleteKEYRecordsRequestType) (*DeleteKEYRecordsResponseType, error) {
	response := new(DeleteKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneKEYRecord(request *DeleteOneKEYRecordRequestType) (*DeleteOneKEYRecordResponseType, error) {
	response := new(DeleteOneKEYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateKXRecord(request *CreateKXRecordRequestType) (*CreateKXRecordResponseType, error) {
	response := new(CreateKXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneKXRecord(request *GetOneKXRecordRequestType) (*GetOneKXRecordResponseType, error) {
	response := new(GetOneKXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetKXRecords(request *GetKXRecordsRequestType) (*GetKXRecordsResponseType, error) {
	response := new(GetKXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateKXRecord(request *UpdateKXRecordRequestType) (*UpdateKXRecordResponseType, error) {
	response := new(UpdateKXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteKXRecords(request *DeleteKXRecordsRequestType) (*DeleteKXRecordsResponseType, error) {
	response := new(DeleteKXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneKXRecord(request *DeleteOneKXRecordRequestType) (*DeleteOneKXRecordResponseType, error) {
	response := new(DeleteOneKXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateLOCRecord(request *CreateLOCRecordRequestType) (*CreateLOCRecordResponseType, error) {
	response := new(CreateLOCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneLOCRecord(request *GetOneLOCRecordRequestType) (*GetOneLOCRecordResponseType, error) {
	response := new(GetOneLOCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetLOCRecords(request *GetLOCRecordsRequestType) (*GetLOCRecordsResponseType, error) {
	response := new(GetLOCRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateLOCRecord(request *UpdateLOCRecordRequestType) (*UpdateLOCRecordResponseType, error) {
	response := new(UpdateLOCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteLOCRecords(request *DeleteLOCRecordsRequestType) (*DeleteLOCRecordsResponseType, error) {
	response := new(DeleteLOCRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneLOCRecord(request *DeleteOneLOCRecordRequestType) (*DeleteOneLOCRecordResponseType, error) {
	response := new(DeleteOneLOCRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateMXRecord(request *CreateMXRecordRequestType) (*CreateMXRecordResponseType, error) {
	response := new(CreateMXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneMXRecord(request *GetOneMXRecordRequestType) (*GetOneMXRecordResponseType, error) {
	response := new(GetOneMXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetMXRecords(request *GetMXRecordsRequestType) (*GetMXRecordsResponseType, error) {
	response := new(GetMXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateMXRecord(request *UpdateMXRecordRequestType) (*UpdateMXRecordResponseType, error) {
	response := new(UpdateMXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteMXRecords(request *DeleteMXRecordsRequestType) (*DeleteMXRecordsResponseType, error) {
	response := new(DeleteMXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneMXRecord(request *DeleteOneMXRecordRequestType) (*DeleteOneMXRecordResponseType, error) {
	response := new(DeleteOneMXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateNAPTRRecord(request *CreateNAPTRRecordRequestType) (*CreateNAPTRRecordResponseType, error) {
	response := new(CreateNAPTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneNAPTRRecord(request *GetOneNAPTRRecordRequestType) (*GetOneNAPTRRecordResponseType, error) {
	response := new(GetOneNAPTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetNAPTRRecords(request *GetNAPTRRecordsRequestType) (*GetNAPTRRecordsResponseType, error) {
	response := new(GetNAPTRRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateNAPTRRecord(request *UpdateNAPTRRecordRequestType) (*UpdateNAPTRRecordResponseType, error) {
	response := new(UpdateNAPTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteNAPTRRecords(request *DeleteNAPTRRecordsRequestType) (*DeleteNAPTRRecordsResponseType, error) {
	response := new(DeleteNAPTRRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneNAPTRRecord(request *DeleteOneNAPTRRecordRequestType) (*DeleteOneNAPTRRecordResponseType, error) {
	response := new(DeleteOneNAPTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateNSAPRecord(request *CreateNSAPRecordRequestType) (*CreateNSAPRecordResponseType, error) {
	response := new(CreateNSAPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneNSAPRecord(request *GetOneNSAPRecordRequestType) (*GetOneNSAPRecordResponseType, error) {
	response := new(GetOneNSAPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetNSAPRecords(request *GetNSAPRecordsRequestType) (*GetNSAPRecordsResponseType, error) {
	response := new(GetNSAPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateNSAPRecord(request *UpdateNSAPRecordRequestType) (*UpdateNSAPRecordResponseType, error) {
	response := new(UpdateNSAPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteNSAPRecords(request *DeleteNSAPRecordsRequestType) (*DeleteNSAPRecordsResponseType, error) {
	response := new(DeleteNSAPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneNSAPRecord(request *DeleteOneNSAPRecordRequestType) (*DeleteOneNSAPRecordResponseType, error) {
	response := new(DeleteOneNSAPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreatePOLICYRecord(request *CreatePOLICYRecordRequestType) (*CreatePOLICYRecordResponseType, error) {
	response := new(CreatePOLICYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOnePOLICYRecord(request *GetOnePOLICYRecordRequestType) (*GetOnePOLICYRecordResponseType, error) {
	response := new(GetOnePOLICYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetPOLICYRecords(request *GetPOLICYRecordsRequestType) (*GetPOLICYRecordsResponseType, error) {
	response := new(GetPOLICYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdatePOLICYRecord(request *UpdatePOLICYRecordRequestType) (*UpdatePOLICYRecordResponseType, error) {
	response := new(UpdatePOLICYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeletePOLICYRecords(request *DeletePOLICYRecordsRequestType) (*DeletePOLICYRecordsResponseType, error) {
	response := new(DeletePOLICYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOnePOLICYRecord(request *DeleteOnePOLICYRecordRequestType) (*DeleteOnePOLICYRecordResponseType, error) {
	response := new(DeleteOnePOLICYRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreatePTRRecord(request *CreatePTRRecordRequestType) (*CreatePTRRecordResponseType, error) {
	response := new(CreatePTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOnePTRRecord(request *GetOnePTRRecordRequestType) (*GetOnePTRRecordResponseType, error) {
	response := new(GetOnePTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetPTRRecords(request *GetPTRRecordsRequestType) (*GetPTRRecordsResponseType, error) {
	response := new(GetPTRRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdatePTRRecord(request *UpdatePTRRecordRequestType) (*UpdatePTRRecordResponseType, error) {
	response := new(UpdatePTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeletePTRRecords(request *DeletePTRRecordsRequestType) (*DeletePTRRecordsResponseType, error) {
	response := new(DeletePTRRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOnePTRRecord(request *DeleteOnePTRRecordRequestType) (*DeleteOnePTRRecordResponseType, error) {
	response := new(DeleteOnePTRRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreatePXRecord(request *CreatePXRecordRequestType) (*CreatePXRecordResponseType, error) {
	response := new(CreatePXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOnePXRecord(request *GetOnePXRecordRequestType) (*GetOnePXRecordResponseType, error) {
	response := new(GetOnePXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetPXRecords(request *GetPXRecordsRequestType) (*GetPXRecordsResponseType, error) {
	response := new(GetPXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdatePXRecord(request *UpdatePXRecordRequestType) (*UpdatePXRecordResponseType, error) {
	response := new(UpdatePXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeletePXRecords(request *DeletePXRecordsRequestType) (*DeletePXRecordsResponseType, error) {
	response := new(DeletePXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOnePXRecord(request *DeleteOnePXRecordRequestType) (*DeleteOnePXRecordResponseType, error) {
	response := new(DeleteOnePXRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateRPRecord(request *CreateRPRecordRequestType) (*CreateRPRecordResponseType, error) {
	response := new(CreateRPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneRPRecord(request *GetOneRPRecordRequestType) (*GetOneRPRecordResponseType, error) {
	response := new(GetOneRPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetRPRecords(request *GetRPRecordsRequestType) (*GetRPRecordsResponseType, error) {
	response := new(GetRPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateRPRecord(request *UpdateRPRecordRequestType) (*UpdateRPRecordResponseType, error) {
	response := new(UpdateRPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteRPRecords(request *DeleteRPRecordsRequestType) (*DeleteRPRecordsResponseType, error) {
	response := new(DeleteRPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneRPRecord(request *DeleteOneRPRecordRequestType) (*DeleteOneRPRecordResponseType, error) {
	response := new(DeleteOneRPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateSPFRecord(request *CreateSPFRecordRequestType) (*CreateSPFRecordResponseType, error) {
	response := new(CreateSPFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneSPFRecord(request *GetOneSPFRecordRequestType) (*GetOneSPFRecordResponseType, error) {
	response := new(GetOneSPFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetSPFRecords(request *GetSPFRecordsRequestType) (*GetSPFRecordsResponseType, error) {
	response := new(GetSPFRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateSPFRecord(request *UpdateSPFRecordRequestType) (*UpdateSPFRecordResponseType, error) {
	response := new(UpdateSPFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteSPFRecords(request *DeleteSPFRecordsRequestType) (*DeleteSPFRecordsResponseType, error) {
	response := new(DeleteSPFRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneSPFRecord(request *DeleteOneSPFRecordRequestType) (*DeleteOneSPFRecordResponseType, error) {
	response := new(DeleteOneSPFRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateSRVRecord(request *CreateSRVRecordRequestType) (*CreateSRVRecordResponseType, error) {
	response := new(CreateSRVRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneSRVRecord(request *GetOneSRVRecordRequestType) (*GetOneSRVRecordResponseType, error) {
	response := new(GetOneSRVRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetSRVRecords(request *GetSRVRecordsRequestType) (*GetSRVRecordsResponseType, error) {
	response := new(GetSRVRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateSRVRecord(request *UpdateSRVRecordRequestType) (*UpdateSRVRecordResponseType, error) {
	response := new(UpdateSRVRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteSRVRecords(request *DeleteSRVRecordsRequestType) (*DeleteSRVRecordsResponseType, error) {
	response := new(DeleteSRVRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneSRVRecord(request *DeleteOneSRVRecordRequestType) (*DeleteOneSRVRecordResponseType, error) {
	response := new(DeleteOneSRVRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateSSHFPRecord(request *CreateSSHFPRecordRequestType) (*CreateSSHFPRecordResponseType, error) {
	response := new(CreateSSHFPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneSSHFPRecord(request *GetOneSSHFPRecordRequestType) (*GetOneSSHFPRecordResponseType, error) {
	response := new(GetOneSSHFPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetSSHFPRecords(request *GetSSHFPRecordsRequestType) (*GetSSHFPRecordsResponseType, error) {
	response := new(GetSSHFPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateSSHFPRecord(request *UpdateSSHFPRecordRequestType) (*UpdateSSHFPRecordResponseType, error) {
	response := new(UpdateSSHFPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteSSHFPRecords(request *DeleteSSHFPRecordsRequestType) (*DeleteSSHFPRecordsResponseType, error) {
	response := new(DeleteSSHFPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneSSHFPRecord(request *DeleteOneSSHFPRecordRequestType) (*DeleteOneSSHFPRecordResponseType, error) {
	response := new(DeleteOneSSHFPRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateTLSARecord(request *CreateTLSARecordRequestType) (*CreateTLSARecordResponseType, error) {
	response := new(CreateTLSARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneTLSARecord(request *GetOneTLSARecordRequestType) (*GetOneTLSARecordResponseType, error) {
	response := new(GetOneTLSARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetTLSARecords(request *GetTLSARecordsRequestType) (*GetTLSARecordsResponseType, error) {
	response := new(GetTLSARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateTLSARecord(request *UpdateTLSARecordRequestType) (*UpdateTLSARecordResponseType, error) {
	response := new(UpdateTLSARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteTLSARecords(request *DeleteTLSARecordsRequestType) (*DeleteTLSARecordsResponseType, error) {
	response := new(DeleteTLSARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneTLSARecord(request *DeleteOneTLSARecordRequestType) (*DeleteOneTLSARecordResponseType, error) {
	response := new(DeleteOneTLSARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateTXTRecord(request *CreateTXTRecordRequestType) (*CreateTXTRecordResponseType, error) {
	response := new(CreateTXTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneTXTRecord(request *GetOneTXTRecordRequestType) (*GetOneTXTRecordResponseType, error) {
	response := new(GetOneTXTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetTXTRecords(request *GetTXTRecordsRequestType) (*GetTXTRecordsResponseType, error) {
	response := new(GetTXTRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateTXTRecord(request *UpdateTXTRecordRequestType) (*UpdateTXTRecordResponseType, error) {
	response := new(UpdateTXTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteTXTRecords(request *DeleteTXTRecordsRequestType) (*DeleteTXTRecordsResponseType, error) {
	response := new(DeleteTXTRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneTXTRecord(request *DeleteOneTXTRecordRequestType) (*DeleteOneTXTRecordResponseType, error) {
	response := new(DeleteOneTXTRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneSOARecord(request *GetOneSOARecordRequestType) (*GetOneSOARecordResponseType, error) {
	response := new(GetOneSOARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetSOARecords(request *GetSOARecordsRequestType) (*GetSOARecordsResponseType, error) {
	response := new(GetSOARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateSOARecord(request *UpdateSOARecordRequestType) (*UpdateSOARecordResponseType, error) {
	response := new(UpdateSOARecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateNSRecord(request *CreateNSRecordRequestType) (*CreateNSRecordResponseType, error) {
	response := new(CreateNSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneNSRecord(request *GetOneNSRecordRequestType) (*GetOneNSRecordResponseType, error) {
	response := new(GetOneNSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetNSRecords(request *GetNSRecordsRequestType) (*GetNSRecordsResponseType, error) {
	response := new(GetNSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateNSRecord(request *UpdateNSRecordRequestType) (*UpdateNSRecordResponseType, error) {
	response := new(UpdateNSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteNSRecords(request *DeleteNSRecordsRequestType) (*DeleteNSRecordsResponseType, error) {
	response := new(DeleteNSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneNSRecord(request *DeleteOneNSRecordRequestType) (*DeleteOneNSRecordResponseType, error) {
	response := new(DeleteOneNSRecordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceARecords(request *ReplaceARecordsRequestType) (*ReplaceARecordsResponseType, error) {
	response := new(ReplaceARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceAAAARecords(request *ReplaceAAAARecordsRequestType) (*ReplaceAAAARecordsResponseType, error) {
	response := new(ReplaceAAAARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceALIASRecords(request *ReplaceALIASRecordsRequestType) (*ReplaceALIASRecordsResponseType, error) {
	response := new(ReplaceALIASRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceCAARecords(request *ReplaceCAARecordsRequestType) (*ReplaceCAARecordsResponseType, error) {
	response := new(ReplaceCAARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceCDNSKEYRecords(request *ReplaceCDNSKEYRecordsRequestType) (*ReplaceCDNSKEYRecordsResponseType, error) {
	response := new(ReplaceCDNSKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceCDSRecords(request *ReplaceCDSRecordsRequestType) (*ReplaceCDSRecordsResponseType, error) {
	response := new(ReplaceCDSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceCERTRecords(request *ReplaceCERTRecordsRequestType) (*ReplaceCERTRecordsResponseType, error) {
	response := new(ReplaceCERTRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceCNAMERecords(request *ReplaceCNAMERecordsRequestType) (*ReplaceCNAMERecordsResponseType, error) {
	response := new(ReplaceCNAMERecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceCSYNCRecords(request *ReplaceCSYNCRecordsRequestType) (*ReplaceCSYNCRecordsResponseType, error) {
	response := new(ReplaceCSYNCRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceDHCIDRecords(request *ReplaceDHCIDRecordsRequestType) (*ReplaceDHCIDRecordsResponseType, error) {
	response := new(ReplaceDHCIDRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceDNAMERecords(request *ReplaceDNAMERecordsRequestType) (*ReplaceDNAMERecordsResponseType, error) {
	response := new(ReplaceDNAMERecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceDNSKEYRecords(request *ReplaceDNSKEYRecordsRequestType) (*ReplaceDNSKEYRecordsResponseType, error) {
	response := new(ReplaceDNSKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceDSRecords(request *ReplaceDSRecordsRequestType) (*ReplaceDSRecordsResponseType, error) {
	response := new(ReplaceDSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceIPSECKEYRecords(request *ReplaceIPSECKEYRecordsRequestType) (*ReplaceIPSECKEYRecordsResponseType, error) {
	response := new(ReplaceIPSECKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceKEYRecords(request *ReplaceKEYRecordsRequestType) (*ReplaceKEYRecordsResponseType, error) {
	response := new(ReplaceKEYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceKXRecords(request *ReplaceKXRecordsRequestType) (*ReplaceKXRecordsResponseType, error) {
	response := new(ReplaceKXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceLOCRecords(request *ReplaceLOCRecordsRequestType) (*ReplaceLOCRecordsResponseType, error) {
	response := new(ReplaceLOCRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceMXRecords(request *ReplaceMXRecordsRequestType) (*ReplaceMXRecordsResponseType, error) {
	response := new(ReplaceMXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceNAPTRRecords(request *ReplaceNAPTRRecordsRequestType) (*ReplaceNAPTRRecordsResponseType, error) {
	response := new(ReplaceNAPTRRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceNSAPRecords(request *ReplaceNSAPRecordsRequestType) (*ReplaceNSAPRecordsResponseType, error) {
	response := new(ReplaceNSAPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplacePOLICYRecords(request *ReplacePOLICYRecordsRequestType) (*ReplacePOLICYRecordsResponseType, error) {
	response := new(ReplacePOLICYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplacePTRRecords(request *ReplacePTRRecordsRequestType) (*ReplacePTRRecordsResponseType, error) {
	response := new(ReplacePTRRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplacePXRecords(request *ReplacePXRecordsRequestType) (*ReplacePXRecordsResponseType, error) {
	response := new(ReplacePXRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceRPRecords(request *ReplaceRPRecordsRequestType) (*ReplaceRPRecordsResponseType, error) {
	response := new(ReplaceRPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceSPFRecords(request *ReplaceSPFRecordsRequestType) (*ReplaceSPFRecordsResponseType, error) {
	response := new(ReplaceSPFRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceSRVRecords(request *ReplaceSRVRecordsRequestType) (*ReplaceSRVRecordsResponseType, error) {
	response := new(ReplaceSRVRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceSSHFPRecords(request *ReplaceSSHFPRecordsRequestType) (*ReplaceSSHFPRecordsResponseType, error) {
	response := new(ReplaceSSHFPRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceTLSARecords(request *ReplaceTLSARecordsRequestType) (*ReplaceTLSARecordsResponseType, error) {
	response := new(ReplaceTLSARecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceTXTRecords(request *ReplaceTXTRecordsRequestType) (*ReplaceTXTRecordsResponseType, error) {
	response := new(ReplaceTXTRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ReplaceNSRecords(request *ReplaceNSRecordsRequestType) (*ReplaceNSRecordsResponseType, error) {
	response := new(ReplaceNSRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetANYRecords(request *GetANYRecordsRequestType) (*GetANYRecordsResponseType, error) {
	response := new(GetANYRecordsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetAllAliasQNames(request *GetAllAliasQNamesRequestType) (*GetAllAliasQNamesResponseType, error) {
	response := new(GetAllAliasQNamesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneUser(request *GetOneUserRequestType) (*GetOneUserResponseType, error) {
	response := new(GetOneUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneUser(request *DeleteOneUserRequestType) (*DeleteOneUserResponseType, error) {
	response := new(DeleteOneUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateUser(request *CreateUserRequestType) (*CreateUserResponseType, error) {
	response := new(CreateUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateUser(request *UpdateUserRequestType) (*UpdateUserResponseType, error) {
	response := new(UpdateUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetUsers(request *GetUsersRequestType) (*GetUsersResponseType, error) {
	response := new(GetUsersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetUpdateUsers(request *GetUpdateUsersRequestType) (*GetUpdateUsersResponseType, error) {
	response := new(GetUpdateUsersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateUpdateUser(request *UpdateUpdateUserRequestType) (*UpdateUpdateUserResponseType, error) {
	response := new(UpdateUpdateUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneUpdateUser(request *DeleteOneUpdateUserRequestType) (*DeleteOneUpdateUserResponseType, error) {
	response := new(DeleteOneUpdateUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateUserPassword(request *UpdateUserPasswordRequestType) (*UpdateUserPasswordResponseType, error) {
	response := new(UpdateUserPasswordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) BlockUser(request *BlockUserRequestType) (*BlockUserResponseType, error) {
	response := new(BlockUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UnblockUser(request *UnblockUserRequestType) (*UnblockUserResponseType, error) {
	response := new(UnblockUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateContact(request *CreateContactRequestType) (*CreateContactResponseType, error) {
	response := new(CreateContactResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneContact(request *GetOneContactRequestType) (*GetOneContactResponseType, error) {
	response := new(GetOneContactResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetContacts(request *GetContactsRequestType) (*GetContactsResponseType, error) {
	response := new(GetContactsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneContact(request *DeleteOneContactRequestType) (*DeleteOneContactResponseType, error) {
	response := new(DeleteOneContactResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateContact(request *UpdateContactRequestType) (*UpdateContactResponseType, error) {
	response := new(UpdateContactResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCustomer(request *CreateCustomerRequestType) (*CreateCustomerResponseType, error) {
	response := new(CreateCustomerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCustomer(request *UpdateCustomerRequestType) (*UpdateCustomerResponseType, error) {
	response := new(UpdateCustomerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCustomer(request *GetOneCustomerRequestType) (*GetOneCustomerResponseType, error) {
	response := new(GetOneCustomerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCustomers(request *GetCustomersRequestType) (*GetCustomersResponseType, error) {
	response := new(GetCustomersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCustomer(request *DeleteOneCustomerRequestType) (*DeleteOneCustomerResponseType, error) {
	response := new(DeleteOneCustomerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCustomerPrefs(request *GetCustomerPrefsRequestType) (*GetCustomerPrefsResponseType, error) {
	response := new(GetCustomerPrefsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetCustomerPrefs(request *SetCustomerPrefsRequestType) (*SetCustomerPrefsResponseType, error) {
	response := new(SetCustomerPrefsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCustomerIPACL(request *GetCustomerIPACLRequestType) (*GetCustomerIPACLResponseType, error) {
	response := new(GetCustomerIPACLResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetCustomerIPACL(request *SetCustomerIPACLRequestType) (*SetCustomerIPACLResponseType, error) {
	response := new(SetCustomerIPACLResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCustomerOracleMetadata(request *CreateCustomerOracleMetadataRequestType) (*CreateCustomerOracleMetadataResponseType, error) {
	response := new(CreateCustomerOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCustomerOracleMetadata(request *UpdateCustomerOracleMetadataRequestType) (*UpdateCustomerOracleMetadataResponseType, error) {
	response := new(UpdateCustomerOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCustomerOracleMetadata(request *GetCustomerOracleMetadataRequestType) (*GetCustomerOracleMetadataResponseType, error) {
	response := new(GetCustomerOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteCustomerOracleMetadata(request *DeleteCustomerOracleMetadataRequestType) (*DeleteCustomerOracleMetadataResponseType, error) {
	response := new(DeleteCustomerOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateZoneOracleMetadata(request *CreateZoneOracleMetadataRequestType) (*CreateZoneOracleMetadataResponseType, error) {
	response := new(CreateZoneOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateZoneOracleMetadata(request *UpdateZoneOracleMetadataRequestType) (*UpdateZoneOracleMetadataResponseType, error) {
	response := new(UpdateZoneOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetZoneOracleMetadata(request *GetZoneOracleMetadataRequestType) (*GetZoneOracleMetadataResponseType, error) {
	response := new(GetZoneOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteZoneOracleMetadata(request *DeleteZoneOracleMetadataRequestType) (*DeleteZoneOracleMetadataResponseType, error) {
	response := new(DeleteZoneOracleMetadataResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDDNS(request *CreateDDNSRequestType) (*CreateDDNSResponseType, error) {
	response := new(CreateDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDDNS(request *GetOneDDNSRequestType) (*GetOneDDNSResponseType, error) {
	response := new(GetOneDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDDNSs(request *GetDDNSsRequestType) (*GetDDNSsResponseType, error) {
	response := new(GetDDNSsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDDNS(request *UpdateDDNSRequestType) (*UpdateDDNSResponseType, error) {
	response := new(UpdateDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDDNS(request *DeleteOneDDNSRequestType) (*DeleteOneDDNSResponseType, error) {
	response := new(DeleteOneDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateDDNS(request *ActivateDDNSRequestType) (*ActivateDDNSResponseType, error) {
	response := new(ActivateDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateDDNS(request *DeactivateDDNSRequestType) (*DeactivateDDNSResponseType, error) {
	response := new(DeactivateDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ResetDDNS(request *ResetDDNSRequestType) (*ResetDDNSResponseType, error) {
	response := new(ResetDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetUpdateUserPassword(request *GetUpdateUserPasswordRequestType) (*GetUpdateUserPasswordResponseType, error) {
	response := new(GetUpdateUserPasswordResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDDNSHost(request *CreateDDNSHostRequestType) (*CreateDDNSHostResponseType, error) {
	response := new(CreateDDNSHostResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateUpdateUser(request *CreateUpdateUserRequestType) (*CreateUpdateUserResponseType, error) {
	response := new(CreateUpdateUserResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) AddDDNS(request *AddDDNSRequestType) (*AddDDNSResponseType, error) {
	response := new(AddDDNSResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateFailover(request *CreateFailoverRequestType) (*CreateFailoverResponseType, error) {
	response := new(CreateFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneFailover(request *GetOneFailoverRequestType) (*GetOneFailoverResponseType, error) {
	response := new(GetOneFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetFailovers(request *GetFailoversRequestType) (*GetFailoversResponseType, error) {
	response := new(GetFailoversResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateFailover(request *UpdateFailoverRequestType) (*UpdateFailoverResponseType, error) {
	response := new(UpdateFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneFailover(request *DeleteOneFailoverRequestType) (*DeleteOneFailoverResponseType, error) {
	response := new(DeleteOneFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateFailover(request *ActivateFailoverRequestType) (*ActivateFailoverResponseType, error) {
	response := new(ActivateFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateFailover(request *DeactivateFailoverRequestType) (*DeactivateFailoverResponseType, error) {
	response := new(DeactivateFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverFailover(request *RecoverFailoverRequestType) (*RecoverFailoverResponseType, error) {
	response := new(RecoverFailoverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateLoadBalance(request *CreateLoadBalanceRequestType) (*CreateLoadBalanceResponseType, error) {
	response := new(CreateLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneLoadBalance(request *GetOneLoadBalanceRequestType) (*GetOneLoadBalanceResponseType, error) {
	response := new(GetOneLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetLoadBalances(request *GetLoadBalancesRequestType) (*GetLoadBalancesResponseType, error) {
	response := new(GetLoadBalancesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateLoadBalance(request *UpdateLoadBalanceRequestType) (*UpdateLoadBalanceResponseType, error) {
	response := new(UpdateLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneLoadBalance(request *DeleteOneLoadBalanceRequestType) (*DeleteOneLoadBalanceResponseType, error) {
	response := new(DeleteOneLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateLoadBalance(request *ActivateLoadBalanceRequestType) (*ActivateLoadBalanceResponseType, error) {
	response := new(ActivateLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateLoadBalance(request *DeactivateLoadBalanceRequestType) (*DeactivateLoadBalanceResponseType, error) {
	response := new(DeactivateLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverLoadBalance(request *RecoverLoadBalanceRequestType) (*RecoverLoadBalanceResponseType, error) {
	response := new(RecoverLoadBalanceResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverLoadBalanceIP(request *RecoverLoadBalanceIPRequestType) (*RecoverLoadBalanceIPResponseType, error) {
	response := new(RecoverLoadBalanceIPResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateLoadBalancePoolEntry(request *CreateLoadBalancePoolEntryRequestType) (*CreateLoadBalancePoolEntryResponseType, error) {
	response := new(CreateLoadBalancePoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateLoadBalancePoolEntry(request *UpdateLoadBalancePoolEntryRequestType) (*UpdateLoadBalancePoolEntryResponseType, error) {
	response := new(UpdateLoadBalancePoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneLoadBalancePoolEntry(request *GetOneLoadBalancePoolEntryRequestType) (*GetOneLoadBalancePoolEntryResponseType, error) {
	response := new(GetOneLoadBalancePoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetLoadBalancePoolEntries(request *GetLoadBalancePoolEntriesRequestType) (*GetLoadBalancePoolEntriesResponseType, error) {
	response := new(GetLoadBalancePoolEntriesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneLoadBalancePoolEntry(request *DeleteOneLoadBalancePoolEntryRequestType) (*DeleteOneLoadBalancePoolEntryResponseType, error) {
	response := new(DeleteOneLoadBalancePoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateGSLB(request *CreateGSLBRequestType) (*CreateGSLBResponseType, error) {
	response := new(CreateGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneGSLB(request *GetOneGSLBRequestType) (*GetOneGSLBResponseType, error) {
	response := new(GetOneGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetGSLBs(request *GetGSLBsRequestType) (*GetGSLBsResponseType, error) {
	response := new(GetGSLBsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateGSLB(request *UpdateGSLBRequestType) (*UpdateGSLBResponseType, error) {
	response := new(UpdateGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneGSLB(request *DeleteOneGSLBRequestType) (*DeleteOneGSLBResponseType, error) {
	response := new(DeleteOneGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateGSLB(request *ActivateGSLBRequestType) (*ActivateGSLBResponseType, error) {
	response := new(ActivateGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateGSLB(request *DeactivateGSLBRequestType) (*DeactivateGSLBResponseType, error) {
	response := new(DeactivateGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverGSLB(request *RecoverGSLBRequestType) (*RecoverGSLBResponseType, error) {
	response := new(RecoverGSLBResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverGSLBIP(request *RecoverGSLBIPRequestType) (*RecoverGSLBIPResponseType, error) {
	response := new(RecoverGSLBIPResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateGSLBRegion(request *CreateGSLBRegionRequestType) (*CreateGSLBRegionResponseType, error) {
	response := new(CreateGSLBRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneGSLBRegion(request *GetOneGSLBRegionRequestType) (*GetOneGSLBRegionResponseType, error) {
	response := new(GetOneGSLBRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetGSLBRegions(request *GetGSLBRegionsRequestType) (*GetGSLBRegionsResponseType, error) {
	response := new(GetGSLBRegionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateGSLBRegion(request *UpdateGSLBRegionRequestType) (*UpdateGSLBRegionResponseType, error) {
	response := new(UpdateGSLBRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneGSLBRegion(request *DeleteOneGSLBRegionRequestType) (*DeleteOneGSLBRegionResponseType, error) {
	response := new(DeleteOneGSLBRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateGSLBRegionPoolEntry(request *CreateGSLBRegionPoolEntryRequestType) (*CreateGSLBRegionPoolEntryResponseType, error) {
	response := new(CreateGSLBRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateGSLBRegionPoolEntry(request *UpdateGSLBRegionPoolEntryRequestType) (*UpdateGSLBRegionPoolEntryResponseType, error) {
	response := new(UpdateGSLBRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneGSLBRegionPoolEntry(request *GetOneGSLBRegionPoolEntryRequestType) (*GetOneGSLBRegionPoolEntryResponseType, error) {
	response := new(GetOneGSLBRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetGSLBRegionPoolEntries(request *GetGSLBRegionPoolEntriesRequestType) (*GetGSLBRegionPoolEntriesResponseType, error) {
	response := new(GetGSLBRegionPoolEntriesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneGSLBRegionPoolEntry(request *DeleteOneGSLBRegionPoolEntryRequestType) (*DeleteOneGSLBRegionPoolEntryResponseType, error) {
	response := new(DeleteOneGSLBRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateRTTM(request *CreateRTTMRequestType) (*CreateRTTMResponseType, error) {
	response := new(CreateRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneRTTM(request *GetOneRTTMRequestType) (*GetOneRTTMResponseType, error) {
	response := new(GetOneRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetRTTMs(request *GetRTTMsRequestType) (*GetRTTMsResponseType, error) {
	response := new(GetRTTMsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateRTTM(request *UpdateRTTMRequestType) (*UpdateRTTMResponseType, error) {
	response := new(UpdateRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneRTTM(request *DeleteOneRTTMRequestType) (*DeleteOneRTTMResponseType, error) {
	response := new(DeleteOneRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateRTTM(request *ActivateRTTMRequestType) (*ActivateRTTMResponseType, error) {
	response := new(ActivateRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateRTTM(request *DeactivateRTTMRequestType) (*DeactivateRTTMResponseType, error) {
	response := new(DeactivateRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverRTTM(request *RecoverRTTMRequestType) (*RecoverRTTMResponseType, error) {
	response := new(RecoverRTTMResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverRTTMIP(request *RecoverRTTMIPRequestType) (*RecoverRTTMIPResponseType, error) {
	response := new(RecoverRTTMIPResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetRTTMLogs(request *GetRTTMLogsRequestType) (*GetRTTMLogsResponseType, error) {
	response := new(GetRTTMLogsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetRTTMRRSets(request *GetRTTMRRSetsRequestType) (*GetRTTMRRSetsResponseType, error) {
	response := new(GetRTTMRRSetsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateRTTMRegion(request *CreateRTTMRegionRequestType) (*CreateRTTMRegionResponseType, error) {
	response := new(CreateRTTMRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneRTTMRegion(request *GetOneRTTMRegionRequestType) (*GetOneRTTMRegionResponseType, error) {
	response := new(GetOneRTTMRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetRTTMRegions(request *GetRTTMRegionsRequestType) (*GetRTTMRegionsResponseType, error) {
	response := new(GetRTTMRegionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateRTTMRegion(request *UpdateRTTMRegionRequestType) (*UpdateRTTMRegionResponseType, error) {
	response := new(UpdateRTTMRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneRTTMRegion(request *DeleteOneRTTMRegionRequestType) (*DeleteOneRTTMRegionResponseType, error) {
	response := new(DeleteOneRTTMRegionResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateRTTMRegionPoolEntry(request *CreateRTTMRegionPoolEntryRequestType) (*CreateRTTMRegionPoolEntryResponseType, error) {
	response := new(CreateRTTMRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateRTTMRegionPoolEntry(request *UpdateRTTMRegionPoolEntryRequestType) (*UpdateRTTMRegionPoolEntryResponseType, error) {
	response := new(UpdateRTTMRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneRTTMRegionPoolEntry(request *GetOneRTTMRegionPoolEntryRequestType) (*GetOneRTTMRegionPoolEntryResponseType, error) {
	response := new(GetOneRTTMRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetRTTMRegionPoolEntries(request *GetRTTMRegionPoolEntriesRequestType) (*GetRTTMRegionPoolEntriesResponseType, error) {
	response := new(GetRTTMRegionPoolEntriesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneRTTMRegionPoolEntry(request *DeleteOneRTTMRegionPoolEntryRequestType) (*DeleteOneRTTMRegionPoolEntryResponseType, error) {
	response := new(DeleteOneRTTMRegionPoolEntryResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateHTTPRedirect(request *CreateHTTPRedirectRequestType) (*CreateHTTPRedirectResponseType, error) {
	response := new(CreateHTTPRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneHTTPRedirect(request *GetOneHTTPRedirectRequestType) (*GetOneHTTPRedirectResponseType, error) {
	response := new(GetOneHTTPRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetHTTPRedirects(request *GetHTTPRedirectsRequestType) (*GetHTTPRedirectsResponseType, error) {
	response := new(GetHTTPRedirectsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateHTTPRedirect(request *UpdateHTTPRedirectRequestType) (*UpdateHTTPRedirectResponseType, error) {
	response := new(UpdateHTTPRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneHTTPRedirect(request *DeleteOneHTTPRedirectRequestType) (*DeleteOneHTTPRedirectResponseType, error) {
	response := new(DeleteOneHTTPRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateAdvRedirectRule(request *CreateAdvRedirectRuleRequestType) (*CreateAdvRedirectRuleResponseType, error) {
	response := new(CreateAdvRedirectRuleResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateAdvRedirectRule(request *UpdateAdvRedirectRuleRequestType) (*UpdateAdvRedirectRuleResponseType, error) {
	response := new(UpdateAdvRedirectRuleResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneAdvRedirectRule(request *GetOneAdvRedirectRuleRequestType) (*GetOneAdvRedirectRuleResponseType, error) {
	response := new(GetOneAdvRedirectRuleResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetAdvRedirectRules(request *GetAdvRedirectRulesRequestType) (*GetAdvRedirectRulesResponseType, error) {
	response := new(GetAdvRedirectRulesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneAdvRedirectRule(request *DeleteOneAdvRedirectRuleRequestType) (*DeleteOneAdvRedirectRuleResponseType, error) {
	response := new(DeleteOneAdvRedirectRuleResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateAdvRedirect(request *CreateAdvRedirectRequestType) (*CreateAdvRedirectResponseType, error) {
	response := new(CreateAdvRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneAdvRedirect(request *GetOneAdvRedirectRequestType) (*GetOneAdvRedirectResponseType, error) {
	response := new(GetOneAdvRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetAdvRedirects(request *GetAdvRedirectsRequestType) (*GetAdvRedirectsResponseType, error) {
	response := new(GetAdvRedirectsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateAdvRedirect(request *UpdateAdvRedirectRequestType) (*UpdateAdvRedirectResponseType, error) {
	response := new(UpdateAdvRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneAdvRedirect(request *DeleteOneAdvRedirectRequestType) (*DeleteOneAdvRedirectResponseType, error) {
	response := new(DeleteOneAdvRedirectResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateCDNManager(request *CreateCDNManagerRequestType) (*CreateCDNManagerResponseType, error) {
	response := new(CreateCDNManagerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneCDNManager(request *GetOneCDNManagerRequestType) (*GetOneCDNManagerResponseType, error) {
	response := new(GetOneCDNManagerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetCDNManagers(request *GetCDNManagersRequestType) (*GetCDNManagersResponseType, error) {
	response := new(GetCDNManagersResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateCDNManager(request *UpdateCDNManagerRequestType) (*UpdateCDNManagerResponseType, error) {
	response := new(UpdateCDNManagerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneCDNManager(request *DeleteOneCDNManagerRequestType) (*DeleteOneCDNManagerResponseType, error) {
	response := new(DeleteOneCDNManagerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) FailoverCDNManager(request *FailoverCDNManagerRequestType) (*FailoverCDNManagerResponseType, error) {
	response := new(FailoverCDNManagerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RecoverCDNManager(request *RecoverCDNManagerRequestType) (*RecoverCDNManagerResponseType, error) {
	response := new(RecoverCDNManagerResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetNodeList(request *GetNodeListRequestType) (*GetNodeListResponseType, error) {
	response := new(GetNodeListResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) PublishZone(request *PublishZoneRequestType) (*PublishZoneResponseType, error) {
	response := new(PublishZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) PruneZone(request *PruneZoneRequestType) (*PruneZoneResponseType, error) {
	response := new(PruneZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) FreezeZone(request *FreezeZoneRequestType) (*FreezeZoneResponseType, error) {
	response := new(FreezeZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ThawZone(request *ThawZoneRequestType) (*ThawZoneResponseType, error) {
	response := new(ThawZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) RestoreZone(request *RestoreZoneRequestType) (*RestoreZoneResponseType, error) {
	response := new(RestoreZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) BlockZone(request *BlockZoneRequestType) (*BlockZoneResponseType, error) {
	response := new(BlockZoneResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteZoneChangeset(request *DeleteZoneChangesetRequestType) (*DeleteZoneChangesetResponseType, error) {
	response := new(DeleteZoneChangesetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetZoneChangeset(request *GetZoneChangesetRequestType) (*GetZoneChangesetResponseType, error) {
	response := new(GetZoneChangesetResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetZoneNotes(request *GetZoneNotesRequestType) (*GetZoneNotesResponseType, error) {
	response := new(GetZoneNotesResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UploadZoneFile(request *UploadZoneFileRequestType) (*UploadZoneFileResponseType, error) {
	response := new(UploadZoneFileResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) TransferZoneIn(request *TransferZoneInRequestType) (*TransferZoneInResponseType, error) {
	response := new(TransferZoneInResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetTransferStatus(request *GetTransferStatusRequestType) (*GetTransferStatusResponseType, error) {
	response := new(GetTransferStatusResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetZoneConfigOptions(request *GetZoneConfigOptionsRequestType) (*GetZoneConfigOptionsResponseType, error) {
	response := new(GetZoneConfigOptionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SetZoneConfigOptions(request *SetZoneConfigOptionsRequestType) (*SetZoneConfigOptionsResponseType, error) {
	response := new(SetZoneConfigOptionsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateIPTrack(request *CreateIPTrackRequestType) (*CreateIPTrackResponseType, error) {
	response := new(CreateIPTrackResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneIPTrack(request *GetOneIPTrackRequestType) (*GetOneIPTrackResponseType, error) {
	response := new(GetOneIPTrackResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetIPTracks(request *GetIPTracksRequestType) (*GetIPTracksResponseType, error) {
	response := new(GetIPTracksResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateIPTrack(request *UpdateIPTrackRequestType) (*UpdateIPTrackResponseType, error) {
	response := new(UpdateIPTrackResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneIPTrack(request *DeleteOneIPTrackRequestType) (*DeleteOneIPTrackResponseType, error) {
	response := new(DeleteOneIPTrackResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateIPTrack(request *ActivateIPTrackRequestType) (*ActivateIPTrackResponseType, error) {
	response := new(ActivateIPTrackResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateIPTrack(request *DeactivateIPTrackRequestType) (*DeactivateIPTrackResponseType, error) {
	response := new(DeactivateIPTrackResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateDNSSEC(request *CreateDNSSECRequestType) (*CreateDNSSECResponseType, error) {
	response := new(CreateDNSSECResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneDNSSEC(request *GetOneDNSSECRequestType) (*GetOneDNSSECResponseType, error) {
	response := new(GetOneDNSSECResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDNSSECs(request *GetDNSSECsRequestType) (*GetDNSSECsResponseType, error) {
	response := new(GetDNSSECsResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateDNSSEC(request *UpdateDNSSECRequestType) (*UpdateDNSSECResponseType, error) {
	response := new(UpdateDNSSECResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneDNSSEC(request *DeleteOneDNSSECRequestType) (*DeleteOneDNSSECResponseType, error) {
	response := new(DeleteOneDNSSECResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) ActivateDNSSEC(request *ActivateDNSSECRequestType) (*ActivateDNSSECResponseType, error) {
	response := new(ActivateDNSSECResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeactivateDNSSEC(request *DeactivateDNSSECRequestType) (*DeactivateDNSSECResponseType, error) {
	response := new(DeactivateDNSSECResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetDNSSECTimeline(request *GetDNSSECTimelineRequestType) (*GetDNSSECTimelineResponseType, error) {
	response := new(GetDNSSECTimelineResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetTasks(request *GetTasksRequestType) (*GetTasksResponseType, error) {
	response := new(GetTasksResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneTask(request *GetOneTaskRequestType) (*GetOneTaskResponseType, error) {
	response := new(GetOneTaskResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CancelTask(request *CancelTaskRequestType) (*CancelTaskResponseType, error) {
	response := new(CancelTaskResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) CreateExtNameserver(request *CreateExtNameserverRequestType) (*CreateExtNameserverResponseType, error) {
	response := new(CreateExtNameserverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetOneExtNameserver(request *GetOneExtNameserverRequestType) (*GetOneExtNameserverResponseType, error) {
	response := new(GetOneExtNameserverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetExtNameservers(request *GetExtNameserversRequestType) (*GetExtNameserversResponseType, error) {
	response := new(GetExtNameserversResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) UpdateExtNameserver(request *UpdateExtNameserverRequestType) (*UpdateExtNameserverResponseType, error) {
	response := new(UpdateExtNameserverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) DeleteOneExtNameserver(request *DeleteOneExtNameserverRequestType) (*DeleteOneExtNameserverResponseType, error) {
	response := new(DeleteOneExtNameserverResponseType)
	err := service.client.Call("https://api2.dynect.net/SOAP/", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
