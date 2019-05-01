## Description
This is a straightforward implementation for a RESTFul API in Golang using Mongodb with simple caching using goroutines.

## Installation
After cloning the repo, run the following command to download the required packages:

  ```dep ensure -update```
  
After runing the server for the first time, the database as well as the collection will be created, but will still be empty, you can populate the collection with sample documents from `people.json` using `mongoimport` command:

```mongoimport --db people --collection peopleCollection --file people.json```
