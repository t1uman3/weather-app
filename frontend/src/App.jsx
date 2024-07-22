import { useState } from "react";
import axios from "axios";

function App() {
  const [city, setCity] = useState("");
  const [weather, setWeather] = useState(null);

  const getWeather = async () => {
    try {
      const response = await axios.post("http://127.0.0.1:8000/weather", {
        city,
      });
      setWeather(response.data);
      сonsole.log(response.data);
    } catch (error) {
      console.error("error fetching the weather", error);
    }
  };

  return (
    <div className="body m-4 min-h-screen bg-gray-100 flex flex-col items-center">
        <div className="m-4 flex items-center">
            <img className="weather_logo m-10"src="/img/weather_logo.png" alt="weather_logo"/>
            <h1 className=" text-3xl font-bold">Weather App</h1>
        </div>
      <div>
        <input
          type="text"
          value={city}
          onChange={(e) => setCity(e.target.value)}
          className="px-4 py-2 border rounded shadow"
          placeholder="Enter city name"
        />
        <button
          onClick={getWeather}
          className="greenButton ml-2 px-4 py-2 text-white rounded shadow"
        >
          Get Weather
        </button>
      </div>
      {weather && (
        <div className="Card mt-4 p-4 bg-white rounded shadow">
          <h2 className="text-2xl items-center font-bold">{weather.city}</h2>
          <p className="text-lg">Temperature: {weather.temperature} °C</p>
          <p className="text-lg">Weather: {weather.weather}</p>
        </div>
      )}
    </div>
  );
}

export default App;
