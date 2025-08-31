# Artisanal Tea

Artisanal Tea is a modern React frontend for the Artisanal Kettle backend, providing a user-friendly interface to manage, select, and execute remote commands on services and environments. It is designed for DevOps, SREs, and platform engineers who want a simple, beautiful UI for infrastructure operations.

## Features

- **Service & Environment Selection:**
  - Dropdowns populated dynamically from the backend API.
- **Command Submission:**
  - Enter and submit commands to selected services/environments.
- **Live Response Display:**
  - See command output and error messages in real time.
- **Environment Variable Configuration:**
  - API URLs and settings are loaded from `.env` for easy deployment.
- **Reusable Components:**
  - Modular React components for dropdowns, search, command bar, and buttons.
- **Robust Error Handling:**
  - User-friendly error messages and loading states.
- **Beautiful, Modern UI:**
  - Clean, responsive design for desktop and mobile.

## How It Works

- Fetches available services and environments from the Artisanal Kettle backend.
- Lets users select a service, environment, and command, then submit to the backend.
- Displays the backend's response (or error) in a styled response box.

## Setup & Usage

1. **Clone the repository**
2. **Install dependencies:**
   ```sh
   cd artisanal-tea
   npm install
   ```
3. **Configure environment variables:**
   - Copy `.env.example` to `.env` and set the backend API URLs as needed.
4. **Run the app:**
   ```sh
   npm start
   ```
   The app will be available at [http://localhost:3000](http://localhost:3000)

## Project Structure

- `src/`
  - `components/` — Reusable UI components (dropdown, searchbar, navbar, command bar, button)
  - `utils/` — API utility functions
  - `App.tsx` — Main app logic and state
  - `App.css` — Main styles

## Example .env

```
REACT_APP_API_BASE_URL=http://localhost:8080
```

## API Integration
- Works out of the box with the Artisanal Kettle backend.
- All API endpoints and error handling are environment-variable driven for flexibility.

## Contributing
Pull requests and issues are welcome! Please open an issue to discuss your feature or bugfix idea.

## License
MIT
