+++
categories = ["servers"]
date = "2016-02-19T14:40:27Z"
description = ""
draft = false
image = "/img/bgs/cctv-bg.jpg"
tags = ["ssl", "tls", "letsencrypt", "http2", "nginx", "ubuntu", "servers"]
title = "Multiple site free ssl with http2 nginx"

+++
There still exists this myth that you require a single IP address per vhost for an ssl certificate. It was perpetuated because the encryption handshake used to happen before the resource required could be ascertained, thanks to some pretty serious security vulnerabilites SSL itself is pretty much a dead duck and the modern TLS has superceded it. One of the great things about TLS is SNI:

>Server Name Indication (SNI) is an extension to the TLS computer networking protocol[1] by which a client indicates which hostname it is attempting to connect to at the start of the handshaking process. This allows a server to present multiple certificates on the same IP address and TCP port number and hence allows multiple secure (HTTPS) websites (or any other Service over TLS) to be served off the same IP address without requiring all those sites to use the same certificate. It is the conceptual equivalent to HTTP/1.1 name-based virtual hosting, but for HTTPS. [more here](https://en.wikipedia.org/wiki/Server_Name_Indication)

## But secure websites require buying a certificate which is expensive!

Haven't you heared of [Let's Encrypt](https://letsencrypt.org/) yet?

>Let’s Encrypt is a new Certificate Authority: 
>It’s free, automated, and open. 
>In Public Beta

## So what's http2?

You may have already heard of the magic google did with SPDY well that and even more features are being rolled into HTTP2, and you can have your website future proofed and supporting these features right now.

>HTTP/2 (originally named HTTP/2.0) is the second major version of the HTTP network protocol used by the World Wide Web. It is based on SPDY. HTTP/2 was developed by the Hypertext Transfer Protocol working group (httpbis, where bis means "repeat" or "twice") of the Internet Engineering Task Force. [More Here](https://en.wikipedia.org/wiki/HTTP/2)

## Let's do this!

Okay for the purpose of this post, I will show you how to get sorted with ubuntu (15.04). It is relatively straight forward to do the same for other distrubutions following the links provided. But if you are having difficulty please write your questions in the comments below.

At the time of writing the stable version in the nginx repositories is not a high enough version for http2 (It became available in [version 1.9.5](https://www.nginx.com/blog/nginx-1-9-5/)), so you need to add the mainline repositories to your apt (though depending on when you are reading this that could have changed).

Instructions for adding the repositories can be found [here](http://nginx.org/en/linux_packages.html#mainline). But I'll include them for the sake of brevity.

{{< highlight console >}}
wget -q http://nginx.org/keys/nginx_signing.key -O- | sudo apt-key add -
{{< /highlight >}}

Then add the following to /etc/apt/sources.list.d/nginx.list:

{{< highlight console >}}
deb http://nginx.org/packages/mainline/ubuntu/ codename nginx
deb-src http://nginx.org/packages/mainline/ubuntu/ codename nginx
{{< /highlight >}}

Where codename should be your release codename, wily in my case. Then to install:

{{< highlight console >}}
apt-get update
apt-get install nginx
{{< /highlight >}}

Okay after doing its thing you should now have nginx at a new enough version for http2.

# Now for a free signed cert.

I followed [this](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-14-04) tutorial from digital ocean (great tutorials) to get the certificate up and running with the following extra bits.

* Make sure you have all the vhosts you want to host with nginx (you won't create a certificate for every domain there really is no need, you can reuse the one)

* When you are creating seperate vhost configs in nginx (doing thiis is a bit beyond the scope of this artile but if you need help comment below), add the following directive to the vhost (to point the letsencrypt requests back to the default sites):

{{< highlight console >}}
 location /.well-known/ {
        alias /var/www/html/.well-known/;
        allow all;
    }
{{< /highlight >}}

Then to enable http2 for each of the vhosts add the keyword at the top of your server block where your listen is:

{{< highlight console >}}
    listen 443 ssl http2;
{{< /highlight >}}

# Welcome to the future!
