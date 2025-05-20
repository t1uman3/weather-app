![Weather Logo](frontend/public/img/imgonline-com-ua-Resize-51kpq72uaL.png)

# Weather App

This project uses Go backend (Echo framework) and React frontend.

Live demo: [Weather App](https://t1uman3.github.io/weather-app)

To start the project you need
- running Docker
- use a terminal that supports Unix commands such as Git Bash
- create `.env` file with your OpenWeather API key

## Run Locally

Clone the project

```bash
  git clone https://github.com/t1uman3/weather-app.git
```

Go to the project directory

```bash
  cd weather-app
```

Make the script executable

```bash
  chmod +x deploy.sh
```

Create `.env` file in the root directory with your OpenWeather API key:

```
WEATHER_API_KEY=your_openweather_api_key_here
PORT=8000
```

Start the application:

```bash
  ./deploy.sh
```

## Deployment

The frontend is deployed to GitHub Pages. To deploy:

1. Make sure you have the latest changes committed
2. Run the deployment command:
```bash
cd frontend
npm run deploy
```

This will build the project and deploy it to the `gh-pages` branch.

## Go Backend 

The project uses Go with the Echo framework for the backend. The structure is:

- `api/` - HTTP request handlers
  - `weather.go` - weather handlers
  - `favorite.go` - favorite cities handlers
- `model/` - data models 
  - `weather.go` - weather data model
  - `favorite.go` - favorite cities model
- `service/` - business logic
  - `weather.go` - weather service for API requests
  - `favorite.go` - favorite cities service
- `cmd/` - application entry point
  - `main.go` - main application file

## API Endpoints

- `GET /` - server health check
- `POST /weather` - get weather data for a city
- `GET /favorites` - get list of favorite cities
- `POST /favorite` - add a city to favorites
- `DELETE /favorite/:id` - remove a city from favorites
