package main

import (
	"fmt"
	"time"
)

// DiagramBuilder creates diagrams from C4 documentation
type DiagramBuilder struct {
	Doc    *C4Doc
	Theme  string // light, dark, neutral
	Layout string // horizontal, vertical, layered
}

// Build creates nodes and edges based on C4 level
func (db *DiagramBuilder) Build() ([]Node, []Edge) {
	switch db.Doc.Level {
	case "L1":
		return db.buildL1()
	case "L2":
		return db.buildL2()
	case "L3":
		return db.buildL3()
	case "L4":
		return db.buildL4()
	default:
		return db.buildL2() // default to L2
	}
}

// buildL1 creates context diagram
func (db *DiagramBuilder) buildL1() ([]Node, []Edge) {
	nodes := []Node{}
	edges := []Edge{}

	// Central system
	systemID := NodeID("system", "asstBackend API")
	nodes = append(nodes, Node{
		ID:    systemID,
		Label: "asstBackend API\n[Software System]\n\nProvides chat-based AI assistant interactions",
		Style: "rounded=1;fillColor=#0b5fff;fontColor=#ffffff;fontSize=14;fontStyle=1",
		X:     300,
		Y:     200,
		W:     200,
		H:     120,
	})

	// Actors
	actorY := 50
	actors := []string{"End Users", "Developers", "System Admins"}
	for i, actor := range actors {
		actorID := NodeID("actor", actor)
		nodes = append(nodes, Node{
			ID:    actorID,
			Label: actor,
			Style: "shape=umlActor;verticalLabelPosition=bottom;verticalAlign=top;fillColor=#666666;fontColor=#333333;",
			X:     100 + i*250,
			Y:     actorY,
			W:     40,
			H:     60,
		})
		edges = append(edges, Edge{
			ID:     EdgeID(actorID, systemID),
			Source: actorID,
			Target: systemID,
			Label:  "Uses",
			Style:  "rounded=0;endArrow=block;",
		})
	}

	// External systems
	extSystems := []struct {
		name  string
		desc  string
		style string
	}{
		{"OpenAI GPT", "LLM Provider", "shape=cylinder;fillColor=#10B981;fontColor=#ffffff;"},
		{"PostgreSQL", "Database", "shape=cylinder;fillColor=#2b6cb0;fontColor=#ffffff;"},
		{"Redis Cache", "Cache Layer", "shape=cylinder;fillColor=#DC2626;fontColor=#ffffff;"},
	}

	extY := 400
	for i, ext := range extSystems {
		extID := NodeID("external", ext.name)
		nodes = append(nodes, Node{
			ID:    extID,
			Label: ext.name + "\n[" + ext.desc + "]",
			Style: ext.style,
			X:     150 + i*200,
			Y:     extY,
			W:     150,
			H:     80,
		})
		edges = append(edges, Edge{
			ID:     EdgeID(systemID, extID),
			Source: systemID,
			Target: extID,
			Label:  "Reads/Writes",
			Style:  "rounded=0;dashed=1;",
		})
	}

	return nodes, edges
}

// buildL2 creates container diagram
func (db *DiagramBuilder) buildL2() ([]Node, []Edge) {
	nodes := []Node{}
	edges := []Edge{}

	// Frontend app (external)
	frontendID := NodeID("external", "Frontend App")
	nodes = append(nodes, Node{
		ID:    frontendID,
		Label: "Web Frontend\n[Single Page App]",
		Style: "rounded=1;fillColor=#666666;fontColor=#ffffff;",
		X:     320,
		Y:     30,
		W:     160,
		H:     80,
	})

	// API Gateway / Web Server
	apiID := NodeID("container", "API Gateway")
	nodes = append(nodes, Node{
		ID:    apiID,
		Label: "API Gateway\n[Go + Gin]\n\nHTTP REST API, routing, middleware",
		Style: "rounded=1;fillColor=#0b5fff;fontColor=#ffffff;fontSize=12;",
		X:     300,
		Y:     180,
		W:     200,
		H:     100,
	})

	edges = append(edges, Edge{
		ID:     EdgeID(frontendID, apiID),
		Source: frontendID,
		Target: apiID,
		Label:  "JSON/HTTPS",
		Style:  "rounded=0;endArrow=block;",
	})

	// Database
	dbID := NodeID("container", "PostgreSQL")
	nodes = append(nodes, Node{
		ID:    dbID,
		Label: "PostgreSQL\n[Database]\n\nUsers, chats, messages",
		Style: "shape=cylinder;fillColor=#2b6cb0;fontColor=#ffffff;",
		X:     100,
		Y:     350,
		W:     160,
		H:     100,
	})

	edges = append(edges, Edge{
		ID:     EdgeID(apiID, dbID),
		Source: apiID,
		Target: dbID,
		Label:  "SQL queries",
		Style:  "rounded=0;",
	})

	// LLM Providers (external)
	llmID := NodeID("external", "LLM Providers")
	nodes = append(nodes, Node{
		ID:    llmID,
		Label: "LLM Providers\n[External API]\n\nOpenAI, Anthropic",
		Style: "rounded=1;fillColor=#10B981;fontColor=#ffffff;",
		X:     500,
		Y:     350,
		W:     180,
		H:     100,
	})

	edges = append(edges, Edge{
		ID:     EdgeID(apiID, llmID),
		Source: apiID,
		Target: llmID,
		Label:  "HTTPS/JSON",
		Style:  "rounded=0;dashed=1;",
	})

	// Redis (future)
	redisID := NodeID("container", "Redis Cache")
	nodes = append(nodes, Node{
		ID:    redisID,
		Label: "Redis Cache\n[Cache]\n\nSessions, rate limiting",
		Style: "shape=cylinder;fillColor=#DC2626;fontColor=#ffffff;",
		X:     320,
		Y:     500,
		W:     160,
		H:     80,
	})

	edges = append(edges, Edge{
		ID:     EdgeID(apiID, redisID),
		Source: apiID,
		Target: redisID,
		Label:  "Redis protocol",
		Style:  "rounded=0;dashed=1;",
	})

	return nodes, edges
}

// buildL3 creates component diagram
func (db *DiagramBuilder) buildL3() ([]Node, []Edge) {
	nodes := []Node{}
	edges := []Edge{}

	baseY := 50
	layerSpacing := 180

	// Handlers layer
	handlers := []string{"UserHandler", "ChatHandler", "MessageHandler"}
	for i, h := range handlers {
		hID := NodeID("component", h)
		nodes = append(nodes, Node{
			ID:    hID,
			Label: h + "\n[HTTP Handler]",
			Style: "rounded=1;fillColor=#0b5fff;fontColor=#ffffff;",
			X:     100 + i*220,
			Y:     baseY,
			W:     180,
			H:     70,
		})
	}

	// Middleware
	mwX := 700
	middleware := []string{"AuthMiddleware", "CORS"}
	for i, mw := range middleware {
		mwID := NodeID("component", mw)
		nodes = append(nodes, Node{
			ID:    mwID,
			Label: mw + "\n[Middleware]",
			Style: "rounded=1;fillColor=#7C3AED;fontColor=#ffffff;",
			X:     mwX,
			Y:     baseY + i*90,
			W:     150,
			H:     70,
		})
	}

	// Services layer
	services := []string{"UserService", "ChatService", "MessageService"}
	for i, s := range services {
		sID := NodeID("component", s)
		nodes = append(nodes, Node{
			ID:    sID,
			Label: s + "\n[Business Logic]",
			Style: "rounded=1;fillColor=#F59E0B;fontColor=#ffffff;",
			X:     100 + i*220,
			Y:     baseY + layerSpacing,
			W:     180,
			H:     70,
		})

		// Connect handler to service
		hID := NodeID("component", handlers[i])
		edges = append(edges, Edge{
			ID:     EdgeID(hID, sID),
			Source: hID,
			Target: sID,
			Label:  "calls",
			Style:  "rounded=0;endArrow=block;",
		})
	}

	// Repositories layer
	repos := []string{"UserRepository", "ChatRepository", "MessageRepository"}
	for i, r := range repos {
		rID := NodeID("component", r)
		nodes = append(nodes, Node{
			ID:    rID,
			Label: r + "\n[Data Access]",
			Style: "rounded=1;fillColor=#10B981;fontColor=#ffffff;",
			X:     100 + i*220,
			Y:     baseY + layerSpacing*2,
			W:     180,
			H:     70,
		})

		// Connect service to repository
		sID := NodeID("component", services[i])
		edges = append(edges, Edge{
			ID:     EdgeID(sID, rID),
			Source: sID,
			Target: rID,
			Label:  "uses",
			Style:  "rounded=0;endArrow=block;",
		})
	}

	// Database
	dbID := NodeID("database", "PostgreSQL")
	nodes = append(nodes, Node{
		ID:    dbID,
		Label: "PostgreSQL\n[Database]",
		Style: "shape=cylinder;fillColor=#2b6cb0;fontColor=#ffffff;",
		X:     300,
		Y:     baseY + layerSpacing*3,
		W:     180,
		H:     100,
	})

	// Connect all repos to database
	for _, r := range repos {
		rID := NodeID("component", r)
		edges = append(edges, Edge{
			ID:     EdgeID(rID, dbID),
			Source: rID,
			Target: dbID,
			Label:  "SQL",
			Style:  "rounded=0;",
		})
	}

	return nodes, edges
}

// buildL4 creates code/API diagram
func (db *DiagramBuilder) buildL4() ([]Node, []Edge) {
	nodes := []Node{}
	edges := []Edge{}

	// API Endpoints grouped
	endpoints := []struct {
		group string
		apis  []string
	}{
		{"Auth", []string{"POST /auth/register", "POST /auth/login"}},
		{"Chats", []string{"POST /chats", "GET /chats", "GET /chats/{id}", "PUT /chats/{id}", "DELETE /chats/{id}"}},
		{"Messages", []string{"POST /chats/{id}/messages", "GET /chats/{id}/messages"}},
	}

	baseX := 50
	baseY := 50
	groupSpacing := 250

	for gi, group := range endpoints {
		// Group container
		groupID := NodeID("api-group", group.group)
		groupHeight := len(group.apis)*60 + 40
		nodes = append(nodes, Node{
			ID:    groupID,
			Label: group.group + " APIs",
			Style: "rounded=1;fillColor=#E5E7EB;strokeColor=#6B7280;fontColor=#1F2937;dashed=1;",
			X:     baseX + gi*groupSpacing,
			Y:     baseY,
			W:     220,
			H:     groupHeight,
		})

		// Individual endpoints
		for i, api := range group.apis {
			endpointID := NodeID("endpoint", api)
			nodes = append(nodes, Node{
				ID:    endpointID,
				Label: api,
				Style: "rounded=0;fillColor=#0b5fff;fontColor=#ffffff;fontSize=10;",
				X:     baseX + gi*groupSpacing + 10,
				Y:     baseY + 30 + i*60,
				W:     200,
				H:     50,
			})
		}
	}

	// Data models
	models := []string{"User", "Chat", "Message"}
	modelY := 400
	for i, model := range models {
		modelID := NodeID("model", model)
		nodes = append(nodes, Node{
			ID:    modelID,
			Label: model + "\n[Data Model]",
			Style: "rounded=0;fillColor=#F59E0B;fontColor=#ffffff;",
			X:     100 + i*250,
			Y:     modelY,
			W:     180,
			H:     80,
		})
	}

	// Database schema
	schemaID := NodeID("schema", "Database Schema")
	nodes = append(nodes, Node{
		ID:    schemaID,
		Label: "PostgreSQL Schema\nusers | chats | messages",
		Style: "shape=cylinder;fillColor=#2b6cb0;fontColor=#ffffff;",
		X:     300,
		Y:     550,
		W:     200,
		H:     100,
	})

	// Connect models to schema
	for _, model := range models {
		modelID := NodeID("model", model)
		edges = append(edges, Edge{
			ID:     EdgeID(modelID, schemaID),
			Source: modelID,
			Target: schemaID,
			Label:  "maps to",
			Style:  "rounded=0;dashed=1;",
		})
	}

	return nodes, edges
}

// GenerateMetadata creates metadata content for a diagram
func GenerateMetadata(prompt, level, outputPath string, sources []string, assumptions []string) string {
	meta := fmt.Sprintf(`# Diagram Generation Metadata

**Generated**: %s
**Generator Version**: v1.0.0
**C4 Level**: %s

## Source Prompt
%s

## Source Files
`, time.Now().UTC().Format(time.RFC3339), level, prompt)

	for _, s := range sources {
		meta += fmt.Sprintf("- %s\n", s)
	}

	if len(assumptions) > 0 {
		meta += "\n## Assumptions & Mappings\n"
		for _, a := range assumptions {
			meta += fmt.Sprintf("- %s\n", a)
		}
	}

	meta += fmt.Sprintf("\n## Output\n- %s\n", outputPath)
	meta += "\n## Usage\nOpen the .drawio file in https://app.diagrams.net or desktop draw.io application for editing.\n"
	
	return meta
}
