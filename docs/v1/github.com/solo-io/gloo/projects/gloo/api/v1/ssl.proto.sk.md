
---
title: "ssl.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `gloo.solo.io` 
##### Types:


- [SslConfig](#SslConfig)
- [SSLFiles](#SSLFiles)
- [UpstreamSslConfig](#UpstreamSslConfig)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/ssl.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/ssl.proto)





---
### <a name="SslConfig">SslConfig</a>

 
SslConfig contains the options necessary to configure a virtual host or listener to use TLS

```yaml
"secret_ref": .core.solo.io.ResourceRef
"ssl_files": .gloo.solo.io.SSLFiles
"sni_domains": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `secret_ref` | [.core.solo.io.ResourceRef](../../../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) | * SecretRef contains the secret ref to a gloo secret containing the following structure: { "tls.crt": <ca chain data...>, "tls.key": <private key data...> } |  |
| `ssl_files` | [.gloo.solo.io.SSLFiles](../ssl.proto.sk#SSLFiles) | SSLFiles reference paths to certificates which are local to the proxy |  |
| `sni_domains` | `[]string` | optional. the SNI domains that should be considered for TLS connections |  |




---
### <a name="SSLFiles">SSLFiles</a>

 
SSLFiles reference paths to certificates which can be read by the proxy off of its local filesystem

```yaml
"tls_cert": string
"tls_key": string
"root_ca": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `tls_cert` | `string` |  |  |
| `tls_key` | `string` |  |  |
| `root_ca` | `string` | for client cert validation. optional |  |




---
### <a name="UpstreamSslConfig">UpstreamSslConfig</a>

 
SslConfig contains the options necessary to configure a virtual host or listener to use TLS

```yaml
"secret_ref": .core.solo.io.ResourceRef
"ssl_files": .gloo.solo.io.SSLFiles
"sni": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `secret_ref` | [.core.solo.io.ResourceRef](../../../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) |  |  |
| `ssl_files` | [.gloo.solo.io.SSLFiles](../ssl.proto.sk#SSLFiles) | SSLFiles reference paths to certificates which are local to the proxy |  |
| `sni` | `string` | optional. the SNI domains that should be considered for TLS connections |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
