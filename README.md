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

### Configuration by variable environement

The configuration of the Database is write directly in the code, We
really need define it by ENV variable

 * MongoDB Host
 * MongoDB Port
 * MongoDB Username
 * MongoDB Password
 * MongoDB Database

### Add API to add some URL

Add API to add new check url. The format will be :

```
POST /url, { :name => xxx }
```

If the URL is valid return the URL format in JSON with code 201
If the URL is not pass in params, return a code 400, and a JSON of this
error
```
{
  error: 'you need pass name args'
}
```
If the URL is already save on the database return a code 400 with a JSON
error

```
{
  error: 'url already check'
}
```

### Add API to get information about an URL

We need get all information about url check. An API get all of this
information. The URL must be to

```
GET /url/:id
```

The return if the URL exist is the url JSON response with HTTP status
code 200

Change the format of the URL JSON response with the new more JSON-api
related :

```
{
  urls:[{
    id: '213eq2',
    src: 'http://google.com'
    last_status: 200
    last_checked_at: "Tue, 01 Jan 2013 04:51:39 +0000"
    href: 'http:/xxx/urls/213eq2'
  }]
}
```
Note : the date is on format RFC 2822

In the same way, the API GET /urls need change is format in correspond
to the GET /urls/x response.

## Add API to get all check information about an URL

an URL can get all informations about check doing :

```
GET /urls/:id/checks
```

The return is in format like :

```
{
  urls:[{
    id: '213eq2',
    src: 'http://google.com'
    last_status: 200
    last_checked_at: "Tue, 01 Jan 2013 04:51:39 +0000"
    href: 'http:/xxx/urls/213eq2'
  }],
  links: {
    checks.url: 'http:/xxx/urls/{checks.url}'
  }
  checks: [
    {
      id: '453qa3',
      status: 200,
      checked_at: "Tue, 01 Jan 2013 04:51:39 +0000",
      links: {
        'url': '213eq2',
      }
    },
    {
      id: '453qa4',
      status: 200,
      checked_at: "Tue, 01 Jan 2013 04:49:39 +0000",
      links: {
        'url': '213eq2',
      }
    }
  ]
}
```

The return is limit to 100 checks information
