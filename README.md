# Meson Operator

Eventual goal is to build the Meson operator in Go which abstracts an internal platform for hosting opnioniated applications.


### PROJECT file

One important file of note is the PROJECT file. All the next commands we run use the information in this file.

### Manager

Quick look of the `main.go` file shows the code that initializes and runs the Manager. The manager is responsible for registering the scheme for all custom resource API definitions along with running controllers and webhooks.


### Create an API & Controller

```bash
# Use the cli to boostrap the api and controller
operator-sdk create api --group=network.meson --version=v1alpha1 --kind=Meson
```

Run below to update the generated code whenever the *_types.go files are modified.

Under the hood below command runs the controller-gen utility to implement `runtime.Object` interface automatically for our type.


```bash
make generate
```

Below would generate the CRDs automatically

```bash
make manifests
```