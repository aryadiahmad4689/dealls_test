# Dating App Backend

## Overview
I wanna to disclaimer i never using tinder before, therefore i only follow the requirement to create this app/service.

This repository contains the backend implementation for a Dating Mobile App, similar to Tinder/Bumble. It's designed to handle user interactions and data processing for functionalities like sign-up, login, swiping, and premium features. The project is written in [Golang] and demonstrates RESTful API practices.

## Features
- **User Authentication**: Supports sign up and login functionalities.
- **Profile Swiping**: Users can view, swipe left (pass), and swipe right (like) on other dating profiles with a daily limit of 10 swipes.
- **Swipe Quota**: Profiles do not repeat within the same day.
- **Premium Packages**: Users can purchase premium packages to unlock features such as unlimited swipes and a verified user label.

## Tech Stack
- **Language**: [golang] i choose becauser is xfaster, mature and big community.
- **Database**: [sqlite3] - i choose because is easy to using for tecnical test like this, but for production i think is not good choice, i recomended to use [postgres] for database relational.
- **Managemen** : [makefile] i choose because is easily to run and migrate,this also complies with https://12factor.net/admin-processes
- **Envirotment**: [.env] i choose is easily manage envirotmenr,this also complies with https://12factor.net/config
- **Sofware Architecture** : [layered architecture] I chose it because it is easy to develop and is modular and meets the principles of clean architecture https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html


## Getting Started

### Prerequisites
- Golang 1.18 up
- VsCode

### Installation
1. Clone the repository:
   ```bash
   git clone github.com/aryadiahmad4689/dealls_test.git
2. Enter the main folder 

3. Copy the `.env.example ` to `.env`

4. Field full envirotment

5. Run `make migrate` this is to migrate database

6. Run `go mod tidy` to get dependency

7. Run `make run` to running apps

8. Congrats app is up

### Postman
You can import file json `Dating_Apps.json` to postman.
- i only handle positif test
- before you test in postman you must to create envirotmen and add variable `AUTH_USER`


### Testing
For unit test i just handle two layer (repository and usecase). I decided not to handle endpoint and handler because of time constraints. Please understand
 - Unit test `make test`
 - Coverage uni test `make test-coverage`

### Diagram
To see diagram you must to put code to plantuml online or install extension vscode PlantUML
 - ERD : https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/ERD.puml
 - Diagram Sequence
    - sign_in :  https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagram_sec_sign_in.puml
    - sign_up :  https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagram_sec_sign_up.puml
    - get_dating : https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagaram_sec_get_dating.puml
    - get_package :  https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagram_sec_get_package.puml
    - payment :  https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagram_sec_payment.puml
    - swipe_left :  https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagram_sec_swipe_left.puml
    - swipe_right :  https://github.com/aryadiahmad4689/dealls_test/blob/master/diagram/diagram_sec_swipe_right.puml