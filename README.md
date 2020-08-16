# HSelect [![Build Status](https://travis-ci.com/gptankit/hselect.svg?branch=master)](https://travis-ci.com/gptankit/hselect) [![GoDoc](https://godoc.org/github.com/gptankit/hselect?status.svg)](https://pkg.go.dev/github.com/gptankit/hselect?tab=overview)

**hselect** is a command line tool for selecting an active service among a group of errored and non-errored services (mirrors). It uses harmonic dispatch algorithm (https://github.com/gptankit/harmonic) for service selection and also updates harmonic with service errors if dial fails for the selected service.

Here are the steps to use *hselect* - </br>

**Install**

<pre>
$ mkdir hselect && git clone https://github.com/gptankit/hselect hselect/
$ cd hselect
$ make
$ make install
</pre>

**Run**

First, add *hselect* command to PATH - 
<pre>
$ export PATH=/usr/local/hselect/bin:$PATH
</pre>

Now, run *hselect* with **-e** flag that expects a comma separated list of services - 

<pre>
$ sudo hselect -e https://goproxy.io,https://proxy.golang.org,https://gocenter.io
https://proxy.golang.org
</pre>

You can optionally run *hselect* with **-v** flag for verbose output - 

<pre>
$ sudo hselect -v -e https://goproxy.io,https://proxy.golang.org,https://gocenter.io
Parsing endpoints...Done.
Initializing cluster state...Done.
Populating errors...Done.
Selecting service...Done -> https://goproxy.io
Dialing service https://goproxy.io...Success.
Connection successful to https://goproxy.io
https://goproxy.io
</pre>

Output of *hselect* (when used without **-v**) can also be piped to *curl* or *wget* - 

<pre>
$ sudo hselect -e https://proxy.golang.org,https://goproxy.io,https://gocenter.io | xargs wget
</pre>

Feel free to play around and post feedbacks
