# go-clean-simple-monolith
# ===============================
# Dev.Dockerfile — instructions
# ===============================
#
# Usage with docker-compose:
#
# 1. Build and run dev environment:
#    GITHUB_USER=your_user GITHUB_TOKEN=your_token docker-compose up --build
#
# 2. For M1 (arm64):
#    PLATFORM=arm64 GITHUB_USER=your_user GITHUB_TOKEN=your_token docker-compose up --build
#
# 3. For amd64 (default):
#    PLATFORM=amd64 GITHUB_USER=your_user GITHUB_TOKEN=your_token docker-compose up --build
#
# 4. For debugging:
#    Connect to port 8012 (api) or 8013 (worker) using GoLand/VSCode (dlv headless)
#
# 5. Hot-reload:
#    Any changes in .go files will automatically rebuild and restart the service (via reflex)
#
# 6. CMD for launch is set in docker-compose.yml (can be changed for any service)
#
# ===============================