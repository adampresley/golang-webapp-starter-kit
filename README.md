Go Webapp Starter Kit
=====================

This project is a starter kit for building web applications or RESTful services
in Google Go (#golang). It has the following features baked in.

* Terminate/Interrupt signal handler
* Flags for IP/Port binding
* Middleware for managing request flow
* Context to handle data needed throughout request life-cycle

This kit uses a lot of libraries.

* [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
* [Gorilla Context](http://www.gorillatoolkit.org/pkg/context)
* [Alice](https://github.com/justinas/alice)
* [GoHttpService](https://github.com/adampresley/GoHttpService)
* [Logging](https://github.com/adampresley/logging)

Getting Started
---------------
Clone this project and copy the contents to your directory structure for Go.
Perform a search and replace for anything ```github.com/adampresley/golang-webapp-starter-kit```
and replace it with whatever your package structure is.

Now open the file **/bower.json** and modify it to match your project's criteria.
This may include the project name, authors, license, and dependencies. Save your
changes. You should also make similar changes in the file **/package.json**.

Once that is done you will need to perform a few things on the command line.
The instructions below assume you have the following tools.

* NodeJS/NPM
* Bower

```bash
$ bower install
$ npm install
$ cd www/assets/promiscuous
$ node ./build/build.js
```


License
-------
The MIT License (MIT)

Copyright (c) 2015 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
