# Portfolio CLI

<p align="center">
  <img src="https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/portfolio.jpg" alt="Portfolio CLI Banner">
</p>

<p align="center">
  <a href="https://go.dev/doc/go1.20"><img src="https://img.shields.io/badge/Go-1.20-blue.svg" alt="Go Version"></a>
  <a href="https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License"></a>
  <a href="https://github.com/AkifhanIlgaz/portfolio-cli/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome"></a>
</p>

**Portfolio CLI** is a command-line application that allows you to easily track your cryptocurrency portfolio. It fetches real-time price data using the CoinGecko API and stores your portfolio locally using a BoltDB database.

---

## ✨ Features

-   **Portfolio Management:** View your assets and total balance.
-   **Real-Time Prices:** Track current cryptocurrency prices via the CoinGecko API.
-   **Transaction History:** List all your buy and sell transactions.
-   **Easy to Use:** Manage your portfolio with simple and clear commands.
-   **Local Database:** All your data is stored securely on your computer.
-   **Currency Support:** View prices in USD and TRY.

---

## 🚀 Installation

Clone the project and install it on your computer.

```bash
git clone https://github.com/AkifhanIlgaz/portfolio-cli.git
cd portfolio-cli
go install
```

After the installation is complete, you can check if the application is working correctly with the following command:

```bash
portfolio -h
```

![Help Menu](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/portfolio-h.jpg)

> **Note:** If you get a `command not found` error, you may need to add Go's bin directory to your system's PATH. You can solve this issue by adding the following line to your `~/.bashrc` or `~/.zshrc` file:
>
> ```bash
> export PATH="~/go/bin:$PATH"
> ```

---

## 🛠️ Usage

### View Portfolio

Displays your assets and total balance in USD. You can also view it in TRY with the `-t` flag.

```bash
portfolio
```

![Portfolio](https.github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/portfolio.jpg)

### Get Real-Time Prices

Shows the current price of the specified cryptocurrencies.

```bash
portfolio price <asset1> <asset2>
# Example
portfolio price bitcoin ethereum
```

![Prices](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/price_bitcoin_ethereum.jpg)

### Manage Transactions

#### List All Transactions

```bash
portfolio tx
```

![Transactions](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/tx.jpg)

#### Add a New Transaction

```bash
portfolio tx add
```

![Add Transaction](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/tx_add.jpg)

#### Edit a Transaction

Edits the transaction with the specified ID. You can press `Enter` without typing anything for fields you don't want to change.

```bash
portfolio tx edit <transactionID>
```

![Edit Transaction](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/readme-imgs/tx_edit.jpg)

#### Delete a Transaction

```bash
portfolio tx delete <transactionID>
```

### Delete an Asset

Removes an asset and all its associated transactions from your portfolio.

```bash
portfolio delete <asset>
```

---

## 🤝 Contributing

Contributions help make the project better! Feel free to share your ideas, bug reports, or new feature requests on the [Issues](https://github.com/AkifhanIlgaz/portfolio-cli/issues) page. Pull Requests are always welcome.

1.  Fork the project.
2.  Create a new branch (`git checkout -b feature/new-feature`).
3.  Commit your changes (`git commit -m '✨ Added new feature'`).
4.  Push your branch (`git push origin feature/new-feature`).
5.  Create a Pull Request.

---

## 📄 License

This project is licensed under the [MIT License](https://github.com/AkifhanIlgaz/portfolio-cli/blob/main/LICENSE).
