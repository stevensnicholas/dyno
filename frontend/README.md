Frontend
========

This repository contains code for Dyno web application.
The Dyno web application prompts users to log in with their GitHub account to use Dyno service. When the users is logged in, they will have access to Dashboard. Dashboard page will display test results.

Prerequisites
=============

Install npm version 8.13.2 and Node.js version 16.15.1 or above. Download npm from https://docs.npmjs.com/downloading-and-installing-node-js-and-npm and Node.js from https://nodejs.org/en/download/

User Manual
===========
1. Change to frontend directory: `cd comp9447-team2/frontend`
2. Install the required dependencies: `npm install`
3. Run the React app: `npm start`

## Installation

```
npm install
```
Installs all modules listed as dependencires in package.json. \
Check you are in the frontend directory before run the command.

## Run

```
npm start
```

Runs the app in the development mode.\
The page will reload if you make edits.\
You will also see any lint errors in the console.

## Test

```
npm run test
```

Launches the test runner in the interactive watch mode.

## Build

```
npm run build
```

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.\
The build is minified and the filenames include the hashes.

## Lint

```
npm run lint
```

Runs eslint with the ruleset specified in `eslint.js`.
