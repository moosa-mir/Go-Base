# Domain-Driven Design (DDD) Architecture

This document outlines the new Domain-Driven Design architecture for the EDRGO project.

## Overview

The codebase has been refactored to follow Domain-Driven Design principles, organizing code around business domains rather than technical layers. This approach provides better separation of concerns, improved maintainability, and clearer business logic organization.

## Domain Structure

### 1. User Domain (`internal/domain/user/`)
**Core Business Concept**: User management and authentication

**Entities**:
- `User`: Core user entity with profile information
- `UserStatus`: Enum for user account status

**Services**:
- `UserService`: Business logic for user operations
- Registration, authentication, profile management

**Repository**:
- `UserRepository`: Data access contract for users

**Errors**:
- Domain-specific error types for user operations

### 2. Product Domain (`internal/domain/product/`)
**Core Business Concept**: Product catalog and inventory

**Entities**:
- `Product`: Core product entity with details and pricing
- `ProductCategory`: Product categorization

**Services**:
- `ProductService`: Business logic for product operations
- Product creation, search, inventory management

**Repository**:
- `ProductRepository`: Data access contract for products

**Errors**:
- Domain-specific error types for product operations

### 3. Order Domain (`internal/domain/order/`)
**Core Business Concept**: Order management and processing

**Entities**:
- `Order`: Core order entity with items and status
- `OrderItem`: Individual items within an order
- `OrderStatus`: Order lifecycle status
- `OrderItemStatus`: Individual item status

**Services**:
- `OrderService`: Business logic for order operations
- Order creation, status updates, order history

**Repository**:
- `OrderRepository`: Data access contract for orders

**Errors**:
- Domain-specific error types for order operations

### 4. Payment Domain (`internal/domain/payment/`)
**Core Business Concept**: Payment processing and management

**Entities**:
- `Payment`: Core payment entity
- `PaymentStatus`: Payment lifecycle status
- `PaymentMethod`: Supported payment methods

**Services**:
- `PaymentService`: Business logic for payment operations
- Payment processing, refunds, payment history

**Repository**:
- `PaymentRepository`: Data access contract for payments

**Errors**:
- Domain-specific error types for payment operations

### 5. Wallet Domain (`internal/domain/wallet/`)
**Core Business Concept**: Digital wallet management

**Entities**:
- `Wallet`: Core wallet entity with balance

**Services**:
- `WalletService`: Business logic for wallet operations
- Balance management, transactions

**Repository**:
- `WalletRepository`: Data access contract for wallets

**Errors**:
- Domain-specific error types for wallet operations

### 6. Basket Domain (`internal/domain/basket/`)
**Core Business Concept**: Shopping cart management

**Entities**:
- `Basket`: Core basket entity
- `BasketItem`: Items in the shopping cart

**Services**:
- `BasketService`: Business logic for basket operations
- Add/remove items, quantity management

**Repository**:
- `BasketRepository`: Data access contract for baskets

**Errors**:
- Domain-specific error types for basket operations

### 7. Seller Domain (`internal/domain/seller/`)
**Core Business Concept**: Seller management

**Entities**:
- `Seller`: Core seller entity
- `SellerStatus`: Seller account status

**Services**:
- `SellerService`: Business logic for seller operations
- Seller registration, profile management

**Repository**:
- `SellerRepository`: Data access contract for sellers

**Errors**:
- Domain-specific error types for seller operations

## DDD Principles Applied

### 1. Ubiquitous Language
- Domain terminology is consistently used throughout the codebase
- Business concepts are clearly defined and named
- Technical implementation details are hidden behind domain abstractions

### 2. Bounded Contexts
- Each domain is a separate bounded context
- Clear boundaries between different business areas
- Minimal coupling between domains

### 3. Entities and Value Objects
- **Entities**: Objects with identity (User, Product, Order, etc.)
- **Value Objects**: Immutable objects without identity (Status enums, etc.)
- Rich domain models with behavior, not just data

### 4. Domain Services
- Business logic is encapsulated in domain services
- Services coordinate between entities and repositories
- Complex business rules are centralized

### 5. Repository Pattern
- Data access is abstracted through repository interfaces
- Domain logic is separated from data access concerns
- Easy to swap implementations (e.g., for testing)

### 6. Domain Events (Future Enhancement)
- Events can be added to capture domain state changes
- Enables loose coupling between domains
- Supports event-driven architecture

## Benefits of DDD Architecture

### 1. **Business Focus**
- Code organization reflects business domains
- Easier for business stakeholders to understand
- Business rules are clearly expressed

### 2. **Maintainability**
- Changes to business logic are localized to specific domains
- Reduced coupling between different parts of the system
- Easier to test individual domains

### 3. **Scalability**
- Domains can be developed and deployed independently
- Clear boundaries enable team autonomy
- Easier to add new features within existing domains

### 4. **Testability**
- Domain logic can be tested in isolation
- Repository interfaces enable easy mocking
- Business rules are clearly defined and testable

### 5. **Flexibility**
- Easy to change implementations without affecting business logic
- Repository pattern enables different data storage strategies
- Domain services can be enhanced without breaking existing code

## Migration Strategy

### Phase 1: Domain Structure (Current)
- ✅ Created domain entities and services
- ✅ Defined repository interfaces
- ✅ Established domain-specific errors

### Phase 2: Repository Implementation
- 🔄 Implement repository interfaces with existing database code
- 🔄 Migrate existing handlers to use domain services
- 🔄 Update database layer to work with new domain models

### Phase 3: Application Layer
- ⏳ Create application services that coordinate between domains
- ⏳ Update HTTP handlers to use application services
- ⏳ Implement cross-domain business logic

### Phase 4: Infrastructure
- ⏳ Add domain events for cross-domain communication
- ⏳ Implement caching strategies
- ⏳ Add monitoring and logging

## Next Steps

1. **Implement Repository Layer**: Create concrete implementations of repository interfaces
2. **Update Handlers**: Migrate existing HTTP handlers to use domain services
3. **Add Application Services**: Create application layer services for complex operations
4. **Implement Domain Events**: Add event-driven communication between domains
5. **Add Integration Tests**: Ensure domains work together correctly

## File Structure

```
internal/
├── domain/
│   ├── user/
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── product/
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── order/
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── payment/
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── wallet/
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── basket/
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── repository.go
│   │   └── service.go
│   └── seller/
│       ├── entity.go
│       ├── errors.go
│       ├── repository.go
│       └── service.go
```

This DDD architecture provides a solid foundation for building a maintainable, scalable, and business-focused application.



