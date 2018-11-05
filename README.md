# REST-API
Dies ist die REST-API für unser Schulungssystems mit der sich Schulungen anzeigen lassen und mit der eine Schulung gebucht werden kann. Folgende Funktionalitäten werden angeboten:

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



## Anzeige der Schulungen in einem bestimmten Zeitraum

**URL** : `/trainings/{start}/{stop}`

**Params** : 
{start} Anfang des Zeitraums als Uhrzeit Bsp: "12:20"
{stop} Ende des Zeitraums als Uhrzeit Bsp: "15:30"

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response
**Code** : `200 OK`



## Anzeige der Termine zu einer bestimmten Schulung

**URL** : `/trainings/{start}/{stop}`

**Params** : 
{start} Anfang des Zeitraums als Uhrzeit Bsp: "12:20"
{stop} Ende des Zeitraums als Uhrzeit Bsp: "15:30"

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
]
```
## Notes
Sind zu dem Zeitraums keine Schulungen vorhanden, kommt eine leerer HTML-Body zurück.

