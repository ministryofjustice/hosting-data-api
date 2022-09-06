# Hosting data API

[![repo standards badge](https://img.shields.io/badge/dynamic/json?color=blue&style=for-the-badge&logo=github&label=MoJ%20Compliant&query=%24.data%5B%3F%28%40.name%20%3D%3D%20%22hosting-data-api%22%29%5D.status&url=https%3A%2F%2Foperations-engineering-reports.cloud-platform.service.justice.gov.uk%2Fgithub_repositories)](https://operations-engineering-reports.cloud-platform.service.justice.gov.uk/github_repositories#hosting-data-api "Link to report")

This repository contains the code to run the Hosting data API, which powers [Hosting data](https://hosting-data.apps.live.cloud-platform.service.justice.gov.uk/).

## Running this locally

To run this locally, build the Docker image and run it:

```sh
$ docker build -t hosting-data-api .
$ docker run -p 8080:8080 hosting-data-api
```

Then, visit [https://localhost:8080](https://localhost:8080).
