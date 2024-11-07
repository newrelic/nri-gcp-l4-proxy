<a href="https://opensource.newrelic.com/oss-category/#new-relic-experimental"><picture><source media="(prefers-color-scheme: dark)" srcset="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/dark/Experimental.png"><source media="(prefers-color-scheme: light)" srcset="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/Experimental.png"><img alt="New Relic Open Source experimental project banner." src="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/Experimental.png"></picture></a>

# New Relic Integration for Google Cloud L4 Proxy

New Relic Infraestructure Agent integration for Google Cloud L4 Proxy metrics.

## Value

|Metrics | Events | Logs | Traces | Visualization | Automation |
|:-:|:-:|:-:|:-:|:-:|:-:|
|:white_check_mark:|:x:|:x:|:x:|:x:|:x:|

### List of Metrics

| Name | Type | Description |
|-|-|-|
| `gcp.l4_proxy.new_connections` | Count | Number of connections that were openend. |
| `gcp.l4_proxy.closed_connections` | Count | Number of connections that were terminated. |
| `gcp.l4_proxy.egress_bytes` | Count | Number of bytes sent from VM to client using proxy. |
| `gcp.l4_proxy.ingress_bytes` | Count | Number of bytes sent from client to VM using proxy. |

For more information about `l4_proxy` load balancer metrics, check out the official GCP [documentation](https://cloud.google.com/load-balancing/docs/metrics).

## Installation

- Download the pre-generated binaries (TODO) or go through the [Building](#building) section first.
- Place the binary file (`nri-gcp-l4-proxy`) in `/var/db/newrelic-infra/custom-integrations/`.
- Copy the [sample configuration](./gcp-l4-proxy-config.yml) to `/etc/newrelic-infra/integrations.d/`.

GCP requirements:

- A service account in the same project of the L4 proxy.
- Configure the service account with JWT authentication and get the key. [More info](https://developers.google.com/identity/protocols/oauth2/service-account#creatinganaccount).
- Enable the *Monitoring Viewer* role for the service account. [More info](https://cloud.google.com/iam/docs/grant-role-console).

New Relic requirements:

- Infrastructure Agent installed and configured. [More info](https://docs.newrelic.com/docs/infrastructure/infrastructure-agent/linux-installation/package-manager-install/).

## Usage

Once configured, the New Relic Infrastructure agent will automatically run the integration periodically and will ingest data. To visualize the generated metrics:

- Open [NROne](https://one.newrelic.com).
- In the left menu select "Metrics & Events", in the top selector click "Metrics", and filter by "gcp.l4_proxy".

## Building

Prerequisites:

- Go lang 1.21 or higher.
- GNU Make.

Steps:

- Clone this repo and open a terminal in the folder where you cloned it.
- Run `make`.

The generated binaries will be located at `cmd/bin/`. Choose the one for your operating system and architecture.

## Support

New Relic hosts and moderates an online forum where customers can interact with New Relic employees as well as other customers to get help and share best practices. Like all official New Relic open source projects, there's a related Community topic in the New Relic Explorers Hub.

## Contributing

We encourage your contributions to improve New Relic Integration for Google Cloud L4 Proxy! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project.
If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company,  please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

## License

New Relic Integration for Google Cloud L4 Proxy is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.
