# Go Web Server

This is a minimal web server written in Go to

1. serve as an example of a basic web API and web file server
2. provide boilerplate for very basic microservices
3. add to any web projects that require a minimal web server on any platform (such as Phaser.io game development).

Go's single, runnable files make it ideal for this usage as a minimal
web server that will run on anything without additional installation
or configuration. Just download the binary for your platform and
run it.  Here's how we made it and what it does step by step. (We
won't get into the details of how Hyper Text Transfer Protocol works
for now).

## Install

There is no installation step. Just download the binary file from the
directory matching your platform and run it.

## Usage

Really, really basic, by design:

```
serve
serve www
serve 8080
serve www 8080
serve 8080 www
```
