# CF DynDNS

This is a simple script to update a Cloudflare DNS record with your current IP address.

## Usage

Create a config file in /etc/cf-dyndns/cf-dyndns.yaml with the following contents:

```
zone: example.com
sub_zone: sub.example.com
api_key: <your api key>
daemon: 360s
```

Then run the script:

```
cf-dyndns start
```

You can also install it as a systemctl service:

```
cf-dyndns install
systemctl daemon-reload
systemctl enable cf-dyndns
systemctl start cf-dyndns
```

## Api Key
You will need to create an API key in Cloudflare with the following permissions:

```
Zone.Zone.Read, Zone.DNS.Edit
```