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

An Account is simply a collection of Users and Projects.

```
GET /accounts
GET /accounts/:id
POST /accounts
PATCH /accounts/:id
DELETE /accounts/:id
```


### Users

A User is a single login that has permission to view and change Projects within
an Account.

```
GET /users
GET /users/:id
POST /users
PATCH /users/:id
DELETE /users/:id
```


### Projects

A Project is a collection of Persons.

```
GET /projects
GET /projects/:id
POST /projects
PATCH /projects/:id
DELETE /projects/:id
```


### Persons

A Person is an entity that can have events tracked against them. All Person data
is stored inside Sky. They are created automatically through the tracking API.

```
# Retrieve the current state of a person.
GET /projects/:id/persons/:id

# Delete a person from the project.
DELETE /projects/:id/persons:id
```

Persons are just a collection of events so you can see the events associated with a Person:

```
# Retrieve a list of all events associated with a person.
GET /projects/:id/persons/:id/events

# Retrieve a single event at a given point in time associated with a person.
GET /projects/:id/persons/:id/events/:timestamp

# Update an event at a given timestamp.
PATCH /projects/:id/persons/:id/events/:timestamp

# Delete an event at a given timestamp.
DELETE /projects/:id/persons/:id/events/:timestamp
```


### Tracking API

Persons and their state only exist as a set of events. To create a person or to
change a person you must use the tracking API.

```
POST /track
```


