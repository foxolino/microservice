#Aufgabe

Ein Anbieter von Schulungen möchte seine Produkte online anbieten. Dazu benötigt er eine Web-Anwendung, auf der sich potentielle Kunden über das Schulungsangebot informieren und Schulungen buchen können. Jede Schulung wird an verschiedenen Terminen angeboten. An jedem Termin stehen 10 Teilnehmerplätze zur Verfügung, die von potentiellen Kunden einzeln gebucht werden können.

Software-architektonisch ist ein Frontend in Javascript angedacht, das über eine REST-API mit einem Backend kommuniziert.

1. REST-API-Design
Designe eine REST-API, mit der die folgenden Frontend-Funktionalitäten abgebildet werden können:
* Übersicht über alle angebotenen Schulungen (Name der Schulung, Beschreibung, Name des Dozenten, Preis, ...)
* Anzeige der Schulungen in einem bestimmten Zeitraum
* Anzeige der Termine zu einer bestimmten Schulung
* Anlegen/Verändern einer neuen Schulung
* Anlegen/Verändern eines neuen Termins für eine Schulung
* Buchung einer Schulung an einem bestimmten Termin

Diese Aufgabe ist reine Konzeption. Es geht nicht um das Schreiben von Quellcode, sondern um die Definition der API.

2. Umsetzung Server-Seite
Programmiere die Server-Seite der API in einer Programmiersprache Deiner Wahl. Baue hierbei Tests für die API ein. Setze mindestens einen lesenden und einen modifizierenden REST-Call um. Treffe eine sinnvolle Auswahl der oben genannten Funktionalitäten, um das Prinzip der Umsetzung zu zeigen. Es ist nicht notwendig, alle oben genannte Funktionalität umzusetzen. Die Daten müssen nicht über das Beenden der Anwendung hinaus persistiert werden.

3. Umsetzung Client-Seite
Erstelle in einer Frontend-Technologie Deiner Wahl einen Client, mit dem exemplarisch ein bis zwei REST-Calls abgesetzt werden können. Bei dieser Aufgabe geht es nicht um eine grafisch schicke UI, sondern um einen PoC des Zusammenspiels zwischen Client und Server.

4. Security
Was würdest Du in dieser Anwendung zum Thema Security machen? Was muss warum abgesichert werden und welche Konzepte und Technologien würdest Du dafür einsetzen?
Diese Aufgabe ist wieder reine Konzeption. Eine Umsetzung ist nicht erforderlich.


## REST-API

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
Sind noch keine Schulungen vorhanden, kommt eine leerer HTML-Body zurück

## Anzeige der Schulungen in einem bestimmten Zeitraum

**URL** : `/trainings/{start}/{stop}`
**Params** : {start} Anfang des Zeitraums als Uhrzeit Bsp: "12:20"
{stop} Ende des Zeitraums als Uhrzeit Bsp: "15:30"
**Method** : `GET`
**Auth required** : NO
**Permissions required** : None

## Success Response
**Code** : `200 OK`
