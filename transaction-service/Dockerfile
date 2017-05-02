FROM gcr.io/google_appengine/python

RUN virtualenv /env

# Setting these environment variables are the same as running
# source /env/bin/activate.
ENV VIRTUAL_ENV -p python3.5 /env
ENV PATH /env/bin:$PATH

COPY requirements.txt /app/
RUN pip install --requirement /app/requirements.txt
COPY . /app/

ENTRYPOINT []

CMD ["python", "/app/server.py"]
