from fastapi import FastAPI, Depends, HTTPException
from pydantic import BaseModel
from sqlalchemy.orm import Session
from models import Base, FavoriteCity, SessionLocal, engine
import requests
from settings import OPENWEATHER_API_KEY, OPENWEATHER_URL
from fastapi.middleware.cors import CORSMiddleware


app = FastAPI(
    title = "Weather App"
)

class CityRequest(BaseModel):
    city: str
class FavoriteCityRequest(BaseModel):
    city: str
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

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
        response.raise_for_status()
        weather_data = response.json()
        return {
            "city": city,
            "temperature": weather_data["main"]["temp"],
            "weather": weather_data["weather"][0]["description"]
        }
    except requests.exceptions.RequestException as e:
        raise HTTPException(status_code=500, detail=f"Error fetching weather data: {e}")

@app.post("/favorite")
def add_favorite_city(favorite_city: FavoriteCityRequest, db: Session = Depends(get_db)):
    db_city = db.query(FavoriteCity).filter(FavoriteCity.city == favorite_city.city).first()
    if db_city:
        raise HTTPException(status_code=400, detail="City already in favorites")
    new_city = FavoriteCity(city=favorite_city.city)
    db.add(new_city)
    db.commit()
    db.refresh(new_city)
    return new_city

@app.get("/favorites")
def get_favorite_cities(db: Session = Depends(get_db)):
    return db.query(FavoriteCity).all()

@app.delete("/favorite/{city_id}")
def delete_favorite_city(city_id: int, db: Session = Depends(get_db)):
    db_city = db.query(FavoriteCity).filter(FavoriteCity.id == city_id).first()
    if not db_city:
        raise HTTPException(status_code=404, detail="City not found")
    db.delete(db_city)
    db.commit()
    return {"detail": "City deleted"}


# Заголовки CORS
origins = [
    "http://localhost:3000"
    "http://127.0.0.1:3000"
    "http://localhost:5173",
    "http://127.0.0.1:5173",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allows all origins
    allow_credentials=True,
    allow_methods=["*"],  # Allows all methods
    allow_headers=["*"], )




