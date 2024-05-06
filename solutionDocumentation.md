# Kii interview Challenge

# Challenge Description

Step 1: run the kii validator.  Instructions are all on the GitHub repo for Kii

Step 2:  Do the whole exercise again EXCEPT do not replace the genesis file and do not add a persistent peer (essentially, do not connect your validator to kii testnet).  Just have your validator run locally on your computer/server - effectively running your own blockchain consisting of one validator on your local machine.

Step 3: once you have your validator running on your own machine (not connected to kii testnet), add a "blog" module to the code:

https://docs.ignite.com/v0.27/guide/blog/intro

- this is an ignite cli tutorial on how to add a simple module to the existing blockchain

Step 4: demo that the module works.  If you aren’t successful, show what you have done and what kind of errors you are running into.  We will ask you how you interpreted the errors and what kind of debugging you have done with it.

Some helpful resources that you can optionally read (will give you context to the tools you are using for the test #1)

cosmos sdk -

https://docs.cosmos.network/main/learn/intro/overview

ignite cli -

https://docs.cosmos.network/main/learn/intro/overview

Tendermint:

https://docs.tendermint.com/v0.34/introduction/what-is-tendermint.html

# Solving Process

## 1. Copy the Kii Repo and set up

Clone the following repo and install the dependencies (check the Ignite version, must be **v0.27.0**), then follow the GitHub documentation about the Kii validator running locally.

[Configuration file | Ignite Cli Docs](https://docs.ignite.com/references/config)

## 2. Follow the Ignite tutorial about creating a Module

[In-depth tutorial | Ignite Cli Docs](https://docs.ignite.com/v0.27/guide/blog/intro)

## 3. Test the Local Chain an CRUD operations with the Blog module

Run the following commands for testing the Blog module code and manipulate the local Chain, using the Ignite CLI.

### Start local Chain

In the project’s folder runthe following command to start the chain

```docker
ignite chain serve
```

## Create a post in the local chain

```docker
kiichaind tx kiichain create-post hello world --from andresWallet
```

## Check the created post

```yaml
kiichaind q kiichain show-post 0
```

### List all posts in pagination

```yaml
kiichaind q kiichain list-post
```

### Delete the post

```yaml
kiichaind tx kiichain delete-post 0 --from andresWallet
```

# What was the error?

The error was the currency in the start operation when the `ignite chain init` command was launched.

```yaml
version: 1
accounts:
  - name: andresWallet
    coins: [54000000000000tkii] // Notice tkii denomination is wrong
genesis:
  chain_id: kiiventador
  app_state:
    staking:
      params:
        bond_denom: ukii // Notice the blockchain's denomination is ukii
validators:
  - name: andresWallet
    bonded: 100000000000tkii // Notice tkii denomination is wrong
    app:
      pruning: "default"
    config:
      moniker: "andresValidator"
    client:
      output: "json"
```

## How did I find the error?

When I ran the command `ignite chain init` I received the error: “failed to execute message; message index: 0: invalid coin denomination: got **tkii**, expected **ukii**: invalid request”. And checking where the tkii and the **ukii** are, I noticed the error because in some places the tkii is used, whereas ukii is just in the config file. That means an error in the genesis needs to be fixed.

# Troubleshooting

## 1. Error: unknown command "genesis" for "kiichaind”

If you are trying to run the following command: `ignite chain init` and you receive the error “ unknown command "genesis" for "kiichaind” “, Verify you have the property ignite version, as Kii’s Github list, you need the specific ignite version **v0.27.1.**

Therefore, for checking your current version run the command: `ignite version`, if your version isn’t 0.27.1 remove the current version and install the desired one using the command `sudo rm $(which ignite)` and delete the .ignite folder in your home folder.

## 2. Cannot run /home/ubuntu/go/bin/kiichaind

If you try to run the command `ignite chain serve`, the following error is thrown ”panic: failed to execute DeliverTx. failed to execute message; message index: 0: invalid coin denomination: got tkii, expected ukii: invalid request. “ Means a transaction was trying to be executed with the wrong denomination token from the config.yml file. Check the **bond_denom** param in config.yml file and compare it with the transaction denomination.

## 3. Error: post failed: Post "[http://localhost:26657](http://localhost:26657/)": dial tcp 127.0.0.1:26657: connect: connection refused

This error means the local chain isn’t running, you should verify the desired chain is running. To start serving the local chain run the command `ignite chain serve`

### 4. Error: <addressName>.info: key not found

This error means you’re trying to make a transaction from a wallet that doesn’t exist in the blockchain. Verify the wallet account address.
