package powerstudio

import (
	"strconv"
	"time"
)

const (
	HTTTP = "http://"

	URIAllDevices           = "/services/user/devices.xml"
	URIDevicesInfo          = "/services/user/deviceInfo.xml"
	URIDevicesSelectionInfo = "/services/devices/devicesSelectionInfo.xml"
	URIVarInfo              = "/services/user/varInfo.xml"
	URIVarValue             = "/services/user/values.xml"
	URIRecord               = "/services/user/records.xml"
)

// ParseParameters parse list ids and list vars to parameters accept power studio.
func ParseParameters(ids, vars []string) []map[string]interface{} {
	parameters := make([]map[string]interface{}, 0)

	for _, id := range ids {
		parameters = append(parameters, map[string]interface{}{"id": id})
	}

	for _, variable := range vars {
		parameters = append(parameters, map[string]interface{}{"var": variable})
	}

	return parameters
}

// ParseDateToPsFormat convert date type time.time to format power studio.
func ParseDateToPsFormat(time time.Time) string {
	return strconv.Itoa(time.Day()) + strconv.Itoa(int(time.Month())) + strconv.Itoa(time.Year())
}
