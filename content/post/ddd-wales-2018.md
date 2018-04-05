+++
tags = [
    "ci",
    "cd",
    "kotlin",
    "event sourcing",
    "branching",
    "k8s",
    "kubernetes",
    "docker",
    "azure",
    "cloud",
    "techhub",
    "swansea",
    "conferences",
    "development",
]
categories = [
    "Development",
    "Conferences",
]
image = "/img/bgs/balmer-developer.jpg" #optional image - "/img/about-bg.jpg" is the default
description = "Developer Developer Developer Wales"
draft = true
date = "2018-03-30T21:24:47Z"
comments = true
title = "DDD Wales 2018"
+++

![Developer Developer Developer Wales 2018]
(https://www.dddwales.com/Media/Default/Logos/DDD-Wales-Dragon-Logo.png)

Last Saturday I attended [Developer Developer Developer Wales](https://www.dddwales.com/) with Nigel, Dai, Pete and we bumped into James Alfie. 
<blockquote class="embedly-card"><h4><a href="https://www.linkedin.com/feed/update/urn:li:activity:6383242695481135104">At Developer Developer Developer Wales with Nigel Craven, David Duncan, Peter Walsh  and nice to see James Alfei #DDDWales</a></h4><img src="/img/post-img/dddwales.jpg"></blockquote>
<script async src="//cdn.embedly.com/widgets/platform.js" charset="UTF-8"></script>
# So what is DDD Wales?

>The UK's premier FREE Developer conference is coming to Wales! Since its inception in 2005, DDD has spread across the UK and throughout other parts of the world, as far afield as Melbourne and Perth. This time though it is heading to Swansea!
>
>DDD is a free one day technical event for developers. It’s been just over 10 years since DDD was first formed and as a conference it has generated many spin-offs, as well as nurturing talented speakers, who have gone on to become Microsoft Most Valuable Professionals, Microsoft FTEs and to present at National and International Conferences.
>
>DDD was always envisioned as a day of learning, discussion, contribution and involvement in the community, from across the UK. The goal of DDD is to provide free technical education with the added benefit of the networking possibilities with peers and the development of relationships across the .NET Industry!

Based on the title of the conference, and the Microsoft connection I am inclined to believed it has something to do with Steve Balmer: {{< youtube Vhh_GeBPOhs >}}

I have to say I have been to paid conferences that were not as good as this event was, it was amazing, I'll go back next year and maybe speak at one one day. I made some brief notes about the talks that I saw for posterity.

## An Introduction to Kotlin - [Kevin Jones](https://twitter.com/kevinrjones)

This was a great introduction for a language that has piqued my interest ever since I heard about it on a [Software Engineering Daily episode](https://softwareengineeringdaily.com/2017/01/26/kotlin-with-hadi-hariri/) which at the time I thought it was a little odd but after Google announcing first class citizingship at Google IO I started to pay a lot more attention. It was a shame that the slot was so short as I was really enjoying the talk. It really does look like Java without the bit I really dislike about Java, all the boilerplate. Kevin mentioned about kotlin-native which sounded worth a look if I get into it a bit. I was most interested in its functional features and wondered about using it as a replacement to scala for some Spark jobs thought it was worth at least a play. I've never heard of the Elvis Operator before which I got me a chuckle and during the presentation I made a note to read "Effective Java" which I have a copy of somewhere. Kevin has some courses on pluralsight which I hace queued up to my pluralsight playlist.

## Kuburnetes for .NET developers - [Shahid Iqbal](https://twitter.com/shahiddev)

I've been sorta semi following K8s developments from the sideline but only really peripharily, but I have so much new tech front and centre I've not really really had time to really look. I was concerned that this talk was going to be too heavy on the .NET and a bit light on the Kubernetes, I couldnt have been further from the truth. This was a great talk and helped me understand a lot about the tech and how to potentially roll things out. I've played a little with docker swarm previously and this seemed as easy as that to start messing with, which surprised me as I'd thought it was a little bit inaccesible to start with. Minicube seemed a nice way to mess about with things. Being a big AWS user it was nice to see inside the azure console, for the first time and to see what a completely managed k8s looks like, maybe what EKS will be like.

![more resources](/img/post-img/2018/2018-04-05%2017.12.08.jpg)

## releaseflow - [Chris Cundill](https://twitter.com/chriscundill)

This was probably my favourite talk of the conference, probably because it was/is very topical to my team at the moment, we came very close to implementing Git flow.

![Woke:S]
(/img/post-img/2018/2018-04-05%2017.19.41.jpg)

Ironically the branching method we currently resemble the most is what Chris is advocating with [releaseflow](http://releaseflow.org/) with a few caveats but it was nice to see that mostly our thinking has been on point. I snapped a picture of some books to add to the ever extending reading list:

![Books]
(/img/post-img/2018/2018-04-05%2017.11.44.jpg)
