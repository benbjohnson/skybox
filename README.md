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


## Roadmap

TODO

