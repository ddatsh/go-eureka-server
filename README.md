for local test

force push  frequently

# Go Eureka Server

```
post http://127.0.0.1:8761/eureka/apps/DEMO-SERVICE

{
    "instance": {
        "instanceId": "demo:demo-service:8080",
        "hostName": "192.168.1.100",
        "app": "DEMO-SERVICE",
        "ipAddr": "192.168.1.100",
        "status": "UP",
        "overriddenStatus": "UNKNOWN",
        "port": {
            "$": 8080,
            "@enabled": "true"
        },
        "securePort": {
            "$": 443,
            "@enabled": "false"
        },
        "countryId": 1,
        "dataCenterInfo": {
            "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
            "name": "MyOwn"
        },
        "leaseInfo": {
            "renewalIntervalInSecs": 30,
            "durationInSecs": 90,
            "registrationTimestamp": 0,
            "lastRenewalTimestamp": 0,
            "evictionTimestamp": 0,
            "serviceUpTimestamp": 0
        },
        "metadata": {
            "management.port": "8080"
        },
        "homePageUrl": "http://192.168.1.100:8082/",
        "statusPageUrl": "http://192.168.1.100:8082/actuator/info",
        "healthCheckUrl": "http://192.168.1.100:8082/actuator/health",
        "vipAddress": "demo-service",
        "secureVipAddress": "demo-service",
        "isCoordinatingDiscoveryServer": "false",
        "lastUpdatedTimestamp": "1574399263218",
        "lastDirtyTimestamp": "1574399271310"
    }
}

get http://127.0.0.1:8761/eureka/apps
```




# Go Config Server

```
http://127.0.0.1:8888/demo-service.yml
http://127.0.0.1:8888/demo-service.yaml
http://127.0.0.1:8888/demo-service.json

http://127.0.0.1:8888/demo-service-dev.yml

http://127.0.0.1:8888/demo-service/dev
```

