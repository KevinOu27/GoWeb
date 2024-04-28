# Go-based Web Crawler

This repository contains a Go-based web crawler that utilizes the Colly framework to scrape information related to intelligent systems and robotics from specified Wikipedia pages. The aim is to demonstrate the effectiveness of Go in building efficient web crawlers comparable to Python/Scrapy solutions.

## Table of Contents

- [Setup and Installation](#setup-and-installation)
- [Running the Crawler](#running-the-crawler)
- [Libraries Used](#libraries-used)

## Setup and Installation

### Prerequisites

Ensure you have Go installed on your machine. If not, download and install it from the [official Go website](https://golang.org/dl/).

### Setting Up Your Project

To set up your project locally, follow these steps:

1. Clone this repository:
    ```bash
    git clone https://github.com/KevinOu27/GoWeb.git
    cd goweb
    ```

2. Initialize the module (if cloning did not bring the `go.mod` file):
    ```bash
    go mod init goweb
    ```

3. Install Colly and other necessary packages:
    ```bash
    go get -u github.com/gocolly/colly/v2
    go get -u github.com/gocolly/colly/v2/extensions
    ```

## Running the Crawler

To run the crawler, use the following command:

```bash
go run crawler.go

## Libraries Used
Colly: Highly efficient web scraping framework for Go, used for making HTTP requests and parsing HTML.
