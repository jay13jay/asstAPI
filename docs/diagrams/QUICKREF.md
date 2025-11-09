# Architecture Diagrams - Quick Reference

This document provides a quick overview of the architecture diagrams available for the asstBackend API project.

## 📋 Summary

Four comprehensive architecture diagrams have been generated from the C4 documentation, covering all levels from system context to code-level details.

| Level | Name | Description | Nodes | Edges | File |
|-------|------|-------------|-------|-------|------|
| **L1** | Context | System landscape with external actors and systems | 7 | 6 | [asstBackend-l1-context.drawio](asstBackend-l1-context.drawio) |
| **L2** | Container | High-level architecture with API, DB, and external services | 5 | 4 | [asstBackend-l2-containers.drawio](asstBackend-l2-containers.drawio) |
| **L3** | Component | Internal components: handlers, services, repositories | 12 | 9 | [asstBackend-l3-components.drawio](asstBackend-l3-components.drawio) |
| **L4** | Code | API endpoints, data models, and database schema | 16 | 3 | [asstBackend-l4-code.drawio](asstBackend-l4-code.drawio) |

## 🎯 Quick Actions

### View Diagrams Online
Click any link below to open the diagram in your browser (no installation required):

- **[View L1 Context →](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l1-context.drawio)**
- **[View L2 Container →](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l2-containers.drawio)**
- **[View L3 Component →](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l3-components.drawio)**
- **[View L4 Code →](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l4-code.drawio)**

### Edit Diagrams
1. Download any `.drawio` file
2. Go to [app.diagrams.net](https://app.diagrams.net)
3. Click **File → Open from → Device**
4. Select the downloaded file

### Regenerate Diagrams
```bash
# Build the generator
go build -o bin/diag-generator ./docs/diagrams/generator

# Regenerate all diagrams
./bin/diag-generator --c4doc=docs/c4/L1-context.md --level=L1 --out=docs/diagrams/asstBackend-l1-context.drawio
./bin/diag-generator --c4doc=docs/c4/L2-container.md --level=L2 --out=docs/diagrams/asstBackend-l2-containers.drawio
./bin/diag-generator --c4doc=docs/c4/L3-component.md --level=L3 --out=docs/diagrams/asstBackend-l3-components.drawio
./bin/diag-generator --c4doc=docs/c4/L4-code.md --level=L4 --out=docs/diagrams/asstBackend-l4-code.drawio
```

## 📐 Diagram Details

### L1 - System Context
**Purpose**: Shows the big picture - how the system fits into its environment

**Key Elements**:
- Three types of users: End Users, Developers, System Administrators
- Central asstBackend API system
- External systems: OpenAI GPT, PostgreSQL, Redis Cache
- User-to-system and system-to-system relationships

**Use Cases**:
- Executive presentations
- Onboarding new team members
- Understanding system boundaries
- Identifying external dependencies

---

### L2 - Container Architecture
**Purpose**: Zooms into the system to show major technical building blocks

**Key Elements**:
- Web Frontend (Single Page App) - external user interface
- API Gateway (Go + Gin) - main application container
- PostgreSQL Database - data persistence
- LLM Providers (OpenAI, Anthropic) - external AI services
- Redis Cache - session and rate limiting (future)

**Use Cases**:
- Infrastructure planning
- Deployment architecture design
- Technology stack decisions
- Scalability planning

---

### L3 - Component Breakdown
**Purpose**: Shows internal structure of the API Gateway container

**Key Elements**:
- **Handlers Layer**: UserHandler, ChatHandler, MessageHandler
- **Middleware**: AuthMiddleware, CORS
- **Services Layer**: UserService, ChatService, MessageService
- **Repositories Layer**: UserRepository, ChatRepository, MessageRepository
- **Database**: PostgreSQL with SQL connections

**Use Cases**:
- Code organization understanding
- Developer onboarding
- Identifying code ownership
- Planning refactoring efforts

---

### L4 - Code & API Specifications
**Purpose**: Implementation-level details for developers

**Key Elements**:
- **Auth APIs**: Registration and login endpoints
- **Chat APIs**: CRUD operations for chat management
- **Message APIs**: Message creation and retrieval
- **Data Models**: User, Chat, Message entities
- **Database Schema**: Table structures and relationships

**Use Cases**:
- API integration
- Frontend development
- Database schema understanding
- Testing and validation

## 🎨 Color Coding

Diagrams follow standard C4 color conventions:

| Color | Meaning | Example |
|-------|---------|---------|
| ![#0b5fff](https://via.placeholder.com/15/0b5fff/000000?text=+) Blue | Internal containers/components | API Gateway, Services |
| ![#666666](https://via.placeholder.com/15/666666/000000?text=+) Gray | External actors/systems | Users, Frontend App |
| ![#2b6cb0](https://via.placeholder.com/15/2b6cb0/000000?text=+) Dark Blue | Databases | PostgreSQL |
| ![#10B981](https://via.placeholder.com/15/10B981/000000?text=+) Green | External services | LLM Providers |
| ![#DC2626](https://via.placeholder.com/15/DC2626/000000?text=+) Red | Cache/Queue | Redis Cache |
| ![#F59E0B](https://via.placeholder.com/15/F59E0B/000000?text=+) Orange | Business logic | Services |
| ![#7C3AED](https://via.placeholder.com/15/7C3AED/000000?text=+) Purple | Middleware | Auth, CORS |

## 🔄 Keeping Diagrams Updated

### When to Update
- Architecture changes (new components, containers, or systems)
- Technology stack changes
- API endpoint additions or modifications
- Data model changes

### How to Update
1. Update the relevant C4 documentation (`docs/c4/L*.md`)
2. Regenerate the affected diagram(s) using the generator
3. Open in diagrams.net for manual refinements if needed
4. Commit both documentation and diagram files together

### Workflow
```bash
# 1. Edit C4 doc
vim docs/c4/L2-container.md

# 2. Regenerate diagram
./bin/diag-generator --c4doc=docs/c4/L2-container.md --level=L2 --out=docs/diagrams/asstBackend-l2-containers.drawio

# 3. Review and refine in diagrams.net (optional)
# Open the file at https://app.diagrams.net

# 4. Commit changes
git add docs/c4/L2-container.md docs/diagrams/asstBackend-l2-containers.drawio*
git commit -m "Update L2 container architecture"
```

## 📚 Additional Resources

- [Full Diagrams Documentation](README.md)
- [C4 Framework Documentation](../c4/README.md)
- [diagrams.net User Guide](https://www.diagrams.net/doc/)
- [C4 Model](https://c4model.com/)

## ✅ Validation

All diagrams have been validated:
- ✓ Valid XML format
- ✓ Compatible with diagrams.net/draw.io
- ✓ Deterministic IDs for version control
- ✓ Metadata files included
- ✓ Consistent styling and conventions

---

**Last Generated**: 2025-11-09  
**Generator Version**: v1.0.0  
**Source**: docs/c4/*.md
