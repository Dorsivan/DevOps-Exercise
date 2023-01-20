# DevOps-Exercise

This repository contains a simple go script that exposes an api at `localhost:8080`, you may run it using (you will of course need to install go first)
```
go run .
```
Alongside this script we have the associated Dockerfile, which describes the image that will be created for our small application. If you were to run this image as a **container** and check at `localhost:8080`, you will still be able to access the application running inside the container.

## The Task

### Step A

You need to write a simple script in a programming language of your choice that call our application through `localhost:8080` and print out the titles of the objects gathered.

### Step B

Your script needs to run in a different container (that it's dockerfile you will write), and access the first container (on which the provided go script exists) directly.

**Important Notes**

 - Your container needs to access the second container directly, and not question the host for the api
 - The container you create in Step B should perform all of the tasks right as it starts running, without any manual input from your side besides running it.

## Some Tips

 - I would suggest taking the time to read carefully about Docker and Containers.
 - If you do not have a unix based system on hand, you can try and install WSL. Regardless, you are not expected to work directly on Windows.
 - None of the files provided require a change for this exercise to be completed, not even the Dockerfile. However, simply running the image may not be enough. 
 - Using docker-compose isn't necessary, but is recommended, and would serve for a much better solution.
 - Just completing this exercise isn't enough, you will also be required to show an understanding of what you had written.
