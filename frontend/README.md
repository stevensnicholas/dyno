Frontend
========
This repository contains code for the Dyno web application. The web application allows users to view their fuzzing test results from dashboard. It focuses on providing visualisation tool that helps users to identify potential vulnerabilities residing in their APIs. The application consists of two modules, viewing test result and authentication. 

## Dashboard
A user can view their test history. It will display the test ID, OpenAPI file name which the user used for the test and the date of the test. When a user clicks any item from the test history section, the dashboard page will display the detailed result of the test. The dashboard page will also provide other user-friendly feature such as a link to the user's GitHub repository that is connected to Dyno or displaying the most frequent errors.

## Login
The dashboard page is allowed only for authenticated users. In order to use the dashboard page, a user is required to login with their GitHub account. Once the user is logged in, the navigation in the web application will provide an access to the dashboard page.

Prerequisites
==============
Install npm version 8.13.2 and Node.js version 16.15.1 or above. Download npm from https://docs.npmjs.com/downloading-and-installing-node-js-and-npm and Node.js from https://nodejs.org/en/download/

User Manual
===========
1. Change to "frontend" directory: `cd frontend`
2. Install the required dependencies: `npm ci`
3. Run the React app: `npm start`

Other npm commands
===================
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

## Fix

```
npm run fix
```

Fix eslint issues