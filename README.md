# boosted-rpc

Show when you're out an about on a Boosted board on your Discord RPC

I use this with the Ride app for iOS and the IFTTT integration. Effectively, it's a basic GET to `/connected` and `/disconnected`
In the IFTTT app, I set up two applets for both events. The first is to receive the connected event and then send it back out to this server, the second is the same for the disconnected event.

This is a really basic application and getting it up and running is extremely easy if you have golang installed, you can even keep the client ID that is already hard coded since I've already uploaded some presence images to Discord.
