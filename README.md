JSON API for getting Luas (Dublin light rail), times, fares and geo-coded data.

This endpoint relies on http://luasforecasts.rpa.ie .

Url: /stops

```javascript
{
    "lines":[
        {
            "name":"Luas Red Line",
            "stops":[
                {
                    "abrev":"TPT",
                    "isParkRide":false,
                    "isCycleRide":false,
                    "pronunciation":"The Point",
                    "lat":53.34835,
                    "long":-6.2292585
                },
                ...
            ]
        },
        {
            "name":"Luas Green Line",
            "stops":[
                {
                    "abrev":"BRO",
                    "isParkRide":false,
                    "isCycleRide":false,
                    "pronunciation":"Broombridge",
                    "lat":53.37224,
                    "long":-6.2976847
                },
                ...
            ]
        }
    ]
}
```

Url: /stops/{stop}

```javascript
{
    "created":"2019-04-01T19:33:16",
    "stop":"Broombridge",
    "stopAbv":"BRO",
    "message":"Green Line services operating normally",
    "directions":[
        {
            "name":"Inbound",
            "trams":[
                {
                    "dueMins":"",
                    "destination":"No trams forecast"
                }
            ]
        },
        {
            "name":"Outbound",
            "trams":[
                {
                    "dueMins":"7",
                    "destination":"Sandyford"
                },
                {
                    "dueMins":"19",
                    "destination":"Sandyford"
                }
            ]
        }
    ]
}
```

Url: /fares?from={from}&to={to}&children=true

```javascript
{
    "created":"2019-04-01T18:34:44",
    "params":{
        "from":"BRO",
        "to":"PHI",
        "adults":"1",
        "children":"0"
    },
    "result":{
        "peak":"2.10",
        "offpeak":"2.10",
        "zonesTravelled":"1"
    }
}
```