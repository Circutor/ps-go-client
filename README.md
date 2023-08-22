# PS GO client

Library that contains calls to the PowerStudio API

## Index

* [Installation](#installation)
* [Example](#example)
* [PsAllDevices](#PsAllDevices)
* [PsDeviceInfo](#PsDeviceInfo)
* [PsDevicesSelectionInfo](#PsDevicesSelectionInfo)
* [PsVarInfo](#PsVarInfo)
* [PsVarValue](#PsVarValue)
* [PsRecords](#PsRecords)

## Installation <a name="installation"></a>

     go get https://github.com/Circutor/ps-go-client

## Example <a name="example"></a>

```go
// Init logger.
// Example config logger with zap logger.
cfg := zap.Config{
    Encoding:    "console",
    Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
    OutputPaths: []string{"stdout"},
    EncoderConfig: zapcore.EncoderConfig{
		MessageKey:   "msg",
        LevelKey:     "level",
        TimeKey:      "time",
        CallerKey:    "caller",
        EncodeCaller: zapcore.ShortCallerEncoder,
        EncodeTime:   zapcore.ISO8601TimeEncoder,
    },
}

zapLogger, _ := cfg.Build()
sugaredLogger := zapLogger.Sugar()

// Adapter method info.
infoLogger := logger.Func(sugaredLogger.Info)

// Init logger pattern.
newLogger := logger.NewLogAdapter(infoLogger)

// ps methods. 
ps := powerStudioAPI.NewPowerStudio("localhost", "username", "password", newLogger)

// If the ps does not have authentication, the username and password values will be empty
ps := powerStudioAPI.NewPowerStudio("localhost", "", "", newLogger)

// If the tps is in another address, the address will be like this
ps := powerStudioAPI.NewPowerStudio("hostURL", "", "", newLogger)

// get list of devices.
devices, err := ps.PsAllDevices()

// get device info.
devicesInfo, err := ps.PsDeviceInfo([]string{"deviceName1", "deviceNameN"})

// get devices selection info.
devicesSelectionInfo, err := ps.PsDevicesSelectionInfo()

// get description var from device id or var name.
varsInfo, err := ps.PsVarInfo([]string{"deviceName1", "deviceNameN"}, []string{"varName1", "varNameN"})

// get value var from device id or var name.
varsValue, err := ps.PsVarValue([]string{"deviceName1", "deviceNameN"}, []string{"varName1", "varNameN"})

// get value records var name.
timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

records, err := PsRecords(timeBegin, timeEnd, 0, []string{"varName1", "varNameN"})
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

## Method `PsDevicesSelectionInfo()` <a name="PsDevicesSelectionInfo"></a>

Return a devices selection information.

* URI API
    * `http://<host>/services/devices/devicesSelectionInfo.xml`
  * Response
    ```go
    type DevicesSelectionInfo struct {
	       XMLName               xml.Name `xml:"devicesSelectionInfo"`
	       Text                  string   `xml:",chardata"`
	       DeviceID              string   `xml:"deviceId"`
	       CalculatedVariablesID string   `xml:"calculatedVariablesId"`
	       EventID               string   `xml:"eventId"`
	       Devices               struct {
		          Text   string `xml:",chardata"`
		          Device []struct {
			             Chardata  string `xml:",chardata"`
			             Enabled   string `xml:"enabled"`
			             CanSelect string `xml:"canSelect"`
			             Visible   string `xml:"visible"`
			             Event     string `xml:"event"`
			             Scada     string `xml:"scada"`
			             Report    string `xml:"report"`
			             Logger    struct {
				               Text  string `xml:",chardata"`
				               GSTD  string `xml:"GSTD"`
				               GARM  string `xml:"GARM"`
				               GHEVQ string `xml:"GHEVQ"`
				               GDEVQ string `xml:"GDEVQ"`
				               TEVE  string `xml:"TEVE"`
			             } `xml:"logger"`
			             ID    string `xml:"id"`
			             Type  string `xml:"type"`
			             Image struct {
				               Text    string `xml:",chardata"`
				               ImageID string `xml:"imageId"`
			             } `xml:"image"`
			             GUID          string `xml:"guid"`
			             Discriminable string `xml:"discriminable"`
			             Forced        string `xml:"forced"`
			             Text          struct {
				               Text    string `xml:",chardata"`
				               TextID  string `xml:"textId"`
				               TextStr string `xml:"textStr"`
			             } `xml:"text"`
			             DirectVars struct {
				               Text string `xml:",chardata"`
				               TEVE string `xml:"TEVE"`
			             } `xml:"directVars"`
			             DirectVarsFlags struct {
				               Text string `xml:",chardata"`
				               TEVE string `xml:"TEVE"`
			             } `xml:"directVarsFlags"`
		          } `xml:"device"`
	       } `xml:"devices"`
	       Root struct {
		          Text  string `xml:",chardata"`
		          Name  string `xml:"name"`
		          Group []struct {
			             Text  string `xml:",chardata"`
			             Name  string `xml:"name"`
			             Group []struct {
				                Text   string `xml:",chardata"`
				                Name   string `xml:"name"`
				                Device []struct {
					                  Text string `xml:",chardata"`
					                  ID   string `xml:"id"`
				                } `xml:"device"`
			             } `xml:"group"`
			             Device []struct {
				                Text string `xml:",chardata"`
				                ID   string `xml:"id"`
			             } `xml:"device"`
		          } `xml:"group"`
		          Device []struct {
			             Text string `xml:",chardata"`
			             ID   string `xml:"id"`
		          } `xml:"device"`
	       } `xml:"root"`
	       Image []struct {
		          Text    string `xml:",chardata"`
		          ImageID string `xml:"imageId"`
		          Image64 string `xml:"image64"`
	       } `xml:"image"`
	       Discriminators string `xml:"discriminators"`
	       Loggers        struct {
		          Text   string `xml:",chardata"`
		          Logger []struct {
			             Chardata string `xml:",chardata"`
			             Type     string `xml:"type"`
			             Text     struct {
				                Text   string `xml:",chardata"`
				                TextID string `xml:"textId"`
			             } `xml:"text"`
			             AllowDiscriminator string `xml:"allowDiscriminator"`
		          } `xml:"logger"`
	       } `xml:"loggers"`
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

## Method `PsRecords(begin, end time.time, period int, parameters []map[string]interface{}` <a name="PsRecords"></a>

Returns records value.

* URI API
    * `http://<host>/services/user/records.xml`
    * Parameters
        * **begin**: `?begin=DDMMYYYYHHMMSS`
        * **end**: `?end=DDMMYYYYHHMMSS`
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
		         Field    []struct {
			          Text  string `xml:",chardata"`
			          ID    string `xml:"id"`
			          Value string `xml:"value"`
		         } `xml:"field"`
		         FieldComplex []struct {
			          Text  string `xml:",chardata"`
			          ID    string `xml:"id"`
			          Value string `xml:"value"`
			          Flags string `xml:"flags"`
		         } `xml:"fieldComplex"`
		         FieldARM []struct {
			          Text    string `xml:",chardata"`
			          ID      string `xml:"id"`
			          Element []struct {
				            Text     string `xml:",chardata"`
				            Harmonic string `xml:"harmonic"`
				            Value    string `xml:"value"`
			          } `xml:"element"`
		         } `xml:"fieldARM"`
		         FieldFO []struct {
			          Text    string `xml:",chardata"`
			          ID      string `xml:"id"`
			          Element []struct {
				            Text  string `xml:",chardata"`
				            Msec  string `xml:"msec"`
				            Value string `xml:"value"`
			          } `xml:"element"`
		         } `xml:"fieldFO"`
		         FieldEVQ []struct {
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