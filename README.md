# PiBakeAPIv2
PiBakeAPI version 2 beta Non-Sensical Edition

Before, we used the SSH technology to get temperatures from the Raspberry Pi to our server. It looks good on paper
but in practice, we knew there was some flaws. After the project was done, this repository was not touched. Until now!

This piece of software is a RESTful API that sits on some server and hands out and takes in requests as any other RESTful API would. This would supersede the original RESTful API buit it was never completed. That API was built in PHP and is right [here](https://github.com/pibake/PiBakeAPI). Nothing really happened to it, unfortunately.

Please note that this software is in a beta state as of now. I would like to have in a stable state some time in future!

## Installation

If you have Go installed, run the following:

`go get github.com/pibake/PiBakeAPIv2`

Then, navigate to that directory and run the following:

`go install`

Given that you have your `$GOPATH` set and it is already in your `$PATH`, it should be in your path.

## Running

Given that you have your `$GOPATH` and it is already in your `$PATH`, run the following:

`PiBakeAPIv2`

If you would like to do something while the API is up, I would advise you to run `tmux` or `screen` as it doesn't
give a prompt back.

## Troubleshooting

It shouldn't crash. If it happens to crash, [email me!](mailto:wjmiller2016@gmail.com)

## Contributing

Please read [contributing](https://github.com/pibake/PiBakeAPIv2/blob/master/CONTRIBUTING.md) file if you are interested in contributing!