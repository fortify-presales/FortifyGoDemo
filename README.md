# Fortify Go Demo

This is an example insecure [Go](https://go.dev/) application that can be used for the demonstration of Application Security testing tools, such as [OpenText Application Security](https://www.opentext.com/products/application-security). 

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

To scan the application using a local OpenText Static Application Security Testing (Fortify) install:

```
make sast-scan
```

and the file `FortifyGoDemo.fpr` should be created. You can view the results using `auditworkbench`:

```
auditworkbench FortifyGoDemo.fpr
```

---

Kevin A. Lee (kadraman) - klee2@opentext.com