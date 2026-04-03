<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Architectural Engineering Guide
  ### Designing and Managing Distributed Systems with Precision
  
  [![License](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Status](https://img.shields.io/badge/Status-Production--Grade--Draft-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Architecture is about managing complexity and enforcing autonomy through decoupled engineering."**

  ---
</div>

## Introduction: The Philosophy of Microservices

Microservices architecture is the evolution of software engineering to handle extreme scale and complexity. It moves away from the monolithic "all-in-one" model towards a cellular structure where each service is an independent, autonomous unit specialized in a specific business domain.

This repository serves as a focused roadmap to mastering distributed systems, focusing on reliability, scalability, and maintainability.

---

## Chapter 1: The Monolith Crisis and Evolution

Software starts simple. A single repo, a single database. However, as organizations scale, the monolith becomes a bottleneck:
- **Tight Coupling**: A change in one module breaks another.
- **Scaling Issues**: You must scale the entire application even if only one feature is under load.
- **Operational Drag**: Huge build times and high risk of catastrophic failure.

Microservices solve this by decoupling teams and technology, allowing for independent deployment and specialized scaling.

---

## Chapter 2: The Art of Decomposition (DDD)

Breaking a system is a strategic decision. We use **Domain-Driven Design (DDD)** to identify **Bounded Contexts**.
- **Catalog Service**: Owns product identity and pricing.
- **Order Service**: Owns order lifecycle and fulfillment.

In this project, we have already implemented `services/product-service` and `services/order-service` as distinct entities to demonstrate this separation.

---

## Chapter 3: Communication and Protocols

How do independent services talk? We balance performance and flexibility.

1.  **gRPC (Synchronous)**:
    - **Usage**: Internal, service-to-service calls.
    - **Why**: Protocol Buffers provide high-performance, strongly-typed binary communication. ✨
2.  **Messaging (Asynchronous)**:
    - **Usage**: Event-driven architecture.
    - **Why**: Decouples services so they can fail and recover independently without bringing down the system.

---

## Chapter 4: Data Management Strategies

The golden rule: **Each service owns its own database.**
- **Saga Pattern**: Managing transactions across services without global locks.
- **CQRS**: Separating Read and Write models for maximum query performance.

---

## Chapter 5: Operational Excellence

Building the service is only half the battle. Running it at scale requires:
- **Graceful Shutdown**: Handling OS signals to close connections cleanly.
- **Health Checks**: Ensuring dependencies are ready before traffic starts flowing.
- **Service Mesh**: Managing complex network topologies outside of the application code.

---

## Laboratory: Hands-on Execution

The repository is equipped with a `Makefile` to orchestrate the entire development environment:

```bash
# Start Infrastructure (PostgreSQL, RabbitMQ, Redis)
make up

# Run Product Service (Go)
make run-product

# Run Order Service (Go)
make run-order
```

---

## Architecture Roadmap

| Phase | Module | Focus | Status |
| :--- | :--- | :--- | :---: |
| 1 | Foundations | Paradigm Shift | ![100%](https://geps.dev/progress/100) |
| 2 | Clean Architecture | Go Project Layout | ![100%](https://geps.dev/progress/100) |
| 3 | Communication | gRPC Protocols | ![100%](https://geps.dev/progress/100) |
| 4 | Containerization | Docker & Networks | ![100%](https://geps.dev/progress/100) |
| 5 | API Gateway | Security & Routing | ![20%](https://geps.dev/progress/20) |
| 6 | Messaging | Kafka/RabbitMQ | ![10%](https://geps.dev/progress/10) |

<div align="center">
  <br/>
  <sub>Engineering Excellence | **arch-yunus**</sub>
</div>
