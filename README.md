# PS GO client

Library that contains calls to the PowerStudio API

## Instance library

```go
// ps methods. 
ps := powerStudioAPI.NewPowerStudio("localhost")

// get list of devices.
devices, err := ps.PsAllDevices()

// get device info.
devicesInfo, err := ps.PsDeviceInfo([]map[string]interface{}{
    {"id", "deviceName1"}
    {"id", "deviceNameN"}
})
```

## Method `PsAllDevices()`

Returns the list of configured devices.

* URI API
    * `http://<host>/services/user/devices.xml`
* Response
    ```go
    type Devices struct {
        XMLName xml.Name `xml:"devices"`
        Text    string   `xml:",chardata"`
        ID      []string `xml:"id"`
    }
  ```

## Method `PsDeviceInfo()`

* URI API
    * `http://<host>/services/user/deviceInfo.xml`
    * Parameters
        * **id**: `?id=deviceName-1?id=DeviceName-n`
* Response
    ```go
    type DevicesInfo struct {
          XMLName xml.Name `xml:"devices"`
	      Text    string   `xml:",chardata"`
	      Device  []struct {
		        Text            string   `xml:",chardata"`
		        ID              string   `xml:"id"`
		        Type            string   `xml:"type"`
		        TypeDescription string   `xml:"typeDescription"`
		        Var             []string `xml:"var"`
		        SerialNumber    string   `xml:"serialNumber"`
		        Modules         struct {
			          Text   string `xml:",chardata"`
                      Module struct {
				            Text  string `xml:",chardata"`
				            StNum string `xml:"stNum"`
				            Model string `xml:"model"`
			          } `xml:"module"`
		        } `xml:"modules"`
	      } `xml:"device"`
    }
  ```