# pizza-delivery

- Note: This project was created for learning purposes, specifically to practice working with RabbitMQ, Docker, Websocket, the Ionic JavaScript framework and Air is a live reloading tool designed for Golang developers.

- Pizza Delivery is a modern and efficient food ordering system that allows customers to order pizzas online with real-time tracking. 

✔ WebSockets for Real-Time Communication
Maintains a persistent connection with the client.
Pushes real-time updates (e.g., order status changes, delivery tracking).

✔ RabbitMQ for Message Processing
Handles background tasks asynchronously.
Ensures reliable communication between services (e.g., backend → kitchen system → delivery).
 
🏗️ Tech Stack

Backend: Golang (Gin/GORM)

Frontend: JavaScript (Ionic Framework)

Messaging: RabbitMQ

Real-time order updates: Websocket

Deployment: Docker

📦 Installation & Setup 
# Clone the repository
git clone git@github.com:abdukarimxalilov/pizza-delivery.git
cd pizza-delivery

# Set up environment variables
cp .env.example .env

# Run backend (Golang)
air install && air

# Run frontend (Ionic)
npm install && ionic serve

🤝 Contributing

Fork the repository.

Create a new branch: git checkout -b feature-name

Commit changes: git commit -m 'Add new feature'

Push to the branch: git push origin feature-name

Open a Pull Request.

📝 License

This project is licensed under the MIT License.

