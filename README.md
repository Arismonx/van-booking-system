# Van Booking System Backend (Go)
This repository contains the backend API for a comprehensive van booking and management system, developed using **Go (Golang)**. The system facilitates seamless interactions between passengers, drivers, and administrators, offering functionalities for van scheduling, seat reservations, and payment processing.

Designed with a **monolithic architecture** for ease of development and maintenance as a solo project, the backend is structured into clear, distinct modules (Handler, Service, Repository, Model) to ensure a clean, maintainable, and scalable codebase.


## Key Features

* **User Management & Authentication:** Secure user registration, login, and role-based access control (for passengers and drivers) using **JWT (JSON Web Tokens)** and **bcrypt** for password hashing.
  
* **Driver Management:** Dedicated functionalities for drivers to manage their profiles, including linking to their primary user accounts.
  
* **Van Management:** Drivers can add, view, and update their fleet of vans, including details like license plates, models, and seating capacity.
  
* **Route Management:** Define and manage various travel routes (origin and destination points) within the system.
  
* **Schedule Management:** Drivers can create, view, and update detailed van schedules, specifying departure times, routes, prices per seat, and managing available seats.
  
* **Booking System:** Passengers can search for available van schedules, reserve seats, and manage their bookings.
  
* **Payment Processing:** Integration with a mock payment gateway for handling booking payments and managing payment statuses.
  
* **Robust Database Design:** Utilizes **PostgreSQL** with a well-defined schema, including `users`, `employee`, `vans`, `routes`, `schedules`, `bookings`, and `payments` tables, designed for data integrity and efficient querying.
  
* **RESTful API:** Provides a clear, well-structured RESTful API for seamless integration with a frontend application.
  
* **Dockerized Environment:** Includes `Dockerfile` and `docker-compose.yml` for easy setup and deployment of the Go application and PostgreSQL database in development and production environments.

## Technologies Used
* **Go (Golang):** Core programming language.
* **Gin Gonic:** High-performance web framework for building APIs.
* **GORM:** Elegant ORM (Object-Relational Mapping) for database interactions with PostgreSQL.
* **PostgreSQL:** Robust relational database for data storage.
* **Viper:** For efficient configuration management.
* **Go-Playground/Validator:** For declarative request payload validation.
* **Bcrypt:** For secure password hashing.
* **Go-JWT:** For JSON Web Token (JWT) handling.
* **Logrus:** For structured and customizable logging.
* **Docker:** For containerization and environment consistency.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/Arismonx/van-booking-system/blob/main/LICENSE) file for details.
