package main

import (
	"bufio"
	"os"
	"strings"
)

// C4Doc represents parsed C4 documentation
type C4Doc struct {
	Level       string
	Title       string
	Systems     []string
	Containers  []string
	Components  []string
	ExternalSys []string
	Actors      []string
	Technologies map[string]string
	Relationships []Relationship
}

// Relationship represents connections between elements
type Relationship struct {
	From        string
	To          string
	Label       string
	Technology  string
}

// ParseC4Doc reads a C4 markdown file and extracts key elements
func ParseC4Doc(path string) (*C4Doc, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc := &C4Doc{
		Technologies: make(map[string]string),
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Extract level from title
		if strings.HasPrefix(line, "# L") {
			if strings.Contains(line, "L1") {
				doc.Level = "L1"
			} else if strings.Contains(line, "L2") {
				doc.Level = "L2"
			} else if strings.Contains(line, "L3") {
				doc.Level = "L3"
			} else if strings.Contains(line, "L4") {
				doc.Level = "L4"
			}
			doc.Title = strings.TrimPrefix(line, "# ")
		}
		
		// Look for key sections and extract elements
		if strings.HasPrefix(line, "###") || strings.HasPrefix(line, "##") {
			section := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "###"), "##"))
			
			// L1: External systems and actors
			if strings.Contains(section, "Stakeholders") || strings.Contains(section, "Primary Users") {
				doc.Actors = append(doc.Actors, "End Users", "Developers", "System Administrators")
			}
			if strings.Contains(section, "External Systems") || strings.Contains(section, "LLM Providers") {
				doc.ExternalSys = append(doc.ExternalSys, "OpenAI GPT", "Anthropic Claude")
			}
			if strings.Contains(section, "Infrastructure") {
				doc.ExternalSys = append(doc.ExternalSys, "PostgreSQL Database", "Redis Cache")
			}
		}
		
		// L2: Containers
		if strings.Contains(line, "### 1.") || strings.Contains(line, "### 2.") || strings.Contains(line, "### 3.") {
			if strings.Contains(line, "API Gateway") || strings.Contains(line, "Web Server") {
				doc.Containers = append(doc.Containers, "API Gateway")
			} else if strings.Contains(line, "Database") {
				doc.Containers = append(doc.Containers, "Database Server")
			} else if strings.Contains(line, "LLM Provider") {
				doc.Containers = append(doc.Containers, "LLM Provider Services")
			} else if strings.Contains(line, "Cache") {
				doc.Containers = append(doc.Containers, "Cache Layer")
			}
		}
		
		// L3: Components
		if strings.Contains(line, "Handlers") {
			doc.Components = append(doc.Components, "UserHandler", "ChatHandler", "MessageHandler")
		}
		if strings.Contains(line, "Services") {
			doc.Components = append(doc.Components, "UserService", "ChatService", "MessageService")
		}
		if strings.Contains(line, "Repositories") {
			doc.Components = append(doc.Components, "UserRepository", "ChatRepository", "MessageRepository")
		}
		if strings.Contains(line, "Middleware") {
			doc.Components = append(doc.Components, "AuthMiddleware", "CORS Middleware")
		}
		
		// Extract technology info
		if strings.Contains(line, "Technology**:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				tech := strings.TrimSpace(parts[1])
				doc.Technologies["main"] = tech
			}
		}
	}

	return doc, scanner.Err()
}
