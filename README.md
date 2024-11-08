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
| `gcp.l4_proxy.new_connections` | count | Number of connections that were openend. |
| `gcp.l4_proxy.closed_connections` | count | Number of connections that were terminated. |
| `gcp.l4_proxy.egress_bytes` | count | Number of bytes sent from VM to client using proxy. |
| `gcp.l4_proxy.ingress_bytes` | count | Number of bytes sent from client to VM using proxy. |

For more information about `l4_proxy` load balancer metrics, check out the official GCP [documentation](https://cloud.google.com/load-balancing/docs/metrics).

### List of Attributes

| Name | Type | Description |
|-|-|-|
| `project_id` | string | The identifier of the Google Cloud project associated with this resource, such as 'my-project'. |
| `network_name` | string | The name of the customer network in which the Load Balancer resides. |
| `region` | string | The region under which the Load Balancer is defined. |
| `load_balancing_scheme` | string | The load balancing scheme associated with the forwarding rule, one of [INTERNAL_MANAGED, EXTERNAL_MANAGED]. |
| `protocol` | string | The protocol associated with the traffic processed by the proxy, one of [TCP, UDP, SSL, UNKNOWN]. |
| `forwarding_rule_name` | string | The name of the forwarding rule. |
| `target_proxy_name` | string | The name of the target proxy. |
| `backend_target_name` | string | The name of the backend target or service. |
| `backend_target_type` | string | The type of the backend target, one of ['BACKEND_SERVICE'; 'UNKNOWN' - if the backend wasn't assigned]. |
| `backend_name` | string | The name of the backend group. Can be '' if the backend wasn't assigned. |
| `backend_type` | string | The type of the backend group, one of ['INSTANCE_GROUP'; 'NETWORK_ENDPOINT_GROUP'; 'UNKNOWN' - if the backend wasn't assigned]. |
| `backend_scope` | string | The scope of the backend group. Can be 'UNKNOWN' if the backend wasn't assigned. |
| `backend_scope_type` | string | The type of the scope of the backend group, one of ['ZONE'; 'REGION'; 'UNKNOWN' - in case the backend wasn't assigned].  |

For more information about `l4_proxy_rule` load balancer metric attributes, check out the official GCP [documentation](https://cloud.google.com/monitoring/api/resources#tag_l4_proxy_rule).

## Installation and Setup

GCP prerequisites:

- A service account in the same project of the L4 proxy.
- Configure the service account with JWT authentication and get the key. [More info](https://developers.google.com/identity/protocols/oauth2/service-account#creatinganaccount).
- Enable the *Monitoring Viewer* role for the service account. [More info](https://cloud.google.com/iam/docs/grant-role-console).

New Relic prerequisites:

- Infrastructure Agent installed and configured. [More info](https://docs.newrelic.com/docs/infrastructure/infrastructure-agent/linux-installation/package-manager-install/).

Steps:

1. Download the [pre-generated binaries](https://github.com/newrelic/nri-gcp-l4-proxy/releases) or go through the [Building](#building) section first.
2. Place the binary file (`nri-gcp-l4-proxy`) in `/var/db/newrelic-infra/custom-integrations/`.
3. Copy the [sample configuration](./gcp-l4-proxy-config.yml) to `/etc/newrelic-infra/integrations.d/`.
4. Edit the config file and set the appropiate values for fields `interval`, `timeout`, `NAME`, `FILE_PATH`, and `SINCE`.

## Usage

Once configured, the New Relic Infrastructure agent will automatically run the integration periodically and will ingest data. To visualize the generated metrics:

1. Open [NROne](https://one.newrelic.com).
2. In the left menu select "Metrics & Events", in the top selector click "Metrics", and filter by "gcp.l4_proxy".

## Building

Prerequisites:

- Go lang 1.21 or higher.
- GNU Make.

Steps:

1. Clone this repo and open a terminal in the folder where you cloned it.
2. Run `make`.

The generated binaries will be located at `cmd/bin/`. Choose the one for your operating system and architecture.

## Known limitations

- No metric deduplication. In certain cases (e.g. infra agent failure or reboot), duplicated data might be sent.

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
