# MockServer
A Mock Server capable of creating API endpoints dynamically for use with local testing and debugging

---

# Usage
Pull the image, set the port map, and add the volume for the directory where the schema configuration files are located. By default, the server will check for configuration files located in the `/app/data` directory. This can be changed via the env variable `CONFIG_PATH`, if desired.

    image: blankdev117/gomockserver:0.0.1
    volumes:
    - "./Mocks/apis:/app/data"
    networks:
    - network
    ports:
    - "7000:8080"

---

# Api Configuration

The schema is how the configuration file, json or otherwise, should be organized to be interpreted by the mock server processor.

Currently Supported File Types: `.json`

### V1

Configuration files:
- _can_ indicate what version of schema is being used, though a missing version indicates `v1`
- _must_ include a `routes` object, with their listed methods and responses as subobjects

Route configurations:
- _can_ use wild cards(`*`) for generic response catch alls
- _can_ include url path subsitutions in the responses
- _must_ include a status code for every method defined
- _can_ include a response body object that will be returned as json to the requesting client

Notes:
 - Substitutions are case sensitive and the format is `{var}` for the path and `{{var}}` for the response bodies
 - Substitutions are only effective up to a wild card for response bodies

An example setup `.json` file is defined below, but may be different depending on the type of parser being used (i.e. json, yml, etc.):

```json
{
    "routes": {
        "/*": {
            "get": {
                "statusCode": 200,
                "body": {
                    "Success": "Amazing!"
                }
            }
        },
        "/{tenantId}/hooray/{id}": {
            "post": {
                "statusCode": 201,
                "body": {
                    "id": "{{id}}",
                    "tenantId": "{{tenantId}}",
                    "IsGood": true
                }
            },
            "get": {
                "statusCode": 200
            }
        }
    }
}
```

In the above example, 3 routes will be generated:
 - A wild card route that will return a 200 status code and body for any `GET` that does not have a more specific route defined
 - 2 routes for the url `/{tenantId}/hooray/{id}`:
   - `POST`: returns 201 status code with a body that will have the `tenantId` and `id` fields replaced by the respective url parameters
   - `GET`: reeturns a 200 status code with an empty body
