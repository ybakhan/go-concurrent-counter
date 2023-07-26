## Go Concurrency Coding Challenge - Contiguous Counter
A basic HTTP server allows a local counter value to be incremented, set, and retrieved.  Your task is to ensure that this counter is incremented in sequential order.  This must happen regardless of how many counter increment requests are made in parallel.  To enable testing concurrent requests, an Apache Bench script has been included in this project.  So this challenge doubles as an introduction to concurrency testing.

## Using Apache Bench
Apache bench is pre-installed on most *nix, Unix, or Unix-like operating systems such as MacOS.  It's main purpose is to test HTTP servers.  In this use case, we'll use ab to test concurrent requests.  The -n parameter represents the number of requests, and the -c parameter is how many parallel requests will be made at a time.  The concurrency count can't exceed the number of requests, and it's recommended that the concurrent requests divide evenly into the total number of requests.  To normalize the challenge, use the script as-is for your final proof, but feel free to play around with the ab parameters.  Keep in mind that a large number of concurrent requests will overflow your local http server based on the configured maximum TCP connections or the file descriptor limit.  On some operating systems, this is difficult to change. If there's enough interest in this subject, a tutorial will be provided on how to change these limits in MacOS.

## Permissions for the shell script
You'll need to run 

```
chmod 755 ./ab-benchmark.sh
```

in order to execute this script with the appropriate permissions.

## Challenge Rules
You may use any means necessary to achieve an accurate counter increment, but the main() function must be left untouched. Using a new library (e.g. gin or gorilla mux) to rewrite the http server and endpoint handling is also not allowed.  The ab tool configured to make parallel requests must be the only thing used to make the increment requests to the HTTP server. Create a branch for your solution and have at it! 

## Bonus Points
Add a decrement endpoint to the HTTP server, and work out how to manage requests that are potentially incrementing and decrementing the counter at the same time.  Going for bonus points will allow you to add 

```http.HandleFunc("/decrement", dec)```

to the main() function, but all other rules still apply.

### *~ Omnia concors.*



