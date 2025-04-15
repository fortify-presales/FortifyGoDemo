# Fortify Go Demo

This is an example insecure [Go](https://go.dev/) application that can be used for the demonstration of Application Security testing tools, such as [OpenText Application Security](https://www.opentext.com/products/application-security). It is designed for testing
that security issues can be found when using different Go "routers", e.g. servemux, gorilla, gin, chi, echo.

Pre-requisities
---------------

To demonstrate this application you will need to have a Linux/UNIX or WSL instance with following installed and configured:

  - [OpenText Static Application Security Testing (Fortify)](https://www.opentext.com/products/static-application-security-testing)
  - [Go 1.23.0](https://go.dev/doc/install) or later
  - [GNU Make](https://www.gnu.org/software/make/)
  - [curl](https://curl.se/) for testing (optional)

Please note: it should possible to run and test the application on Windows but this has not been tested.

Run Application
---------------

The application is a server based API, it is not that interesting but it you want to try it out then:

```
make run
```

and in another console window try:

```
curl -X GET -v http://localhost:8080/api/v1/ping?hostname=localhost
curl -X GET -v http://localhost:8080/api/v1/download/12345
curl -X POST -H 'Content-Type: application/json' -d '{"hostname":"localhost"}' -v http://localhost:8080/api/v1/ping

 ```

Scan Application
----------------

To scan the application using a local OpenText Static Application Security Testing (Fortify) instance edit the
`Makefile` and change the following line:

```
GOROUTER := servemux
```

to the Go router that you want to use, then run a SAST scan using:

```
make sast-scan
```

The file `FortifyGoDemo.fpr` should be created. You can view the results using `auditworkbench`:

```
auditworkbench FortifyGoDemo.fpr
```

For each Go router, 3 results should be found, for example with "servemux":

```
Issue counts by category:

 "Command Injection" => 1 Issues
 "JSON Injection" => 1 Issues
 "Path Manipulation" => 1 Issues

Total for all categories => 3 Issues
```

An updated rule in `etc\example-custom-rules.xml` has been provided for "servemux" to make this work - no other rules have been
implemented as yet.

---

Kevin A. Lee (kadraman) - klee2@opentext.com