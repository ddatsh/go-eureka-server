package model

import "encoding/xml"

// StringMap is a map[string]string.
type StringMap map[string]string

// StringMap marshals into XML.
func (s StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	return e.Flush()
}

type MetaData struct {
	//Map   map[string]string
	//Map   StringMap
	XMLName xml.Name
	Value   string `xml:",chardata"`
	Class string
}

type Applications struct {
	XMLName           xml.Name `xml:"applications"`
	VersionsDelta int           `xml:"versions__delta" json:"versions__delta"`
	AppsHashcode  string        `xml:"apps__hashcode" json:"apps__hashcode"`
	Applications  []Application `xml:"application,omitempty" json:"application"`
}

type Application struct {
	Name      string         `xml:"name"`
	Instances []InstanceInfo `xml:"instance"`
}

type Instance struct {
	Instance InstanceInfo `xml:"instance" json:"instance"`
}

type Port struct {
	Port    int  `xml:",chardata" json:"$"`
	Enabled string `xml:"enabled,attr" json:"@enabled"`
}

type InstanceInfo struct {
	App                           string          `xml:"app" json:"app"`
	HostName                      string          `xml:"hostName" json:"hostName"`
	HomePageUrl                   string          `xml:"homePageUrl,omitempty" json:"homePageUrl,omitempty"`
	StatusPageUrl                 string          `xml:"statusPageUrl" json:"statusPageUrl"`
	HealthCheckUrl                string          `xml:"healthCheckUrl,omitempty" json:"healthCheckUrl,omitempty"`
	IpAddr                        string          `xml:"ipAddr" json:"ipAddr"`
	VipAddress                    string          `xml:"vipAddress" json:"vipAddress"`
	SecureVipAddress              string          `xml:"secureVipAddress,omitempty" json:"secureVipAddress,omitempty"`
	Status                        string          `xml:"status" json:"status"`
	Port                          Port           `xml:"port,omitempty" json:"port,omitempty"`
	SecurePort                    Port           `xml:"securePort,omitempty" json:"securePort,omitempty"`
	DataCenterInfo                DataCenterInfo `xml:"dataCenterInfo" json:"dataCenterInfo"`
	LeaseInfo                     LeaseInfo      `xml:"leaseInfo,omitempty" json:"leaseInfo,omitempty"`
	//Metadata                      *MetaData       `xml:"metadata,omitempty" json:"metadata,omitempty"`
	Metadata                      StringMap       `xml:"metadata,omitempty" json:"metadata,omitempty"`
	IsCoordinatingDiscoveryServer string            `xml:"isCoordinatingDiscoveryServer,omitempty" json:"isCoordinatingDiscoveryServer,omitempty"`
	LastUpdatedTimestamp          string             `xml:"lastUpdatedTimestamp,omitempty" json:"lastUpdatedTimestamp,omitempty"`
	LastDirtyTimestamp            string             `xml:"lastDirtyTimestamp,omitempty" json:"lastDirtyTimestamp,omitempty"`
	ActionType                    string          `xml:"actionType,omitempty" json:"actionType,omitempty"`
	OverriddenStatus              string          `xml:"overriddenstatus,omitempty" json:"overriddenStatus,omitempty"`
	CountryId                     int             `xml:"countryId,omitempty" json:"countryId,omitempty"`
	InstanceID                    string          `xml:"instanceId,omitempty" json:"instanceId,omitempty"`
}
type DataCenterInfo struct {
	Name     string              `xml:"name" json:"name"`
	Class    string              `xml:"class,attr" json:"@class"`
	Metadata *DataCenterMetadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
}

type DataCenterMetadata struct {
	AmiLaunchIndex   string `xml:"ami-launch-index,omitempty" json:"ami-launch-index,omitempty"`
	LocalHostname    string `xml:"local-hostname,omitempty" json:"local-hostname,omitempty"`
	AvailabilityZone string `xml:"availability-zone,omitempty" json:"availability-zone,omitempty"`
	InstanceId       string `xml:"instance-id,omitempty" json:"instance-id,omitempty"`
	PublicIpv4       string `xml:"public-ipv4,omitempty" json:"public-ipv4,omitempty"`
	PublicHostname   string `xml:"public-hostname,omitempty" json:"public-hostname,omitempty"`
	AmiManifestPath  string `xml:"ami-manifest-path,omitempty" json:"ami-manifest-path,omitempty"`
	LocalIpv4        string `xml:"local-ipv4,omitempty" json:"local-ipv4,omitempty"`
	Hostname         string `xml:"hostname,omitempty" json:"hostname,omitempty"`
	AmiId            string `xml:"ami-id,omitempty" json:"ami-id,omitempty"`
	InstanceType     string `xml:"instance-type,omitempty" json:"instance-type,omitempty"`
}

type LeaseInfo struct {
	EvictionDurationInSecs uint `xml:"evictionDurationInSecs,omitempty" json:"evictionDurationInSecs,omitempty"`
	RenewalIntervalInSecs  int  `xml:"renewalIntervalInSecs,omitempty" json:"renewalIntervalInSecs,omitempty"`
	DurationInSecs         int  `xml:"durationInSecs,omitempty" json:"durationInSecs,omitempty"`
	RegistrationTimestamp  int  `xml:"registrationTimestamp,omitempty" json:"registrationTimestamp,omitempty"`
	LastRenewalTimestamp   int  `xml:"lastRenewalTimestamp,omitempty" json:"lastRenewalTimestamp,omitempty"`
	EvictionTimestamp      int  `xml:"evictionTimestamp,omitempty" json:"evictionTimestamp,omitempty"`
	ServiceUpTimestamp     int  `xml:"serviceUpTimestamp,omitempty" json:"serviceUpTimestamp,omitempty"`
}
