# React frontend service

## Run

#### Run locally

Install dependencies and start the app
```bash
npm install
npm start
```

Open [localhost:3000](http://localhost:3000)

#### Run Docker container

```bash
docker build -t react-frontend .
docker run --publish 5000:5000 --name react-frontend-test --rm react-frontend
```

Open [localhost:5000](http://localhost:5000)

## Credits

The React app was bootstraped with [create-react-app](https://github.com/facebookincubator/create-react-app).
