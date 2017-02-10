
+++
title = "ElixiRas"
slug = "ElixiRas"
draft = false
date = "2016-04-14T19:11:21.000Z"

+++
I recently watch [this great talk](https://www.youtube.com/watch?v=kpzQrFC55q4) about embedded Elixir and got super excited about the [Nerves Project](http://nerves-project.org/) which is a suite of tools for same.
I did some hacking earlier this week using Nerves and Bake on my main machine to get things running. And I got `blinky`, the "Hello World" of embedded, to build and run using Bake. ==Mind Blown!!==

I decided that I wanted to actually use Raspbian for my prototyping, on advice from folks in the #nerves room on [Elixir-Lang Slack](https://elixir-lang.slack.com). Getting `Erlang` and `Elixir` built on the PI has been quite a challenge, with many false starts and frustrations. But thanks, again, to folks on Slack, I got it done. I decided to record the process for posterity, so the following is a tutorial of all the relevant steps.

---
> Note: The intention here is to produce an img specifically for doing embedded elixir programming on the pi that you can break out any time you want.

1. Start with a fresh Raspbian install and get internet and others basics like SSH, vim, etc. configured.
2. Follow steps here: http://elinux.org/Erlang
  Building from source seems to be the most reliable way to get a good install. However, it's __mad slow__ so be warned.
3. Install a precompiled Elixir. I used these [instructions](https://github.com/mmmries/pi-alarm-clock#install-precompiled-elixir).
4. After you've verified that everything is working (my acid test was doing a `mix compile` for `elixir_ale`) __back up__ your SD card using `dd if=/dev/rdiskx of=/path/to/image bs=1m`. You'll thank yourself later.

**And that's it!** From there you have a working RasPI image you can use for all your elixir embedded development. Enjoy and get in contact if you have any questions.


