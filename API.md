# JSON Requests and Responses

- [JSON Requests and Responses](#json-requests-and-responses)
  - [Requests's Body that Sent to Server](#requestss-body-that-sent-to-server)
  - [Response from Server](#response-from-server)
  - [Server API](#server-api)
    - [Authorization](#authorization)
      - [POST   /auth/signin](#post---authsignin)
      - [POST   /auth/change\_password](#post---authchange_password)
      - [GET     /auth/signout](#get-----authsignout)
    - [Users](#users)
      - [POST   /user](#post---user)
      - [GET    /user/:id](#get----userid)
      - [GET    /user/all?\_limit=10\&\_page=2](#get----userall_limit10_page2)
      - [POST  /user/change\_password](#post--userchange_password)
      - [PATCH  /user/:id](#patch--userid)
      - [DELETE /user/:id](#delete-userid)
    - [User Groups](#user-groups)
      - [POST   /group](#post---group)
      - [GET    /group/:id](#get----groupid)
      - [GET    /group/all?\_limit=10\&\_page=2](#get----groupall_limit10_page2)
      - [PATCH  /group/:id](#patch--groupid)
      - [DELETE /group/:id](#delete-groupid)
      - [POST   /group/user/:id](#post---groupuserid)
      - [GET    /group/user/:id?\_limit=10\&\_page=2](#get----groupuserid_limit10_page2)
      - [DELETE /group/user/:id](#delete-groupuserid)
    - [User/Group Permissions](#usergroup-permissions)
      - [GET    /perm/all?\_limit=10\&\_page=2](#get----permall_limit10_page2)
      - [POST   /perm/user/:id](#post---permuserid)
      - [GET    /perm/user/:id?\_limit=10\&\_page=2](#get----permuserid_limit10_page2)
      - [DELETE /perm/user/:id](#delete-permuserid)
      - [POST   /perm/group/:id](#post---permgroupid)
      - [GET    /perm/group/:id?\_limit=10\&\_page=2](#get----permgroupid_limit10_page2)
      - [DELETE /perm/group/:id](#delete-permgroupid)
    - [Info](#info)
      - [POST   /info](#post---info)
      - [GET    /info/:id](#get----infoid)
      - [GET    /info/all?\_limit=10\&\_page=2](#get----infoall_limit10_page2)
      - [PATCH  /info/:id](#patch--infoid)
      - [DELETE /info/:id](#delete-infoid)
    - [Video](#video)
      - [POST   /video](#post---video)
      - [GET    /video/:id](#get----videoid)
      - [GET    /video/all?\_limit=10\&\_page=2](#get----videoall_limit10_page2)
      - [PATCH  /video/:id](#patch--videoid)
      - [DELETE /video/:id](#delete-videoid)
    - [Videostreams](#videostreams)
      - [GET    /stream/get/:id](#get----streamgetid)
      - [GET    /stream/get/all?\_limit=10\&\_page=2](#get----streamgetall_limit10_page2)


## Requests's Body that Sent to Server

``` JSON
{
    "groupIds": [0, 1]
}
```

The passed list "[0, 1]" must only be filled with integers >= 0, and must not be empty. (Italics indicate a comment, such as this one.)

## Response from Server

``` JSON
{
  "message": "No groups available."
}
```

A message about the successful (non-accidental) completion of the action.

Or such:

``` JSON
{
  "error": {
    "errCode": 203,
    "statement": "User with entered username is exist."
  }
}
```

Error message in case of failure of any action with the error code and its interpretation.

## Server API

### Authorization

POST   /auth/signin\
POST   /auth/change_password\
GET    /auth/signout

#### POST   /auth/signin

The user logged in.

``` JSON
{
    "username": "admin",
    "password": "admin"
}
```

The "username" and "password" fields are required.

``` JSON
{
  "message": "You have successfully signed-in."
}
```

You are logged in.

#### POST   /auth/change_password

User changes his password.

``` JSON
{
    "username": "admin",
    "password": "admin1"
}
```

The "username" and "password" fields are required.

``` JSON
{
  "message": "You have successfully changed password."
}
```

You have changed your password.

#### GET     /auth/signout

User logs out.

Empty request body.

``` JSON
{
  "message": "You have successfully signed-out."
}
```

You are logged out.

### Users

POST   /user\
GET    /user/:id\
GET    /user/all\
POST   /user/change_password\
PATCH  /user/:id\
DELETE /user/:id

#### POST   /user

Create user.

``` JSON
{
    "username": "admin2",
    "password": "admin2",
    "firstName": "Ivan",
    "lastName": "Ivanov"
}
```

The "username" and "password" fields are required.

``` JSON
{
  "message": "User created."
}
```

User has been created.

#### GET    /user/:id

Get user with id from request.

Empty request body.

``` JSON
{
  "id": 1,
  "username": "admin2",
  "password": "614240232425311c142b2d01aa34e9a36bde480645a57fd69e14155dacfab5a3f9257b77fdc8d8",
  "isActive": true,
  "isSuperuser": false,
  "isStaff": false,
  "firstName": "Ivan",
  "lastName": "Ivanov",
  "joiningDate": "2022-05-19T14:44:15.207391+03:00",
  "lastLogin": "2022-05-19T14:44:15.207391+03:00"
}
```

The user has been received.
	
#### GET    /user/all?_limit=10&_page=2

Get information about all users in the system.
	
You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "0": {
    "id": 0,
    "username": "admin",
    "password": "614240232425318c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
    "isActive": true,
    "isSuperuser": true,
    "isStaff": false,
    "firstName": "",
    "lastName": "",
    "joiningDate": "2022-05-11T09:32:41.115644+03:00",
    "lastLogin": "2022-05-19T14:44:11.863332+03:00"
  },
  "1": {
    "id": 1,
    "username": "admin2",
    "password": "614240232425311c142b2d01aa34e9a36bde480645a57fd69e14155dacfab5a3f9257b77fdc8d8",
    "isActive": true,
    "isSuperuser": false,
    "isStaff": false,
    "firstName": "Ivan",
    "lastName": "Ivanov",
    "joiningDate": "2022-05-19T14:44:15.207391+03:00",
    "lastLogin": "2022-05-19T14:44:15.207391+03:00"
  }
}
```

All users of the system were received.

#### POST  /user/change_password

Change password for a specific user.

``` JSON
{
    "username": "vasya",
    "password": "vasya1"
}
```

The "username" and "password" fields are required.

``` JSON
{
  "message": "User password changed."
}
```

User password has been changed.

#### PATCH  /user/:id

Partially change user c id from request.

``` JSON
{
    "username": "admin2_edit",
    "password": "admin2_edit",
    "isActive": false,
    "isSuperuser": true,
    "isStaff": false,
    "firstName": "Петр",
    "lastName": "Петров"
}
```

Only those fields are changed, the values of which are not equal to "", or are not specified at all.

``` JSON
{
  "message": "User partially updated."
}
```

The user has been partially updated.

#### DELETE /user/:id

Remove user with id from request.

Empty request body

``` JSON
{
  "message": "User partially updated."
}
```

User has been deleted.

### User Groups

POST   /group\
GET    /group/:id\
GET    /group/all\
PATCH  /group/:id\
DELETE /group/:id\
POST   /group/user/:id\
GET    /group/user/:id\
DELETE /group/user/:id

#### POST   /group

To create a group.

``` JSON
{
    "name": "New Group"
}
```

The "name" field is required.

``` JSON
{
  "message": "Group created."
}
```

The group was created.

#### GET    /group/:id

Get group with id from request.

Empty request body.

``` JSON
{
  "id": 2,
  "name": "New Group"
}
```

The group was received.

#### GET    /group/all?_limit=10&_page=2

Get information about all groups in the system.
	
You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "0": {
    "id": 0,
    "name": "Info Writers"
  }
}
{
  "1": {
    "id": 1,
    "name": "Video Watchers"
  }
}
{
  "2": {
    "id": 2,
    "name": "New Group"
  }
}
```

All groups received.

#### PATCH  /group/:id

Partially change group c id from request.

``` JSON
{
    "name": "Some Group"
}
```

Only those fields are changed, the values of which are not equal to "", or are not specified at all.

``` JSON
{
  "message": "Group partially updated."
}
```

Group partially changed.

#### DELETE /group/:id

Remove group with id from request.

Empty request body.

``` JSON
{
  "message": "Group deleted."
}
```

The group has been deleted.

#### POST   /group/user/:id

Create an association of a user with an id from a request with groups whose id list is passed in the request body.

``` JSON
{
    "groupIds": [0, 1]
}
```

The passed list "[0, 1]" must only be filled with integers >= 0, and must not be empty.

``` JSON
{
  "message": "User groups set."
}
```

User has been added to groups.

#### GET    /group/user/:id?_limit=10&_page=2

Get a list of id groups to which the user with the id from the request is bound.
	
You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "groupIds": [
    0,
    1
  ]
}
```

The list of id groups in which this user exists was obtained.
	
#### DELETE /group/user/:id

Remove the association of a user with an id from a request with a list of group ids passed in the request body.

``` JSON
{
    "groupIds": [0, 1]
}
```

The passed list "[0, 1]" must only be filled with integers >= 0, and must not be empty.

``` JSON
{
  "message": "User groups deleted."
}
```

User associations with the listed groups have been removed.

### User/Group Permissions

GET    /perm/all\
POST   /perm/user/:id\
GET    /perm/user/:id\
DELETE /perm/user/:id\
POST   /perm/group/:id\
GET    /perm/group/:id\
DELETE /perm/group/:id

#### GET    /perm/all?_limit=10&_page=2

Get information about all permissions in the system.

You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "0": {
    "id": 0,
    "name": "Can create a User",
    "codeName": "post_user"
  },
  "1": {
    "id": 1,
    "name": "Can get a User",
    "codeName": "get_user"
  },
  "2": {
    "id": 2,
    "name": "Can get all of the Users",
    "codeName": "get_all_users"
  },
  "3": {
    "id": 3,
    "name": "Can partially update a User",
    "codeName": "patch_user"
  },
  "4": {
    "id": 4,
    "name": "Can delete a User",
    "codeName": "delete_user"
  }
}
```

All permissions have been obtained for viewing.

#### POST   /perm/user/:id

Create a user id association from a request with a list of permission ids.

``` JSON
{
    "permIds": [5, 6, 7]
}
```

The passed list "[5, 6, 7]" must only be filled with integers >= 0, and must not be empty.

``` JSON
{
  "message": "User permissions set."
}
```

Permissions have been given to the user.

#### GET    /perm/user/:id?_limit=10&_page=2

Get list of all permissions associated with user with id from request.
	
Empty request body.

``` JSON
{
  "permIds": [5, 6, 7]
}
```

A list of data for the user with id from the permissions request has been received. If the user does not have personal permissions, "null" is sent instead of the list.

#### DELETE /perm/user/:id

Remove permission id list associated with user with id from request.

``` JSON
{
    "permIds": [5, 6, 7]
}
```

The passed list "[5, 6, 7]" must only be filled with integers >= 0, and must not be empty.

``` JSON
{
  "message": "User permissions deleted."
}
```

User permissions from the list have been removed from the user with id from the request.

#### POST   /perm/group/:id

Create a group association with id from a request with a list of permission ids.

``` JSON
{
    "permIds": [4, 2, 1, 3]
}
```

The passed list "[4, 2, 1, 3]" should only be filled with integers >= 0, and should not be empty.

``` JSON
{
  "message": "Group permissions set."
}
```

Permissions have been given to the group.

#### GET    /perm/group/:id?_limit=10&_page=2

Get a list of all permissions associated with a group with id from a request.

Empty request body.

``` JSON
{
  "permIds": [4, 2, 1, 3]
}
```

A list of data for the group with id from the permissions request has been received. In case the group does not have permissions, "null" is sent instead of the list.

#### DELETE /perm/group/:id

Remove permission id list associated with group with id from request.

``` JSON
{
    "permIds": [4, 2, 1, 3]
}
```

The passed list "[4, 2, 1, 3]" should only be filled with integers >= 0, and should not be empty.

``` JSON
{
  "message": "Group permissions deleted."
}
```

Group permissions from the list have been removed from the group with id from the request.

### Info

POST   /info\
GET    /info/:id\
GET    /info/all\
PATCH  /info/:id\
DELETE /info/:id

#### POST   /info

Create info.

``` JSON
{
    "stream": "dfmgldmnldndnfgdkfgn",
    "startPeriod": "2022-05-18 16:13:51.973640+03",
    "stopPeriod": "2022-05-18 16:13:51.973640+03",
    "timeLife": "2022-05-18 16:13:51.973640+03"
}
```

The "stream" field is required.

``` JSON
{
  "message": "Info created."
}
```

Info was created.

#### GET    /info/:id

Get info.

Empty request body.

``` JSON
{
  "id": 1,
  "createDate": "2022-05-19T16:37:43.258895+03:00",
  "stream": "dfmgldmnldndnfgdkfgn",
  "startPeriod": "2022-05-18T16:13:51.97364+03:00",
  "stopPeriod": "2022-05-18T16:13:51.97364+03:00",
  "timeLife": "2022-05-18T16:13:51.97364+03:00",
  "userId": 0
}
```

Info has been received.

#### GET    /info/all?_limit=10&_page=2

Get all records info in the system.

You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "1": {
    "id": 1,
    "createDate": "2022-05-19T16:37:43.258895+03:00",
    "stream": "dfmgldmnldndnfgdkfgn",
    "startPeriod": "2022-05-18T16:13:51.97364+03:00",
    "stopPeriod": "2022-05-18T16:13:51.97364+03:00",
    "timeLife": "2022-05-18T16:13:51.97364+03:00",
    "userId": 0
  },
  "2": {
    "id": 2,
    "creationDate": "2022-05-19T16:41:02.697301+03:00",
    "stream": "wernwernjwnerjknwerkjnw",
    "startPeriod": "2022-05-18T16:13:51.97364+03:00",
    "stopPeriod": "2022-05-18T16:13:51.97364+03:00",
    "timeLife": "2022-05-18T16:13:51.97364+03:00",
    "userId": 0
  }
}
```

All info has been received.

#### PATCH  /info/:id

Partially change info.

``` JSON
{
    "stream": "qwerty"
    "createDate": "2022-04-19T16:41:02.697301+03:00",
    "startPeriod": "2022-01-18T16:13:51.97364+03:00",
    "stopPeriod": "2022-02-18T16:13:51.97364+03:00",
    "timeLife": "2022-03-18T16:13:51.97364+03:00",
    "userId": 0
}
```

Only those fields are changed, the values of which are not equal to "", or are not specified at all. "userId" does not change.

``` JSON
{
  "message": "Info partially updated."
}
```

Info has been partially updated.

#### DELETE /info/:id

Delete info.

Empty request body.

``` JSON
{
  "message": "Info deleted."
}
```

Info has been removed.

### Video

POST   /video\
GET    /video/:id\
GET    /video/all\
PATCH  /video/:id\
DELETE /video/:id

#### POST   /video

Create video.

``` JSON
{
    "url": "https://somewebsite.com/video/knsdkjnfsflsnkn",
    "file": "Video1.webm",
    "infoId": 1
}
```

The fields "url" and "fileName" are required.

``` JSON
{
  "message": "Video created."
}
```

Video was created.

#### GET    /video/:id

Get video.

Empty request body.

``` JSON
{
  "id": 1,
  "url": "https://somewebsite.com/video/knsdkjnfsflsnkn",
  "file": "Video1.webm",
  "createDate": "2022-05-19T16:49:52.504174+03:00",
  "infoId": 1,
  "userId": 0
}
```

Video has been received.

#### GET    /video/all?_limit=10&_page=2

Get all video recordings in the system.
	
You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "1": {
    "id": 1,
    "url": "https://somewebsite.com/video/knsdkjnfsflsnkn",
    "file": "Video1.webm",
    "createDate": "2022-05-19T16:49:52.504174+03:00"
    "infoId": 1,
    "userId": 0
  },
  "2": {
    "id": 2,
    "url": "https://somewebsite.com/video/meqjqnrkjqnr",
    "file": "asd.webm",
    "createDate": "2022-05-19T16:54:17.666412+03:00"
    "infoId": 2,
    "userId": 0
  }
}
```

All videos in the system have been received.

#### PATCH  /video/:id

Partially change the video.

``` JSON
{
    "url": "https://somewebsite.com/video/qwerty",
    "file": "qwerty.webm",
    "createDate": "2022-01-19T16:54:17.666412+03:00"
    "infoId": 1,
    "userId": 0
}
```

Only those fields are changed, the values of which are not equal to "", or are not specified at all. "userId" does not change.

``` JSON
{
  "message": "Video partially updated."
}
```

The video has been partially updated.

#### DELETE /video/:id

Delete video.

Empty request body.

``` JSON
{
  "message": "Video deleted."
}
```

The video was removed.

### Videostreams

GET    /stream/get/:id\
GET    /stream/get/all

#### GET    /stream/get/:id

Get stream.

Empty request body.

``` JSON
{
  "id": 1,
  "stream": "some stream",
  "dateTime": "2022-01-19T16:54:17.666412+03:00",
  "statePublic": 1,
  "statusPublic": 1,
  "statusRecord": 1,
  "pathStream": "ksdfkmsfmskfmklsmf"
}
```

Video has been received.

#### GET    /stream/get/all?_limit=10&_page=2

Get all stream records in the system.

You can specify the number of items to display through the parameter after the request: _limit=<value> and display the same amount (or the remainder) on another page, via the parameter: _page=<value>.
If the page parameter is missing, the “first” page is always displayed, if no parameter is specified, the total number of elements will be displayed (default is 100).

Empty request body.

``` JSON
{
  "1": {
    "id": 1,
    "stream": "some stream 1",
    "dateTime": "2022-01-19T16:54:17.666412+03:00",
    "statePublic": 1,
    "statusPublic": 1,
    "statusRecord": 1,
    "pathStream": "ksdfkmsfmskfmklsmf"
  },
  "2": {
    "id": 2,
    "stream": "some stream 2",
    "dateTime": "2022-01-20T16:54:17.666412+03:00",
    "statePublic": 1,
    "statusPublic": 1,
    "statusRecord": 1,
    "pathStream": "asdadandkjnqkwnekqne"
  }
}
```

All streams in the system have been received.
