# PS GO client

Library that contains calls to the PowerStudio API

## Index

* [Installation](#installation)
* [Example](#example)
* [PsAllDevices](#PsAllDevices)
* [PsDeviceInfo](#PsDeviceInfo)
* [PsVarInfo](#PsVarInfo)
* [PsVarValue](#PsVarValue)
* [PsRecords](#PsRecords)

## Installation <a name="installation"></a>

     go get https://github.com/Circutor/ps-go-client

## Example <a name="example"></a>

```go
// ps methods. 
ps := powerStudioAPI.NewPowerStudio("localhost")

// get list of devices.
devices, err := ps.PsAllDevices()

// get device info.
devicesInfo, err := ps.PsDeviceInfo([]{"deviceName1", "...", "deviceNameN"})

// get description var from device id or var name.
varsInfo, err := ps.PsVarInfo([]{"deviceName1", "...", "deviceNameN"}, []{"varName1", "...", "varNameN"})

// get value var from device id or var name.
varsValue, err := ps.PsVarValue([]{"deviceName1", "...", "deviceNameN"}, []{"varName1", "...", "varNameN"})

// get value records var name.
records, err := PsRecords("18102022", "18102022", 0,[]map[string]interface{}{
    {"var", "varName1"}
    {"var", "varNameN"}
})
```

## Method `PsAllDevices()` <a name="PsAllDevices"></a>

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

## Method `PsDeviceInfo(parameters []map[string]interface{})` <a name="PsDeviceInfo"></a>

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

## Method `PsVarInfo(parameters []map[string]interface{})` <a name="PsVarInfo"></a>

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

## Method `PsVarValue(parameters []map[string]interface{})` <a name="PsVarValue"></a>

Returns variable value.

* URI API
    * `http://<host>/services/user/values.xml`
    * Parameters
        * **var**: `?var=deviceVar-1?var=DeviceVar-n`
        * **id**: `?id=deviceName`
            * If parameter content `id` return all variables from device.
* Response
    ```go
    type Values struct {
	      XMLName  xml.Name `xml:"values"`
	      Text     string   `xml:",chardata"`
	      Variable []struct {
		        Text      string `xml:",chardata"`
		        ID        string `xml:"id"`
		        Value     string `xml:"value"`
		        TextValue string `xml:"textValue"`
	      } `xml:"variable"`
    }
  ```

## Method `PsRecords(begin, end string, period int, parameters []map[string]interface{}` <a name="PsRecords"></a>

Returns records value.

* URI API
    * `http://<host>/services/user/records.xml`
    * Parameters
        * **begin**: `?begin=DDMMYYYY`
        * **end**: `?end=DDMMYYYY`
        * **period**: `?period=vuale`: Default value 0.
        * **var**: `?var=deviceVar-1?var=DeviceVar-n`
* Response
    ```go
    type RecordGroup struct {
	      XMLName xml.Name `xml:"recordGroup"`
	      Text    string   `xml:",chardata"`
	      Period  string   `xml:"period"`
	      Record  []struct {
		         Text     string `xml:",chardata"`
		         DateTime string `xml:"dateTime"`
		         Field    struct {
			          Text  string `xml:",chardata"`
			          ID    string `xml:"id"`
			          Value string `xml:"value"`
		         } `xml:"field"`
		         FieldComplex struct {
			          Text  string `xml:",chardata"`
			          ID    string `xml:"id"`
			          Value string `xml:"value"`
			          Flags string `xml:"flags"`
		         } `xml:"fieldComplex"`
		         FieldARM struct {
			          Text    string `xml:",chardata"`
			          ID      string `xml:"id"`
			          Element []struct {
				            Text     string `xml:",chardata"`
				            Harmonic string `xml:"harmonic"`
				            Value    string `xml:"value"`
			          } `xml:"element"`
		         } `xml:"fieldARM"`
		         FieldFO struct {
			          Text    string `xml:",chardata"`
			          ID      string `xml:"id"`
			          Element []struct {
				            Text  string `xml:",chardata"`
				            Msec  string `xml:"msec"`
				            Value string `xml:"value"`
			          } `xml:"element"`
		         } `xml:"fieldFO"`
		         FieldEVQ struct {
			          Text             string `xml:",chardata"`
			          ID               string `xml:"id"`
			          Value            string `xml:"value"`
			          Phase            string `xml:"phase"`
			          Duration         string `xml:"duration"`
			          AverageValue     string `xml:"averageValue"`
			          PreviousValue    string `xml:"previousValue"`
			          EventType        string `xml:"eventType"`
			          EndForced        string `xml:"endForced"`
			          SemicycleVoltage []struct {
				            Text  string `xml:",chardata"`
				            Date  string `xml:"date"`
				            Value string `xml:"value"`
			          } `xml:"semicycleVoltage"`
		         } `xml:"fieldEVQ"`
	      } `xml:"record"`
    }
  ```