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

```bash
# Show assets and total balance in USD. You can convert to TRY by `-t` flag
portfolio
```

---

```bash
# Show the current price of cryptocurrencies in USD by default. You can convert to TRY by `-t` flag
portfolio price <asset> <asset2>
# Example
portfolio price bitcoin ethereum
```

![help](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/price-bitcoin-ethereum.jpg)

---
