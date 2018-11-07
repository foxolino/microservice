# REST-API
Dies ist die REST-API für unser Schulungssystems mit der sich Schulungen anzeigen, verändern lassen und gebucht werden können. Folgende Funktionalitäten werden angeboten:
<!-- ====================================================================== -->
## Übersicht über alle angebotenen Schulungen (Name der Schulung, Beschreibung, Name des Dozenten, Preis, ...) 

**URL** : `/trainings`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`

**Content examples**

Für nur eine vorhandene Schulung mit der Beschreibung Technische Informatik
```json
[
    {
    "Id":11,
    "name":"Tech Inf",
    "description":"Technische Informatik",
    "Lecturer":{
        "Id":101,
        "Name":"Norbert Jung",
        "Age":55,
        "EMail":"jung@fbrs.de"},
    "Time":"0000-01-01T10:30:00Z",
    "Price":250
    }
]
```
## Notes
Sind noch keine Schulungen vorhanden, kommt eine leerer HTML-Body zurück.


<!-- ====================================================================== -->
## Anzeige der Schulungen in einem bestimmten Zeitraum

**URL** : `/trainings/{start}/{stop}`

**Params** :

{start} Anfang des Zeitraums als Uhrzeit Bsp: "11:20"

{stop} Ende des Zeitraums als Uhrzeit Bsp: "12:30"

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`

**Content examples**

```json
[
    {
        "Id":11,
        "name":"Tech Inf",
        "description":"Technische Informatik",
        "Lecturer":{
            "Id":101,
            "Name":"Norbert Jung",
            "Age":55,
            "EMail":"jung@fbrs.de"},
        "Time":"0000-01-01T10:30:00Z",
        "Price":250
    }
    {
        "Id":12,
        "name":"Theo Inf",
        "description":"Theoretische Informatik",
        "Lecturer":{
            "Id":102,
            "Name":"Kurt-Ulrich Witt",
            "Age":59,
            "EMail":"witt@fbrs.de"},
        "Time":"0000-01-01T11:30:00Z",
        "Price":400
    }
]
```
## Notes
Ist zu dem Zeitraums keine Schulungen vorhanden, kommt eine leerer HTML-Body zurück.

<!-- ====================================================================== -->
## Anzeige der Termine zu einer bestimmten Schulung

**URL** : `/trainings/{name}`

**Params** : 
{name} Name der Schulung Bsp: "Technische Informatik"

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`


**Content examples**

```json
[
    {
        "Id":11,
        "name":"Tech Inf",
        "description":"Technische Informatik",
        "Lecturer":{
            "Id":101,
            "Name":"Norbert Jung",
            "Age":55,
            "EMail":"jung@fbrs.de"},
        "Time":"0000-01-01T10:30:00Z",
        "Price":250
    }
    {
        "Id":11,
        "name":"Tech Inf",
        "description":"Technische Informatik",
        "Lecturer":{
            "Id":101,
            "Name":"Norbert Jung",
            "Age":55,
            "EMail":"jung@fbrs.de"},
        "Time":"0000-01-01T16:30:00Z",
        "Price":250
    }
]
```
## Notes
Ist zu dem Zeitraums keine Schulungen vorhanden, kommt eine leerer HTML-Body zurück.


<!-- ====================================================================== -->
## Verändern einer Schulung
**URL** : `/trainings/{id}`

**Params** : 
{id} ID der Schulung

**Method** : `POST`

**Content-Type** : Application/json

```json
{
"Id":11,
"name":"Tech Inf",
"description":"Technische Informatik",
"Lecturer":{
    "Id":101,
    "Name":"Norbert Jung",
    "Age":55,
    "EMail":"jung@fbrs.de"},
"Time":"0000-01-01T10:30:00Z",
"Price":250
}
```
**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`

## Error Response
**Code** : `404 Not Found`

<!-- ====================================================================== -->
## Löschen einer Schulung
**URL** : `/trainings/{id}`

**Params** : 
{id} ID der Schulung

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`

## Error Response
**Code** `404 Not Found`
<!-- ====================================================================== -->
## Buchung einer Schulung an einem bestimmten Termin
**URL** : `/book/{id}`

**Params** : 
{id} ID der Schulung

**Method** : `POST`

**Content-Type** : Application/json

```json
{
"name":"Hans Müller",
"address": "Oxfordstr. 77, 53111 Bonn",
"payment": "..."
}
```
**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`

## Error Response
**Code** `404 Not Found`

## Notes
Z.Z. wird nur das anonyme Buchen einer Schulung unterstützt. 
