# portfolio-cli

CLI app to keep track of your crypto portfolio. This app uses CoinGecko API for price info and BoltDB as database.

## How to install

Clone this repository

```bash
git clone https://github.com/AkifhanIlgaz/portfolio-cli.git
```

Install binary

```bash
go install
```

After this step, you should be able to use the app.Type this command to your terminal to make sure everything is okay.

```bash
portfolio -h
```

![help](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/portfolio-h.jpg)

If you get `Command not found error`, there is something wrong with your path. I have solved this issue by adding this line to `~/.bashrc`.

```bash
export PATH="~/go/bin:$PATH"
```

## How to use app

> Show assets and total balance in USD. You can convert to TRY by `-t` flag

```bash
portfolio
```

![portfolio](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/portfolio1.jpg)

---

> Show the current price of cryptocurrencies in USD by default. You can convert to TRY by `-t` flag

```bash
portfolio price <asset> <asset2>
# Example
portfolio price bitcoin ethereum
```

![help](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/price_bitcoin_ethereum.jpg)

---

> Show all transactions

```bash
portfolio tx <asset> --type <txType>
```

![tx](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/portfolio.jpg)

---

> Add new transaction

```bash
portfolio tx add
```

![tx-add](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/tx_add.jpg)

---

> Edit a transaction

```bash
portfolio tx edit <txID>
# If you press enter without typing anything, the value for that field doesn't change
```

![tx-edit](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/tx_edit.jpg)

---

> Delete a transaction

```bash
portfolio tx delete <txID>
```

---

> Delete an asset

```bash
portfolio delete <asset>
```
