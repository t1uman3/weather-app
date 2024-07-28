import React, { useState } from 'react';
import axios from 'axios';

const FavoritesComponent = () => {
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
        <div className="m-4 min-h-screen bg-gray-100 flex flex-col items-center">

            <h2 className=" text-3xl font-bold">Weather on Favorite Cities</h2>
            <button className="greenButton ml-2 px-4 py-2 text-white rounded shadow"
                    onClick={fetchFavorites}>Show Favorite Cities
            </button>
            <div>
                {favorites.map((favCity) => (
                    <div className="Card mt-4 p-4 bg-white rounded shadow"
                         key={favCity.id}>
                        <h2 className="text-2xl items-center font-bold">{favCity.city}</h2>
                        {weatherData[favCity.city] && (
                            <div>
                                <p className="text-lg">Temperature: {weatherData[favCity.city].temperature}°C</p>
                                <p className="text-lg">Weather: {weatherData[favCity.city].weather}</p>
                            </div>

                        )}
                     <button
                            className="redButton px-4 py-2 text-white rounded shadow"
                            onClick={() => deleteFavoriteCity(favCity.id)}>
                            Delete
                     </button>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default FavoritesComponent;
