# Hosting data API

[![repo standards badge](https://img.shields.io/badge/dynamic/json?color=blue&style=for-the-badge&logo=github&label=MoJ%20Compliant&query=%24.data%5B%3F%28%40.name%20%3D%3D%20%22hosting-data-api%22%29%5D.status&url=https%3A%2F%2Foperations-engineering-reports.cloud-platform.service.justice.gov.uk%2Fgithub_repositories)](https://operations-engineering-reports.cloud-platform.service.justice.gov.uk/github_repositories#hosting-data-api "Link to report")

This repository contains the code to run the Hosting data API, which powers [Hosting data](https://hosting-data.apps.live.cloud-platform.service.justice.gov.uk/).

## Requirements

This repository makes heavy use of preprocessing CSV data (such as AWS Cost and Usage reports), which generate JSON files to use as static data in the API.

You will need to install:

- [csvkit](https://formulae.brew.sh/formula/csvkit#default)
- [aws-cli](https://formulae.brew.sh/formula/awscli#default)

You will also need read-only access to the [AWS root account](https://github.com/ministryofjustice/aws-root-account) to fetch data from AWS Cost and Usage reports.

## Running this locally

To run this locally, build the Docker image and run it:

```sh
$ docker build -t hosting-data-api .
$ docker run -p 8080:8080 hosting-data-api
```

Then, visit [https://localhost:8080](https://localhost:8080).
