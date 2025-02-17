---
title: Prerequisites
description: Everything you need to get started with these guides.
---

import { NumscriptBlock } from 'react-numscript-codeblock';

# Do this first: Prerequisites for using the guides

The guides in this section all involve the movement of money to and from various accounts. In the real world, money should only enter your ledger from the `@world` account when real money moves into real accounts that you are tracking. However, to run the demonstrations in the guides, we need to seed the example accounts with funds.

**This is not realistic** and you should as a general rule never arbitrarily create money without a corresponding movement of money in the real world!

First, create a file called `seed.num` with the following contents:

<NumscriptBlock script={`send [COIN 6000] (
  source = @world
  destination = {
    1/6 to @centralbank
    1/6 to {
      50% to @player:donnameagle
      remaining to @player:donnameagle:chest
    }
    1/6 to @player:tomhaverford
    1/6 to @player:leslieknope
    1/6 to {
      50/1000 to @player:andydwyer
      remaining to @player:andydwyer:chest
    }
    remaining to @player:annperkins
  }
)`}></NumscriptBlock>


:::info Wow, that's really complex! I don't understand.
Broadly speaking, this Numscript will place 1000 coins in all the accounts we will use in these guides.

But really, right now, **don't worry about understanding this complicated bit of Numscript**. This script is only necessary to set up the Dunshire ledger we'll use in the following guides, and there is no need to fully understand it yet.
:::

Then run it with

```shell
fctl ledger transactions num seed.num
```

You shouldn't need to, but you can always run this script as many times as necessary to ensure that all accounts are topped up with funds.

## VSCode extension

Do you use VSCode? Would you like to have Numscript syntax highlighting? We have a [Numscript VSCode extension](https://marketplace.visualstudio.com/items?itemName=numary.numscript) to enable Numscript syntax highlighting.

## Resetting the Dunshire ledger

If you find that you want a clean slate to start from, you can reset the Dunshire ledger in its entirety. If you followed the [Hello World tutorial](/ledger/get-started/hello-world/), then Formance Ledger is using the default SQLite storage method. Thus you can delete the Dunshire ledger by deleting the corresponding SQLite file, which is stored in `~/.numary/data/numary_dunshire.db`.
