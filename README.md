# boosted-rpc

Show when you're out an about on a Boosted board on your Discord RPC

I use this with the Ride app for iOS and the IFTTT integration. Effectively, it's a basic POST to `/` with the JSON body of the following, where `status` is either `connected` or `disconnected`.

```json
{ "status": "<status>" }
```

In the IFTTT app, I set up two applets for both events. One is when I connect to my board on the Ride app (webhook in, webhook out with connect status) and the other is for disconnecting (same again, with disconnect status).

This is a really basic application and getting it up and running is extremely easy if you have golang installed, you can even keep the client ID that is already
hard coded since I've already uploaded some presence images to Discord.
