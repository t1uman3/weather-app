import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Favorites from "./Favorites.jsx";
import "./Weather.css";

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://127.0.0.1:8000';

const Weather = () => {
    const [city, setCity] = useState("");
    const [weather, setWeather] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [isFavorite, setIsFavorite] = useState(false);
    const [favoriteId, setFavoriteId] = useState(null);
    const [favoritesUpdated, setFavoritesUpdated] = useState(false);

    // Check if city is in favorites only when weather data changes or after adding to favorites
    useEffect(() => {
        if (weather) {
            checkIfFavorite();
        }
    }, [weather, favoritesUpdated]);

    // Fetch all favorites to check against the current city
    const checkIfFavorite = async () => {
        if (!weather) return;
        
        try {
            const response = await axios.get(`${API_BASE_URL}/favorites`);
            const favsList = response.data;
            const favorite = favsList.find(fav => fav.city.toLowerCase() === weather.city.toLowerCase());
            
            if (favorite) {
                setIsFavorite(true);
                setFavoriteId(favorite.id);
            } else {
                setIsFavorite(false);
                setFavoriteId(null);
            }
        } catch (error) {
            console.error('Error checking favorite status:', error);
        }
    };

    const handleKeyPress = (e) => {
        if (e.key === 'Enter') {
            getWeather();
        }
    };

    const getWeather = async () => {
        if (!city.trim()) {
            setError("Please enter a city name");
            return;
        }

        setLoading(true);
        setError(null);
        setIsFavorite(false);
        setFavoriteId(null);
        
        try {
            const response = await axios.post(`${API_BASE_URL}/weather`, {
                city,
            });
            setWeather(response.data);
            setError(null);
        } catch (error) {
            console.error("Error fetching weather data:", error);
            setError(
                error.response?.data?.error || 
                "Failed to fetch weather data. Please try again."
            );
            setWeather(null);
        } finally {
            setLoading(false);
        }
    };

    const toggleFavorite = async () => {
        if (!weather) return;
        
        try {
            if (isFavorite && favoriteId) {
                // Удаляем из избранного
                await axios.delete(`${API_BASE_URL}/favorite/${favoriteId}`);
                setIsFavorite(false);
                setFavoriteId(null);
            } else {
                // Добавляем в избранное
                await axios.post(`${API_BASE_URL}/favorite`, { city: weather.city });
                setIsFavorite(true);
                // Не устанавливаем ID здесь, так как он будет получен при следующей проверке
            }
            // Обновляем список избранного
            setFavoritesUpdated(prev => !prev);
        } catch (error) {
            console.error('Error toggling favorite status:', error);
            setError('Failed to update favorites. Please try again.');
        }
    };

    return (
        <div className="weather-container">
            <section className="search-section">
                <div className="search-box">
                    <input
                        type="text"
                        value={city}
                        onChange={(e) => setCity(e.target.value)}
                        onKeyPress={handleKeyPress}
                        className="search-input"
                        placeholder="Enter city name..."
                        disabled={loading}
                    />
                    <button
                        onClick={getWeather}
                        className="search-button"
                        disabled={loading}
                    >
                        {loading ? (
                            <span className="loading-spinner"></span>
                        ) : (
                            'Get Weather'
                        )}
                    </button>
                </div>
                
                {error && <div className="error-message">{error}</div>}
            </section>

            {loading && !weather && (
                <div className="loading-container">
                    <div className="loading-spinner large"></div>
                    <p>Loading weather data...</p>
                </div>
            )}

            {weather && (
                <section className="weather-card">
                    <div className="weather-header">
                        <h2>{weather.city}</h2>
                        <button 
                            className="favorite-button" 
                            onClick={toggleFavorite}
                            title={isFavorite ? "Remove from favorites" : "Add to favorites"}
                        >
                            <img 
                                src={isFavorite ? "/img/liked.svg" : "/img/heart-outline.svg"} 
                                alt={isFavorite ? "Remove from favorites" : "Add to favorites"} 
                            />
                        </button>
                    </div>
                    
                    <div className="weather-body">
                        <div className="temperature">
                            <span className="temp-value">{Math.round(weather.temperature)}</span>
                            <span className="temp-unit">°C</span>
                        </div>
                        
                        <div className="weather-info">
                            <p className="weather-description">{weather.weather}</p>
                            
                            {weather.humidity !== undefined && (
                                <div className="weather-detail">
                                    <span className="detail-label">Humidity:</span>
                                    <span className="detail-value">{weather.humidity}%</span>
                                </div>
                            )}
                            
                            {weather.wind_speed !== undefined && (
                                <div className="weather-detail">
                                    <span className="detail-label">Wind:</span>
                                    <span className="detail-value">{weather.wind_speed} m/s</span>
                                </div>
                            )}
                        </div>
                    </div>
                </section>
            )}

            <section className="favorites-section">
                <Favorites updateTrigger={favoritesUpdated} />
            </section>
        </div>
    );
};

export default Weather;