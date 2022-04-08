## MVFlights (Maldives Flights Schedule) Golang playground

Example:

```xml

<Arrivals>
    <UpdateTime>20220409 04:47</UpdateTime>
    <Flight>
        <AirLineID>GF</AirLineID>
        <AirLineName>GULF AIR</AirLineName>
        <FlightID>GF144</FlightID>
        <Route1>BAHRAIN</Route1>
        <Route2></Route2>
        <Route3></Route3>
        <Route4></Route4>
        <Route5></Route5>
        <Route6></Route6>
        <Scheduled>06:05</Scheduled>
        <Estimated>05:51</Estimated>
        <Status></Status>
        <Gate></Gate>
        <Terminal></Terminal>
        <CheckinArea></CheckinArea>
        <Date>20220409</Date>
        <PrimaryFlight></PrimaryFlight>
        <CarrierType>I</CarrierType>
    </Flight>
</Arrivals>
```

to

```json
[
    {
        "UpdateTime": "xxx",
        "Flight": [
            {
                "AirLineID": "GF",
                "AirLineName": "GULF AIR",
                "FlightID": "GF144",
                "Route1": "BAHRAIN",
                "Route2": " ",
                "Route3": " ",
                "Route4": " ",
                "Route5": " ",
                "Route6": " ",
                "Scheduled": "06:05",
                "Estimated": "05:51",
                "Status": " ",
                "Gate": " ",
                "Terminal": " ",
                "CheckinArea": " ",
                "Date": "20220409",
                "PrimaryFlight": " ",
                "CarrierType": "I"
            }
        ]
    }
]
```

### Install dependencies

`go get ./...`

### Run

`go run main.go`

### Debug

(in IntelliJ IDEA)

1. Mark the break point in the IDE
2. Right click **main.go**
3. Click **Debug go 'build main.go'**
4. Allow & start debugging.

### Testing

`go test -v`