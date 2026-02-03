🐊 Gator: Blog Aggregator

A CLI tool for aggregating and following RSS feeds. Built with Go and PostgreSQL.
🛠 Prerequisites & Installation
1. Install Go

Ensure you have Go installed (version 1.21+ recommended).

    Download: go.dev/dl

    Verify: Run go version in your terminal.

2. Install PostgreSQL

You need a running PostgreSQL instance to store users, feeds, and posts.

macOS (via Homebrew):

    brew install postgresql@15
    brew services start postgresql@15

Ubuntu/Debian:

    sudo apt update
    sudo apt install postgresql postgresql-contrib
    sudo systemctl start postgresql

Database Setup: Create a database named gator:

    createdb gator

3. Install the Program

Clone the repository and install the binary to your $GOPATH/bin:
Bash

    go install github.com/hugermuger/gator@latest

💻 Command Reference

The program uses a custom command registry to map CLI arguments to Go functions. Here is what each command does:
Public Commands

These can be run by anyone without being logged in.

    login <username>: Sets the current user in the .gatorconfig.json file.

    register <username>: Creates a new user in the database and logs them in.

    reset: Caution! Wipes all users and feeds from the database (useful for development).

    users: Lists all registered users in the system.

    agg <time_between_reqs>: Starts the aggregator service. It fetches feeds periodically (e.g., 1m, 1h).

    feeds: Displays all feeds currently in the database along with the names of the users who added them.

Authenticated Commands (Middleware Protected)

If you aren't logged in, the program will exit with an error before even trying to run the command.

    addfeed <name> <url>: Adds a new RSS feed to the system and automatically follows it.

    follow <url>: Creates a "follow" relationship between the current user and an existing feed.

    following: Lists all feeds the current user is currently following.

    unfollow <url>: Removes the follow relationship for a specific feed.

    browse [limit]: Shows the latest posts from all feeds the user follows. Defaults to 2 posts if no limit is provided.
