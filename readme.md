# Dating App Backend

## Overview
I wanna to disclaimer i never using tinder before, therefore i only follow the requirement to create this app/service.

This repository contains the backend implementation for a Dating Mobile App, similar to Tinder/Bumble. It's designed to handle user interactions and data processing for functionalities like sign-up, login, swiping, and payment/purchasing. The project is written in [Golang] and demonstrates RESTful API practices.

## Features
- **User Authentication**: Supports sign up and login functionalities.
- **Profile Swiping**: Users can view, swipe left (pass), and swipe right (like) on other dating profiles with a daily limit of 10 swipes.
- **Swipe Quota**: Profiles do not repeat within the same day.
- **Premium Packages**: Users can purchase premium packages to unlock features such as unlimited swipes and a verified user label.

## Tech Stack
- **Language**: [golang] i choose becauser is faster, mature and big community.
- **Database**: [sqlite3] - i choose because is easy to using for tecnical test like this, but for production i think is not good choice, i recomended to use [postgres] for database relational.
- **Managemen** : [makefile] i choose because is easily to run and migrate,this also complies with https://12factor.net/admin-processes
- **Envirotment**: [.env] i choose is easily manage envirotment,this also complies with https://12factor.net/config
- **Sofware Architecture** : [layered architecture] I chose it because it is easy to develop and is modular and meets the principles of clean architecture https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html


## Getting Started

### Prerequisites
- Golang 1.18 up
- VsCode
- PlantUML vscode extension
- Postman

### Detail And Structure
- `cmd`: Main of service
- `src`: Content and logic of the application
  - `app`: Main configuration settings of the application
    - `db`: To create configuration and connection to the database
  - `middleware`: To create middleware and manage people who log in to the application
  - `modules`: Business logic and core application

### Dependencies
- github.com/dgrijalva/jwt-go v3.2.0+incompatible
- github.com/go-chi/chi v1.5.5
- github.com/go-chi/cors v1.2.1
- github.com/jabardigitalservice/golog v0.0.8
- github.com/joho/godotenv v1.5.1
- github.com/mattn/go-sqlite3 v1.14.19
- github.com/spf13/viper v1.18.2
- github.com/stretchr/testify v1.8.4
- golang.org/x/crypto v0.16.0
- gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
- gopkg.in/go-playground/assert.v1 v1.2.1

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/aryadiahmad4689/dealls_test.git 
   or 
   git clone github.com/aryadiahmad4689/dealls_test.git
2. Enter the main folder 

3. Copy the `.env.example ` to `.env`

4. Field full envirotment

5. Run `make migrate` this is to migrate database

6. Run `go mod tidy` to get dependency

7. Run `make run` to running apps

8. Congrats app is up

### Url Api And Method
 - `/v1/users/sign-up` method `post`
 - `/v1/users/sign-in` method `post`
 - `/v1/swipes/get-dating` Method `get`
 - `/v1/swipes/right/{id}` Method `put`
 - `/v1/swipes/left/{id}` Method `put`
 - `/v1/packages` Method `get`
 - `/v1/subscriptions/payment` Method `post`

### Postman
You can import file json `Dating_Apps.json` to postman.
- i only handle positif test
- before you test in postman you must to create envirotmen and add variable `AUTH_USER`


### Testing
For unit test i just handle two layer (repository and usecase). I decided not to handle endpoint and handler because of time constraints. Please understand
 - Unit test run `make test`
 - Coverage uni test run `make test-coverage`

### Linting
 - VsCode go extension

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

### Advice
I understand that my system definitely has a lot of gaps and there are still many things I haven't done, therefore I want to suggest several things in the application that can be added
 - for databases you should use `postgres`
 - create validation
 - complete unit test
 - add log in application
 - should using transaction
 - mapping error database

## Thanks