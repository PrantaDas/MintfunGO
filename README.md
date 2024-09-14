## MintfunGo.

Mintfun Bot is a web3 nft minting bot that [https://mint.fun](https://mint.fun) mints the NFT from the reference website depending on multiple conditions. It basically trgets the Tokens that are on [Base](https://docs.base.org/) chain and the Tokens that are not reported by the community. The bot search for the lowest price and minimum minting quantity of the collection and after satisfying the conditions the complete the transaction and save the collection info with the transaction ID to prevent the duplication of minting process for the same collection.

### Table of Contents
- [Features](#features)
- [How it works](#how-it-works)
- [License](#license)

## Features:
* The bot starts only 2 days in a week through a scheduler.
* Target for the whietelisted NFTs.
* Purchase the NFTs that are free to mint or have the lowest cost with different quantity.
* Save the transaction into the DB for avoiding duplicate transactions.

### How it works
* The bot gathers collection information from a api `https://mint.fun/api/mintfun/feed/free?range=30m&chain=8453` provided by the `mintfun` with 60 seconds intervals.

* Then it filter outs the Blacklisted tokens from the collection and returns the rest of the data.
* Then we iterate over the collection and send it for further processing where we gather necessary information for completig the transactiona and the after successful processing we save the collection info in `DB`.

### License
[MIT](https://github.com/PrantaDas/MintfunGO/blob/main/LICENSE)