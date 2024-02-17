# Anime List Web Application

This is a simple web application for keeping track of your favorite anime titles.

## Usage

To run this web application locally, make sure you have Node.js and npm installed on your machine. Then, follow these steps:

1. Clone this repository to your local machine.
2. Navigate to the project directory in your terminal.
3. Run `npm install` to install the necessary dependencies.
4. After the installation is complete, you can use the following npm scripts:
   - `cd ./vue`: Go to `vue` directory.
   - `npm run dev`: Start the development server.
   - `npm run build`: Build the production-ready files.
   - `npm run lint`: Lint the codebase for errors and warnings.
   - `npm run preview`: Preview the production build locally.
   - `npm run format`: Format the codebase to ensure consistency.

## Building Backend Application (Go)

If you've made changes to the backend application written in Go and you need to rebuild it, follow these steps:

1. Navigate to the `go` directory in your terminal.
2. Run the following command to build the backend application:
`go build -o ./api.exe api.go`

This will compile your Go code and produce an executable file named `api.exe`.

## Features

- Add anime titles to your list.
- Remove anime titles from your list.
- View a list of all anime titles added.
