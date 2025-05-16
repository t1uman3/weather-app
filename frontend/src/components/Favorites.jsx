import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';
import './Favorites.css';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://127.0.0.1:8000';

const Favorites = ({ updateTrigger }) => {
    const [favorites, setFavorites] = useState([]);
    const [weatherData, setWeatherData] = useState({});
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [isRefreshing, setIsRefreshing] = useState(false);

    // Мемоизируем функцию fetchFavorites, чтобы не создавать новую при каждом рендере
    const fetchFavorites = useCallback(async () => {
        if (isRefreshing) return; // Предотвращаем параллельные запросы
        
        setLoading(true);
        setError(null);
        
        try {
            const response = await axios.get(`${API_BASE_URL}/favorites`);
            const favoriteCities = response.data;
            setFavorites(favoriteCities);
            
            if (favoriteCities.length > 0) {
                await fetchWeatherForFavorites(favoriteCities);
            } else {
                setWeatherData({});
            }
        } catch (error) {
            console.error('Error fetching favorite cities:', error);
            setError('Failed to load favorite cities');
        } finally {
            setLoading(false);
        }
    }, [isRefreshing]);

    // Используем useEffect с мемоизированной функцией fetchFavorites
    useEffect(() => {
        fetchFavorites();
    }, [fetchFavorites, updateTrigger]);

    const refreshFavorites = async () => {
        if (isRefreshing) return;
        
        setIsRefreshing(true);
        await fetchFavorites();
        setIsRefreshing(false);
    };

    const fetchWeatherForFavorites = async (favoriteCities) => {
        try {
            // Ограничиваем количество одновременных запросов до 5
            const batchSize = 5;
            const weatherData = {};
            
            // Обработка запросов батчами
            for (let i = 0; i < favoriteCities.length; i += batchSize) {
                const batch = favoriteCities.slice(i, i + batchSize);
                const promises = batch.map(city =>
                    axios.post(`${API_BASE_URL}/weather`, { city: city.city })
                        .then(response => {
                            if (response && response.data) {
                                weatherData[response.data.city] = response.data;
                            }
                            return response;
                        })
                        .catch(err => {
                            console.error(`Error fetching weather for ${city.city}:`, err);
                            return null;
                        })
                );
                
                await Promise.all(promises);
            }
            
            setWeatherData(weatherData);
        } catch (error) {
            console.error('Error fetching weather for favorite cities:', error);
            setError('Failed to load weather data for some cities');
        }
    };

    const deleteFavoriteCity = async (cityId, cityName) => {
        try {
            await axios.delete(`${API_BASE_URL}/favorite/${cityId}`);
            
            // Update local state
            setFavorites(favorites.filter(city => city.id !== cityId));
            
            // Update weather data
            const newWeatherData = { ...weatherData };
            if (newWeatherData[cityName]) {
                delete newWeatherData[cityName];
                setWeatherData(newWeatherData);
            }
        } catch (error) {
            console.error('Error deleting favorite city:', error);
            setError('Failed to delete city from favorites');
        }
    };

    if (loading && favorites.length === 0) {
        return (
            <div className="favorites-container">
                <div className="favorites-header">
                    <h2>Favorite Cities</h2>
                </div>
                <div className="loading-container">
                    <div className="loading-spinner"></div>
                    <p>Loading favorites...</p>
                </div>
            </div>
        );
    }

    if (favorites.length === 0 && !loading) {
        return (
            <div className="favorites-container">
                <div className="favorites-header">
                    <h2>Favorite Cities</h2>
                    <button 
                        className="refresh-button" 
                        onClick={refreshFavorites} 
                        disabled={isRefreshing}
                        title="Refresh favorites"
                    >
                        <span className={isRefreshing ? 'loading' : ''}>↻</span>
                    </button>
                </div>
                
                <div className="no-favorites">
                    <p>You don't have any favorite cities yet.</p>
                    <p>Search for a city above and add it to favorites.</p>
                </div>
            </div>
        );
    }

    return (
        <div className="favorites-container">
            <div className="favorites-header">
                <h2>Favorite Cities</h2>
                <button 
                    className="refresh-button" 
                    onClick={refreshFavorites} 
                    disabled={isRefreshing}
                    title="Refresh favorites"
                >
                    <span className={isRefreshing ? 'loading' : ''}>↻</span>
                </button>
            </div>

            {error && <div className="error-message">{error}</div>}
            
            <div className="favorites-grid">
                {favorites.map((favCity) => (
                    <div className="favorite-card" key={favCity.id}>
                        <div className="favorite-card-header">
                            <h3>{favCity.city}</h3>
                            <button 
                                className="delete-button" 
                                onClick={() => deleteFavoriteCity(favCity.id, favCity.city)}
                                title="Remove from favorites"
                            >
                                ×
                            </button>
                        </div>
                        
                        {weatherData[favCity.city] ? (
                            <div className="favorite-card-body">
                                <div className="favorite-temp">
                                    {Math.round(weatherData[favCity.city].temperature)}°C
                                </div>
                                <div className="favorite-description">
                                    {weatherData[favCity.city].weather}
                                </div>
                                
                                {weatherData[favCity.city].humidity !== undefined && (
                                    <div className="favorite-detail">
                                        <span>Humidity: {weatherData[favCity.city].humidity}%</span>
                                    </div>
                                )}
                                
                                {weatherData[favCity.city].wind_speed !== undefined && (
                                    <div className="favorite-detail">
                                        <span>Wind: {weatherData[favCity.city].wind_speed} m/s</span>
                                    </div>
                                )}
                            </div>
                        ) : (
                            <div className="favorite-card-loading">
                                <div className="loading-spinner small"></div>
                                <span>Loading weather data...</span>
                            </div>
                        )}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Favorites;
