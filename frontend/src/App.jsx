import React from "react";
import Weather from "./components/Weather.jsx";
import "./App.css";

function App() {
    return (
        <div className="app-container">
            <header className="app-header">
                <div className="logo-container">
                    <img className="weather-logo" src="/img/weather_logo.png" alt="Weather App Logo"/>
                    <h1>Weather App</h1>
                </div>
            </header>
            <main className="app-content">
                <Weather />
            </main>
            <footer className="app-footer">
                <p>Weather data provided by OpenWeather API</p>
            </footer>
        </div>
    );
}

export default App;
