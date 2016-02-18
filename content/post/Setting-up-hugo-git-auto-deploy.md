+++
Description = ""
Tags = ["go", "golang", "hugo", "digital ocean", "ubuntu", "git"]
date = "2016-02-18T12:29:13Z"
title = "Setting up a hugo based blog/website on ubuntu with auto deploy via Git"

+++

I followed the following pages: https://www.digitalocean.com/community/tutorials/how-to-install-and-use-hugo-a-static-site-generator-on-ubuntu-14-04

I didn't want to user a package, my prefference was to compile from source: https://gohugo.io/overview/installing/ however I got the following message while trying that:
{{< highlight console >}}
go: missing Mercurial command. See https://golang.org/s/gogetcmd
package bitbucket.org/pkg/inflect: exec: "hg": executable file not found in $PATH
{{< / highlight >}}
Thanks to http://kaiq.me/2015/12/23/go/go-get-tools/ and google translate I realised Mercurial was a dependancy

so
{{< highlight console >}}
sudo apt-get install mercurial
{{< / highlight >}}
and then

{{< highlight console >}}
go get -u -v github.com/spf13/hugo
{{< / highlight >}}

got me cooking again.


In the DO tutorial it says to run 

{{< highlight console >}}
sudo hugo genautocomplete
{{< / highlight >}}

this seems to be factored out to 

{{< highlight console >}}
sudo hugo gen autocomplete
{{< / highlight >}}

but because I compiled from source hugo isn't in sudo environment path so 

{{< highlight console >}}
sudo ~/go-workspace/bin/hugo gen autocomplete
{{< / highlight >}}

[~/go-workspace is my $GOPATH so change to yours]


While cloning the themes I decided hugo-themes was a better name for the directory in my home dir but of course in later steps when I symlink the name of the link is themes.

{{< highlight console >}}
 ln -s ../hugo-themes/ themes
{{< / highlight >}}

It seems that some other switches are different from the DO tutorial (-themes= in git cloned version is -t )
