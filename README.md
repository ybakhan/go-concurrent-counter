## Go Concurrency Coding Challenge - Contiguous Counter

A basic HTTP REST API allows a local counter value to be set and retrieved.  Your task is to ensure that this counter is incremented in sequential order.  This must happen regardless of how many requests to increment the counter are called in parallel.  To enable testing concurrent requests, an Apache Bench script has been included in this project.  So this challenge doubles as an introduction to concurrency testing.

## Using Apache Bench

Apache bench ins pre-installed on most *nix, Unix, or Unix-like operating systems such as MacOS.  It's main purpose is to test HTTP servers.  In this use case, we'll use ab to test concurrent requests.  The -n parameter represents the number of requests, and the -c parameter is how many parallel requests will be made at a time.  The concurrency count can't exceed the number of requests, and it's recommended that the concurrent requests divide evenly into the total number of requests.  To normalize the challenge, use the script as-is for your final proof, but feel free to play around with the ab parameters.  Keep in mind that a large number of concurrent requests will overflow your http server given that it doesn't properly scale and/or you haven't implemented your concurrency solution.

## Permissions for the shell script
You'll need to run 

```
chmod 755 ./ab-benchmark.sh
```

in order to execute this script with the appropriate permissions.

## Challenge Rules

You may use any means necessary to achieve an accurate counter increment. The ab tool must be the only thing used to make the requests to the http server, and also prove that true counter concurrency has been achieved by configuring ab to make parallel requests. Create a branch for your solution and have at it! 

## Bonus Points
Add decrement endpoints to the API server, and work out how to manage requests that are potentially incrementing and decrementing the counter at the same time or in a race condition.

### *~ Omnia concors.*



