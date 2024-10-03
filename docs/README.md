## Endpoints

### User Register
<p>Method: <b>POST</b></p>
<p>URL: http://localhost:8000/user</p>
<p>Body: </p>

```
{
    "name": "User Postman",
    "username": "userpostman",
    "email": "user@postman.com",
    "password": "123456",
    "birth_date": "1997-10-10"
}
```


Response: 
```
{
    "id": 2,
    "name": "User Postman",
    "username": "userpostman",
    "email": "user@postman.com",
    "password": "123456",
    "birthDate": "1997-10-10T00:00:00Z",
    "createdAt": "2024-10-03T18:30:39.499582345Z",
    "updatedAt": "2024-10-03T18:30:39.499582345Z",
    "deletedAt": null
}
```
### Billing Address Register
<p>Method: <b>POST</b></p>
<p>URL: http://localhost:8000/billing_address</p>
<p>Body: </p>

```
{
  "street": "123 Main Street",
  "number": "456",
  "zip_code": "10001",
  "city": "New York",
  "country": "US",
  "state": "NY",
  "user_id": "1"
}
```

Response: 
```
{
    "id": 1,
    "street": "123 Main Street",
    "number": "456",
    "zipCode": "10001",
    "city": "New York",
    "country": "US",
    "state": "NY",
    "userId": 1,
    "createdAt": "0001-01-01T00:00:00Z",
    "updatedAt": "0001-01-01T00:00:00Z",
    "deletedAt": null
}
```

### Card Register
<p>Method: <b>POST</b></p>
<p>URL: http://localhost:8000/card</p>
<p>Body: </p>

```
{
    "card_name": "API Gateway Card",
    "number": "1234567890123456",
    "holder_name": "Jhon Doe",
    "expiry_month": 12,
    "expiry_year": 2025,
    "cvv": 123,
    "user_id": 1,
    "address_id": 1
}
```

Response: 
```
{
    "id": 5,
    "cardName": "API Gateway Card",
    "number": "1234567890123456",
    "holderName": "Jhon Doe",
    "expiryMonth": 12,
    "expiryYear": 2025,
    "cvv": 123,
    "userId": 1,
    "addressId": 1,
    "createdAt": "0001-01-01T00:00:00Z",
    "updatedAt": "0001-01-01T00:00:00Z",
    "deletedAt": null
}
```