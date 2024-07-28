from sqlalchemy import Column, Integer, String, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

DATABASE_URL = "sqlite:///./test.db"

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

class FavoriteCity(Base):
    __tablename__ = "favorite_cities"

    id = Column(Integer, primary_key=True, index=True)
    city = Column(String, unique=True, index=True)

Base.metadata.create_all(bind=engine)

