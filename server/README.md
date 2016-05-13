## External Package Dependencies


The following are the server's external package dependencies.

These packages are needed for Google cloud storage:

go get -u golang.org/x/oauth2
go get -u google.golang.org/cloud/storage
go get -u google.golang.org/appengine/ ...

The following package is needed to generate unique filenames in the cloud based upon "version 4" unique IDs:

go get -u github.com/twinj/uuid