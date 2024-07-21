from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import requests
from settings import OPENWEATHER_API_KEY, OPENWEATHER_URL
from fastapi.middleware.cors import CORSMiddleware


app = FastAPI(
    title = "Weather App"
)

class CityRequest(BaseModel):
    city: str

@app.get("/")
def main():
    return {
        "message": "backend is ready",
        "version": "0.1.0"
    }

@app.post("/weather")
def get_weather(city_request: CityRequest):
    city = city_request.city
    req_parameters = {
        "q": city,
        "appid": OPENWEATHER_API_KEY,
        "units": "metric"
    }

    try:
        response = requests.get(OPENWEATHER_URL, params=req_parameters)
        response.raise_for_status()  # Проверяем статус ответа
        weather_data = response.json()
        return {
            "city": city,
            "temperature": weather_data["main"]["temp"],
            "weather": weather_data["weather"][0]["description"]
        }
    except requests.exceptions.RequestException as e:
        raise HTTPException(status_code=500, detail=f"Error fetching weather data: {e}")

origins = [
    "http://localhost:5173",
    "http://127.0.0.1:5173",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allows all origins
    allow_credentials=True,
    allow_methods=["*"],  # Allows all methods
    allow_headers=["*"], )




