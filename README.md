# simple-go-server

Simple go http server for use on GCP Cloud Run

GCP Cloud Run requires your service to take a environment variable named `PORT` and serve traffic on it.

This server by default serves on 80, but will accept the `PORT` environment variable and use that to listen instead if it is present.
