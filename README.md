httpruncmd
==========

A dead simple webserver that will run a shell command if the good URL is called.

We use it to be called from github after a commit to start the build and deploy process.


./httpruncmd -cmd=/bin/date -path=toto -port=8000 then url http://localhost:8000/toto/
