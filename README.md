# About
Who doesn't know it. You're in an 'eh' meeting and you are looking for a #relatable XKCD image about the currently discussed topic.
Be the coolest guy in the meeting, run this script, search for a term and receive eternal fame and about 5 seconds of humour.

# Usage
```
go build main.go
./main -s=DNS
-------------------------------------------------
1361
https://imgs.xkcd.com/comics/google_announcement.png

[Person stands behind a podium labeled Google.]
Person: The rumors are true. Google will be shutting down Plus--
Person: along with Hangouts, Photos, Voice, Docs, Drive, Maps, Gmail, Chrome, Android, and Search--
Person: to focus on our core project:
Person: the 8.8.8.8 DNS server.

{{Title text: The less popular 8.8.4.4 is slated for discontinuation.}}
```

Running this code for the first time populates the JSON file, so run it before the meeting or else you'll have to wait like 3 minutes.

# License
IDK man do whatever you want with it, it's an exercise from the fantastic [GO Programming Language](https://www.gopl.io) book which i heartly recommend to anyone trying to learn Go.
