Skybox
======

Skybox is an open source user track and funnel analysis application.
It provides a simple, drop-in JavaScript snippet that lets you to automatically track visits to your site.
You can then go to the application and build funnels on the fly to see what your users are doing.


## Getting Started

If you want to try out Skybox, you can sign up for an account at: [http://skyboxanalytics.com](http://skyboxanalytics.com)


### Running it on your server

To install Skybox, you'll first need to get the [Sky behavioral analytics database](http://github.com/skydb/sky) installed.
The easiest way to do this is to run it using Docker:

```sh
$ docker pull skydb/sky-llvm
$ docker -t -i -p 8585:8585 -v ~/sky:/var/lib/sky:rw skydb/sky-llvm 
```

Then in a separate window, download the latest version of skybox and run it.

```sh
$ wget https://github.com/skybox/skybox/releases/download/v0.1.0/skybox_v0.1.0_linux_amd64.tar.gz
$ tar zxvf skybox_v0.1.0_linux_amd64.tar.gz
$ cd skybox_v0.1.0_linux_amd64
$ ./skybox --data-dir ~/skybox
Listening on http://localhost:7000
```

Now navigate to [http://localhost:7000](http://localhost:7000) and you can create an account.
Once you're logged in you'll see the JavaScript snippet you'll need to paste onto your web site.


## Events & Funnel Analysis

TODO


## Roadmap

TODO

