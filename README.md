## Overview

Hello, my name is Saknovskyi Oleksii, this is my solution to the test given by ContentQuo. This app makes `GET` request to the API https://catfact.ninja/breeds to get list of cat breeds (you can set the amount of cat breeds that app will GET from the API using `-limit` flag) and creates the file out.json.
That out.json file stores breeds grouped by country. Breeds in country groups sorted by their coat.

## How to run

Clone the repository: 

    git clone https://github.com/sosshik/catninja-api-test

Run the app(you can set the limit flag to any positive amount, by default it's 100): 

    go run main.go -limit=50