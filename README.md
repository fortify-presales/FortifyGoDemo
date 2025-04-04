# Fortify Go Demo

This is an example insecure [Go](https://www.mulesoft.com/platform/enterprise-integration) application that can beused for the demonstration of Application Security testing tools - such as [OpenText Application Security](https://www.opentext.com/products/application-security). 

Pre-requisities
---------------

To demonstrate this application you will need the following installed and configured:

  - Fortify Static Code Analyzer
  - Go 1.23.0 or later
  - GNU Make
  - curl for testing (optional)

Run Application
---------------

The application is a server based API, it is not that interesting to run but it you want to try it out execute

```
make run
```

and then in another termina:

```
curl -v http://localhost:8080/api/v1/ping?hostname=localhost
 ```

Scan Application
----------------

To scan the application using a local Fortify Static Code Analyzer instance, execute:

```
make sast-scan
```

and you should see the results displayed on the console.

---

Kevin A. Lee (kadraman) - klee2@opentext.com