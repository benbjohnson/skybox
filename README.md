Skybox
======

Skybox is an open source user tracking and funnel analysis application.
It provides a simple, drop-in JavaScript snippet that lets you automatically track visits to your site.
You can then use the Skybox application to build funnels on the fly to see what your users are doing.


## Getting Started

If you want to try out Skybox, you can sign up at:

[http://skyboxanalytics.com](http://skyboxanalytics.com)


### Running it on your server

To install Skybox, you'll first need to install [SkyDB](http://github.com/skydb/sky).
The easiest way to do this is to run it using [Docker](https://www.docker.io/):

```sh
$ docker pull skydb/sky-llvm
$ docker -t -i -p 8585:8585 -v ~/sky:/var/lib/sky:rw skydb/sky-llvm 
```

Then in a separate window, download the latest version of skybox and run it.

```sh
$ wget https://github.com/skybox/skybox/releases/download/v0.1.0/skyboxd_v0.1.0_linux_amd64.tar.gz
$ tar zxvf skyboxd_v0.1.0_linux_amd64.tar.gz
$ cd skyboxd_v0.1.0_linux_amd64
$ ./skyboxd --data-dir ~/skybox
Listening on http://localhost:7000
```

Navigate to [http://localhost:7000](http://localhost:7000) and you can create an account.
Once you're logged in you'll see the JavaScript snippet you'll need to paste onto your web site.


## Events

Skybox works by tracking events associated with individual users.
These events track several pieces of information:

- Channel: The JavaScript snippet always sets this to `web`. Once the API is available then this can be defined by the API consumer.

- Resource: For web sites, this is the type of page being accessed. The JavaScript snippet automatically generalizes the URL path to create the resource. By default, any numeric sections of the path will be removed. For example, a visit to `/users/5/friends/182` will be converted to `/users/:id/friends/:id`.

- Action: The JavaScript snippet uses `view` by default but this will be definable in the future. For example, a single resource might have multiple events occur with different actions (e.g. `click`, `form submit`, etc).

- Domain: The hostname of the web site being visited.

- Path: The original URL path accessed.

Currently the Skybox interface only works with the resource when defining funnels but the funnel language will be opened up in the near future.


## Funnel Analysis

The most basic tool for behavioral analysis is the funnel.
A funnel defines multiple steps that a user is expected to complete.

For example, a conversion funnel might look something like this:

```
/index.html -> /pricing -> /signup
```

By defining the expected behavior we can use Skybox to tell us how many people made it to each step along the way:

```
    500
  ███████
  ███████         200
  ███████       ███████         50
  ███████       ███████       ███████
  ███████       ███████       ███████
/index.html -> /pricing   ->  /signup
```

Here we can see that 500 people went to our home page, only 200 of those people made it to the pricing page, and finally only 50 people actually made it to the sign up page. By finding where users drop off along the way we can find problem areas in our user experience and fix them.


## Roadmap

The aim of Skybox is to provide really useful data analysis and visualization.
Many tools throw on a ton of useless features and eye candy but Skybox is meant to be simple.
Because of this goal, we want to add features as people need them.
If you'd like to see a feature added please add a Github issue to discuss it.

Some of the features we're thinking about adding include:

* REST API - You should own your data and do whatever you want with it.

* Interactive funnel analysis - Drill into charts on the fly to understand behavior better.

* Reverse funnel analysis - Start from an event and work backwards to see what events led up to it.

* Cohort analysis - See how different groups perform over time.

Please let us know your thoughts!

Skybox is built on top of SkyDB so there's a powerful query engine to take advantage of.
We want to make sure that Skybox is providing the best view into that data that it can.
