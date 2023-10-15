# Go Lang Parking system POC
Example that shows sample POC for parking system in Golang projects.

## API:
### GET /parking-spaces
Returns all parking details 

##### Example Output:
```
[
    {
        "id": 1,
        "level": 1,
        "type": "Compact",
        "occupied": false,
        "reserver": false,
        "ReservedBy": "",
        "ReservationID": ""
    }
]
```
### POST /reserve-park
used to reserve the parking slot usign id

##### Example Input:
```
{
    "space_id":2,
    "recerved_by":"sakthivel"
}
```
##### Example Output:
If parking available and reserved successfully
```
{
    "recervation_id": "99c1940e-faf4-47e9-9797-dd37f1d94129",
    "message": "Parking reserved successfully"
}
```
If parking is not available
```
{
    "error": "Parking space is already occupied"
}
```
If parking is already reserved
```
{
    "error": "Parking space is already reserved by other."
}
```

### POST /park
used to park the vehical using available id or with id and recervation_id

##### Example Input:
```
{
    "space_id":1
}
```
```
{
    "space_id":2,
    "reservation_id":"99c1940e-faf4-47e9-9797-dd37f1d94129"
}
```

##### Example Output:
If parking available
```
{
    "message": "Car parked successfully"
}
```
If parking is not available
```
{
    "error": "Parking space is already occupied"
}
```
If parking is already reserved
```
{
    "error": "Parking space is already reserved by other."
}
```

## API:
### GET /unpark
Returns all parking details 

##### Example Input:
```
{
    "space_id":1
}
```

##### Example Output:
```
{
    "error": "Parking space is already unparked"
}
```
```
{
    "message": "Car unparked successfully"
}
```