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

// get description var from device id.
varsInfo, err := ps.PsVarInfo([]map[string]interface{}{
    {"id", "deviceName1"}
    {"id", "deviceNameN"}
})

// get description var from var name.
varsInfo, err := ps.PsVarInfo([]map[string]interface{}{
    {"var", "varName1"}
    {"var", "varNameN"}
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

Return a devices information.

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

## Method `PsVarInfo()`

Returns variable information.

* URI API
    * `http://<host>/services/user/varInfo.xml`
    * Parameters
        * **var**: `?var=deviceVar-1?var=DeviceVar-n`
        * **id**: `?id=deviceName`
          * If parameter content `id` return all variables from device.
* Response
    ```go
    type VarInfo struct {
          XMLName xml.Name `xml:"varInfo"`
          Text    string   `xml:",chardata"`
          Var     []struct {
		            Text         string `xml:",chardata"`
		            ID           string `xml:"id"`
		            IdEx         string `xml:"idEx"`
		            Title        string `xml:"title"`
		            HasValue     string `xml:"hasValue"`
		            HasLogger    string `xml:"hasLogger"`
		            HasForced    string `xml:"hasForced"`
		            SampleMode   string `xml:"sampleMode"`
		            MeasureUnits string `xml:"measureUnits"`
		            UnitsFactor  string `xml:"unitsFactor"`
		            Decimals     string `xml:"decimals"`
		            VarType      string `xml:"varType"`
		            ValueInfo    struct {
			              Text     string `xml:",chardata"`
			              CtrlType string `xml:"ctrlType"`
			              Type     string `xml:"type"`
		            } `xml:"valueInfo"`
	         } `xml:"var"`
    }
  ```