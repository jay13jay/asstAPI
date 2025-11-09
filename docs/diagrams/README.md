# Architecture Diagrams

This directory contains visual representations of the asstBackend API architecture, following the C4 framework for software architecture documentation.

## 📊 Available Diagrams

### L1 - System Context
**File**: [`asstBackend-l1-context.drawio`](asstBackend-l1-context.drawio)

Shows the big picture of how the asstBackend API fits into the overall system landscape, including:
- External actors (End Users, Developers, System Administrators)
- The asstBackend API system
- External systems (OpenAI GPT, PostgreSQL, Redis Cache)
- High-level interactions between actors and systems

**View in diagrams.net**: [Open L1 Context Diagram](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l1-context.drawio)

### L2 - Container Architecture
**File**: [`asstBackend-l2-containers.drawio`](asstBackend-l2-containers.drawio)

Zooms into the asstBackend API system to show the high-level technical building blocks:
- API Gateway / Web Server (Go + Gin)
- PostgreSQL Database
- LLM Provider Services (External)
- Redis Cache Layer (Future)
- Data flow between containers

**View in diagrams.net**: [Open L2 Container Diagram](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l2-containers.drawio)

### L3 - Component Breakdown
**File**: [`asstBackend-l3-components.drawio`](asstBackend-l3-components.drawio)

Decomposes the API Gateway container into its major components:
- **Handlers Layer**: UserHandler, ChatHandler, MessageHandler
- **Middleware**: AuthMiddleware, CORS
- **Services Layer**: UserService, ChatService, MessageService
- **Repositories Layer**: UserRepository, ChatRepository, MessageRepository
- **Database**: PostgreSQL
- Component interactions and dependencies

**View in diagrams.net**: [Open L3 Component Diagram](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l3-components.drawio)

### L4 - Code & API Specifications
**File**: [`asstBackend-l4-code.drawio`](asstBackend-l4-code.drawio)

Shows implementation-level details:
- API endpoint groups (Auth, Chats, Messages)
- Individual REST endpoints
- Data models (User, Chat, Message)
- Database schema mapping
- Request/response flows

**View in diagrams.net**: [Open L4 Code Diagram](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l4-code.drawio)

## 🎨 Viewing & Editing Diagrams

### Online Viewing
All diagrams can be opened directly in your browser using diagrams.net:
1. Go to [app.diagrams.net](https://app.diagrams.net)
2. Click **File** → **Open from** → **Device**
3. Select any `.drawio` file from this directory

### Desktop Application
For the best editing experience, download the desktop app:
- **Download**: [draw.io Desktop](https://github.com/jgraph/drawio-desktop/releases)
- **Platforms**: Windows, macOS, Linux

### VS Code Extension
Edit diagrams directly in VS Code:
- **Extension**: [Draw.io Integration](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio)
- **Install**: Search for "Draw.io Integration" in VS Code Extensions

## 🔧 Regenerating Diagrams

The diagrams are generated from the C4 documentation in [`docs/c4/`](../c4/) using a custom generator tool.

### Prerequisites
- Go 1.21 or higher
- Access to the repository root

### Build the Generator
```bash
cd /path/to/asstAPI
go build -o bin/diag-generator ./docs/diagrams/generator
```

### Generate Individual Diagrams

**L1 Context Diagram:**
```bash
./bin/diag-generator \
  --c4doc=docs/c4/L1-context.md \
  --level=L1 \
  --out=docs/diagrams/asstBackend-l1-context.drawio
```

**L2 Container Diagram:**
```bash
./bin/diag-generator \
  --c4doc=docs/c4/L2-container.md \
  --level=L2 \
  --out=docs/diagrams/asstBackend-l2-containers.drawio
```

**L3 Component Diagram:**
```bash
./bin/diag-generator \
  --c4doc=docs/c4/L3-component.md \
  --level=L3 \
  --out=docs/diagrams/asstBackend-l3-components.drawio
```

**L4 Code Diagram:**
```bash
./bin/diag-generator \
  --c4doc=docs/c4/L4-code.md \
  --level=L4 \
  --out=docs/diagrams/asstBackend-l4-code.drawio
```

### Generate All Diagrams
```bash
# From repository root
./bin/diag-generator --c4doc=docs/c4/L1-context.md --level=L1 --out=docs/diagrams/asstBackend-l1-context.drawio
./bin/diag-generator --c4doc=docs/c4/L2-container.md --level=L2 --out=docs/diagrams/asstBackend-l2-containers.drawio
./bin/diag-generator --c4doc=docs/c4/L3-component.md --level=L3 --out=docs/diagrams/asstBackend-l3-components.drawio
./bin/diag-generator --c4doc=docs/c4/L4-code.md --level=L4 --out=docs/diagrams/asstBackend-l4-code.drawio
```

### Generator Options
```
--c4doc string      Path to C4 markdown document
--level string      C4 level: L1, L2, L3, or L4
--out string        Output .drawio file path
--theme string      Diagram theme: light, dark, neutral (default "light")
--layout string     Layout: horizontal, vertical, layered (default "vertical")
--dry-run           Print summary without writing files
```

## 📝 Metadata Files

Each diagram has an accompanying `.meta.md` file containing:
- Generation timestamp
- Generator version
- Source C4 documentation files
- Assumptions and mappings made during generation
- Usage instructions

Example: `asstBackend-l1-context.drawio.meta.md`

## 🎯 Customization & Manual Edits

While diagrams are generated from C4 documentation, manual refinements are encouraged:

### Layout Adjustments
- **Node positioning**: Drag nodes to optimal locations
- **Edge routing**: Adjust connector paths for clarity
- **Spacing**: Use grid alignment (10px grid) for clean layout

### Styling Enhancements
- **Colors**: Follow C4 color conventions or customize for clarity
- **Labels**: Add detailed descriptions to nodes and edges
- **Icons**: Add icons from the diagram.net library for visual appeal

### Best Practices
1. **Keep it synchronized**: After manual edits, document changes in the diagram's notes
2. **Export images**: Generate PNG/SVG exports for use in presentations and documentation
3. **Version control**: Commit both `.drawio` and `.meta.md` files
4. **Consistency**: Use the same styles and conventions across all diagrams

## 🖼️ Exporting Diagrams

### Export from diagrams.net
1. Open the diagram
2. **File** → **Export as** → Choose format:
   - **PNG**: For embedding in documentation
   - **SVG**: For scalable web graphics
   - **PDF**: For print-quality documents

### Batch Export (Future)
A batch export script will be added to generate PNG/SVG exports of all diagrams automatically.

## 🔄 Keeping Diagrams Updated

Diagrams should be regenerated when:
- C4 documentation is significantly updated
- New components, containers, or systems are added
- Architecture relationships change
- Technology choices are modified

**Workflow:**
1. Update the relevant C4 documentation (`docs/c4/L*.md`)
2. Regenerate the affected diagram(s)
3. Review and make manual adjustments if needed
4. Commit both C4 docs and diagrams together

## 📚 Related Documentation

- [C4 Framework Documentation](../c4/README.md)
- [L1 Context Documentation](../c4/L1-context.md)
- [L2 Container Documentation](../c4/L2-container.md)
- [L3 Component Documentation](../c4/L3-component.md)
- [L4 Code Documentation](../c4/L4-code.md)

## 🛠️ Generator Implementation

The diagram generator is located at [`generator/`](generator/) and includes:
- `main.go`: CLI entry point
- `gen.go`: Core XML generation functions
- `builder.go`: C4-level specific builders
- `c4parser.go`: C4 markdown parser
- `gen_test.go`: Unit tests

### Architecture
The generator follows a builder pattern:
1. Parse C4 markdown documentation
2. Build level-specific nodes and edges
3. Generate mxGraph XML format
4. Write `.drawio` and `.meta.md` files

## 🤝 Contributing

When contributing changes to diagrams:
1. Update C4 documentation first
2. Regenerate affected diagrams
3. Review visual output in diagrams.net
4. Make any necessary manual refinements
5. Export preview images if making significant changes
6. Commit all changed files together

## 📄 License

These diagrams are part of the asstBackend API project and follow the same MIT License.

---

**Note**: Diagrams are living documents and should evolve with the codebase. Regular reviews and updates ensure they remain accurate and valuable for the team.
