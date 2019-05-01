# GO RESTFul API with Caching
## Description
This is a straightforward implementation for a RESTFul API in Golang using Mongodb with simple caching using goroutines.

## Installation
Clone the repo inside your `%GOPATH%/src` folder, then run the following command to download the required packages:

  ```dep ensure -update```
  
After getting the packages ready, you can create your database and collection by running the server and performing your first create request, or you can create the database and the collection and populate them with sample documents from `people.json` using `mongoimport` command:

```mongoimport --db people --collection peopleCollection --file people.json```

You can change all the configuration settings in `config.toml` file
