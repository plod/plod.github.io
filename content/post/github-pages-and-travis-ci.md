+++
categories = ["Hosting", "github", "ci"]
date = "2017-01-09T22:38:12Z"
description = ""
draft = false
image= "img/bgs/bg-travisci.jpg"
tags = ["hosting", "hugo", "templates", "themes", "development", "travis"]
title = "Github Pages and Travis Ci"

+++

I decided that I wanted to move away from self hosted for this website, mostly for resilience (I am often tinkering with ~~this~~that server), though I also wanted to play some with Travis CI to help contributing with more FOSS and playing with different technologies. 

I followed the tutorial [here](http://rcoedo.com/post/hugo-static-site-generator/) paying special attention to the note at the bottom. In a very recent version of hugo there seems to be a problem with:

> Failed to normalize URL string. Returning in =

[This](https://discuss.gohugo.io/t/started-getting-failed-to-normalize-url-string-returning-in/5034/7) helped me with that, although it didnt quite work with my hugo/travis workflow, so I fixed it like this [.travis.yml](https://github.com/plod/plod.github.io/blob/source/.travis.yml).

I also then decided to use cloudflare to get https although I haven't quite mastered the autoredirect to https.

##### Background image a derivitive of [this](https://www.flickr.com/photos/bethscupham/7364320130/in/photolist-cdL4NJ-5tYhfL-7KsKqx-63KGHV-6a7o3M-8FBc7A-7vvR4Y-7x5yiZ-8SX3pi-9iimcr-8T18gC-8SX3tv-8ThRyW-qDbGGp-8SX3A6-6Vy9HP-8PaDRY-bBxMw9-9zewGu-a2cHLb-dYUuzD-4kPT8n-9nyXf2-gsd9Nz-j5wA7p-gscWdY-8F2xeK-avnoWQ-6cuh21-fFpwaM-oXnSij-pnzH9i-dSeqaV-dM7Wnz-a2cJyf-a2cKpj-6cuh5b-8FuZRs-oZwHyq-4HMXAY-aqyWws-pmv94-o3kSbN-o3aw4x-nKYhY1-8H1SkE-wvHibn-wtoGSS-wv2qdL-xabn6N) work by [Beth Scupham](https://www.flickr.com/photos/bethscupham/) [Licence](https://creativecommons.org/licenses/by/2.0/)