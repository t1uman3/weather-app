import React, { useState } from 'react';
import axios from 'axios';

const Favorites = () => {
    const [favorites, setFavorites] = useState([]);
    const [weatherData, setWeatherData] = useState({});

    const fetchFavorites = async () => {
        try {
            const response = await axios.get('http://127.0.0.1:8000/favorites');
            const favoriteCities = response.data;
            setFavorites(favoriteCities);
            fetchWeatherForFavorites(favoriteCities);
        } catch (error) {
            console.error('Error fetching favorite cities:', error);
        }
    };

    const fetchWeatherForFavorites = async (favoriteCities) => {
        try {
            const weatherPromises = favoriteCities.map(city =>
                axios.post('http://127.0.0.1:8000/weather', { city: city.city })
            );
            const weatherResponses = await Promise.all(weatherPromises);
            const weatherData = weatherResponses.reduce((acc, response) => {
                acc[response.data.city] = response.data;
                return acc;
            }, {});
            setWeatherData(weatherData);
        } catch (error) {
            console.error('Error fetching weather for favorite cities:', error);
        }
    };

    const deleteFavoriteCity = async (cityId) => {
        try {
            await axios.delete(`http://127.0.0.1:8000/favorite/${cityId}`);
            setFavorites(favorites.filter(city => city.id !== cityId));
            const newWeatherData = { ...weatherData };
            delete newWeatherData[cityId];
            setWeatherData(newWeatherData);
        } catch (error) {
            console.error('Error deleting favorite city:', error);
        }
    };

    return (
        <div>
            <div className="m-4 flex items-center">
                <img className="cursor-pointer size-5" src="/img/liked.svg" alt="favorite"
                     onClick={fetchFavorites}></img>
                <h2 className="text-2xl items-center m-5 font-bold">Favorite Cities</h2>
            </div>
            <div className="favorites-container flex flex-wrap mr-2">
                {favorites.map((favCity) => (
                    <div className="Card mt-2 mr-4 p-4 bg-white rounded shadow"
                         key={favCity.id}>
                    <h2 className="text-2xl items-center font-bold">{favCity.city}</h2>
                        {weatherData[favCity.city] && (
                            <div>
                                <p className="text-lg">Temperature: {weatherData[favCity.city].temperature}°C</p>
                                <p className="text-lg">Weather: {weatherData[favCity.city].weather}</p>
                            </div>

                        )}

                        <button
                            className="redButton px-2 py-2 text-white rounded shadow"
                            onClick={() => deleteFavoriteCity(favCity.id)}>
                            Delete
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Favorites;
