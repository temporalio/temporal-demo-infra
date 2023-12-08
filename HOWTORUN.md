# temporal-demo-infra
## How To Run

### 0. Prerequisites

1. [nodejs and npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
1. [golang](https://go.dev/doc/install)
1. A Temporal Namespace
   1. [Temporal Cloud](https://temporal.io/cloud)
   1. [Temporal Dev-Server (local)](https://github.com/temporalio/cli#start-the-server)
1. A [Pulumi](https://www.pulumi.com/) account.
    1. Follow the steps [here, too](https://www.pulumi.com/docs/using-pulumi/automation-api/getting-started-automation-api/) since this demo uses the Pulumi automation API.
1. AWS Access with an AWS Profile that has privileges to create and list S3 Buckets.

### 1. Install The Things

`make deps`

All four applications need an `.env` file to find resources.
Each has either an `env.template`, `env.cloud.template`, `env.local.template` to help setup.

Which template you use is determined by whether you are using a Temporal Cloud Namespace or the Temporal CLI dev server (`localhost`).

The four paths where ensure an .env file _must_ be are:
- `(root)/.env`
- `/bff/.env`
- `/ui/.env`
- `/provisioning_aws/.env`

### 2. Run The Applications

There are four applications to run.

#### Temporal Application Worker
_Golang_

```
make domain
```

#### Temporal `AWS Provisioning` Pulumi Activity Worker
_TypeScript_ 

```
make aws
```

#### API Server (BFF)
_Golang_

```
make api
```

#### UI Server

```
make ui
```
_Visit the site at http://localhost:5173_