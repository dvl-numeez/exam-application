# exam-application

## API Requests

This document provides examples of `curl` commands for various operations with the application API.

## Fetching All Applications

To fetch all applications with a filter applied (e.g., by city):

```bash
curl -X POST  http://localhost:5000/getallapplications
```

## Fetching All Applications with a Filter

To fetch all applications with a filter applied (e.g., by city):

```bash
curl -X POST -d '{"city":"Melbourne"}' http://localhost:5000/getallapplications
```

## Making a New Application

To create a new application with the required fields:

```bash
curl -X POST http://localhost:5000/makeapplication \
-H "Content-Type: application/json" \
-d '{
    "firstName": "Mr",
    "middleName": "Virat",
    "lastName": "Kohli",
    "fatherFirstName": "Asif",
    "fatherLastName": "Baloch",
    "fatherMiddleName": "Khan",
    "gender": "MaLe",
    "city": "Melbourne",
    "district": "Ahmedabad",
    "stateOfDomicile": "Gujarat",
    "houseNo": "1",
    "village": "Modasa",
    "rollNumber": "355678692",
    "state": "Gujarat",
    "address": "1-Alfaruk society near Mohammadi park, Jivraj Park",
    "pincode": 380051,
    "yearOfPassing": "2023",
    "boardName": "Gujarat State Board",
    "homeDistrict": "Ahmedabad",
    "dob": "2001-09-24T07:44:09.000Z"
}'
```

## Updating an Application

To update an application by specifying the ID in the query parameter and the fields to update in the request body:

```bash
curl -X POST -d '{"firstName":"Numeez Khan Sahab"}' "http://localhost:5000/updateapplication?id=5a01368b-7559-42fb-8f17-0ab802f1bc64"
```

## Deleting an Application by ID

To delete an application by specifying the ID in the request body:

```bash
curl -X POST -d '{"id":"5a01368b-7559-42fb-8f17-0ab802f1bc64"}' http://localhost:5000/deleteapplication
```
