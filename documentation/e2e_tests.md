# E2E tests

- Implemented with [Cypress](https://www.cypress.io)
- E2E-tests utilize MongoDB’s collection SlowersTest, which is is automatically created and cleared before running the tests.
- Tutorial used for writing tests with Cypress can be found from FullStack Open’s course [material](https://fullstackopen.com/en/part5/end_to_end_testing_cypress)

## How to run E2E-tests

1. Inside ```frontend``` directory, install the dependencies using the command ```npm install``` and create a ```.env``` file inside the backend directory following instructions in README.
2. Run the tests either in interactive mode or command-line mode.
    1. **Option 1 (recommended)**: To run E2E-tests in interactive test runner mode, use command ```npm run e2e``` inside the frontend directory. 
    2. **Option 2**: Tests can be run from the command line using the command ```npx cypress run``` while the backend is running with the command ```ENV=test go run .``` and frontend with the command ```npm run dev```.

### Running E2E-tests in the interactive test runner mode

1. If ran using the command `npm run e2e`, the [script](https://github.com/Slowers-Team/Slowers-App/blob/486e601aab91a22354b13165a5d03b5c6ab2be3e/frontend/package.json#L14) will automatically start both backend and frontend, and open Cypress with the command ```cypress open```. 
2. Select ”E2E Testing”.
   ![Kuvankaappaus 2024-11-18 kello 13 49 02 ip](https://github.com/user-attachments/assets/f7e21a14-fa86-4676-b3bb-2bf88c6f70d8)

3. Choose your preferred browser, and press ”Start E2E Testing in ___”.
  ![Kuvankaappaus 2024-11-18 kello 13 50 06 ip](https://github.com/user-attachments/assets/03ceeea8-c522-4701-8b02-18cd0fc15af1)

4. Click ”slowers_app.cy.js”
![Kuvankaappaus 2024-11-18 kello 13 51 54 ip](https://github.com/user-attachments/assets/c2419446-2744-44bd-bded-a5e5e3dc1c97)

5. The tests are now running, and you’ll be able to see passing and failing tests.
![Kuvankaappaus 2024-11-18 kello 13 55 15 ip](https://github.com/user-attachments/assets/a841957c-5b29-45be-b24e-8d512b8b7510)


## Notes

- E2E-tests currently face issues when the frontend is run in production mode. It is recommended to only run E2E-tests while running frontend in development mode.
- Cypress seems to use the system's default browser language. Changing this default is not straightforward, and this affects tests requiring language-specific validation, such as error messages during user registration.
