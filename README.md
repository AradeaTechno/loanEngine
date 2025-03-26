    ## LOAN ENGINE WITH GO
    This project is base structure to create a RestFul using Go. In this project I tried to solve loan mechanism process.

    ---

    ## Features

    - STAFF
    - BORROWER
    - LOAN

    --- 

    ## Accessibility
    Except `REGISTER` and `LOGIN`, all endpoints need `BEARER TOKEN AUTHORIZATION` or will return `401`. Use Bearer token you will get from `/login` endpoint.

    ## Requirements
    Ensure your system meets the following requirements:
    - PHP >= 8.3
    - Composer >= 2.x
    - Laravel >= 12.x
    - MySQL >= 5.8

    ## Installation

    ### Step 1: Clone the Repository
    ```bash
    git clone https://github.com/AradeaTechno/be-laravel.git
    cd be-laravel
    ```

    ## Step 2: Install Dependencies

    Run the following command to install PHP dependencies:
    ```bash
    composer install
    ```

    If your project uses Laravel Mix, install Node.js dependencies:
    ```bash
    npm install
    ```

    ## Step 3: Weather API key
    Create **[weatherapi](https://www.weatherapi.com/)** account and get your weather API key 

    ## Step 4: Environment Configuration
    Create the .env file from the example:
    ```bash
    cp .env.example .env
    ```
    Update the .env file with your environment variables:
    ```bash
    APP_NAME=JuiceBox
    APP_ENV=local
    APP_KEY=base64:XQQXTbxoKGmfMfA2JgF1ae4acMHrEFGvU2KpjWcqce4=
    APP_DEBUG=true
    APP_URL=http://localhost

    APP_LOCALE=en
    APP_FALLBACK_LOCALE=en
    APP_FAKER_LOCALE=en_US

    APP_MAINTENANCE_DRIVER=file
    # APP_MAINTENANCE_STORE=database

    PHP_CLI_SERVER_WORKERS=4

    BCRYPT_ROUNDS=12

    LOG_CHANNEL=stack
    LOG_STACK=single
    LOG_DEPRECATIONS_CHANNEL=null
    LOG_LEVEL=debug

    DB_CONNECTION=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_DATABASE=your_db_name
    DB_USERNAME=your_db_user
    DB_PASSWORD=your_db_pass

    SESSION_DRIVER=database
    SESSION_LIFETIME=120
    SESSION_ENCRYPT=false
    SESSION_PATH=/
    SESSION_DOMAIN=null

    BROADCAST_CONNECTION=log
    FILESYSTEM_DISK=local
    QUEUE_CONNECTION=database

    CACHE_STORE=database
    CACHE_DRIVER=file
    # CACHE_PREFIX=

    MEMCACHED_HOST=127.0.0.1

    REDIS_CLIENT=phpredis
    REDIS_HOST=127.0.0.1
    REDIS_PASSWORD=null
    REDIS_PORT=6379

    MAIL_MAILER=smtp
    MAIL_HOST=your_smtp_host
    MAIL_PORT=587
    MAIL_USERNAME=your_smtp_user@example.com
    MAIL_PASSWORD=your_smtp_pass
    MAIL_ENCRYPTION=tls
    MAIL_FROM_ADDRESS="your_smtp_user@example.com"
    MAIL_FROM_NAME="${APP_NAME}"

    AWS_ACCESS_KEY_ID=
    AWS_SECRET_ACCESS_KEY=
    AWS_DEFAULT_REGION=us-east-1
    AWS_BUCKET=
    AWS_USE_PATH_STYLE_ENDPOINT=false

    VITE_APP_NAME="${APP_NAME}"

    # WEATHER API
    WEATHERAPI_KEY=Your_weather_api_key

    ```

    ## Step 5: Generate Application Key
    Run the following command to generate the application key:
    ```bash
    php artisan key:generate
    ```

    ## Step 6: Configure Database
    Set up your database and run the migrations:
    ```bash
    php artisan migrate
    ```

    ## Step 7: Configure Cache
    ```bash
    chmod -R 775 storage/framework/cache
    chmod -R 775 storage/framework/sessions
    chmod -R 775 storage/logs
    ```

    ## USAGE
    Start the Local Server
    Run the following command to start the server:
    ```bash
    php artisan serve
    ```
    Visit the application in your browser at http://127.0.0.1:8000

    ## ALTERNATIVE TESTING
    You can use postman to test all endpoinst contains in this project. I also include `JuiceBoxApi.postman_collection.json` and you can import into your postman app. Import the collection into your postman and it ready to be used. 

    ## Test Weather Service API
    Use the endpoint /api/weather to fetch current weather data for a default or specified location (By default is set to Perth):
    ```bash
    GET /api/weather?location=Perth
    ```

    ## New User Registration and Email Notification
    When a user registers via the /api/register endpoint, a welcome email is sent automatically. Example:
    ```bash
    POST /api/register
    Content-Type: application/json

    {
        "name": "John Doe",
        "email": "john.doe@example.com",
        "password": "securepassword"
    }
    ```

    ## Email Queue
    Emails are sent asynchronously via Laravel queues. Ensure the queue worker is running:
    ```bash
    php artisan queue:work
    ```

    ## Manually dispatch email command
    In this project I included manual artisan command to send email for your user using this command:
    ```bash
    php artisan email:send-welcome <user_id>
    ```
    For example
    ```bash
    php artisan email:send-welcome 13
    ```
    This Command wil return:
    ```bash
    Welcome email dispatched to user: your_targeted_user_email@example.com
    ```

    ## Background Jobs
    ## Weather Updates
    Set up the Laravel scheduler to run the weather update job every 15 minutes:

    1. Add the following cron job:
    ```bash
    */15 * * * * php /path-to-your-project/artisan schedule:run >> /dev/null 2>&1
    ```
    2. Start scheduler:
    ```bash
    php artisan schedule:work
    ```

    ## Testing
    Run Unit and Feature Tests
    Ensure all tests pass:

    To run all tests at once;
    ```bash
    php artisan --filter=ApiEndpointsTest
    ```

    Should Return:
    ```bash
    PASS  Tests\Feature\ApiEndpointsTest
    ✓ can create post                                                                                                                                                                                  0.21s  
    ✓ can get all post                                                                                                                                                                                 0.13s  
    ✓ can create user                                                                                                                                                                                  0.13s  
    ✓ can get all user                                                                                                                                                                                 0.07s  
    ✓ can call weather                                                                                                                                                                                 0.12s  

    Tests:    5 passed (180 assertions)
    Duration: 0.69s

    ```

    To run test by name `e.g`:
    ```bash
    php artisan test --filter=ApiEndpointsTest::test_can_get_all_post
    ```

    Should Return:
    ```bash
    PASS  Tests\Feature\ApiEndpointsTest
    ✓ can get all post                                                                                                                                                                                 0.36s  

    Tests:    1 passed (78 assertions)
    Duration: 0.49s

    ```

    ## ADDITIONAL SETTING
    In your `phpunit.xml` file change the following setting to make test run well:
    ```bash
    <env name="DB_CONNECTION" value="mysql"/>
    <env name="DB_DATABASE" value="YourDatabaseName"/>
    ```