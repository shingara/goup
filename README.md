# GoUp

### An open source uptime tracker

GoUp is a tool to help you manage your uptime on your different service.
Answer the question : "My service is up or not"

## Status

Actually, it's still in development and can't really be use actually.

## Development

### Dependencies

All check and all URL to check are save on a MongoDB database. You need
install one to use GoUp.

There are some Go Dependencies

 * Gorilla ( http://www.gorillatoolkit.org/ ) webserver toolkit
 * Mgo ( http://labix.org/mgo ) the MongoDB driver

You can install this dependencies by :

```
go get github.com/gorilla/mux
go get labix.org/v2/mgo
```

You can build the project by :

```
go build
```

A `goup` binarie is produce and you can use it directly

## Usage

When you have the `goup` binary, you can do 2 commands :

 * `goup add http://google.com` The add command allow to save some url
   in MongoDB and can be check by goup
 * `goup` without args, launch the webserver and worker doing some
   request each second on all url save.


This behavior is not the best to have. So you need a lot new feature to
improve this behavior.


## API

GoUp is design to have 2 parts, a worker and an Web API.

### GET /urls

Return the list of all urls save on your system. The return is in JSON
like :

```
[
  {
    Id_: "51c7577e44fd997103000001",
    Name: "http://google.com/"
  }
]
```

## Features in future

The feature need to be implement in this order. This order can change in
future.

 * Configuration by variable environement [#1][]
 * Add API to add some URL [#2][]
 * Add API to get information about an URL [#3][]
 * Add API to get all check information about an URL [#4][]

[#1]: https://github.com/shingara/goup/issues/1
[#2]: https://github.com/shingara/goup/issues/2
[#3]: https://github.com/shingara/goup/issues/3
[#4]: https://github.com/shingara/goup/issues/4
