# PS GO client

Library that contains calls to the PowerStudio API

## Instance library

```go
// ps methods. 
ps := powerStudioAPI.NewPowerStudio("localhost")

// get list of devices.
devices, err := ps.PsAllDevices()
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
