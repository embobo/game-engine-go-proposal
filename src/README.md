# README for the source files

## Architecture Dependencies

For web display of Go capabilities we will be using one of the following options:

* [Node-go](https://www.npmjs.com/package/node-go-require#overview)
* [GopherJS](https://github.com/gopherjs/gopherjs)

Node-go uses GopherJS. Node-go also appears to have better documentation and allows the developer to simply add a wrapper function into the Go struct.

[Node-go example](https://github.com/sagiegurari/node-go-require/blob/master/docs/api.md#nodegorequire--object)

While this is similarly available in GopherJS, the call to create the Go object is different. We have yet to determine the functional difference.

[GopherJS example](https://github.com/gopherjs/gopherjs#providing-library-functions-for-use-in-other-javascript-code)

## Components in Go

|math|physics|game object|
|----|----|----|
|PASSING|PASSSING|INC|


