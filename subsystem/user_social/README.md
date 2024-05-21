# Social system
Handle connection between users:

- Fetch user info from the pov of other users
- Handle connection between user (friend / friend requests)

## Friend system
Connection are stored in ``u_friend_status``, this handle both friend and friend requests.

There are 2 records per connection. The fields are as follow:

- `user_id`: This is the user id to be used when handling request by this user.
- `other_user_id`: This is the user id of the targeted connection.
- `friend_approve_at`: This store the time the friend request is approved, it is nullable but nullable will not be used to check for the friend status.
- `request_status`: This is the status of the friend request:

   - 0: This means that there's no request from `user_id`:

      - It is the default value if there's no connection between any `user_id` and `other_user_id`.
      - If it is stored in the database, then `other_user_id` sent a request to `user_id`.
   - 1: This means that there is a friend request directed from `user_id` to `other_user_id`.
   - 2: this means that the friend request is approved, the users are friends
- `is_request_pending`: This is used to check if the request is pending for the received player:

   - Always false for the sender or friend
   - true for the receiver
- `is_new`: whether the connection is "new":

   - Request will be marked as new
   - Once approved, friend will be marked as new
   - Flag is cleared on loading the friend list

## Limitations

- Each user can have at most n friends, n is decided by the database
- Each user can have at most n/m incoming/outgoing request, n/m is decided by the constant database
