# Crypto Converter

## Description

Crypto Converter is a command-line utility to convert cryptocurrencies using the CoinMarketCap API. 
This project is structured following clean code and SOLID principles to ensure maintainability and scalability.

## Project Structure

```
crypto-converter/
│
├── cmd/
│   └── cryptoconverter/
│       └── main.go
|       └── .env
│
├── internal/
│   └── converter/
│       └── converter.go
│
├── vendor/
│
└── go.mod
```

## Requirements

- Go 1.21
- [CoinMarketCap API key](https://coinmarketcap.com/api/v1/#section/Quick-Start-Guide)

## Setup

**Clone the repository:**

```bash
git clone https://github.com/polkadot21/liteWeightCryptoConverter.git
cd liteWeightCryptoConverter
```

**Install dependencies:**

```bash
go mod tidy
```

**Create a .env file in `cmd/cryptoconverter` with the following content:**

```env
COINMARKETCAP_API_KEY=your_api_key_here
COINMARKETCAP_API_URL=https://pro-api.coinmarketcap.com/v1/tools/price-conversion
```

**Build the project:**

```bash
go build -o bin/cryptoconverter cmd/cryptoconverter/main.go
```

## Usage

```bash
./bin/cryptoconverter <amount> <from_currency> <to_currency>
```

**Example:**

```bash
./bin/cryptoconverter 123.45 USD BTC
```