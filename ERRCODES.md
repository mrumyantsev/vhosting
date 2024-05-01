# Error Codes

100. User with entered username or password is not exist.
A database search was performed for the username entered and the hashed password, but returned no results.

101. Cannot generate token. Error: <debug_error>.
Unable to generate token. This is purely a logical error on the server side. Most likely this will never happen.

102. You must be signed-in for changing password.
You have not been authorized to change your password (your token does not have a copy in the form of a session in the database, or you simply do not have one).

103. Cannot parse token. Error: <debug_error>.
It is impossible to decrypt the token due to an error in its encryption.

104. User with entered username is not exist.
The user with the entered username is not in the database.

105. Password cannot be empty.
The password you entered is an empty string.

106. Cannot update user password. Error: <debug_error>.
A database error will be thrown if the user's password update fails.

107. You must be signed-in for signing-out.
You simply cannot log out of the system until you are logged in.

108. User with this username doesn't exist.
The user on whose behalf you are going to act does not exist in the system.

200. Cannot bind input data. Error: <debug_error>.
Error in the content of the sent data in JSON format in the body of the request to the server. Check if all the entries you have specified match their data types.

201. Username and password cannot be empty.
You left username or password blank so you can't continue.

202. Cannot check user existence. Error: <debug_error>.
A database error will be thrown if the user existence check fails. Under the hood, it can check for the existence of a record in the database by user id or username.

203. User with entered username is exist.
You cannot create a user with the entered username because it is already in use by another user.

204. Cannot create user. Error: <debug_error>.
The DB error will be thrown when the user is created.

205. Cannot convert requested ID to type int. Error: <debug_error>.
If you enter something other than an integer as an identifier (id) in the request URL, you will get this error.

206. User with requested ID is not exist.
The user by the requested ID is not in the database, which means it is not in the system either.

207. Cannot get user. Error: <debug_error>.
A DB error will be thrown when getting the user.

208. Cannot get all users. Error: <debug_error>.
DB error will be thrown when getting all users.

209. Cannot partially update user. Error: <debug_error>.
A database error will be thrown when the user is partially modified.

210. Cannot delete user. Error: <debug_error>.
A database error will be thrown when the user is deleted.

211. You don't have enough permissions.
The server checked your cookies, if there was a token, it decrypted it and searched for the session; if there is a session, then it was checked who you are by the isSuperuser or isStaff flags, and if you did not have one of these 2 flags, the server looked among your personal permissions for one that would allow you to continue. But, as you can see from the error, I did not find it.

212. Cannot check superuser/staff permissions. Error: <debug_error>.
A database error will be thrown when looking for the isSuperuser or isStaff flag.

213. Cannot check personal permission. Error: <debug_error>.
A database error will be thrown when looking for your personal permission for a specific action on the server.

214. User with entered username is not exist.
You have entered a username that does not exist, so you cannot continue.

300. Cannot delete session. Error: <debug_error>.
A database error will be thrown when you try to delete your current session.

301. Cannot create session. Error: <debug_error>.
A database error will be thrown when trying to create a new session.

302. Cannot get session and date. Error: <debug_error>.
A database error will be thrown when looking up your last session and the date it was created.

400. Cannot do logging. Error: <debug_error>.
Such an error cannot be reflected in the logs in the database. It is present only in the console when the programmer forgot to write a constant with the type of a new structure to compare the output of a newly created module in the system.

500. Group name cannot be empty.
When creating a group, you must not leave the name field empty.

501. Cannot check group existence. Error: <debug_error>.
A database error will be thrown when checking for the group's existence.

502. Group with entered name is exist.
You cannot create a group whose name is the same as another already created group.

503. Cannot create group. Error: <debug_error>.
A database error will be thrown when the group is created.

504. Group with requested ID is not exist.
The group you tried to get via id in the request is not in the system.

505. Cannot get group. Error: <debug_error>.
A database error will be thrown when getting the group.

506. Cannot get all groups. Error: <debug_error>.
A database error will be thrown when all groups are retrieved.

507. Cannot partially update group. Error: <debug_error>.
The database error will be displayed when the group is partially updated.

508. Cannot delete group. Error: <debug_error>.
A database error will be thrown when deleting a group.

509. Group IDs cannot be empty.
The passed id list cannot be empty.

510. Cannot set user groups. Error: <debug_error>.
A database error will be thrown when adding a user to groups.

511. Cannot get user groups. Error: <debug_error>.
A database error will be thrown when taking a list of groups the user is a member of.

512. Cannot delete user groups. Error: <debug_error>.
A database error will be displayed when a user is removed from the groups specified in the list.

600. Cannot get all permissions. Error: <debug_error>.
A database error will be thrown when getting a list of all possible action permissions.

601. Permission IDs cannot be empty.
The passed list of permissions cannot be empty.

602. Cannot set user permissions. Error: <debug_error>.
A DB error will be thrown when setting the permission list for a user.

603. Cannot get user permissions. Error: <debug_error>.
A database error will be thrown when getting a list of the user's personal permissions.

604. Cannot delete user permissions. Error: <debug_error>.
A database error will be thrown when the user's permission list is deleted.

605. Cannot set group permissions. Error: <debug_error>.
A database error will be thrown when setting the permission list for groups.

606. Cannot get group permissions. Error: <debug_error>.
A database error will be thrown when getting a list of permissions for a particular group.

607. Cannot delete group permissions. Error: <debug_error>.
A database error will be thrown when the group's permission list is settled.

700. Stream cannot be empty.
The stream field must not be empty.

701. Cannot create info. Error: <debug_error>.
DB error will be thrown when creating info.

702. Cannot check info existence. Error: <debug_error>.
A database error will be thrown when checking for the existence of info.

703. Info with requested ID is not exist.
Info with the requested id does not exist in the system.

704. Cannot get info. Error: <debug_error>.
DB error will be thrown when info is received.

705. Cannot get all infos. Error: <debug_error>.
A database error will be thrown when all info is received.

706. Cannot partially update info. Error: <debug_error>.
The database error will be displayed when the info is partially changed.

707. Cannot delete info. Error: <debug_error>.
DB error will be thrown when deleting info.

800. URL and file name cannot be empty.
The url and fileName fields must be filled in.

801. Cannot create video. Error: <debug_error>.
DB error will be thrown when creating video

802. Cannot check video existence. Error: <debug_error>.
A database error will be thrown when checking if the video exists.

803. Video with requested ID is not present.
The video with the requested id does not exist in the system.

804. Cannot get video. Error: <debug_error>.
DB error will be thrown when receiving video.

805. Cannot get all videos. Error: <debug_error>.
DB error will be thrown when all videos are retrieved.

806. Cannot partially update video. Error: <debug_error>.
The database error will be displayed when the video is partially updated.

807. Cannot delete video. Error: <debug_error>.
DB error will be thrown when deleting video.
