#  Generation Next: Mind the Lacuba

Welcome to the Generation Next road safety project! This project aims to aims to improve road safety and facilitate communication between citizens and authorities responsible for maintaining roads. 

It is a simple mobile app that allows users to report road pits to the responsible authorities by sending an email with the current coordinates and a predefined message.

When a user encounters a road pit, they can tap a button to open an email client pre-populated with the current coordinates and a message requesting the pit's repair. The email recipient can click on a link within the email to open an HTML form, which allows them to confirm the pit has been fixed and remove its coordinates from the database.

## Table of Contents

- [Overview](#overview)
- [Implementation](#implementation)
  - [Frontend](#frontend)
  - [Backend](#backend)
  - [Incident Classification Service](#incident-classification-service)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Overview

The project consists of three main components:

1. A user-friendly mobile app for reporting incidents.
2. A backend server for storing and managing data.

## Implementation

### Frontend

The frontend of our app is developed using [App Inventor](https://appinventor.mit.edu). You can access the App Inventor project files [here](https://mega.nz/file/FXBn1IBS#KwI89DcGjmgiGSbfv7_6894cWiZlxjL_4Mn_rEm5fmo).

### Backend

The backend is implemented using the [Gin Web Framework](https://github.com/gin-gonic/gin) in Go, and MySQL is used as the database. The server is responsible for receiving pit coordinates reports from the frontend, storing them in the database, and managing the data.
