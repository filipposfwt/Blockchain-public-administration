#!/bin/bash

# Make this script executable with: chmod +x docker-commands.sh

# Function to display help
show_help() {
  echo "Docker Commands Helper Script"
  echo ""
  echo "Usage: ./docker-commands.sh [command]"
  echo ""
  echo "Commands:"
  echo "  start       - Build and start all containers"
  echo "  stop        - Stop all containers"
  echo "  restart     - Restart all containers"
  echo "  logs        - Show logs from all containers"
  echo "  logs-f      - Follow logs from all containers"
  echo "  logs-fe     - Show frontend logs"
  echo "  logs-be     - Show backend logs"
  echo "  logs-db     - Show database logs"
  echo "  clean       - Stop and remove containers, networks, and volumes"
  echo "  rebuild     - Rebuild and restart all containers"
  echo "  status      - Show status of containers"
  echo "  help        - Show this help message"
  echo ""
}

# Check if a command was provided
if [ $# -eq 0 ]; then
  show_help
  exit 1
fi

# Process commands
case "$1" in
  start)
    echo "Starting containers..."
    docker-compose up -d
    ;;
  stop)
    echo "Stopping containers..."
    docker-compose down
    ;;
  restart)
    echo "Restarting containers..."
    docker-compose restart
    ;;
  logs)
    echo "Showing logs..."
    docker-compose logs
    ;;
  logs-f)
    echo "Following logs..."
    docker-compose logs -f
    ;;
  logs-fe)
    echo "Showing frontend logs..."
    docker-compose logs frontend
    ;;
  logs-be)
    echo "Showing backend logs..."
    docker-compose logs backend
    ;;
  logs-db)
    echo "Showing database logs..."
    docker-compose logs db
    ;;
  clean)
    echo "Cleaning up containers, networks, and volumes..."
    docker-compose down -v
    ;;
  rebuild)
    echo "Rebuilding and restarting containers..."
    docker-compose down
    docker-compose up -d --build
    ;;
  status)
    echo "Container status:"
    docker-compose ps
    ;;
  help)
    show_help
    ;;
  *)
    echo "Unknown command: $1"
    show_help
    exit 1
    ;;
esac

exit 0 