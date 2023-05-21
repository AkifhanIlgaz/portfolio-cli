## TODO

Create price command to to get price from CLI
Create getPriceBatch(?) function for multiple currencies
Add flags to commands

# DATABASE ACTIONS

## Transaction Actions

✔️ Create a transaction
✔️ Delete a transaction
✔️ Edit a transaction
✔️ Get a transaction
✔️ Get all transactions
✔️ Get all transactions of an asset
✔️ Get all transactions of same type

## Asset Actions

✔️ Get all assets
✔️ Get all crypto assets
✔️ Get all fiat assets
✔️ Delete an asset
✔️ Update an asset

# COBRA ACTIONS

balance => Return current balance. Accept currency type by flag. Default USD
Example => portfolio-cli balance -c try

## Tx Actions

✔️ tx <asset1>, <asset2> => Return all transactions. If asset names are given by additional arguments just return transactions for given assets

✔️ tx get <id>, <id2> => Return transactions for given IDs

✔️ tx edit <id> => Edit #id in an interactive way

✔️ tx delete <id>, <id2> => Delete transactions for IDs

✔️ tx add => Create new transaction in interactive way

## Asset Actions

✔️ delete <asset>, <asset1> => Delete given assets

# PRICE

## Crypto

✔️ Get price via Coingecko API

## TRY-USD Converter

✔️ Convert TRY-USD
