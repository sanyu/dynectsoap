package dynectsoap

type Dynect interface {

	// Error can be either of the following types:
	//
	//   - fault

	GetJob(request *GetJobRequestType) (*GetJobResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* starts a DynectAPI session */
	SessionLogin(request *SessionLoginRequestType) (*SessionLoginResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* ends a DynectAPI session and invalidates the token */
	SessionLogout(request *SessionLogoutRequestType) (*SessionLogoutResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* checks where session and token are still valid */
	SessionIsAlive(request *SessionIsAliveRequestType) (*SessionIsAliveResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* No operation, prevents sessions from timing out */
	SessionKeepAlive(request *SessionKeepAliveRequestType) (*SessionKeepAliveResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Support only; adds permissions from a given customer */
	ScopeIn(request *ScopeInRequestType) (*ScopeInResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Support only; changes permissions to those of some particular user */
	ScopeAs(request *ScopeAsRequestType) (*ScopeAsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Undoes any ScopeIn or ScopeAs, returning to usual permissions */
	Unscope(request *UnscopeRequestType) (*UnscopeResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Retrieves Queries Per Second statistics in CSV format */
	GetQueryStats(request *GetQueryStatsRequestType) (*GetQueryStatsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateGeo(request *CreateGeoRequestType) (*CreateGeoResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateGeo(request *UpdateGeoRequestType) (*UpdateGeoResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetGeos(request *GetGeosRequestType) (*GetGeosResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneGeo(request *GetOneGeoRequestType) (*GetOneGeoResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneGeo(request *DeleteOneGeoRequestType) (*DeleteOneGeoResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateGeo(request *ActivateGeoRequestType) (*ActivateGeoResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateGeo(request *DeactivateGeoRequestType) (*DeactivateGeoResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateGeoRegionGroup(request *CreateGeoRegionGroupRequestType) (*CreateGeoRegionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateGeoRegionGroup(request *UpdateGeoRegionGroupRequestType) (*UpdateGeoRegionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneGeoRegionGroup(request *DeleteOneGeoRegionGroupRequestType) (*DeleteOneGeoRegionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetGeoRegionGroups(request *GetGeoRegionGroupsRequestType) (*GetGeoRegionGroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneGeoRegionGroup(request *GetOneGeoRegionGroupRequestType) (*GetOneGeoRegionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateGeoNode(request *CreateGeoNodeRequestType) (*CreateGeoNodeResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneGeoNode(request *DeleteOneGeoNodeRequestType) (*DeleteOneGeoNodeResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetGeoNodes(request *GetGeoNodesRequestType) (*GetGeoNodesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSF(request *CreateDSFRequestType) (*CreateDSFResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSF(request *UpdateDSFRequestType) (*UpdateDSFResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFs(request *GetDSFsRequestType) (*GetDSFsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFNotifiers(request *GetDSFNotifiersRequestType) (*GetDSFNotifiersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSF(request *DeleteOneDSFRequestType) (*DeleteOneDSFResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSF(request *GetOneDSFRequestType) (*GetOneDSFResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RevertDSF(request *RevertDSFRequestType) (*RevertDSFResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	PublishDSF(request *PublishDSFRequestType) (*PublishDSFResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddDSFNotifier(request *AddDSFNotifierRequestType) (*AddDSFNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveDSFNotifier(request *RemoveDSFNotifierRequestType) (*RemoveDSFNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSFRuleset(request *CreateDSFRulesetRequestType) (*CreateDSFRulesetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFRuleset(request *UpdateDSFRulesetRequestType) (*UpdateDSFRulesetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFRulesets(request *GetDSFRulesetsRequestType) (*GetDSFRulesetsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSFRuleset(request *GetOneDSFRulesetRequestType) (*GetOneDSFRulesetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFRuleset(request *DeleteOneDSFRulesetRequestType) (*DeleteOneDSFRulesetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSFResponsePool(request *CreateDSFResponsePoolRequestType) (*CreateDSFResponsePoolResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFResponsePool(request *UpdateDSFResponsePoolRequestType) (*UpdateDSFResponsePoolResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFResponsePools(request *GetDSFResponsePoolsRequestType) (*GetDSFResponsePoolsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSFResponsePool(request *GetOneDSFResponsePoolRequestType) (*GetOneDSFResponsePoolResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFResponsePool(request *DeleteOneDSFResponsePoolRequestType) (*DeleteOneDSFResponsePoolResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSFRecordSetFailoverChain(request *CreateDSFRecordSetFailoverChainRequestType) (*CreateDSFRecordSetFailoverChainResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFRecordSetFailoverChain(request *UpdateDSFRecordSetFailoverChainRequestType) (*UpdateDSFRecordSetFailoverChainResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFRecordSetFailoverChains(request *GetDSFRecordSetFailoverChainsRequestType) (*GetDSFRecordSetFailoverChainsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSFRecordSetFailoverChain(request *GetOneDSFRecordSetFailoverChainRequestType) (*GetOneDSFRecordSetFailoverChainResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFRecordSetFailoverChain(request *DeleteOneDSFRecordSetFailoverChainRequestType) (*DeleteOneDSFRecordSetFailoverChainResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSFRecordSet(request *CreateDSFRecordSetRequestType) (*CreateDSFRecordSetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFRecordSet(request *UpdateDSFRecordSetRequestType) (*UpdateDSFRecordSetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSFRecordSet(request *GetOneDSFRecordSetRequestType) (*GetOneDSFRecordSetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFRecordSets(request *GetDSFRecordSetsRequestType) (*GetDSFRecordSetsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFRecordSet(request *DeleteOneDSFRecordSetRequestType) (*DeleteOneDSFRecordSetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSFRecord(request *CreateDSFRecordRequestType) (*CreateDSFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFRecord(request *UpdateDSFRecordRequestType) (*UpdateDSFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSFRecord(request *GetOneDSFRecordRequestType) (*GetOneDSFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFRecords(request *GetDSFRecordsRequestType) (*GetDSFRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFRecord(request *DeleteOneDSFRecordRequestType) (*DeleteOneDSFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddDSFNode(request *AddDSFNodeRequestType) (*AddDSFNodeResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFNodes(request *UpdateDSFNodesRequestType) (*UpdateDSFNodesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFNodes(request *GetDSFNodesRequestType) (*GetDSFNodesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFNode(request *DeleteOneDSFNodeRequestType) (*DeleteOneDSFNodeResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDSFMonitor(request *CreateDSFMonitorRequestType) (*CreateDSFMonitorResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateDSFMonitor(request *UpdateDSFMonitorRequestType) (*UpdateDSFMonitorResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneDSFMonitor(request *GetOneDSFMonitorRequestType) (*GetOneDSFMonitorResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFMonitors(request *GetDSFMonitorsRequestType) (*GetDSFMonitorsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneDSFMonitor(request *DeleteOneDSFMonitorRequestType) (*DeleteOneDSFMonitorResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddDSFMonitorNotifier(request *AddDSFMonitorNotifierRequestType) (*AddDSFMonitorNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDSFMonitorSites(request *GetDSFMonitorSitesRequestType) (*GetDSFMonitorSitesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateNotifier(request *CreateNotifierRequestType) (*CreateNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateNotifier(request *UpdateNotifierRequestType) (*UpdateNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneNotifier(request *GetOneNotifierRequestType) (*GetOneNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetNotifiers(request *GetNotifiersRequestType) (*GetNotifiersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneNotifier(request *DeleteOneNotifierRequestType) (*DeleteOneNotifierResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateConfigLimit(request *CreateConfigLimitRequestType) (*CreateConfigLimitResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneConfigLimit(request *GetOneConfigLimitRequestType) (*GetOneConfigLimitResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetConfigLimits(request *GetConfigLimitsRequestType) (*GetConfigLimitsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateConfigLimit(request *UpdateConfigLimitRequestType) (*UpdateConfigLimitResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneConfigLimit(request *DeleteOneConfigLimitRequestType) (*DeleteOneConfigLimitResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new PermissionGroup */
	CreatePermissionGroup(request *CreatePermissionGroupRequestType) (*CreatePermissionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single PermissionGroup */
	GetOnePermissionGroup(request *GetOnePermissionGroupRequestType) (*GetOnePermissionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every PermissionGroup */
	GetPermissionGroups(request *GetPermissionGroupsRequestType) (*GetPermissionGroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single PermissionGroup */
	DeleteOnePermissionGroup(request *DeleteOnePermissionGroupRequestType) (*DeleteOnePermissionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdatePermissionGroup(request *UpdatePermissionGroupRequestType) (*UpdatePermissionGroupResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetCustomerPermissions(request *GetCustomerPermissionsRequestType) (*GetCustomerPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetUserPermissions(request *GetUserPermissionsRequestType) (*GetUserPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CheckPermissions(request *CheckPermissionsRequestType) (*CheckPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddPermissionGroupUsers(request *AddPermissionGroupUsersRequestType) (*AddPermissionGroupUsersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetPermissionGroupUsers(request *SetPermissionGroupUsersRequestType) (*SetPermissionGroupUsersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemovePermissionGroupUsers(request *RemovePermissionGroupUsersRequestType) (*RemovePermissionGroupUsersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddPermissionGroupSubgroups(request *AddPermissionGroupSubgroupsRequestType) (*AddPermissionGroupSubgroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetPermissionGroupSubgroups(request *SetPermissionGroupSubgroupsRequestType) (*SetPermissionGroupSubgroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemovePermissionGroupSubgroups(request *RemovePermissionGroupSubgroupsRequestType) (*RemovePermissionGroupSubgroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddPermissionGroupPermissions(request *AddPermissionGroupPermissionsRequestType) (*AddPermissionGroupPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetPermissionGroupPermissions(request *SetPermissionGroupPermissionsRequestType) (*SetPermissionGroupPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemovePermissionGroupPermissions(request *RemovePermissionGroupPermissionsRequestType) (*RemovePermissionGroupPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddPermissionGroupZones(request *AddPermissionGroupZonesRequestType) (*AddPermissionGroupZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetPermissionGroupZones(request *SetPermissionGroupZonesRequestType) (*SetPermissionGroupZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemovePermissionGroupZones(request *RemovePermissionGroupZonesRequestType) (*RemovePermissionGroupZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddUserGroups(request *AddUserGroupsRequestType) (*AddUserGroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetUserGroups(request *SetUserGroupsRequestType) (*SetUserGroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveUserGroups(request *RemoveUserGroupsRequestType) (*RemoveUserGroupsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddUserZones(request *AddUserZonesRequestType) (*AddUserZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetUserZones(request *SetUserZonesRequestType) (*SetUserZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveUserZones(request *RemoveUserZonesRequestType) (*RemoveUserZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddUserPermissions(request *AddUserPermissionsRequestType) (*AddUserPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetUserPermissions(request *SetUserPermissionsRequestType) (*SetUserPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveUserPermissions(request *RemoveUserPermissionsRequestType) (*RemoveUserPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddUserForbids(request *AddUserForbidsRequestType) (*AddUserForbidsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetUserForbids(request *SetUserForbidsRequestType) (*SetUserForbidsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveUserForbids(request *RemoveUserForbidsRequestType) (*RemoveUserForbidsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddCustomerPermissions(request *AddCustomerPermissionsRequestType) (*AddCustomerPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetCustomerPermissions(request *SetCustomerPermissionsRequestType) (*SetCustomerPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveCustomerPermissions(request *RemoveCustomerPermissionsRequestType) (*RemoveCustomerPermissionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddCustomerForbids(request *AddCustomerForbidsRequestType) (*AddCustomerForbidsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetCustomerForbids(request *SetCustomerForbidsRequestType) (*SetCustomerForbidsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RemoveCustomerForbids(request *RemoveCustomerForbidsRequestType) (*RemoveCustomerForbidsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetHostStatsFlags(request *GetHostStatsFlagsRequestType) (*GetHostStatsFlagsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetHostStatsFlags(request *SetHostStatsFlagsRequestType) (*SetHostStatsFlagsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateTSIGKey(request *CreateTSIGKeyRequestType) (*CreateTSIGKeyResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneTSIGKey(request *GetOneTSIGKeyRequestType) (*GetOneTSIGKeyResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetTSIGKeys(request *GetTSIGKeysRequestType) (*GetTSIGKeysResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateTSIGKey(request *UpdateTSIGKeyRequestType) (*UpdateTSIGKeyResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneTSIGKey(request *DeleteOneTSIGKeyRequestType) (*DeleteOneTSIGKeyResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new Zone */
	CreateZone(request *CreateZoneRequestType) (*CreateZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single Zone */
	GetOneZone(request *GetOneZoneRequestType) (*GetOneZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every Zone */
	GetZones(request *GetZonesRequestType) (*GetZonesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single Zone */
	DeleteOneZone(request *DeleteOneZoneRequestType) (*DeleteOneZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateSecondaryZone(request *CreateSecondaryZoneRequestType) (*CreateSecondaryZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateSecondary(request *UpdateSecondaryRequestType) (*UpdateSecondaryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateSecondary(request *ActivateSecondaryRequestType) (*ActivateSecondaryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateSecondary(request *DeactivateSecondaryRequestType) (*DeactivateSecondaryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RetransferSecondary(request *RetransferSecondaryRequestType) (*RetransferSecondaryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneSecondary(request *GetOneSecondaryRequestType) (*GetOneSecondaryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetSecondaries(request *GetSecondariesRequestType) (*GetSecondariesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetZoneApex(request *GetZoneApexRequestType) (*GetZoneApexResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new A record */
	CreateARecord(request *CreateARecordRequestType) (*CreateARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single A record */
	GetOneARecord(request *GetOneARecordRequestType) (*GetOneARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every A record */
	GetARecords(request *GetARecordsRequestType) (*GetARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single A record */
	UpdateARecord(request *UpdateARecordRequestType) (*UpdateARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every A record */
	DeleteARecords(request *DeleteARecordsRequestType) (*DeleteARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single A record */
	DeleteOneARecord(request *DeleteOneARecordRequestType) (*DeleteOneARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new AAAA record */
	CreateAAAARecord(request *CreateAAAARecordRequestType) (*CreateAAAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single AAAA record */
	GetOneAAAARecord(request *GetOneAAAARecordRequestType) (*GetOneAAAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every AAAA record */
	GetAAAARecords(request *GetAAAARecordsRequestType) (*GetAAAARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single AAAA record */
	UpdateAAAARecord(request *UpdateAAAARecordRequestType) (*UpdateAAAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every AAAA record */
	DeleteAAAARecords(request *DeleteAAAARecordsRequestType) (*DeleteAAAARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single AAAA record */
	DeleteOneAAAARecord(request *DeleteOneAAAARecordRequestType) (*DeleteOneAAAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new ALIAS record */
	CreateALIASRecord(request *CreateALIASRecordRequestType) (*CreateALIASRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single ALIAS record */
	GetOneALIASRecord(request *GetOneALIASRecordRequestType) (*GetOneALIASRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every ALIAS record */
	GetALIASRecords(request *GetALIASRecordsRequestType) (*GetALIASRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single ALIAS record */
	UpdateALIASRecord(request *UpdateALIASRecordRequestType) (*UpdateALIASRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every ALIAS record */
	DeleteALIASRecords(request *DeleteALIASRecordsRequestType) (*DeleteALIASRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single ALIAS record */
	DeleteOneALIASRecord(request *DeleteOneALIASRecordRequestType) (*DeleteOneALIASRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CAA record */
	CreateCAARecord(request *CreateCAARecordRequestType) (*CreateCAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CAA record */
	GetOneCAARecord(request *GetOneCAARecordRequestType) (*GetOneCAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CAA record */
	GetCAARecords(request *GetCAARecordsRequestType) (*GetCAARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CAA record */
	UpdateCAARecord(request *UpdateCAARecordRequestType) (*UpdateCAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every CAA record */
	DeleteCAARecords(request *DeleteCAARecordsRequestType) (*DeleteCAARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CAA record */
	DeleteOneCAARecord(request *DeleteOneCAARecordRequestType) (*DeleteOneCAARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CDNSKEY record */
	CreateCDNSKEYRecord(request *CreateCDNSKEYRecordRequestType) (*CreateCDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CDNSKEY record */
	GetOneCDNSKEYRecord(request *GetOneCDNSKEYRecordRequestType) (*GetOneCDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CDNSKEY record */
	GetCDNSKEYRecords(request *GetCDNSKEYRecordsRequestType) (*GetCDNSKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CDNSKEY record */
	UpdateCDNSKEYRecord(request *UpdateCDNSKEYRecordRequestType) (*UpdateCDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every CDNSKEY record */
	DeleteCDNSKEYRecords(request *DeleteCDNSKEYRecordsRequestType) (*DeleteCDNSKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CDNSKEY record */
	DeleteOneCDNSKEYRecord(request *DeleteOneCDNSKEYRecordRequestType) (*DeleteOneCDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CDS record */
	CreateCDSRecord(request *CreateCDSRecordRequestType) (*CreateCDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CDS record */
	GetOneCDSRecord(request *GetOneCDSRecordRequestType) (*GetOneCDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CDS record */
	GetCDSRecords(request *GetCDSRecordsRequestType) (*GetCDSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CDS record */
	UpdateCDSRecord(request *UpdateCDSRecordRequestType) (*UpdateCDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every CDS record */
	DeleteCDSRecords(request *DeleteCDSRecordsRequestType) (*DeleteCDSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CDS record */
	DeleteOneCDSRecord(request *DeleteOneCDSRecordRequestType) (*DeleteOneCDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CERT record */
	CreateCERTRecord(request *CreateCERTRecordRequestType) (*CreateCERTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CERT record */
	GetOneCERTRecord(request *GetOneCERTRecordRequestType) (*GetOneCERTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CERT record */
	GetCERTRecords(request *GetCERTRecordsRequestType) (*GetCERTRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CERT record */
	UpdateCERTRecord(request *UpdateCERTRecordRequestType) (*UpdateCERTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every CERT record */
	DeleteCERTRecords(request *DeleteCERTRecordsRequestType) (*DeleteCERTRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CERT record */
	DeleteOneCERTRecord(request *DeleteOneCERTRecordRequestType) (*DeleteOneCERTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CNAME record */
	CreateCNAMERecord(request *CreateCNAMERecordRequestType) (*CreateCNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CNAME record */
	GetOneCNAMERecord(request *GetOneCNAMERecordRequestType) (*GetOneCNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CNAME record */
	GetCNAMERecords(request *GetCNAMERecordsRequestType) (*GetCNAMERecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CNAME record */
	UpdateCNAMERecord(request *UpdateCNAMERecordRequestType) (*UpdateCNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every CNAME record */
	DeleteCNAMERecords(request *DeleteCNAMERecordsRequestType) (*DeleteCNAMERecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CNAME record */
	DeleteOneCNAMERecord(request *DeleteOneCNAMERecordRequestType) (*DeleteOneCNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CSYNC record */
	CreateCSYNCRecord(request *CreateCSYNCRecordRequestType) (*CreateCSYNCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CSYNC record */
	GetOneCSYNCRecord(request *GetOneCSYNCRecordRequestType) (*GetOneCSYNCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CSYNC record */
	GetCSYNCRecords(request *GetCSYNCRecordsRequestType) (*GetCSYNCRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CSYNC record */
	UpdateCSYNCRecord(request *UpdateCSYNCRecordRequestType) (*UpdateCSYNCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every CSYNC record */
	DeleteCSYNCRecords(request *DeleteCSYNCRecordsRequestType) (*DeleteCSYNCRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CSYNC record */
	DeleteOneCSYNCRecord(request *DeleteOneCSYNCRecordRequestType) (*DeleteOneCSYNCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new DHCID record */
	CreateDHCIDRecord(request *CreateDHCIDRecordRequestType) (*CreateDHCIDRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single DHCID record */
	GetOneDHCIDRecord(request *GetOneDHCIDRecordRequestType) (*GetOneDHCIDRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every DHCID record */
	GetDHCIDRecords(request *GetDHCIDRecordsRequestType) (*GetDHCIDRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single DHCID record */
	UpdateDHCIDRecord(request *UpdateDHCIDRecordRequestType) (*UpdateDHCIDRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every DHCID record */
	DeleteDHCIDRecords(request *DeleteDHCIDRecordsRequestType) (*DeleteDHCIDRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single DHCID record */
	DeleteOneDHCIDRecord(request *DeleteOneDHCIDRecordRequestType) (*DeleteOneDHCIDRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new DNAME record */
	CreateDNAMERecord(request *CreateDNAMERecordRequestType) (*CreateDNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single DNAME record */
	GetOneDNAMERecord(request *GetOneDNAMERecordRequestType) (*GetOneDNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every DNAME record */
	GetDNAMERecords(request *GetDNAMERecordsRequestType) (*GetDNAMERecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single DNAME record */
	UpdateDNAMERecord(request *UpdateDNAMERecordRequestType) (*UpdateDNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every DNAME record */
	DeleteDNAMERecords(request *DeleteDNAMERecordsRequestType) (*DeleteDNAMERecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single DNAME record */
	DeleteOneDNAMERecord(request *DeleteOneDNAMERecordRequestType) (*DeleteOneDNAMERecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new DNSKEY record */
	CreateDNSKEYRecord(request *CreateDNSKEYRecordRequestType) (*CreateDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single DNSKEY record */
	GetOneDNSKEYRecord(request *GetOneDNSKEYRecordRequestType) (*GetOneDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every DNSKEY record */
	GetDNSKEYRecords(request *GetDNSKEYRecordsRequestType) (*GetDNSKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single DNSKEY record */
	UpdateDNSKEYRecord(request *UpdateDNSKEYRecordRequestType) (*UpdateDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every DNSKEY record */
	DeleteDNSKEYRecords(request *DeleteDNSKEYRecordsRequestType) (*DeleteDNSKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single DNSKEY record */
	DeleteOneDNSKEYRecord(request *DeleteOneDNSKEYRecordRequestType) (*DeleteOneDNSKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new DS record */
	CreateDSRecord(request *CreateDSRecordRequestType) (*CreateDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single DS record */
	GetOneDSRecord(request *GetOneDSRecordRequestType) (*GetOneDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every DS record */
	GetDSRecords(request *GetDSRecordsRequestType) (*GetDSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single DS record */
	UpdateDSRecord(request *UpdateDSRecordRequestType) (*UpdateDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every DS record */
	DeleteDSRecords(request *DeleteDSRecordsRequestType) (*DeleteDSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single DS record */
	DeleteOneDSRecord(request *DeleteOneDSRecordRequestType) (*DeleteOneDSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new IPSECKEY record */
	CreateIPSECKEYRecord(request *CreateIPSECKEYRecordRequestType) (*CreateIPSECKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single IPSECKEY record */
	GetOneIPSECKEYRecord(request *GetOneIPSECKEYRecordRequestType) (*GetOneIPSECKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every IPSECKEY record */
	GetIPSECKEYRecords(request *GetIPSECKEYRecordsRequestType) (*GetIPSECKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single IPSECKEY record */
	UpdateIPSECKEYRecord(request *UpdateIPSECKEYRecordRequestType) (*UpdateIPSECKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every IPSECKEY record */
	DeleteIPSECKEYRecords(request *DeleteIPSECKEYRecordsRequestType) (*DeleteIPSECKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single IPSECKEY record */
	DeleteOneIPSECKEYRecord(request *DeleteOneIPSECKEYRecordRequestType) (*DeleteOneIPSECKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new KEY record */
	CreateKEYRecord(request *CreateKEYRecordRequestType) (*CreateKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single KEY record */
	GetOneKEYRecord(request *GetOneKEYRecordRequestType) (*GetOneKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every KEY record */
	GetKEYRecords(request *GetKEYRecordsRequestType) (*GetKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single KEY record */
	UpdateKEYRecord(request *UpdateKEYRecordRequestType) (*UpdateKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every KEY record */
	DeleteKEYRecords(request *DeleteKEYRecordsRequestType) (*DeleteKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single KEY record */
	DeleteOneKEYRecord(request *DeleteOneKEYRecordRequestType) (*DeleteOneKEYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new KX record */
	CreateKXRecord(request *CreateKXRecordRequestType) (*CreateKXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single KX record */
	GetOneKXRecord(request *GetOneKXRecordRequestType) (*GetOneKXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every KX record */
	GetKXRecords(request *GetKXRecordsRequestType) (*GetKXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single KX record */
	UpdateKXRecord(request *UpdateKXRecordRequestType) (*UpdateKXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every KX record */
	DeleteKXRecords(request *DeleteKXRecordsRequestType) (*DeleteKXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single KX record */
	DeleteOneKXRecord(request *DeleteOneKXRecordRequestType) (*DeleteOneKXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new LOC record */
	CreateLOCRecord(request *CreateLOCRecordRequestType) (*CreateLOCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single LOC record */
	GetOneLOCRecord(request *GetOneLOCRecordRequestType) (*GetOneLOCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every LOC record */
	GetLOCRecords(request *GetLOCRecordsRequestType) (*GetLOCRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single LOC record */
	UpdateLOCRecord(request *UpdateLOCRecordRequestType) (*UpdateLOCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every LOC record */
	DeleteLOCRecords(request *DeleteLOCRecordsRequestType) (*DeleteLOCRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single LOC record */
	DeleteOneLOCRecord(request *DeleteOneLOCRecordRequestType) (*DeleteOneLOCRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new MX record */
	CreateMXRecord(request *CreateMXRecordRequestType) (*CreateMXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single MX record */
	GetOneMXRecord(request *GetOneMXRecordRequestType) (*GetOneMXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every MX record */
	GetMXRecords(request *GetMXRecordsRequestType) (*GetMXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single MX record */
	UpdateMXRecord(request *UpdateMXRecordRequestType) (*UpdateMXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every MX record */
	DeleteMXRecords(request *DeleteMXRecordsRequestType) (*DeleteMXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single MX record */
	DeleteOneMXRecord(request *DeleteOneMXRecordRequestType) (*DeleteOneMXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new NAPTR record */
	CreateNAPTRRecord(request *CreateNAPTRRecordRequestType) (*CreateNAPTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single NAPTR record */
	GetOneNAPTRRecord(request *GetOneNAPTRRecordRequestType) (*GetOneNAPTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every NAPTR record */
	GetNAPTRRecords(request *GetNAPTRRecordsRequestType) (*GetNAPTRRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single NAPTR record */
	UpdateNAPTRRecord(request *UpdateNAPTRRecordRequestType) (*UpdateNAPTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every NAPTR record */
	DeleteNAPTRRecords(request *DeleteNAPTRRecordsRequestType) (*DeleteNAPTRRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single NAPTR record */
	DeleteOneNAPTRRecord(request *DeleteOneNAPTRRecordRequestType) (*DeleteOneNAPTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new NSAP record */
	CreateNSAPRecord(request *CreateNSAPRecordRequestType) (*CreateNSAPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single NSAP record */
	GetOneNSAPRecord(request *GetOneNSAPRecordRequestType) (*GetOneNSAPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every NSAP record */
	GetNSAPRecords(request *GetNSAPRecordsRequestType) (*GetNSAPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single NSAP record */
	UpdateNSAPRecord(request *UpdateNSAPRecordRequestType) (*UpdateNSAPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every NSAP record */
	DeleteNSAPRecords(request *DeleteNSAPRecordsRequestType) (*DeleteNSAPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single NSAP record */
	DeleteOneNSAPRecord(request *DeleteOneNSAPRecordRequestType) (*DeleteOneNSAPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new POLICY record */
	CreatePOLICYRecord(request *CreatePOLICYRecordRequestType) (*CreatePOLICYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single POLICY record */
	GetOnePOLICYRecord(request *GetOnePOLICYRecordRequestType) (*GetOnePOLICYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every POLICY record */
	GetPOLICYRecords(request *GetPOLICYRecordsRequestType) (*GetPOLICYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single POLICY record */
	UpdatePOLICYRecord(request *UpdatePOLICYRecordRequestType) (*UpdatePOLICYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every POLICY record */
	DeletePOLICYRecords(request *DeletePOLICYRecordsRequestType) (*DeletePOLICYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single POLICY record */
	DeleteOnePOLICYRecord(request *DeleteOnePOLICYRecordRequestType) (*DeleteOnePOLICYRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new PTR record */
	CreatePTRRecord(request *CreatePTRRecordRequestType) (*CreatePTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single PTR record */
	GetOnePTRRecord(request *GetOnePTRRecordRequestType) (*GetOnePTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every PTR record */
	GetPTRRecords(request *GetPTRRecordsRequestType) (*GetPTRRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single PTR record */
	UpdatePTRRecord(request *UpdatePTRRecordRequestType) (*UpdatePTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every PTR record */
	DeletePTRRecords(request *DeletePTRRecordsRequestType) (*DeletePTRRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single PTR record */
	DeleteOnePTRRecord(request *DeleteOnePTRRecordRequestType) (*DeleteOnePTRRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new PX record */
	CreatePXRecord(request *CreatePXRecordRequestType) (*CreatePXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single PX record */
	GetOnePXRecord(request *GetOnePXRecordRequestType) (*GetOnePXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every PX record */
	GetPXRecords(request *GetPXRecordsRequestType) (*GetPXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single PX record */
	UpdatePXRecord(request *UpdatePXRecordRequestType) (*UpdatePXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every PX record */
	DeletePXRecords(request *DeletePXRecordsRequestType) (*DeletePXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single PX record */
	DeleteOnePXRecord(request *DeleteOnePXRecordRequestType) (*DeleteOnePXRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new RP record */
	CreateRPRecord(request *CreateRPRecordRequestType) (*CreateRPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single RP record */
	GetOneRPRecord(request *GetOneRPRecordRequestType) (*GetOneRPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every RP record */
	GetRPRecords(request *GetRPRecordsRequestType) (*GetRPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single RP record */
	UpdateRPRecord(request *UpdateRPRecordRequestType) (*UpdateRPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every RP record */
	DeleteRPRecords(request *DeleteRPRecordsRequestType) (*DeleteRPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single RP record */
	DeleteOneRPRecord(request *DeleteOneRPRecordRequestType) (*DeleteOneRPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new SPF record */
	CreateSPFRecord(request *CreateSPFRecordRequestType) (*CreateSPFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single SPF record */
	GetOneSPFRecord(request *GetOneSPFRecordRequestType) (*GetOneSPFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every SPF record */
	GetSPFRecords(request *GetSPFRecordsRequestType) (*GetSPFRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single SPF record */
	UpdateSPFRecord(request *UpdateSPFRecordRequestType) (*UpdateSPFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every SPF record */
	DeleteSPFRecords(request *DeleteSPFRecordsRequestType) (*DeleteSPFRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single SPF record */
	DeleteOneSPFRecord(request *DeleteOneSPFRecordRequestType) (*DeleteOneSPFRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new SRV record */
	CreateSRVRecord(request *CreateSRVRecordRequestType) (*CreateSRVRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single SRV record */
	GetOneSRVRecord(request *GetOneSRVRecordRequestType) (*GetOneSRVRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every SRV record */
	GetSRVRecords(request *GetSRVRecordsRequestType) (*GetSRVRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single SRV record */
	UpdateSRVRecord(request *UpdateSRVRecordRequestType) (*UpdateSRVRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every SRV record */
	DeleteSRVRecords(request *DeleteSRVRecordsRequestType) (*DeleteSRVRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single SRV record */
	DeleteOneSRVRecord(request *DeleteOneSRVRecordRequestType) (*DeleteOneSRVRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new SSHFP record */
	CreateSSHFPRecord(request *CreateSSHFPRecordRequestType) (*CreateSSHFPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single SSHFP record */
	GetOneSSHFPRecord(request *GetOneSSHFPRecordRequestType) (*GetOneSSHFPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every SSHFP record */
	GetSSHFPRecords(request *GetSSHFPRecordsRequestType) (*GetSSHFPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single SSHFP record */
	UpdateSSHFPRecord(request *UpdateSSHFPRecordRequestType) (*UpdateSSHFPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every SSHFP record */
	DeleteSSHFPRecords(request *DeleteSSHFPRecordsRequestType) (*DeleteSSHFPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single SSHFP record */
	DeleteOneSSHFPRecord(request *DeleteOneSSHFPRecordRequestType) (*DeleteOneSSHFPRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new TLSA record */
	CreateTLSARecord(request *CreateTLSARecordRequestType) (*CreateTLSARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single TLSA record */
	GetOneTLSARecord(request *GetOneTLSARecordRequestType) (*GetOneTLSARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every TLSA record */
	GetTLSARecords(request *GetTLSARecordsRequestType) (*GetTLSARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single TLSA record */
	UpdateTLSARecord(request *UpdateTLSARecordRequestType) (*UpdateTLSARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every TLSA record */
	DeleteTLSARecords(request *DeleteTLSARecordsRequestType) (*DeleteTLSARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single TLSA record */
	DeleteOneTLSARecord(request *DeleteOneTLSARecordRequestType) (*DeleteOneTLSARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new TXT record */
	CreateTXTRecord(request *CreateTXTRecordRequestType) (*CreateTXTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single TXT record */
	GetOneTXTRecord(request *GetOneTXTRecordRequestType) (*GetOneTXTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every TXT record */
	GetTXTRecords(request *GetTXTRecordsRequestType) (*GetTXTRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single TXT record */
	UpdateTXTRecord(request *UpdateTXTRecordRequestType) (*UpdateTXTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every TXT record */
	DeleteTXTRecords(request *DeleteTXTRecordsRequestType) (*DeleteTXTRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single TXT record */
	DeleteOneTXTRecord(request *DeleteOneTXTRecordRequestType) (*DeleteOneTXTRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single SOA record */
	GetOneSOARecord(request *GetOneSOARecordRequestType) (*GetOneSOARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every SOA record */
	GetSOARecords(request *GetSOARecordsRequestType) (*GetSOARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateSOARecord(request *UpdateSOARecordRequestType) (*UpdateSOARecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new NS record */
	CreateNSRecord(request *CreateNSRecordRequestType) (*CreateNSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single NS record */
	GetOneNSRecord(request *GetOneNSRecordRequestType) (*GetOneNSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every NS record */
	GetNSRecords(request *GetNSRecordsRequestType) (*GetNSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single NS record */
	UpdateNSRecord(request *UpdateNSRecordRequestType) (*UpdateNSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes every NS record */
	DeleteNSRecords(request *DeleteNSRecordsRequestType) (*DeleteNSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single NS record */
	DeleteOneNSRecord(request *DeleteOneNSRecordRequestType) (*DeleteOneNSRecordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceARecords(request *ReplaceARecordsRequestType) (*ReplaceARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceAAAARecords(request *ReplaceAAAARecordsRequestType) (*ReplaceAAAARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceALIASRecords(request *ReplaceALIASRecordsRequestType) (*ReplaceALIASRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceCAARecords(request *ReplaceCAARecordsRequestType) (*ReplaceCAARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceCDNSKEYRecords(request *ReplaceCDNSKEYRecordsRequestType) (*ReplaceCDNSKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceCDSRecords(request *ReplaceCDSRecordsRequestType) (*ReplaceCDSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceCERTRecords(request *ReplaceCERTRecordsRequestType) (*ReplaceCERTRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceCNAMERecords(request *ReplaceCNAMERecordsRequestType) (*ReplaceCNAMERecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceCSYNCRecords(request *ReplaceCSYNCRecordsRequestType) (*ReplaceCSYNCRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceDHCIDRecords(request *ReplaceDHCIDRecordsRequestType) (*ReplaceDHCIDRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceDNAMERecords(request *ReplaceDNAMERecordsRequestType) (*ReplaceDNAMERecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceDNSKEYRecords(request *ReplaceDNSKEYRecordsRequestType) (*ReplaceDNSKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceDSRecords(request *ReplaceDSRecordsRequestType) (*ReplaceDSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceIPSECKEYRecords(request *ReplaceIPSECKEYRecordsRequestType) (*ReplaceIPSECKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceKEYRecords(request *ReplaceKEYRecordsRequestType) (*ReplaceKEYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceKXRecords(request *ReplaceKXRecordsRequestType) (*ReplaceKXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceLOCRecords(request *ReplaceLOCRecordsRequestType) (*ReplaceLOCRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceMXRecords(request *ReplaceMXRecordsRequestType) (*ReplaceMXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceNAPTRRecords(request *ReplaceNAPTRRecordsRequestType) (*ReplaceNAPTRRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceNSAPRecords(request *ReplaceNSAPRecordsRequestType) (*ReplaceNSAPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplacePOLICYRecords(request *ReplacePOLICYRecordsRequestType) (*ReplacePOLICYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplacePTRRecords(request *ReplacePTRRecordsRequestType) (*ReplacePTRRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplacePXRecords(request *ReplacePXRecordsRequestType) (*ReplacePXRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceRPRecords(request *ReplaceRPRecordsRequestType) (*ReplaceRPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceSPFRecords(request *ReplaceSPFRecordsRequestType) (*ReplaceSPFRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceSRVRecords(request *ReplaceSRVRecordsRequestType) (*ReplaceSRVRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceSSHFPRecords(request *ReplaceSSHFPRecordsRequestType) (*ReplaceSSHFPRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceTLSARecords(request *ReplaceTLSARecordsRequestType) (*ReplaceTLSARecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceTXTRecords(request *ReplaceTXTRecordsRequestType) (*ReplaceTXTRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ReplaceNSRecords(request *ReplaceNSRecordsRequestType) (*ReplaceNSRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetANYRecords(request *GetANYRecordsRequestType) (*GetANYRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetAllRecords(request *GetAllRecordsRequestType) (*GetAllRecordsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetAllAliasQNames(request *GetAllAliasQNamesRequestType) (*GetAllAliasQNamesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single User */
	GetOneUser(request *GetOneUserRequestType) (*GetOneUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single User */
	DeleteOneUser(request *DeleteOneUserRequestType) (*DeleteOneUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateUser(request *CreateUserRequestType) (*CreateUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateUser(request *UpdateUserRequestType) (*UpdateUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetUsers(request *GetUsersRequestType) (*GetUsersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetUpdateUsers(request *GetUpdateUsersRequestType) (*GetUpdateUsersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateUpdateUser(request *UpdateUpdateUserRequestType) (*UpdateUpdateUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneUpdateUser(request *DeleteOneUpdateUserRequestType) (*DeleteOneUpdateUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateUserPassword(request *UpdateUserPasswordRequestType) (*UpdateUserPasswordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	BlockUser(request *BlockUserRequestType) (*BlockUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UnblockUser(request *UnblockUserRequestType) (*UnblockUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new Contact */
	CreateContact(request *CreateContactRequestType) (*CreateContactResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single Contact */
	GetOneContact(request *GetOneContactRequestType) (*GetOneContactResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every Contact */
	GetContacts(request *GetContactsRequestType) (*GetContactsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single Contact */
	DeleteOneContact(request *DeleteOneContactRequestType) (*DeleteOneContactResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateContact(request *UpdateContactRequestType) (*UpdateContactResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateCustomer(request *CreateCustomerRequestType) (*CreateCustomerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateCustomer(request *UpdateCustomerRequestType) (*UpdateCustomerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneCustomer(request *GetOneCustomerRequestType) (*GetOneCustomerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetCustomers(request *GetCustomersRequestType) (*GetCustomersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneCustomer(request *DeleteOneCustomerRequestType) (*DeleteOneCustomerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetCustomerPrefs(request *GetCustomerPrefsRequestType) (*GetCustomerPrefsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetCustomerPrefs(request *SetCustomerPrefsRequestType) (*SetCustomerPrefsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetCustomerIPACL(request *GetCustomerIPACLRequestType) (*GetCustomerIPACLResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetCustomerIPACL(request *SetCustomerIPACLRequestType) (*SetCustomerIPACLResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateCustomerOracleMetadata(request *CreateCustomerOracleMetadataRequestType) (*CreateCustomerOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateCustomerOracleMetadata(request *UpdateCustomerOracleMetadataRequestType) (*UpdateCustomerOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetCustomerOracleMetadata(request *GetCustomerOracleMetadataRequestType) (*GetCustomerOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteCustomerOracleMetadata(request *DeleteCustomerOracleMetadataRequestType) (*DeleteCustomerOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateZoneOracleMetadata(request *CreateZoneOracleMetadataRequestType) (*CreateZoneOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateZoneOracleMetadata(request *UpdateZoneOracleMetadataRequestType) (*UpdateZoneOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetZoneOracleMetadata(request *GetZoneOracleMetadataRequestType) (*GetZoneOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteZoneOracleMetadata(request *DeleteZoneOracleMetadataRequestType) (*DeleteZoneOracleMetadataResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new Dynamic DNS service */
	CreateDDNS(request *CreateDDNSRequestType) (*CreateDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single Dynamic DNS service */
	GetOneDDNS(request *GetOneDDNSRequestType) (*GetOneDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every Dynamic DNS service */
	GetDDNSs(request *GetDDNSsRequestType) (*GetDDNSsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single Dynamic DNS service */
	UpdateDDNS(request *UpdateDDNSRequestType) (*UpdateDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single Dynamic DNS service */
	DeleteOneDDNS(request *DeleteOneDDNSRequestType) (*DeleteOneDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateDDNS(request *ActivateDDNSRequestType) (*ActivateDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateDDNS(request *DeactivateDDNSRequestType) (*DeactivateDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ResetDDNS(request *ResetDDNSRequestType) (*ResetDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetUpdateUserPassword(request *GetUpdateUserPasswordRequestType) (*GetUpdateUserPasswordResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateDDNSHost(request *CreateDDNSHostRequestType) (*CreateDDNSHostResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateUpdateUser(request *CreateUpdateUserRequestType) (*CreateUpdateUserResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	AddDDNS(request *AddDDNSRequestType) (*AddDDNSResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new Simple Failover service */
	CreateFailover(request *CreateFailoverRequestType) (*CreateFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single Simple Failover service */
	GetOneFailover(request *GetOneFailoverRequestType) (*GetOneFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every Simple Failover service */
	GetFailovers(request *GetFailoversRequestType) (*GetFailoversResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single Simple Failover service */
	UpdateFailover(request *UpdateFailoverRequestType) (*UpdateFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single Simple Failover service */
	DeleteOneFailover(request *DeleteOneFailoverRequestType) (*DeleteOneFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateFailover(request *ActivateFailoverRequestType) (*ActivateFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateFailover(request *DeactivateFailoverRequestType) (*DeactivateFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverFailover(request *RecoverFailoverRequestType) (*RecoverFailoverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new basic LoadBalance service */
	CreateLoadBalance(request *CreateLoadBalanceRequestType) (*CreateLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single basic LoadBalance service */
	GetOneLoadBalance(request *GetOneLoadBalanceRequestType) (*GetOneLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every basic LoadBalance service */
	GetLoadBalances(request *GetLoadBalancesRequestType) (*GetLoadBalancesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single basic LoadBalance service */
	UpdateLoadBalance(request *UpdateLoadBalanceRequestType) (*UpdateLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single basic LoadBalance service */
	DeleteOneLoadBalance(request *DeleteOneLoadBalanceRequestType) (*DeleteOneLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateLoadBalance(request *ActivateLoadBalanceRequestType) (*ActivateLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateLoadBalance(request *DeactivateLoadBalanceRequestType) (*DeactivateLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverLoadBalance(request *RecoverLoadBalanceRequestType) (*RecoverLoadBalanceResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverLoadBalanceIP(request *RecoverLoadBalanceIPRequestType) (*RecoverLoadBalanceIPResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateLoadBalancePoolEntry(request *CreateLoadBalancePoolEntryRequestType) (*CreateLoadBalancePoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateLoadBalancePoolEntry(request *UpdateLoadBalancePoolEntryRequestType) (*UpdateLoadBalancePoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneLoadBalancePoolEntry(request *GetOneLoadBalancePoolEntryRequestType) (*GetOneLoadBalancePoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetLoadBalancePoolEntries(request *GetLoadBalancePoolEntriesRequestType) (*GetLoadBalancePoolEntriesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneLoadBalancePoolEntry(request *DeleteOneLoadBalancePoolEntryRequestType) (*DeleteOneLoadBalancePoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new GSLB */
	CreateGSLB(request *CreateGSLBRequestType) (*CreateGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single GSLB */
	GetOneGSLB(request *GetOneGSLBRequestType) (*GetOneGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every GSLB */
	GetGSLBs(request *GetGSLBsRequestType) (*GetGSLBsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single GSLB */
	UpdateGSLB(request *UpdateGSLBRequestType) (*UpdateGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single GSLB */
	DeleteOneGSLB(request *DeleteOneGSLBRequestType) (*DeleteOneGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateGSLB(request *ActivateGSLBRequestType) (*ActivateGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateGSLB(request *DeactivateGSLBRequestType) (*DeactivateGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverGSLB(request *RecoverGSLBRequestType) (*RecoverGSLBResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverGSLBIP(request *RecoverGSLBIPRequestType) (*RecoverGSLBIPResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new GSLBRegion */
	CreateGSLBRegion(request *CreateGSLBRegionRequestType) (*CreateGSLBRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single GSLBRegion */
	GetOneGSLBRegion(request *GetOneGSLBRegionRequestType) (*GetOneGSLBRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every GSLBRegion */
	GetGSLBRegions(request *GetGSLBRegionsRequestType) (*GetGSLBRegionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single GSLBRegion */
	UpdateGSLBRegion(request *UpdateGSLBRegionRequestType) (*UpdateGSLBRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single GSLBRegion */
	DeleteOneGSLBRegion(request *DeleteOneGSLBRegionRequestType) (*DeleteOneGSLBRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateGSLBRegionPoolEntry(request *CreateGSLBRegionPoolEntryRequestType) (*CreateGSLBRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateGSLBRegionPoolEntry(request *UpdateGSLBRegionPoolEntryRequestType) (*UpdateGSLBRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneGSLBRegionPoolEntry(request *GetOneGSLBRegionPoolEntryRequestType) (*GetOneGSLBRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetGSLBRegionPoolEntries(request *GetGSLBRegionPoolEntriesRequestType) (*GetGSLBRegionPoolEntriesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneGSLBRegionPoolEntry(request *DeleteOneGSLBRegionPoolEntryRequestType) (*DeleteOneGSLBRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new RTTM */
	CreateRTTM(request *CreateRTTMRequestType) (*CreateRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single RTTM */
	GetOneRTTM(request *GetOneRTTMRequestType) (*GetOneRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every RTTM */
	GetRTTMs(request *GetRTTMsRequestType) (*GetRTTMsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single RTTM */
	UpdateRTTM(request *UpdateRTTMRequestType) (*UpdateRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single RTTM */
	DeleteOneRTTM(request *DeleteOneRTTMRequestType) (*DeleteOneRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateRTTM(request *ActivateRTTMRequestType) (*ActivateRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateRTTM(request *DeactivateRTTMRequestType) (*DeactivateRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverRTTM(request *RecoverRTTMRequestType) (*RecoverRTTMResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverRTTMIP(request *RecoverRTTMIPRequestType) (*RecoverRTTMIPResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetRTTMLogs(request *GetRTTMLogsRequestType) (*GetRTTMLogsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetRTTMRRSets(request *GetRTTMRRSetsRequestType) (*GetRTTMRRSetsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new RTTMRegion */
	CreateRTTMRegion(request *CreateRTTMRegionRequestType) (*CreateRTTMRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single RTTMRegion */
	GetOneRTTMRegion(request *GetOneRTTMRegionRequestType) (*GetOneRTTMRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every RTTMRegion */
	GetRTTMRegions(request *GetRTTMRegionsRequestType) (*GetRTTMRegionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single RTTMRegion */
	UpdateRTTMRegion(request *UpdateRTTMRegionRequestType) (*UpdateRTTMRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single RTTMRegion */
	DeleteOneRTTMRegion(request *DeleteOneRTTMRegionRequestType) (*DeleteOneRTTMRegionResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateRTTMRegionPoolEntry(request *CreateRTTMRegionPoolEntryRequestType) (*CreateRTTMRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateRTTMRegionPoolEntry(request *UpdateRTTMRegionPoolEntryRequestType) (*UpdateRTTMRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneRTTMRegionPoolEntry(request *GetOneRTTMRegionPoolEntryRequestType) (*GetOneRTTMRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetRTTMRegionPoolEntries(request *GetRTTMRegionPoolEntriesRequestType) (*GetRTTMRegionPoolEntriesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneRTTMRegionPoolEntry(request *DeleteOneRTTMRegionPoolEntryRequestType) (*DeleteOneRTTMRegionPoolEntryResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new HTTPRedirect */
	CreateHTTPRedirect(request *CreateHTTPRedirectRequestType) (*CreateHTTPRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single HTTPRedirect */
	GetOneHTTPRedirect(request *GetOneHTTPRedirectRequestType) (*GetOneHTTPRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every HTTPRedirect */
	GetHTTPRedirects(request *GetHTTPRedirectsRequestType) (*GetHTTPRedirectsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single HTTPRedirect */
	UpdateHTTPRedirect(request *UpdateHTTPRedirectRequestType) (*UpdateHTTPRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single HTTPRedirect */
	DeleteOneHTTPRedirect(request *DeleteOneHTTPRedirectRequestType) (*DeleteOneHTTPRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	CreateAdvRedirectRule(request *CreateAdvRedirectRuleRequestType) (*CreateAdvRedirectRuleResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UpdateAdvRedirectRule(request *UpdateAdvRedirectRuleRequestType) (*UpdateAdvRedirectRuleResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetOneAdvRedirectRule(request *GetOneAdvRedirectRuleRequestType) (*GetOneAdvRedirectRuleResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetAdvRedirectRules(request *GetAdvRedirectRulesRequestType) (*GetAdvRedirectRulesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneAdvRedirectRule(request *DeleteOneAdvRedirectRuleRequestType) (*DeleteOneAdvRedirectRuleResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new AdvRedirect */
	CreateAdvRedirect(request *CreateAdvRedirectRequestType) (*CreateAdvRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single AdvRedirect */
	GetOneAdvRedirect(request *GetOneAdvRedirectRequestType) (*GetOneAdvRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every AdvRedirect */
	GetAdvRedirects(request *GetAdvRedirectsRequestType) (*GetAdvRedirectsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single AdvRedirect */
	UpdateAdvRedirect(request *UpdateAdvRedirectRequestType) (*UpdateAdvRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteOneAdvRedirect(request *DeleteOneAdvRedirectRequestType) (*DeleteOneAdvRedirectResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new CDNManager */
	CreateCDNManager(request *CreateCDNManagerRequestType) (*CreateCDNManagerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single CDNManager */
	GetOneCDNManager(request *GetOneCDNManagerRequestType) (*GetOneCDNManagerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every CDNManager */
	GetCDNManagers(request *GetCDNManagersRequestType) (*GetCDNManagersResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single CDNManager */
	UpdateCDNManager(request *UpdateCDNManagerRequestType) (*UpdateCDNManagerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single CDNManager */
	DeleteOneCDNManager(request *DeleteOneCDNManagerRequestType) (*DeleteOneCDNManagerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	FailoverCDNManager(request *FailoverCDNManagerRequestType) (*FailoverCDNManagerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RecoverCDNManager(request *RecoverCDNManagerRequestType) (*RecoverCDNManagerResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetNodeList(request *GetNodeListRequestType) (*GetNodeListResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	PublishZone(request *PublishZoneRequestType) (*PublishZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	PruneZone(request *PruneZoneRequestType) (*PruneZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	FreezeZone(request *FreezeZoneRequestType) (*FreezeZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ThawZone(request *ThawZoneRequestType) (*ThawZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	RestoreZone(request *RestoreZoneRequestType) (*RestoreZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	BlockZone(request *BlockZoneRequestType) (*BlockZoneResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeleteZoneChangeset(request *DeleteZoneChangesetRequestType) (*DeleteZoneChangesetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetZoneChangeset(request *GetZoneChangesetRequestType) (*GetZoneChangesetResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetZoneNotes(request *GetZoneNotesRequestType) (*GetZoneNotesResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	UploadZoneFile(request *UploadZoneFileRequestType) (*UploadZoneFileResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	TransferZoneIn(request *TransferZoneInRequestType) (*TransferZoneInResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetTransferStatus(request *GetTransferStatusRequestType) (*GetTransferStatusResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetZoneConfigOptions(request *GetZoneConfigOptionsRequestType) (*GetZoneConfigOptionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	SetZoneConfigOptions(request *SetZoneConfigOptionsRequestType) (*SetZoneConfigOptionsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new IPTrack */
	CreateIPTrack(request *CreateIPTrackRequestType) (*CreateIPTrackResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single IPTrack */
	GetOneIPTrack(request *GetOneIPTrackRequestType) (*GetOneIPTrackResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every IPTrack */
	GetIPTracks(request *GetIPTracksRequestType) (*GetIPTracksResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single IPTrack */
	UpdateIPTrack(request *UpdateIPTrackRequestType) (*UpdateIPTrackResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single IPTrack */
	DeleteOneIPTrack(request *DeleteOneIPTrackRequestType) (*DeleteOneIPTrackResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateIPTrack(request *ActivateIPTrackRequestType) (*ActivateIPTrackResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateIPTrack(request *DeactivateIPTrackRequestType) (*DeactivateIPTrackResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new DNSSEC */
	CreateDNSSEC(request *CreateDNSSECRequestType) (*CreateDNSSECResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single DNSSEC */
	GetOneDNSSEC(request *GetOneDNSSECRequestType) (*GetOneDNSSECResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every DNSSEC */
	GetDNSSECs(request *GetDNSSECsRequestType) (*GetDNSSECsResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single DNSSEC */
	UpdateDNSSEC(request *UpdateDNSSECRequestType) (*UpdateDNSSECResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single DNSSEC */
	DeleteOneDNSSEC(request *DeleteOneDNSSECRequestType) (*DeleteOneDNSSECResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	ActivateDNSSEC(request *ActivateDNSSECRequestType) (*ActivateDNSSECResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	DeactivateDNSSEC(request *DeactivateDNSSECRequestType) (*DeactivateDNSSECResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault

	GetDNSSECTimeline(request *GetDNSSECTimelineRequestType) (*GetDNSSECTimelineResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every Task */
	GetTasks(request *GetTasksRequestType) (*GetTasksResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single Task */
	GetOneTask(request *GetOneTaskRequestType) (*GetOneTaskResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Cancels a Task */
	CancelTask(request *CancelTaskRequestType) (*CancelTaskResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Creates a new ExtNameserver */
	CreateExtNameserver(request *CreateExtNameserverRequestType) (*CreateExtNameserverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds single ExtNameserver */
	GetOneExtNameserver(request *GetOneExtNameserverRequestType) (*GetOneExtNameserverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Finds every ExtNameserver */
	GetExtNameservers(request *GetExtNameserversRequestType) (*GetExtNameserversResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Updates a single ExtNameserver */
	UpdateExtNameserver(request *UpdateExtNameserverRequestType) (*UpdateExtNameserverResponseType, error)

	// Error can be either of the following types:
	//
	//   - fault
	/* Deletes a single ExtNameserver */
	DeleteOneExtNameserver(request *DeleteOneExtNameserverRequestType) (*DeleteOneExtNameserverResponseType, error)
}
