# Vjetmemorize

A chart of the tools and applications used is given below.

![App Overview ](./application_overview.png)

To run this code, you will need docker and docker-compose installed on your machine. In the project root, run `docker-compose up`.


## Application Architecture

A summary of the architecture is depicted below [go_clean_arch](https://viblo.asia/p/clean-architecture-Ljy5VMYzlra).

![Application Architecture ](./goCleanArchitecture.png)


## Install golang-migrate CLI

Install golang-migrate CLI [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
### Add Cloud Key

In order to access Google Cloud for storing profile images, you will need to download a service account JSON file to your `account` application folder and call it `serviceAccount.json`. This file will be references in .env.dev.

Instructions for installing the Google Cloud Storage Client and getting this key are found at:

[https://cloud.google.com/storage/docs/reference/libraries](https://cloud.google.com/storage/docs/reference/libraries)