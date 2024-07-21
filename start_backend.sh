
#!/bin/sh

# Start the backend service
cd backend
python -m venv env
source env/bin/activate
pip install -r requirements.txt
uvicorn main:app --reload

