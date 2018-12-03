package dynectsoap

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

var (
	PollingInterval = 2 * time.Second
	ApiUrl          = "https://api2.dynect.net/SOAP/"
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
			err = service.client.Call(ApiUrl, request, response)
			if err != nil {
				return nil, err
			}

			realResponse := response.Data.(JobResponseInterface)
			status := realResponse.status()

			switch status {
			case "incomplete":
				//Loop around again
				log.Debugf("Retrieving Job %v again, status is incomplete", response.Job_id)
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

	err := service.client.Call(ApiUrl, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) GetJob(request *GetJobRequestType) (*GetJobResponseType, error) {
	response := new(GetJobResponseType)

	err := service.client.Call(ApiUrl, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionLogin(request *SessionLoginRequestType) (*SessionLoginResponseType, error) {
	response := new(SessionLoginResponseType)
	err := service.client.Call(ApiUrl, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionLogout(request *SessionLogoutRequestType) (*SessionLogoutResponseType, error) {
	response := new(SessionLogoutResponseType)
	err := service.client.Call(ApiUrl, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionIsAlive(request *SessionIsAliveRequestType) (*SessionIsAliveResponseType, error) {
	response := new(SessionIsAliveResponseType)
	err := service.client.Call(ApiUrl, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dynect) SessionKeepAlive(request *SessionKeepAliveRequestType) (*SessionKeepAliveResponseType, error) {
	response := new(SessionKeepAliveResponseType)
	err := service.client.Call(ApiUrl, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
