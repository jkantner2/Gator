<h1 align="center">
    "Gator"
<h1>

The 'Gator' CLI tool provides the ability to aggregate RSS feeds in terminal 

# Quickstart

## Requirements

- Postgres
- Golang

## Installation

You can install 'Gator' by copying this repository and then running the go install command

a config file titled '.gatorconfig.json' will need to be created in the home directory

This file should contain a dictionary {"db_url":"postgres://username:@localhost:5432/gator?sslmode=disable", "currentUsername": ""}

Ensure you list your username in the postgres connection string

Once the config file is setup with the connection string for the postgres database and the database is running you can begin interacting with the application

A few commands to get started:

- register: <name> | Registers a user in the database
- login: <name> | Login to the user if the user is registered in the database
- addfeed: <name> <url> | add an RSS feed to the database. List the name for the feed then the url for the feed
