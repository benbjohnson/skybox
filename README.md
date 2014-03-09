Skybox
======

Skybox is an open source behavioral analytics application that provides similar
functionality to tools like MixPanel or KISSMetrics. The idea is to provide a
simple API to store all the events about your users and a query interface for
analyzing how they behave.


## Getting Started

To run skybox, you'll just need to run Sky v0.4.0 or higher and execute the
skybox server:

```sh
$ skyboxd
```

An easier to way to get up and running is simply to use the prebuilt Docker container:

```sh
$ docker pull skybox/skyboxd
$ docker run skybox/skyboxd
```

<!-- TODO: Expand on 'Getting Started' -->


## API

### Accounts

An Account is simply a collection of Users and Persons.

```
GET /accounts
GET /accounts/:id
POST /accounts
PATCH /accounts/:id
DELETE /accounts/:id
```


### Users

```
GET /users
GET /users/:id
POST /users
PATCH /users/:id
DELETE /users/:id
```


### Tracking API

Persons and their state only exist as a set of events. To create a person or to
change a person you must use the tracking API.

```
POST /track
```


